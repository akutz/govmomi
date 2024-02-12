/*
Copyright (c) 2023-2024 VMware, Inc. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package model

import (
	"context"
	"flag"
	"fmt"
	"reflect"
	"strings"

	"github.com/vmware/govmomi/govc/cli"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

type tree struct {
}

func init() {
	cli.Register("object.model.properties", &tree{})
}

func (cmd *tree) Register(ctx context.Context, f *flag.FlagSet) {
}

func (cmd *tree) Description() string {
	return `Print the object model's properties.

Examples:
  govc object.model.properties
  govc object.model.properties VirtualMachine`
}

func (cmd *tree) Process(ctx context.Context) error {
	return nil
}

func (cmd *tree) Run(ctx context.Context, f *flag.FlagSet) error {

	typeName := f.Arg(0)
	if typeName == "" {
		return fmt.Errorf("empty type name")
	}

	t, ok := types.TypeFunc()(typeName)
	if !ok {
		if t, ok = mo.TypeFunc()(typeName); !ok {
			return fmt.Errorf("%q not found", typeName)
		}
	}

	return printProperties(
		t,
		nil,
		stack[string]{refs: map[string]struct{}{}},
		getGovmomiTypes())
}

func printProperties(
	t reflect.Type,
	names []string,
	tstck stack[string],
	gtyps []reflect.Type) error {

	tn := fmt.Sprintf("%s.%s", t.PkgPath(), t.Name())

	if tn != "." {
		if tstck.Exists(tn) {
			fmt.Println(strings.Join(names, "."))
			return nil
		}
		tstck.Push(tn)
		defer tstck.Pop()
	}

	if strings.Contains(tn, "Device[].VirtualDevice") {
		return nil
	}

	switch t.Kind() {
	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			tf := t.Field(i)
			if !tf.IsExported() {
				continue
			}
			if tf.Anonymous {
				printProperties(tf.Type, names, tstck, gtyps)
			} else {
				tfn := tf.Name
				if tfk := tf.Type.Kind(); tfk == reflect.Array ||
					tfk == reflect.Slice ||
					tfk == reflect.Map {

					tfn = fmt.Sprintf("%s[]", tfn)
				}
				printProperties(tf.Type, append(names, tfn), tstck, gtyps)
			}
		}
	case reflect.Interface:
		switch t.Name() {
		case "AnyType", "BaseMethodFault":
			fmt.Println(strings.Join(names, "."))
			return nil
		default:
			for i := 0; i < len(gtyps); i++ {
				gt := gtyps[i]
				//fmt.Printf("gt=%s\n", gt.Name())
				if gt.Implements(t) {
					//	fmt.Printf("gt=%s implements %s\n", gt.Name(), t.Name())
					printProperties(gt, append(names, gt.Name()), tstck, gtyps)
				} else if pgt := reflect.New(gt).Type(); pgt.Implements(t) {
					//	fmt.Printf("pgt=%s implements %s\n", pgt.Name(), t.Name())
					printProperties(pgt, append(names, gt.Name()), tstck, gtyps)
				}
			}
		}
	case reflect.Pointer:
		return printProperties(reflect.Zero(t.Elem()).Type(), names, tstck, gtyps)
	case reflect.Array, reflect.Slice:
		return printProperties(reflect.Zero(t.Elem()).Type(), names, tstck, gtyps)
	case reflect.Map:
		return printProperties(reflect.Zero(t.Elem()).Type(), names, tstck, gtyps)
	case reflect.Bool,
		reflect.Float32, reflect.Float64,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.String:

		fmt.Println(strings.Join(names, "."))
	}

	return nil
}

// //go:linkname typelinks reflect.typelinks
// func typelinks() (sections []unsafe.Pointer, offset [][]int32)

// //go:linkname add reflect.add
// func add(p unsafe.Pointer, x uintptr, whySafe string) unsafe.Pointer

// func getGovmomiTypes() []reflect.Type {
// 	var result []reflect.Type
// 	sections, offsets := typelinks()
// 	for i := range sections {
// 		base := sections[i]
// 		for _, offset := range offsets[i] {
// 			typeAddr := add(base, uintptr(offset), "")
// 			typ3 := reflect.TypeOf(*(*any)(unsafe.Pointer(&typeAddr)))
// 			realType := typ3
// 			switch typ3.Kind() {
// 			case reflect.Array, reflect.Chan, reflect.Map, reflect.Pointer, reflect.Slice:
// 				realType = reflect.Zero(typ3.Elem()).Type()
// 			}
// 			if realType.Name() != "" {
// 				result = append(result, realType)
// 			}

// 			// if strings.HasPrefix(typ3.PkgPath(), "github.com/vmware/govmomi") {
// 			// 	realType := reflect.Zero(typ3.Elem()).Type()
// 			// 	fmt.Println(realType)
// 			// 	result = append(result, realType)
// 			// }
// 		}
// 	}
// 	return result
// }

func getGovmomiTypes() []reflect.Type {
	var result []reflect.Type
	for _, t := range types.AllTypes() {
		result = append(result, t)
	}
	for _, t := range mo.AllTypes() {
		result = append(result, t)
	}
	return result
}

type stack[T comparable] struct {
	data []T
	refs map[T]struct{}
}

func (s *stack[T]) Push(t T) {
	s.data = append(s.data, t)
	s.refs[t] = struct{}{}
}

func (s *stack[T]) Pop() {
	if len(s.data) == 0 {
		return
	}
	t := s.data[len(s.data)-1]
	delete(s.refs, t)
	s.data = s.data[:len(s.data)-1]
}

func (s *stack[T]) Exists(t T) bool {
	_, ok := s.refs[t]
	return ok
}
