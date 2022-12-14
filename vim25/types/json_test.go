/*
Copyright (c) 2022-2022 VMware, Inc. All Rights Reserved.

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
	"os"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"

	"github.com/vmware/govmomi/vim25/json"
)

func TestJSONMarshalVirtualMachineConfigSpec(t *testing.T) {
	var w bytes.Buffer
	enc := json.NewEncoder(&w)
	enc.SetIndent("", "  ")
	enc.SetDiscriminator("_typeName", "_value", "_byAddr")

	if err := enc.Encode(VirtualMachineConfigSpec{
		Name: "Hello, world.",
		DeviceChange: []BaseVirtualDeviceConfigSpec{
			&VirtualDeviceConfigSpec{
				Operation:     VirtualDeviceConfigSpecOperationAdd,
				FileOperation: VirtualDeviceConfigSpecFileOperationCreate,
				Device: &VirtualVmxnet3{
					VirtualVmxnet: VirtualVmxnet{
						VirtualEthernetCard: VirtualEthernetCard{
							VirtualDevice: VirtualDevice{
								Key: 3,
							},
							MacAddress: "00:11:22:33:44:55:66:88",
						},
					},
				},
			},
		},
	}); err != nil {
		t.Fatal(err)
	}
	act, exp := w.String(), virtualMachineConfigSpecWithDeviceChangesJSON
	if act != exp {
		t.Errorf("act json != exp json\nact=%s\nexp=%s", act, exp)
	}
}

func TestJSONUnmarshalVirtualMachineConfigSpec(t *testing.T) {
	dec := json.NewDecoder(strings.NewReader(virtualMachineConfigSpecWithVAppConfigJSON))
	dec.SetDiscriminator("_typeName", "_value", "_byAddr", json.DiscriminatorToTypeFunc(TypeFunc()))

	var obj VirtualMachineConfigSpec
	if err := dec.Decode(&obj); err != nil {
		t.Fatal(err)
	}

	var w bytes.Buffer
	enc := json.NewEncoder(&w)
	enc.SetIndent("", "  ")
	enc.SetDiscriminator("_typeName", "_value", "_byAddr")

	if err := enc.Encode(obj); err != nil {
		t.Fatal(err)
	}

	act, exp := w.String(), virtualMachineConfigSpecWithVAppConfigJSON
	if act != exp {
		t.Errorf("act json != exp json\nact=%s\nexp=%s", act, exp)
	}
}

func TestJSONUnmarshalVirtualMachineConfigInfo(t *testing.T) {
	f, err := os.Open("./testdata/virtualMachineConfigInfo.json")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	dec := json.NewDecoder(f)
	dec.SetDiscriminator("_typeName", "_value", "_byValue", json.DiscriminatorToTypeFunc(TypeFunc()))

	var obj VirtualMachineConfigInfo
	if err := dec.Decode(&obj); err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(obj, virtualMachineConfigInfoObjForTestData); diff != "" {
		t.Errorf("mismatched VirtualMachineConfigInfo: %s", diff)
		fmt.Println(diff)
	}
}

const virtualMachineConfigSpecWithDeviceChangesJSON = `{
  "_typeName": "VirtualMachineConfigSpec",
  "name": "Hello, world.",
  "deviceChange": [
    {
      "_typeName": "VirtualDeviceConfigSpec",
      "_byAddr": true,
      "operation": "add",
      "fileOperation": "create",
      "device": {
        "_typeName": "VirtualVmxnet3",
        "_byAddr": true,
        "key": 3,
        "macAddress": "00:11:22:33:44:55:66:88"
      }
    }
  ]
}
`

const virtualMachineConfigSpecWithVAppConfigJSON = `{
  "_typeName": "VirtualMachineConfigSpec",
  "name": "Hello, world.",
  "vAppConfig": {
    "_typeName": "VmConfigSpec",
    "_byAddr": true,
    "product": [
      {
        "_typeName": "VAppProductSpec",
        "operation": "add",
        "info": {
          "_typeName": "VAppProductInfo",
          "key": 1,
          "name": "p1"
        }
      }
    ],
    "installBootRequired": false
  }
}
`

func mustParseTime(layout, value string) time.Time {
	t, err := time.Parse(layout, value)
	if err != nil {
		panic(err)
	}
	return t
}

func addrOfMustParseTime(layout, value string) *time.Time {
	t := mustParseTime(layout, value)
	return &t
}

func addrOfBool(v bool) *bool {
	return &v
}

func addrOfInt32(v int32) *int32 {
	return &v
}

func addrOfInt64(v int64) *int64 {
	return &v
}

var virtualMachineConfigInfoObjForTestData VirtualMachineConfigInfo = VirtualMachineConfigInfo{
	ChangeVersion:         "2022-12-12T11:48:35.473645Z",
	Modified:              mustParseTime(time.RFC3339, "1970-01-01T00:00:00Z"),
	Name:                  "test",
	GuestFullName:         "VMware Photon OS (64-bit)",
	Version:               "vmx-20",
	Uuid:                  "422ca90b-853b-1101-3350-759f747730cc",
	CreateDate:            addrOfMustParseTime(time.RFC3339, "2022-12-12T11:47:24.685785Z"),
	InstanceUuid:          "502cc2a5-1f06-2890-6d70-ba2c55c5c2b7",
	NpivTemporaryDisabled: addrOfBool(true),
	LocationId:            "",
	Template:              false,
	GuestId:               "vmwarePhoton64Guest",
	AlternateGuestName:    "",
	Annotation:            "",
	Files: VirtualMachineFileInfo{
		VmPathName:        "[datastore1] test/test.vmx",
		SnapshotDirectory: "[datastore1] test/",
		SuspendDirectory:  "[datastore1] test/",
		LogDirectory:      "[datastore1] test/",
	},
	Tools: &ToolsConfigInfo{
		ToolsVersion:            0,
		AfterPowerOn:            addrOfBool(true),
		AfterResume:             addrOfBool(true),
		BeforeGuestStandby:      addrOfBool(true),
		BeforeGuestShutdown:     addrOfBool(true),
		BeforeGuestReboot:       nil,
		ToolsUpgradePolicy:      "manual",
		SyncTimeWithHostAllowed: addrOfBool(true),
		SyncTimeWithHost:        addrOfBool(false),
		LastInstallInfo: &ToolsConfigInfoToolsLastInstallInfo{
			Counter: 0,
		},
	},
	Flags: VirtualMachineFlagInfo{
		EnableLogging:            addrOfBool(true),
		UseToe:                   addrOfBool(false),
		RunWithDebugInfo:         addrOfBool(false),
		MonitorType:              "release",
		HtSharing:                "any",
		SnapshotDisabled:         addrOfBool(false),
		SnapshotLocked:           addrOfBool(false),
		DiskUuidEnabled:          addrOfBool(false),
		SnapshotPowerOffBehavior: "powerOff",
		RecordReplayEnabled:      addrOfBool(false),
		FaultToleranceType:       "unset",
		CbrcCacheEnabled:         addrOfBool(false),
		VvtdEnabled:              addrOfBool(false),
		VbsEnabled:               addrOfBool(false),
	},
	DefaultPowerOps: VirtualMachineDefaultPowerOpInfo{
		PowerOffType:        "soft",
		SuspendType:         "hard",
		ResetType:           "soft",
		DefaultPowerOffType: "soft",
		DefaultSuspendType:  "hard",
		DefaultResetType:    "soft",
		StandbyAction:       "checkpoint",
	},
	RebootPowerOff: addrOfBool(false),
	Hardware: VirtualHardware{
		NumCPU:              1,
		NumCoresPerSocket:   1,
		AutoCoresPerSocket:  addrOfBool(true),
		MemoryMB:            2048,
		VirtualICH7MPresent: addrOfBool(false),
		VirtualSMCPresent:   addrOfBool(false),
		Device: []BaseVirtualDevice{
			&VirtualIDEController{
				VirtualController: VirtualController{
					VirtualDevice: VirtualDevice{
						Key: 200,
						DeviceInfo: &Description{
							Label:   "IDE 0",
							Summary: "IDE 0",
						},
					},
					BusNumber: 0,
				},
			},
			&VirtualIDEController{
				VirtualController: VirtualController{
					VirtualDevice: VirtualDevice{
						Key: 201,
						DeviceInfo: &Description{
							Label:   "IDE 1",
							Summary: "IDE 1",
						},
					},
					BusNumber: 1,
				},
			},
			&VirtualPS2Controller{
				VirtualController: VirtualController{
					VirtualDevice: VirtualDevice{
						Key: 300,
						DeviceInfo: &Description{
							Label:   "PS2 controller 0",
							Summary: "PS2 controller 0",
						},
					},
					BusNumber: 0,
					Device:    []int32{600, 700},
				},
			},
			&VirtualPCIController{
				VirtualController: VirtualController{
					VirtualDevice: VirtualDevice{
						Key: 100,
						DeviceInfo: &Description{
							Label:   "PCI controller 0",
							Summary: "PCI controller 0",
						},
					},
					BusNumber: 0,
					Device:    []int32{500, 12000, 14000, 1000, 15000, 4000},
				},
			},
			&VirtualSIOController{
				VirtualController: VirtualController{
					VirtualDevice: VirtualDevice{
						Key: 400,
						DeviceInfo: &Description{
							Label:   "SIO controller 0",
							Summary: "SIO controller 0",
						},
					},
					BusNumber: 0,
				},
			},
			&VirtualKeyboard{
				VirtualDevice: VirtualDevice{
					Key: 600,
					DeviceInfo: &Description{
						Label:   "Keyboard",
						Summary: "Keyboard",
					},
					ControllerKey: 300,
					UnitNumber:    addrOfInt32(0),
				},
			},
		},
		MotherboardLayout:   "i440bxHostBridge",
		SimultaneousThreads: 1,
	},
	CpuAllocation: &ResourceAllocationInfo{
		Reservation:           addrOfInt64(0),
		ExpandableReservation: addrOfBool(false),
		Limit:                 addrOfInt64(-1),
		Shares: &SharesInfo{
			Shares: 1000,
			Level:  SharesLevelNormal,
		},
	},
	MemoryAllocation: &ResourceAllocationInfo{
		Reservation:           addrOfInt64(0),
		ExpandableReservation: addrOfBool(false),
		Limit:                 addrOfInt64(-1),
		Shares: &SharesInfo{
			Shares: 20480,
			Level:  SharesLevelNormal,
		},
	},
	LatencySensitivity: &LatencySensitivity{
		Level: LatencySensitivitySensitivityLevelNormal,
	},
	MemoryHotAddEnabled: addrOfBool(false),
	CpuHotAddEnabled:    addrOfBool(false),
	CpuHotRemoveEnabled: addrOfBool(false),
	ExtraConfig: []BaseOptionValue{
		&OptionValue{
			Key:   "nvram",
			Value: "test.nvram",
		},
	},
}
