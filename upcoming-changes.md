---
title: Noteworthy changes for the next KubeVirt release
---

# Noteworthy changes for the next KubeVirt release

This list contains the noteworthy changes made after the latest KubeVirt release. The community expects these changes to be included in the next Kubevirt release.

> [!WARNING]
> **Please be aware that any of these might be excluded from the next release.**

| Upcoming changes | PR                                                                   | Author                                          |
|------------------|----------------------------------------------------------------------|-------------------------------------------------|
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
