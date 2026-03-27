---
title: Noteworthy changes for the next KubeVirt release
---

# Noteworthy changes for the next KubeVirt release

This list contains the noteworthy changes made after the latest KubeVirt release. The community expects these changes to be included in the next Kubevirt release.

> [!WARNING]
> **Please be aware that any of these might be excluded from the next release.**

| Upcoming changes | PR                                                                   | Author                                          |
|------------------|----------------------------------------------------------------------|-------------------------------------------------|
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
| Graduate ExpandDisk Feature Gate  | [#16604](https://github.com/kubevirt/kubevirt/pull/16604) | [Dsanatar](https://github.com/Dsanatar) |
| Added support for attestation on the Intel TDX Confidential Computing Platform  | [#15958](https://github.com/kubevirt/kubevirt/pull/15958) | [Aseeef](https://github.com/Aseeef) |
| Handle migration during backup according to migration priority  | [#16877](https://github.com/kubevirt/kubevirt/pull/16877) | [ShellyKa13](https://github.com/ShellyKa13) |
| BugFix: VMs requiring enlightenment are now able to be live migrated after a decentralized live migration  | [#16871](https://github.com/kubevirt/kubevirt/pull/16871) | [awels](https://github.com/awels) |
| Subtract non-schedulable nodes from kubevirt_allocatable_nodes  | [#16865](https://github.com/kubevirt/kubevirt/pull/16865) | [machadovilaca](https://github.com/machadovilaca) |
| fix VMExport failure with long PVC names  | [#16846](https://github.com/kubevirt/kubevirt/pull/16846) | [Aneesh-Hegde](https://github.com/Aneesh-Hegde) |
| fix: Prevent stale VMI backup status update when reusing backup names  | [#16399](https://github.com/kubevirt/kubevirt/pull/16399) | [Acedus](https://github.com/Acedus) |
| Allow disabling Velero hooks in virt-launcher via Annotation  | [#16786](https://github.com/kubevirt/kubevirt/pull/16786) | [alromeros](https://github.com/alromeros) |


_This page is updated daily._
