---
title: Noteworthy changes for the next KubeVirt release
---

# Noteworthy changes for the next KubeVirt release

This list contains the noteworthy changes made after the latest KubeVirt release. The community expects these changes to be included in the next Kubevirt release.

> [!WARNING]
> **Please be aware that any of these might be excluded from the next release.**

| Upcoming changes | PR                                                                   | Author                                          |
|------------------|----------------------------------------------------------------------|-------------------------------------------------|
| Add recording rule for report suspected orphaned PVCs  | [#18521](https://github.com/kubevirt/kubevirt/pull/18521) | [Ronilerr](https://github.com/Ronilerr) |
| Graduate LibvirtHooksServerAndClient feature gate to GA  | [#18687](https://github.com/kubevirt/kubevirt/pull/18687) | [Barakmor1](https://github.com/Barakmor1) |
| Graduate VmiMemoryOverheadReport feature gate to GA  | [#18708](https://github.com/kubevirt/kubevirt/pull/18708) | [Barakmor1](https://github.com/Barakmor1) |
| Fix premature domain undefine during slow QEMU startup that left VMs transient and broke hotplug  | [#18545](https://github.com/kubevirt/kubevirt/pull/18545) | [amasolov](https://github.com/amasolov) |
| Backend storage PVC size is increased by 1Gi for VMs with CBT enabled.  | [#18470](https://github.com/kubevirt/kubevirt/pull/18470) | [Acedus](https://github.com/Acedus) |
| BREAKING CHANGE: Modifies internal on-the-wire API fields for the stall detector feature; MigrationStallDetection FG must be disabled for cluster upgrades to succeed.  | [#18303](https://github.com/kubevirt/kubevirt/pull/18303) | [Aseeef](https://github.com/Aseeef) |
| Fixed a timing side channel in the export server token check by switching to a constant-time comparison.  | [#18200](https://github.com/kubevirt/kubevirt/pull/18200) | [naruto-lgtm](https://github.com/naruto-lgtm) |
| Bug fix: Prevent guest OS from ejecting critical PCI devices (disks, memory balloon, RNG, watchdog) via hotplug mechanisms like the Windows "Safely Remove Hardware" tray menu. PCI root ports hosting non-NIC devices are now marked as non-hotpluggable, while NIC ports and empty ports reserved for future hotplug remain available.  | [#17605](https://github.com/kubevirt/kubevirt/pull/17605) | [dasionov](https://github.com/dasionov) |
| Bump hermetic_launcher to v0.0.15 to fix s390x container image builds.  | [#18692](https://github.com/kubevirt/kubevirt/pull/18692) | [vamsikrishna-siddu](https://github.com/vamsikrishna-siddu) |
| Removed the slirp network binding plugin sidecar image. The plugin had no known consumers.  | [#18613](https://github.com/kubevirt/kubevirt/pull/18613) | [orelmisan](https://github.com/orelmisan) |
| Live migration downtime can now be dynamically tuned by using .spec.experimental.downtimeTuning field in MigrationPolicy. When enabled, QEMU's max_downtime is gradually increased from an initial value (default 150ms) to a configurable ceiling (default 1050ms), helping migrations converge with bounded sub-second switchover downtime.  | [#17480](https://github.com/kubevirt/kubevirt/pull/17480) | [michalskrivanek](https://github.com/michalskrivanek) |
| Bug fix: During live migration, domain stats collection no longer blocks on the libvirt job lock for up to 30 seconds. The NOWAIT flag is used to skip locked domains instead of waiting, avoiding lock contention and disruptive error logs.  | [#18362](https://github.com/kubevirt/kubevirt/pull/18362) | [iholder101](https://github.com/iholder101) |
| BugFix: Online expansion scales poorly with many volumes  | [#18311](https://github.com/kubevirt/kubevirt/pull/18311) | [akalenyu](https://github.com/akalenyu) |
| Fix sidecar hooks silently dropping QEMU command-line arguments (qemu:commandline) during XML round-trip  | [#18460](https://github.com/kubevirt/kubevirt/pull/18460) | [dasionov](https://github.com/dasionov) |
| Migrate container image builds from rules_oci to rules_img, enabling hermetic multi-arch image builds with native s390x platform support.  | [#17598](https://github.com/kubevirt/kubevirt/pull/17598) | [vamsikrishna-siddu](https://github.com/vamsikrishna-siddu) |
| RamFB display validation now treats a vGPU whose ramFB is present with its enabled field unset as an enabled display, matching the API default. A VirtualMachine or VirtualMachineInstanceReplicaSet template with more than one such vGPU is now rejected at creation instead of being accepted and failing later when its VMI is created.  | [#18542](https://github.com/kubevirt/kubevirt/pull/18542) | [thc1006](https://github.com/thc1006) |
| BugFix: file descriptor leak in vmexport proxy  | [#18400](https://github.com/kubevirt/kubevirt/pull/18400) | [awels](https://github.com/awels) |
| Decentralized live migration to beta  | [#18168](https://github.com/kubevirt/kubevirt/pull/18168) | [awels](https://github.com/awels) |
| Add support for OCI artifacts (e.g. raw disk images pushed with ORAS) as containerDisk sources via the ImageVolume path. Set the annotation `kubevirt.io/image-volume-skip-digest-resolution: "true"` on the VMI to skip the digest-resolving init containers that are incompatible with non-container OCI artifacts. The disk file discovery now also supports images with the disk at the root directory, not just under /disk.  | [#18491](https://github.com/kubevirt/kubevirt/pull/18491) | [vladikr](https://github.com/vladikr) |
| Fix a goroutine (and connection resource) leak in client-go's AsyncSubresourceHelper when its StreamInterface is consumed via AsConn() instead of Stream() (affects VSOCK, PortForward TCP/UDP, and similar consumers).  | [#18410](https://github.com/kubevirt/kubevirt/pull/18410) | [vikin91](https://github.com/vikin91) |
| Fix CallNodeHooks to evaluate plugin-level Spec.Condition, ensuring consistent behavior between domain hooks and node hooks.  | [#18414](https://github.com/kubevirt/kubevirt/pull/18414) | [iholder101](https://github.com/iholder101) |
| Bug fix: Running VMIs are no longer immediately marked as Failed when virt-handler restarts and temporarily loses the launcher socket connection. A 3-minute grace period (matching the existing Scheduled VMI behavior) now allows the launcher client to reconnect before transitioning the VMI to Failed.  | [#18453](https://github.com/kubevirt/kubevirt/pull/18453) | [iholder101](https://github.com/iholder101) |
| Add namespace label to all KubeVirt alerts for proper OpenShift<br>monitoring routing, silencing, and multi-tenant isolation.  | [#17995](https://github.com/kubevirt/kubevirt/pull/17995) | [sradco](https://github.com/sradco) |
| Fixed missing validation for network interface bindings: interfaces with no binding or multiple bindings are now rejected at admission instead of failing silently at runtime.  | [#18374](https://github.com/kubevirt/kubevirt/pull/18374) | [orelmisan](https://github.com/orelmisan) |
| Bug fix: Expose VFIO cdev devices (/dev/vfio/devices/vfioN) to virt-launcher containers, fixing IOMMUFD-based GPU passthrough on ARM64 systems with SMMUv3.  | [#18300](https://github.com/kubevirt/kubevirt/pull/18300) | [lyarwood](https://github.com/lyarwood) |
| Fixed VMI owner resolution when an attachment pod is deleted while the virt-launcher pod is gone  | [#18222](https://github.com/kubevirt/kubevirt/pull/18222) | [lukashes](https://github.com/lukashes) |
| BugFix: hotplug volume detach deadlock for already removed volumes.  | [#18044](https://github.com/kubevirt/kubevirt/pull/18044) | [maxwmsft](https://github.com/maxwmsft) |
| virt-launcher pods get cleaned up along with migration objects  | [#17707](https://github.com/kubevirt/kubevirt/pull/17707) | [jean-edouard](https://github.com/jean-edouard) |
| BugFix: hotplug volume detach deadlock when VMIs exchange RWO PVCs between attachment pods.  | [#18045](https://github.com/kubevirt/kubevirt/pull/18045) | [maxwmsft](https://github.com/maxwmsft) |
| Network binding plugin sidecars can now infer their registration name using the `NETWORK_BINDING_PLUGIN_NAME` environment variable.  | [#18435](https://github.com/kubevirt/kubevirt/pull/18435) | [orelmisan](https://github.com/orelmisan) |
| N/A  | [#17117](https://github.com/kubevirt/kubevirt/pull/17117) | [suPer8Hu](https://github.com/suPer8Hu) |
| Deprecated old method to set SSH key in a VM.  | [#13881](https://github.com/kubevirt/kubevirt/pull/13881) | [akrejcir](https://github.com/akrejcir) |
| Set unlimited memlock prlimit for VMs requiring memory locking (VFIO, realtime, SEV, explicit memlock), ensuring libvirt's memlock calculation always succeeds in session mode.  | [#17857](https://github.com/kubevirt/kubevirt/pull/17857) | [lyarwood](https://github.com/lyarwood) |
| Require restart for NAD swap for non-migratable VMs  | [#18194](https://github.com/kubevirt/kubevirt/pull/18194) | [frenzyfriday](https://github.com/frenzyfriday) |
| Low virt count alerts are using running rule intead of up rule  | [#18380](https://github.com/kubevirt/kubevirt/pull/18380) | [Ronilerr](https://github.com/Ronilerr) |
| Bug fix: VMs with NFS-backed storage are no longer incorrectly restarted during temporary NFS unavailability (e.g., NFS server failover).  | [#17752](https://github.com/kubevirt/kubevirt/pull/17752) | [iholder101](https://github.com/iholder101) |
| Bugfix: Fix VM restore when backend storage is not snapshottable  | [#17286](https://github.com/kubevirt/kubevirt/pull/17286) | [alromeros](https://github.com/alromeros) |
| Network conformance tests can now be run with a custom network binding plugin via the `--primary-network-binding-plugin` test flag.  | [#18021](https://github.com/kubevirt/kubevirt/pull/18021) | [orelmisan](https://github.com/orelmisan) |
| BugFix: Importer pod rejected by ValidatingAdmissionPolicy 'kubevirt-plugin-sidecar-subpath-policy'  | [#18392](https://github.com/kubevirt/kubevirt/pull/18392) | [akalenyu](https://github.com/akalenyu) |
| The `kubevirt_vmi_guest_os_panic_total` metric is now emitted<br>for all guest panic events, including when the guest recovers internally<br>(e.g. via kdump/crashdump with the ISA pvpanic driver)  | [#18370](https://github.com/kubevirt/kubevirt/pull/18370) | [michalskrivanek](https://github.com/michalskrivanek) |
| Add E2E tests support for DRA feature.  | [#18061](https://github.com/kubevirt/kubevirt/pull/18061) | [Sreeja1725](https://github.com/Sreeja1725) |
| Bug fix: Fixed hugetblfs-dir EmptyDir volume masking the parent hugetlbfs mount at /dev/hugepages, which prevented QEMU from allocating hugepages via -mem-path. The EmptyDir is now explicitly hugepage-backed.  | [#18272](https://github.com/kubevirt/kubevirt/pull/18272) | [laxmi-333](https://github.com/laxmi-333) |
| client-go/log: The `-v` verbosity flag is no longer registered on `flag.CommandLine` during `init()`. External consumers that used `flag.CommandLine.Lookup("v")` to obtain the verbosity flag must switch to `log.VerbosityFlag()`. This eliminates panics when importing `kubevirt.io/client-go` alongside packages that also register a `-v` flag (e.g. glog, klog).  | [#18294](https://github.com/kubevirt/kubevirt/pull/18294) | [guzalv](https://github.com/guzalv) |
| virtctl: addvolume optional --disk-name flag  | [#17936](https://github.com/kubevirt/kubevirt/pull/17936) | [dsanatar](https://github.com/dsanatar) |
| Fixed a delay in VM/VMI deletion that could happen when the VMI controller missed a pod event.  | [#18206](https://github.com/kubevirt/kubevirt/pull/18206) | [lukashes](https://github.com/lukashes) |
| Fix panic in ClusterConfig CRD delete handler when the informer delivers a tombstone object instead of the deleted CRD.  | [#17880](https://github.com/kubevirt/kubevirt/pull/17880) | [Shreesha001](https://github.com/Shreesha001) |
| Support inferFromVolume for VolumeSnapshot-backed DataVolumes, DataVolumeTemplates, and DataSources  | [#18337](https://github.com/kubevirt/kubevirt/pull/18337) | [0xFelix](https://github.com/0xFelix) |
| Set explicit SMMUv3 address capability defaults for Grace I/O Virtualization.  | [#18316](https://github.com/kubevirt/kubevirt/pull/18316) | [fanzhangio](https://github.com/fanzhangio) |
| vep-10: move GPUsWithDRA and HostDevicesWithDRA to beta  | [#18247](https://github.com/kubevirt/kubevirt/pull/18247) | [alaypatel07](https://github.com/alaypatel07) |
| Adding NoLeadingVirtController alert  | [#18154](https://github.com/kubevirt/kubevirt/pull/18154) | [Ronilerr](https://github.com/Ronilerr) |
| BugFix: admin can remove memory dump from their VM  | [#17775](https://github.com/kubevirt/kubevirt/pull/17775) | [akalenyu](https://github.com/akalenyu) |
| Live migration data stream can be compressed now by using .spec.experimental.compression field in MigrationPolicy applied to a set of VMs. Currently only "zstd" algorithm is supported.  | [#17509](https://github.com/kubevirt/kubevirt/pull/17509) | [michalskrivanek](https://github.com/michalskrivanek) |
| Add namespace-aware VSOCK dialing support in virt-handler (VEP 222)  | [#17955](https://github.com/kubevirt/kubevirt/pull/17955) | [akrejcir](https://github.com/akrejcir) |


_This page is updated daily._
