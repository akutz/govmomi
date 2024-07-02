/*
Copyright (c) 2024-2024 VMware, Inc. All Rights Reserved.

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

package simulator

import (
	"context"
	"testing"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/types"
)

func TestVmCompatibilityChecker(t *testing.T) {
	ctx := context.Background()
	m := VPX()
	defer m.Remove()

	err := m.Create()
	if err != nil {
		t.Fatal(err)
	}

	s := m.Service.NewServer()
	defer s.Close()

	c, err := govmomi.NewClient(ctx, s.URL, true)
	if err != nil {
		t.Fatal(err)
	}

	finder := find.NewFinder(c.Client, true)
	datacenter, err := finder.DefaultDatacenter(ctx)
	if err != nil {
		t.Fatalf("default datacenter not found: %s", err)
	}
	finder.SetDatacenter(datacenter)
	vmList, err := finder.VirtualMachineList(ctx, "*")
	if len(vmList) == 0 {
		t.Fatal("vmList == 0")
	}
	vm := vmList[0]
	vmRef := vm.Reference()

	vmcc := object.NewVmCompatibilityChecker(c.Client)

	t.Run("CheckCompatibility", func(t *testing.T) {
		result, err := vmcc.CheckCompatibility(
			ctx,
			vm.Reference(),
			nil,
			nil)

		if err != nil {
			t.Fatal(err)
		}
		if len(result.Error) > 0 {
			t.Fatal("result has errors")
		}
		if len(result.Warning) > 0 {
			t.Fatal("result has warnings")
		}
	})

	t.Run("CheckVmConfig", func(t *testing.T) {
		t.Run("for existing VM", func(t *testing.T) {
			result, err := vmcc.CheckVmConfig(
				ctx,
				types.VirtualMachineConfigSpec{
					NumCPUs: 8,
				},
				&vmRef,
				nil,
				nil)

			if err != nil {
				t.Fatal(err)
			}
			if len(result.Error) > 0 {
				t.Fatal("result has errors")
			}
			if len(result.Warning) > 0 {
				t.Fatal("result has warnings")
			}
		})
		t.Run("for new VM", func(t *testing.T) {
			result, err := vmcc.CheckVmConfig(
				ctx,
				types.VirtualMachineConfigSpec{
					Name:    "my-vm",
					NumCPUs: 8,
				},
				&vmRef,
				nil,
				nil)

			if err != nil {
				t.Fatal(err)
			}
			if len(result.Error) > 0 {
				t.Fatal("result has errors")
			}
			if len(result.Warning) > 0 {
				t.Fatal("result has warnings")
			}
		})
	})
}
