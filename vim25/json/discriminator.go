// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package json

import (
	"reflect"
)

// DiscriminatorToTypeFunc is used to get a reflect.Type from its
// discriminator.
type DiscriminatorToTypeFunc func(discriminator string) (reflect.Type, bool)

func (d *decodeState) getDiscriminatorValue() (reflect.Value, bool) {
	if !d.isDiscriminatorSet() {
		return reflect.Value{}, false
	}

	// Create a temporary decodeState that is initialized with the data from
	// the offset onwards, prefixed with a "{" character. This should provide
	// an enclosed, complex object that can be decoded into the object below.
	dd := &decodeState{}
	dd.init(append([]byte{'{'}, d.data[d.off:]...))
	defer freeScanner(&dd.scan)

	obj := map[string]interface{}{}
	if err := dd.unmarshal(&obj); err != nil {
		return reflect.Value{}, false
	}

	typeName, _ := obj[d.discriminatorTypeFieldName].(string)
	if typeName == "" {
		return reflect.Value{}, false
	}

	t, ok := d.discriminatorToTypeFn(typeName)
	if !ok {
		return reflect.Value{}, false
	}

	tv := reflect.New(t).Elem()

	if dv, ok := obj[d.discriminatorValueFieldName]; ok {
		v := reflect.ValueOf(dv)
		if v.Type() != t {
			v = v.Convert(t)
		}
		tv.Set(v)
	}

	isPtr, _ := obj[d.discriminatorByAddrFieldName].(bool)
	if isPtr {
		tv = tv.Addr()
	}

	return tv, true
}

type discriminatorInterfaceEncoder struct{}

func (d discriminatorInterfaceEncoder) encode(e *encodeState, v reflect.Value, opts encOpts) bool {
	v = v.Elem()

	if v.Kind() == reflect.Ptr && v.Elem().Kind() == reflect.Struct {
		v = v.Elem()
		se := structEncoder{
			fields:   cachedTypeFields(v.Type()),
			typeName: v.Type().Name(),
			byAddr:   true,
		}
		se.encode(e, v, opts)
		return true
	}

	var isPtr bool
	if v.Kind() == reflect.Ptr && isPrimitiveKind(v.Elem().Kind()) {
		isPtr = true
		v = v.Elem()
	}

	if !isPrimitiveKind(v.Kind()) {
		return false
	}

	e.WriteString(`{"`)
	e.WriteString(opts.discriminatorTypeFieldName)
	e.WriteString(`":"`)
	e.WriteString(v.Type().Name())
	e.WriteString(`","`)
	if isPtr {
		e.WriteString(opts.discriminatorByAddrFieldName)
		e.WriteString(`":`)
		e.reflectValue(reflect.ValueOf(true), opts)
		e.WriteString(`,"`)
	}
	e.WriteString(opts.discriminatorValueFieldName)
	e.WriteString(`":`)
	e.reflectValue(v, opts)
	e.WriteByte('}')
	return true
}

func isPrimitiveKind(k reflect.Kind) bool {
	switch k {
	case reflect.Bool,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Uintptr,
		reflect.Float32, reflect.Float64,
		reflect.Complex64, reflect.Complex128,
		reflect.String:
		return true
	}
	return false
}
