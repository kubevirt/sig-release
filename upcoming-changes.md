---
title: Noteworthy changes for the next KubeVirt release
---

# Noteworthy changes for the next KubeVirt release

This list contains the noteworthy changes made after the latest KubeVirt release. The community expects these changes to be included in the next Kubevirt release.

> [!WARNING]
> **Please be aware that any of these might be excluded from the next release.**

| Upcoming changes | PR                                                                   | Author                                          |
|------------------|----------------------------------------------------------------------|-------------------------------------------------|
| virtctl: updated virt-template subcomands to v1beta1  | [#17910](https://github.com/kubevirt/kubevirt/pull/17910) | [jcanocan](https://github.com/jcanocan) |
| BugFix: Storage Live Migration Filesystem-to-Block fails with libvirt error Code 84 pre-creation of storage target for incremental storage migration of disk is not supported  | [#18143](https://github.com/kubevirt/kubevirt/pull/18143) | [akalenyu](https://github.com/akalenyu) |
| VirtualMachineBackups now consist of only 3 conditions: Progressing, Complete, and Failed, with matching reasons to replace the now removed Initializing, ExportInitiated, ExportReady, Aborting and Done.  | [#18090](https://github.com/kubevirt/kubevirt/pull/18090) | [Acedus](https://github.com/Acedus) |
| VEP 250: Add serviceAccountName to VirtualMachineInstance Spec  | [#17861](https://github.com/kubevirt/kubevirt/pull/17861) | [mhenriks](https://github.com/mhenriks) |
| Fix the virtctl image-upload command token expiry issue by refreshing the token.  | [#17960](https://github.com/kubevirt/kubevirt/pull/17960) | [dsanatar](https://github.com/dsanatar) |
| virt-operator: customizeComponents.Patches now apply to the install-strategy job (match it with resourceType "Job" and resourceName "virt-operator-install-strategy"). The job's generated name changed from "kubevirt-<id>-job" to "virt-operator-install-strategy-".  | [#17976](https://github.com/kubevirt/kubevirt/pull/17976) | [ethan-gallant](https://github.com/ethan-gallant) |
| Enable live migration for VMs with SCSI persistent reservations  | [#17505](https://github.com/kubevirt/kubevirt/pull/17505) | [alromeros](https://github.com/alromeros) |
| Add IOMMUFD device plugin to virt-handler behind the IOMMUFD Alpha feature gate. When enabled, virt-controller requests devices.kubevirt.io/iommufd for every launcher pod. Nodes without /dev/iommu (kernel <6.2) will report unhealthy devices, making pods unschedulable there. This feature also emits domain-level <iommufd enabled='yes' fdgroup='iommu'/> and uses virDomainFDAssociate, which require a libvirt/QEMU stack that supports fdgroup-based IOMMUFD, currently libvirt >= 12.2. Enable this gate only on clusters where target nodes and virt-launcher images provide compatible kernel, libvirt, and QEMU support.  | [#17956](https://github.com/kubevirt/kubevirt/pull/17956) | [fossedihelm](https://github.com/fossedihelm) |
| BugFix: Live migration with CBT and RWO backend storage now correctly retains checkpoints post-migration  | [#17993](https://github.com/kubevirt/kubevirt/pull/17993) | [Acedus](https://github.com/Acedus) |
| The RebootPolicy feature was graduated to Beta status.  | [#17623](https://github.com/kubevirt/kubevirt/pull/17623) | [MarSik](https://github.com/MarSik) |
| Libvirt bumped to v11.10  | [#17422](https://github.com/kubevirt/kubevirt/pull/17422) | [jean-edouard](https://github.com/jean-edouard) |
| Add foundation types for Grace IO Virtualization: GraceIOVirtualization feature gate, IOMMU domain XML schema, HostDevice IOMMU/ACPI extensions, and NUMACell pointer change for zero-memory cells.  | [#17891](https://github.com/kubevirt/kubevirt/pull/17891) | [vladikr](https://github.com/vladikr) |
| Added CEL-based domain hook evaluation for the structured plugin system (alpha, behind Plugins feature gate).  | [#17897](https://github.com/kubevirt/kubevirt/pull/17897) | [iholder101](https://github.com/iholder101) |
| - Allow configuring launch security via preference.<br>- Deprecate launchSecurity field in instancetypes.  | [#17551](https://github.com/kubevirt/kubevirt/pull/17551) | [jschintag](https://github.com/jschintag) |
| vep-10: add dra packages to lint configuration and fix linter errors  | [#17923](https://github.com/kubevirt/kubevirt/pull/17923) | [alaypatel07](https://github.com/alaypatel07) |
| Migrated port-forwarding and remote exec from hardcoded SPDY to WebSocket-primary with SPDY fallback, fixing intermittent stream creation failures against Kubernetes 1.31+ clusters where the PortForwardWebsockets feature gate is enabled.  | [#16990](https://github.com/kubevirt/kubevirt/pull/16990) | [lyarwood](https://github.com/lyarwood) |
| Updated virt-template to v0.2.0  | [#18076](https://github.com/kubevirt/kubevirt/pull/18076) | [jcanocan](https://github.com/jcanocan) |
| Fixed a build failure in `pkg/hypervisor` on GOOS/GOARCH combinations outside linux/{amd64,arm64,s390x}, which previously prevented downstream packagers from cross-building virtctl on architectures such as riscv64, ppc64le, and 386. On those architectures, `common.SchedSetScheduler` now returns the new sentinel error `common.ErrUnsupportedRTScheduling`.  | [#18052](https://github.com/kubevirt/kubevirt/pull/18052) | [vidit-bhat](https://github.com/vidit-bhat) |
| virt-operator: fix DeploymentInProgress after toggling optional feature gates  | [#17996](https://github.com/kubevirt/kubevirt/pull/17996) | [ksimon1](https://github.com/ksimon1) |
| Incremental backups now distinguish discarded blocks from data changes by merging base:allocation and qemu:dirty-bitmap contexts, reducing transfer size by skipping dirty extents that read as zero.  | [#17984](https://github.com/kubevirt/kubevirt/pull/17984) | [Acedus](https://github.com/Acedus) |
| Network binding plugins can now read the interface MTU from the network-info downward API annotation.  | [#17845](https://github.com/kubevirt/kubevirt/pull/17845) | [orelmisan](https://github.com/orelmisan) |
| cgroup v1 support is deprecated now, with removal planned for the next release  | [#17809](https://github.com/kubevirt/kubevirt/pull/17809) | [michalskrivanek](https://github.com/michalskrivanek) |
| vep-10: migrate away from k8sv1.PodResourceClaim to kubevirts own type  | [#17797](https://github.com/kubevirt/kubevirt/pull/17797) | [alaypatel07](https://github.com/alaypatel07) |
| Introduce the Plugin CRD (plugin.kubevirt.io/v1alpha1) behind the Plugins feature gate (Alpha), enabling declarative VM extension via domain hooks, node hooks, and admission references (VEP-190).  | [#17790](https://github.com/kubevirt/kubevirt/pull/17790) | [iholder101](https://github.com/iholder101) |
| Use --expand-cpu-features and --supported-cpu-features in node-labeller for<br>complete CPU feature reporting when available  | [#17947](https://github.com/kubevirt/kubevirt/pull/17947) | [Barakmor1](https://github.com/Barakmor1) |
| Add OCI artifact export support behind the OCIExport feature gate. Use `virtctl vmexport download --format=oci` to export a VM as an OCI image layout TAR.  | [#17944](https://github.com/kubevirt/kubevirt/pull/17944) | [0xFelix](https://github.com/0xFelix) |
| Bump github.com/moby/spdystream from v0.5.0 to v0.5.1 to address CVE-2026-35469 (GHSA-pc3f-x583-g7j2).  | [#17989](https://github.com/kubevirt/kubevirt/pull/17989) | [UdayYendva](https://github.com/UdayYendva) |
| Promote `VideoConfig` FG to General Availability  | [#16599](https://github.com/kubevirt/kubevirt/pull/16599) | [dasionov](https://github.com/dasionov) |
| Add alpha support for DRA SR-IOV networks via `spec.networks[].resourceClaim` behind `NetworkDevicesWithDRA`, with webhook validation and virt-launcher hostdev generation from DRA metadata.  | [#17661](https://github.com/kubevirt/kubevirt/pull/17661) | [oshoval](https://github.com/oshoval) |
| Graduate HotplugVolumes and DeclarativeHotplugVolumes to Beta  | [#16932](https://github.com/kubevirt/kubevirt/pull/16932) | [dsanatar](https://github.com/dsanatar) |
| Fixed a target-side hotplug mount leak that could remain after failed or canceled live migration.  | [#17883](https://github.com/kubevirt/kubevirt/pull/17883) | [yaroslavborbat](https://github.com/yaroslavborbat) |
| Add kubevirt_vmi_guest_os_panic_total Prometheus counter metric to track guest OS panic events per VMI, with labels for panic type and hyper-v bugcheck code  | [#17836](https://github.com/kubevirt/kubevirt/pull/17836) | [iholder101](https://github.com/iholder101) |
| Fix symlink traversal in VMExport dir handler  | [#17959](https://github.com/kubevirt/kubevirt/pull/17959) | [mhenriks](https://github.com/mhenriks) |
| Persistent Reservation GA  | [#16674](https://github.com/kubevirt/kubevirt/pull/16674) | [akalenyu](https://github.com/akalenyu) |
| Add GetVMStats unified gRPC RPC for monitoring data collection  | [#17550](https://github.com/kubevirt/kubevirt/pull/17550) | [machadovilaca](https://github.com/machadovilaca) |
| Fix: Handle disks with qcow2 overlay in migration related code  | [#17757](https://github.com/kubevirt/kubevirt/pull/17757) | [Acedus](https://github.com/Acedus) |
| BugFix: endless cycle of attachment pod deletion/creation  | [#17825](https://github.com/kubevirt/kubevirt/pull/17825) | [akalenyu](https://github.com/akalenyu) |
| only populate vmexport with a vm manifest if source is not pvc  | [#17699](https://github.com/kubevirt/kubevirt/pull/17699) | [dsanatar](https://github.com/dsanatar) |
| Fixed a gRPC connection leak in virt-handler's GetLauncherClient that caused unbounded memory growth, socket accumulation, and goroutine leaks when multiple controllers raced to create connections for the same VMI.  | [#17798](https://github.com/kubevirt/kubevirt/pull/17798) | [SamAlber](https://github.com/SamAlber) |
| Updated virt-template to v0.1.8  | [#17644](https://github.com/kubevirt/kubevirt/pull/17644) | [kubevirt-bot](https://github.com/kubevirt-bot) |
| Feat: Record K8s event on VMI when GuestAgentPing probe fails  | [#17272](https://github.com/kubevirt/kubevirt/pull/17272) | [ksimon1](https://github.com/ksimon1) |
| Fixed VMI status reporting the pod's IPv6 address instead of the guest's when using bridge binding on a network with IPv6 IPAM.  | [#17536](https://github.com/kubevirt/kubevirt/pull/17536) | [nirdothan](https://github.com/nirdothan) |
| Enable all Beta feature gates by default. Users can opt out of individual Beta features by adding them to `spec.configuration.developerConfiguration.disabledFeatureGates` in the KubeVirt CR. A feature gate report with all non-GA feature gates will be added to each release's artifacts.  | [#17405](https://github.com/kubevirt/kubevirt/pull/17405) | [iholder101](https://github.com/iholder101) |
| Fixed multi-device VFIO passthrough VMs failing to start with "cannot limit locked memory" by scaling virt-handler's memlock rlimit to account for per-device memory locking, matching libvirt's calculation introduced in v8.7.0.  | [#17805](https://github.com/kubevirt/kubevirt/pull/17805) | [lyarwood](https://github.com/lyarwood) |
| The --additional-launcher-annotations-sync and --additional-launcher-labels-sync flags now support prefix wildcards via a trailing '*' suffix (e.g. 'vendor.io/*'), allowing all matching labels/annotations to be propagated from VM template to VMI and virt-launcher pod without enumerating each key individually.  | [#17496](https://github.com/kubevirt/kubevirt/pull/17496) | [fanzhangio](https://github.com/fanzhangio) |
| Fix PCI address stability across upgrades with v3 hotplug port topology  | [#17029](https://github.com/kubevirt/kubevirt/pull/17029) | [mhenriks](https://github.com/mhenriks) |
| Fix race condition in VM force restart where the pod could remain stuck in Terminating state for the full terminationGracePeriodSeconds instead of terminating promptly. The VMI's terminationGracePeriodSeconds is now patched before triggering the restart, ensuring the short grace period is always honored.  | [#17475](https://github.com/kubevirt/kubevirt/pull/17475) | [samt-ai](https://github.com/samt-ai) |
| N/A  | [#17031](https://github.com/kubevirt/kubevirt/pull/17031) | [suPer8Hu](https://github.com/suPer8Hu) |
| Add container-level SecurityContext to virt-exportproxy and virt-synchronization-controller deployments  | [#17548](https://github.com/kubevirt/kubevirt/pull/17548) | [Barakmor1](https://github.com/Barakmor1) |
| fix: cross-namespace live migration now works on IPv6 clusters  | [#17755](https://github.com/kubevirt/kubevirt/pull/17755) | [dasionov](https://github.com/dasionov) |
| Fixed virt-controller DRA claim rendering for GPU/HostDevice resources by preserving per-device claim/request tuples (including shared claim names with different requests).  | [#17490](https://github.com/kubevirt/kubevirt/pull/17490) | [oshoval](https://github.com/oshoval) |
| Fix VM with PCI hostdev failing to restart after hotplug block volume  | [#17527](https://github.com/kubevirt/kubevirt/pull/17527) | [mhenriks](https://github.com/mhenriks) |
| Bug fix: virt-operator error messages no longer dump entire resource structs via %+v, preventing the KubeVirt CR from exceeding the etcd 3MB object size limit when resource creation fails  | [#17573](https://github.com/kubevirt/kubevirt/pull/17573) | [lyarwood](https://github.com/lyarwood) |
| NA  | [#17691](https://github.com/kubevirt/kubevirt/pull/17691) | [vishnuchalla](https://github.com/vishnuchalla) |
| Graduates LiveUpdateNADRef feature gate  | [#17049](https://github.com/kubevirt/kubevirt/pull/17049) | [frenzyfriday](https://github.com/frenzyfriday) |
| Fixed DHCP failure after live migration followed by guest reboot when using bridge binding without a specified MAC address.  | [#16697](https://github.com/kubevirt/kubevirt/pull/16697) | [dippydocus](https://github.com/dippydocus) |
| Add release 1.8 perf-scale benchmarks data  | [#17599](https://github.com/kubevirt/kubevirt/pull/17599) | [Sreeja1725](https://github.com/Sreeja1725) |
| Bugfix: Added enum validation for the targetReadinessPolicy field for restore resources.  | [#17576](https://github.com/kubevirt/kubevirt/pull/17576) | [ema-aka-young](https://github.com/ema-aka-young) |
| Fixed virt-api truncating deep subresources (vnc/screenshot, sev/*, evacuate/cancel) when constructing SubjectAccessReviews, causing authorization checks against incorrect subresource names.  | [#17571](https://github.com/kubevirt/kubevirt/pull/17571) | [0xFelix](https://github.com/0xFelix) |
| Adding missing metrics, recording rules and alerts for virt components  | [#17015](https://github.com/kubevirt/kubevirt/pull/17015) | [Ronilerr](https://github.com/Ronilerr) |
| Fix VirtualMachineStuckOnNode and VMCannotBeEvicted alerts failing during live migration due to duplicate kubevirt_vmi_info series  | [#17130](https://github.com/kubevirt/kubevirt/pull/17130) | [sradco](https://github.com/sradco) |
| change /var/lib/kubelet mount from Bidirectional to HostToContainer  | [#17488](https://github.com/kubevirt/kubevirt/pull/17488) | [alaypatel07](https://github.com/alaypatel07) |
| VEP-10: bug fixes for DRA Devices to align kubevirt implementation to KEP-5304  | [#17028](https://github.com/kubevirt/kubevirt/pull/17028) | [alaypatel07](https://github.com/alaypatel07) |
| Fixed GuestPanicked event details for non-root virt-launcher  | [#17557](https://github.com/kubevirt/kubevirt/pull/17557) | [dshchedr](https://github.com/dshchedr) |
| Fix: GuestAgentPing liveness/readiness probes no longer cause Kubernetes to restart the virt-launcher pod when the guest agent is temporarily unreachable for a non-fault reason; suppression covers live migration (both pre-copy target and post-copy source) and any intentional or transient VM pause such  as user pause, snapshot, save, or dump.  | [#17235](https://github.com/kubevirt/kubevirt/pull/17235) | [tiraboschi](https://github.com/tiraboschi) |
| Remediate CVE-2026-33186 by bumping grpc to 1.79.3  | [#17497](https://github.com/kubevirt/kubevirt/pull/17497) | [sbiradar10](https://github.com/sbiradar10) |
| new metric kubevirt_vmi_sync_total added in order to track number of times a controller has synced a VMI.  | [#17295](https://github.com/kubevirt/kubevirt/pull/17295) | [avlitman](https://github.com/avlitman) |
| Remove vnc/screenshot from kubevirt.io:edit  | [#17512](https://github.com/kubevirt/kubevirt/pull/17512) | [0xFelix](https://github.com/0xFelix) |
| Snapshot: Add PartialSnapshot indication for excluded volumes  | [#16882](https://github.com/kubevirt/kubevirt/pull/16882) | [alromeros](https://github.com/alromeros) |
| Bug-fix: virt-handler now detects when `domain-notify.sock` is deleted and automatically restarts the notify server.  | [#17398](https://github.com/kubevirt/kubevirt/pull/17398) | [dasionov](https://github.com/dasionov) |
| KubeVirt's PCI device plugin now supports passing of pre-setup VF (vGPU)  | [#16890](https://github.com/kubevirt/kubevirt/pull/16890) | [xpivarc](https://github.com/xpivarc) |
| fix: VirtualMachineBackup printer columns (Type, CheckpointName) now display correctly in kubectl output  | [#17425](https://github.com/kubevirt/kubevirt/pull/17425) | [shubham-pampattiwar](https://github.com/shubham-pampattiwar) |
| preserve annotation for restore pvc  | [#17407](https://github.com/kubevirt/kubevirt/pull/17407) | [dsanatar](https://github.com/dsanatar) |
| fix: correctly handle source resolution for disks with a qcow2 overlay, preventing incorrect disk expansion and wrong cache/IO mode detection.  | [#17297](https://github.com/kubevirt/kubevirt/pull/17297) | [Acedus](https://github.com/Acedus) |
| Improve Unmount() cleanup by processing all entries and preserving failed paths for retry.  | [#16742](https://github.com/kubevirt/kubevirt/pull/16742) | [keinsword](https://github.com/keinsword) |
| VMs with backend storage volume use and report the volume name as `persistent-state-for-this-vm` rather than trying to embed the vm name in the volume name.<br>Persistent TPM, EFI, snapshot, export and CBT features now work with VM names of all lengths up to 63 chars.  | [#16853](https://github.com/kubevirt/kubevirt/pull/16853) | [dankenigsberg](https://github.com/dankenigsberg) |
| Fixes bug in Live NAD Ref Update feature where a VM with no interfaces/networks is unable to start when LiveNADRefUpdate FG is enabled.  | [#17315](https://github.com/kubevirt/kubevirt/pull/17315) | [frenzyfriday](https://github.com/frenzyfriday) |
| Bug fix: sync-controller healthz server and virt-exportserver now respect TLSConfiguration from the KubeVirt CR.  | [#17102](https://github.com/kubevirt/kubevirt/pull/17102) | [Barakmor1](https://github.com/Barakmor1) |
| GA VMExport feature gate  | [#16730](https://github.com/kubevirt/kubevirt/pull/16730) | [alromeros](https://github.com/alromeros) |
| fix hotplug volume status being stuck in Detaching phase  | [#17139](https://github.com/kubevirt/kubevirt/pull/17139) | [dsanatar](https://github.com/dsanatar) |
| The `PanicDevices` feature has graduated to GA and no longer requires the associated feature gate to be enabled.  | [#16514](https://github.com/kubevirt/kubevirt/pull/16514) | [varunrsekar](https://github.com/varunrsekar) |
| Fixed VM startup failure under software emulation when /dev/kvm is absent, caused by cgroup device rules not accounting for the emulation case.  | [#17251](https://github.com/kubevirt/kubevirt/pull/17251) | [iholder101](https://github.com/iholder101) |
| Fixed migration not reporting succeeded when doing compute migration after decentralized live migration  | [#17042](https://github.com/kubevirt/kubevirt/pull/17042) | [awels](https://github.com/awels) |
| OpenApi V3 paths for subresources.kubevirt.io is present  | [#16842](https://github.com/kubevirt/kubevirt/pull/16842) | [ksimon1](https://github.com/ksimon1) |
| multiple recording rules are deprecated in favor of new names, in order to comply with the recording rules naming conventions. kubevirt_vm_created_total recording rule and kubevirt_vm_created_by_pod_total metric are deprecated completely  | [#17065](https://github.com/kubevirt/kubevirt/pull/17065) | [avlitman](https://github.com/avlitman) |
| Maintenance: revert bazel server network change monitoring  | [#17285](https://github.com/kubevirt/kubevirt/pull/17285) | [dhiller](https://github.com/dhiller) |
| Make --vnc-path optional for --vnc-type, allowing VNC viewer binary lookup from $PATH automatically  | [#17071](https://github.com/kubevirt/kubevirt/pull/17071) | [samt-ai](https://github.com/samt-ai) |
| bug-fix: restart virt-handler's domain-notify server on unexpected exit.  | [#17109](https://github.com/kubevirt/kubevirt/pull/17109) | [dasionov](https://github.com/dasionov) |
| Maintenance: make the bazel server container reload on network changes  | [#16885](https://github.com/kubevirt/kubevirt/pull/16885) | [akalenyu](https://github.com/akalenyu) |
| vGPU (mdev) now supports live migration (limited to one device)  | [#16675](https://github.com/kubevirt/kubevirt/pull/16675) | [csomani1](https://github.com/csomani1) |
| Allow multifd (parallel migration threads) with post-copy migration  | [#17106](https://github.com/kubevirt/kubevirt/pull/17106) | [iholder101](https://github.com/iholder101) |
| bugfix: Online snapshots now correctly include live-changes applied to a VM.  | [#15298](https://github.com/kubevirt/kubevirt/pull/15298) | [Acedus](https://github.com/Acedus) |
| Fix virt-handler hotplug mount target record growth by preventing duplicate TargetFile entries.  | [#16737](https://github.com/kubevirt/kubevirt/pull/16737) | [keinsword](https://github.com/keinsword) |
| Fixed an infinite VMI status update loop between virt-controller and virt-handler that occurred when the VMI spec listed the primary network interface after a secondary one.  | [#17041](https://github.com/kubevirt/kubevirt/pull/17041) | [orelmisan](https://github.com/orelmisan) |
| Fixed SMBIOS system information not being visible inside ARM64 guest VMs  | [#16783](https://github.com/kubevirt/kubevirt/pull/16783) | [dasionov](https://github.com/dasionov) |
| VEP-10: Update DRA devices implementation to read from metadata file instead of VMI status  | [#16556](https://github.com/kubevirt/kubevirt/pull/16556) | [alaypatel07](https://github.com/alaypatel07) |
| Introduce incremental backup pull mode support  | [#16930](https://github.com/kubevirt/kubevirt/pull/16930) | [Acedus](https://github.com/Acedus) |
| Add MSHV backend for multi-hypervisor support interfaces.  | [#16927](https://github.com/kubevirt/kubevirt/pull/16927) | [harshitgupta1337](https://github.com/harshitgupta1337) |
| Handle migration with CBT and backup checkpoints  | [#16817](https://github.com/kubevirt/kubevirt/pull/16817) | [ShellyKa13](https://github.com/ShellyKa13) |
| Add PCIe NUMA-aware topology placement for GPU and host devices behind the PCINUMAAwareTopology feature gate (Alpha). When enabled, devices are automatically placed on PCIe expander buses matching their NUMA affinity for improved performance.  | [#16632](https://github.com/kubevirt/kubevirt/pull/16632) | [mresvanis](https://github.com/mresvanis) |
| 'virtctl expose' creates now services with an ownerReference pointing to the exposed resource.  | [#16884](https://github.com/kubevirt/kubevirt/pull/16884) | [EdDev](https://github.com/EdDev) |
| Allows the user to update the NAD reference (networkName) of a network on a running VM through Live Migration.  | [#16412](https://github.com/kubevirt/kubevirt/pull/16412) | [frenzyfriday](https://github.com/frenzyfriday) |
| Expose Memory Overhead on VMI Status behind VmiMemoryOverheadReport feature gate  | [#16746](https://github.com/kubevirt/kubevirt/pull/16746) | [Barakmor1](https://github.com/Barakmor1) |
| Add a new config option to opt-out RBAC aggregation  | [#16350](https://github.com/kubevirt/kubevirt/pull/16350) | [orenc1](https://github.com/orenc1) |
| Use defined deployment number of replicas as base to fire low count alerts  | [#16806](https://github.com/kubevirt/kubevirt/pull/16806) | [machadovilaca](https://github.com/machadovilaca) |
| Graduate ExpandDisk Feature Gate  | [#16604](https://github.com/kubevirt/kubevirt/pull/16604) | [dsanatar](https://github.com/dsanatar) |
| Added support for attestation on the Intel TDX Confidential Computing Platform  | [#15958](https://github.com/kubevirt/kubevirt/pull/15958) | [Aseeef](https://github.com/Aseeef) |
| Handle migration during backup according to migration priority  | [#16877](https://github.com/kubevirt/kubevirt/pull/16877) | [ShellyKa13](https://github.com/ShellyKa13) |
| BugFix: VMs requiring enlightenment are now able to be live migrated after a decentralized live migration  | [#16871](https://github.com/kubevirt/kubevirt/pull/16871) | [awels](https://github.com/awels) |
| Subtract non-schedulable nodes from kubevirt_allocatable_nodes  | [#16865](https://github.com/kubevirt/kubevirt/pull/16865) | [machadovilaca](https://github.com/machadovilaca) |
| fix VMExport failure with long PVC names  | [#16846](https://github.com/kubevirt/kubevirt/pull/16846) | [Aneesh-Hegde](https://github.com/Aneesh-Hegde) |
| fix: Prevent stale VMI backup status update when reusing backup names  | [#16399](https://github.com/kubevirt/kubevirt/pull/16399) | [Acedus](https://github.com/Acedus) |
| Allow disabling Velero hooks in virt-launcher via Annotation  | [#16786](https://github.com/kubevirt/kubevirt/pull/16786) | [alromeros](https://github.com/alromeros) |


_This page is updated daily._
