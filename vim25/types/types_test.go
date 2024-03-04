/*
Copyright (c) 2014-2015 VMware, Inc. All Rights Reserved.

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

package types

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/vmware/govmomi/vim25/xml"
)

func TestManagedObjectReference(t *testing.T) {
	t.Run("Marshal", func(t *testing.T) {
		testCases := []struct {
			name    string
			obj     ManagedObjectReference
			expXML  string
			expJSON string
		}{
			{
				name: "CumServerGUID",
				obj: ManagedObjectReference{
					Type:       "fake",
					Value:      "fake",
					ServerGUID: "fake",
				},
				expXML:  `<ManagedObjectReference type="fake" serverGuid="fake">fake</ManagedObjectReference>`,
				expJSON: `{"_typeName":"ManagedObjectReference","type":"fake","value":"fake","serverGuid":"fake"}`,
			},
			{
				name: "SansServerGUID",
				obj: ManagedObjectReference{
					Type:  "fake",
					Value: "fake",
				},
				expXML:  `<ManagedObjectReference type="fake">fake</ManagedObjectReference>`,
				expJSON: `{"_typeName":"ManagedObjectReference","type":"fake","value":"fake"}`,
			},
		}

		for i := range testCases {
			tc := testCases[i] // capture the test case

			t.Run(tc.name, func(t *testing.T) {
				t.Run("XML", func(t *testing.T) {
					if data, err := xml.Marshal(tc.obj); err != nil {
						t.Errorf("unexpected error: %v", err)
					} else if a, e := string(data), tc.expXML; a != e {
						t.Errorf("unexpected value: a=%q, e=%q", a, e)
					}
				})
				t.Run("JSON", func(t *testing.T) {
					var w bytes.Buffer
					if err := NewJSONEncoder(&w).Encode(tc.obj); err != nil {
						t.Errorf("unexpected error: %v", err)
					}
					assert.JSONEq(t, tc.expJSON, w.String(), "unexpected value")
				})
			})
		}
	})

	t.Run("ToFromString", func(t *testing.T) {
		testCases := []struct {
			name   string
			refStr string
			refObj ManagedObjectReference
			expErr error
		}{
			{
				name:   "EmptyInput",
				expErr: fmt.Errorf(`invalid ref: ""`),
			},
			{
				name:   "MissingValue",
				refStr: "VirtualMachine",
				expErr: fmt.Errorf(`invalid ref: "VirtualMachine"`),
			},
			{
				name:   "MissingType",
				refStr: "vm-123",
				expErr: fmt.Errorf(`invalid ref: "vm-123"`),
			},
			{
				name:   "TooMuchInfo",
				refStr: "VirtualMachine:vm-123:7972e324-1844-4242-8d6b-edd847152da7:InvalidField",
				expErr: fmt.Errorf(`invalid ref: "VirtualMachine:vm-123:7972e324-1844-4242-8d6b-edd847152da7:InvalidField"`),
			},
			{
				name:   "RefSansServerGUID",
				refStr: "VirtualMachine:vm-123",
				refObj: ManagedObjectReference{
					Type:  "VirtualMachine",
					Value: "vm-123",
				},
			},
			{
				name:   "RefCumServerGUID",
				refStr: "VirtualMachine:vm-123:7972e324-1844-4242-8d6b-edd847152da7",
				refObj: ManagedObjectReference{
					ServerGUID: "7972e324-1844-4242-8d6b-edd847152da7",
					Type:       "VirtualMachine",
					Value:      "vm-123",
				},
			},
		}

		for i := range testCases {
			tc := testCases[i]

			t.Run(tc.name, func(t *testing.T) {

				t.Run("FromString", func(t *testing.T) {
					var ref ManagedObjectReference
					ok := ref.FromString(tc.refStr)
					if a, e := ok, tc.expErr; a && e != nil {
						t.Errorf("unexpected error: %v", a)
					} else if !a && e == nil {
						t.Errorf("unexpected failure sans error")
					} else if a && ref.IsEmpty() {
						t.Errorf("invalid result: empty, ok")
					} else if !a && !ref.IsEmpty() {
						t.Errorf("invalid result: !empty, !ok")
					} else if a, e := ref, tc.refObj; a != e {
						t.Errorf("unexpected ref: a=%s, e=%s", a, e)
					}
				})

				t.Run("ParseManagedObjectReference", func(t *testing.T) {
					ref, err := ParseManagedObjectReference(tc.refStr)
					if a, e := err, tc.expErr; a != nil && e == nil {
						t.Errorf("unexpected error: %v", a)
					} else if a == nil && e != nil {
						t.Errorf("expected error did not occur: %s", e)
					} else if a != nil && e != nil {
						if a.Error() != e.Error() {
							t.Errorf("unexpected error: a=%v, e=%v", a, e)
						}
					} else if ref.IsEmpty() {
						t.Errorf("unexpected failure sans error")
					} else if a, e := ref, tc.refObj; a != e {
						t.Errorf("unexpected ref: a=%s, e=%s", a, e)
					}
				})

				t.Run("String", func(t *testing.T) {
					s := tc.refObj.String()
					if tc.expErr == nil {
						if a, e := s, tc.refStr; a != e {
							t.Errorf("unexpected value: a=%q, e=%q", a, e)
						}
					} else if a, e := s, ""; a != e {
						t.Errorf("unexpected value: a=%q, e=%q", a, e)
					}
				})
			})
		}
	})

	t.Run("EncodeDecode", func(t *testing.T) {
		testCases := []struct {
			name   string
			refStr string
			refObj ManagedObjectReference
			expErr error
		}{
			{
				name:   "EmptyInput",
				expErr: fmt.Errorf(`invalid ref: ""`),
			},
			{
				name:   "MissingValue",
				refStr: "VirtualMachine",
				expErr: fmt.Errorf(`invalid ref: "VirtualMachine"`),
			},
			{
				name:   "MissingType",
				refStr: "vm-123",
				expErr: fmt.Errorf(`invalid ref: "vm-123"`),
			},
			{
				name:   "TooMuchInfo",
				refStr: "VirtualMachine.vm-123.7972e324-1844-4242-8d6b-edd847152da7.InvalidField",
				expErr: fmt.Errorf(`invalid ref: "VirtualMachine.vm-123.7972e324-1844-4242-8d6b-edd847152da7.InvalidField"`),
			},
			{
				name:   "RefSansServerGUID",
				refStr: "VirtualMachine.vm-123",
				refObj: ManagedObjectReference{
					Type:  "VirtualMachine",
					Value: "vm-123",
				},
			},
			{
				name:   "RefCumServerGUID",
				refStr: "VirtualMachine.vm-123.7972e324-1844-4242-8d6b-edd847152da7",
				refObj: ManagedObjectReference{
					ServerGUID: "7972e324-1844-4242-8d6b-edd847152da7",
					Type:       "VirtualMachine",
					Value:      "vm-123",
				},
			},
		}

		for i := range testCases {
			tc := testCases[i]

			t.Run(tc.name, func(t *testing.T) {

				t.Run("Encode", func(t *testing.T) {
					s := tc.refObj.Encode()
					if tc.expErr == nil {
						if a, e := s, tc.refStr; a != e {
							t.Errorf("unexpected value: a=%q, e=%q", a, e)
						}
					} else if a, e := s, ""; a != e {
						t.Errorf("unexpected value: a=%q, e=%q", a, e)
					}
				})

				t.Run("Decode", func(t *testing.T) {
					var ref ManagedObjectReference
					ok := ref.Decode(tc.refStr)
					if a, e := ok, tc.expErr; a && e != nil {
						t.Errorf("unexpected error: %v", a)
					} else if !a && e == nil {
						t.Errorf("unexpected failure sans error")
					} else if a && ref.IsEmpty() {
						t.Errorf("invalid result: empty, ok")
					} else if !a && !ref.IsEmpty() {
						t.Errorf("invalid result: !empty, !ok")
					} else if a, e := ref, tc.refObj; a != e {
						t.Errorf("unexpected ref: a=%s, e=%s", a, e)
					}
				})

				t.Run("DecodeManagedObjectReference", func(t *testing.T) {
					ref, err := DecodeManagedObjectReference(tc.refStr)
					if a, e := err, tc.expErr; a != nil && e == nil {
						t.Errorf("unexpected error: %v", a)
					} else if a == nil && e != nil {
						t.Errorf("expected error did not occur: %s", e)
					} else if a != nil && e != nil {
						if a.Error() != e.Error() {
							t.Errorf("unexpected error: a=%v, e=%v", a, e)
						}
					} else if ref.IsEmpty() {
						t.Errorf("unexpected failure sans error")
					} else if a, e := ref, tc.refObj; a != e {
						t.Errorf("unexpected ref: a=%s, e=%s", a, e)
					}
				})
			})
		}
	})
}

func TestVirtualMachineConfigSpec(t *testing.T) {
	spec := VirtualMachineConfigSpec{
		Name:     "vm-001",
		GuestId:  "otherGuest",
		Files:    &VirtualMachineFileInfo{VmPathName: "[datastore1]"},
		NumCPUs:  1,
		MemoryMB: 128,
		DeviceChange: []BaseVirtualDeviceConfigSpec{
			&VirtualDeviceConfigSpec{
				Operation: VirtualDeviceConfigSpecOperationAdd,
				Device: &VirtualLsiLogicController{VirtualSCSIController{
					SharedBus: VirtualSCSISharingNoSharing,
					VirtualController: VirtualController{
						BusNumber: 0,
						VirtualDevice: VirtualDevice{
							Key: 1000,
						},
					},
				}},
			},
			&VirtualDeviceConfigSpec{
				Operation:     VirtualDeviceConfigSpecOperationAdd,
				FileOperation: VirtualDeviceConfigSpecFileOperationCreate,
				Device: &VirtualDisk{
					VirtualDevice: VirtualDevice{
						Key:           0,
						ControllerKey: 1000,
						UnitNumber:    new(int32), // zero default value
						Backing: &VirtualDiskFlatVer2BackingInfo{
							DiskMode:        string(VirtualDiskModePersistent),
							ThinProvisioned: NewBool(true),
							VirtualDeviceFileBackingInfo: VirtualDeviceFileBackingInfo{
								FileName: "[datastore1]",
							},
						},
					},
					CapacityInKB: 4000000,
				},
			},
			&VirtualDeviceConfigSpec{
				Operation: VirtualDeviceConfigSpecOperationAdd,
				Device: &VirtualE1000{VirtualEthernetCard{
					VirtualDevice: VirtualDevice{
						Key: 0,
						DeviceInfo: &Description{
							Label:   "Network Adapter 1",
							Summary: "VM Network",
						},
						Backing: &VirtualEthernetCardNetworkBackingInfo{
							VirtualDeviceDeviceBackingInfo: VirtualDeviceDeviceBackingInfo{
								DeviceName: "VM Network",
							},
						},
					},
					AddressType: string(VirtualEthernetCardMacTypeGenerated),
				}},
			},
		},
		ExtraConfig: []BaseOptionValue{
			&OptionValue{Key: "bios.bootOrder", Value: "ethernet0"},
		},
	}

	_, err := xml.MarshalIndent(spec, "", " ")
	if err != nil {
		t.Fatal(err)
	}
}

func TestVirtualMachineAffinityInfo(t *testing.T) {
	// See https://github.com/vmware/govmomi/issues/1008
	in := VirtualMachineAffinityInfo{
		AffinitySet: []int32{0, 1, 2, 3},
	}

	b, err := xml.Marshal(in)
	if err != nil {
		t.Fatal(err)
	}

	var out VirtualMachineAffinityInfo

	err = xml.Unmarshal(b, &out)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(in, out) {
		t.Errorf("%#v vs %#v", in, out)
	}
}
