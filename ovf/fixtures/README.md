# OVF Fixtures

Fixtures for testing the govmomi OVF package against the [DMTF OVF Specification DSP0243 v2.1.1](https://www.dmtf.org/sites/default/files/standards/documents/DSP0243_2.1.1.pdf).

## Fixture matrix

| Fixture | Purpose | ToConfigSpec |
|---------|---------|--------------|
| `configspec.ovf` | Full VMware-style OVF (Item, Config, ExtraConfig, DeploymentOption, ProductSection, Eula, etc.) | ✓ |
| `properties.ovf` | ProductSection properties, class/instance, categories | ✓ |
| `virtualsystemcollection.ovf` | OVF 2.0 envelope, VirtualSystemCollection, EthernetPortItem, StorageItem (all processed in document order) | ✓ (index 0/1) |
| `minimal.ovf` | Minimal valid OVF: single VirtualSystem, Item-only hardware, one disk, one network | ✓ |
| `empty-disk.ovf` | DiskSection Disk without fileRef (empty disk per spec 9.1); disk Item with HostResource | ✓ |
| `deployment-option-non-default.ovf` | DeploymentOptionSection with default=medium; Items with ovf:configuration (spec 9.8) | ✓ |
| `property-value-per-config.ovf` | Property with Value elements per configuration (spec 9.5.1, 9.8) | ✓ |
| `ovf2-namespace.ovf` | OVF 2.0 envelope namespace (envelope/2); Item-only hardware | ✓ (if parser accepts envelope/2) |
| `file-references.ovf` | References with ovf:compression, ovf:chunkSize (spec 7.1) | ✓ |
| `two-virtual-hardware-sections.ovf` | Multiple VirtualHardwareSection (ovf:id on second); only first is used | ✓ |
| `bound-range-markers.ovf` | Item elements with ovf:bound="min" / "max" (spec 8.4). Range markers are skipped; normal value is used | ✓ |
| `product-section-class-instance.ovf` | ProductSection with ovf:class and ovf:instance (key-value-env, spec 9.5.1) | ✓ |
| `unsupported-resourcetype.ovf` | Item with invalid ResourceType for Strict mode | error |
| `unsupported-resourcesubtype.ovf` | Item with invalid ResourceSubType for Strict mode | error |
| `ubuntu24.10.ovf`, `photon5.ovf`, `ttylinux.ovf` | Real-world export samples | ✓ |

## Gaps vs DSP0243

- **EthernetPortItem / StorageItem**: OVF 2.x `EthernetPortItem` and `StorageItem` are now supported. They are converted to the same internal form as `Item` and processed in document order (NICs from EthernetPortItem, disks from StorageItem with ResourceType 31 LogicalDisk). StorageItem without Parent uses a default SCSI controller.
- **ovf:bound**: Items with `ovf:bound="min"` or `ovf:bound="max"` are range markers (spec 8.4). They are now skipped when building ConfigSpec; the normal item for that InstanceID is used.
- **InstallSection, BootDeviceSection, EnvironmentFilesSection**: Not in the Go envelope struct; unknown sections are dropped during unmarshal.
- **Strings / Msg (localization)**: Not in the Go envelope struct.
