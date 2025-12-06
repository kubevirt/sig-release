---
title: Noteworthy changes for the next KubeVirt release
---

# Noteworthy changes for the next KubeVirt release

This list contains the noteworthy changes made after the latest KubeVirt release. The community expects these changes to be included in the next Kubevirt release.

> [!WARNING]
> **Please be aware that any of these might be excluded from the next release.**

| Upcoming changes | PR                                                                   | Author                                          |
|------------------|----------------------------------------------------------------------|-------------------------------------------------|
| bugfix: prevent cross-vendor migrations  | [#16060](https://github.com/kubevirt/kubevirt/pull/16060) | [dasionov](https://github.com/dasionov) |
| Add event logging for pause and unpause VM operations to align with other VM lifecycle events such as reset  | [#15821](https://github.com/kubevirt/kubevirt/pull/15821) | [SamAlber](https://github.com/SamAlber) |
| VirtualMachinePool now correctly appends index to CloudInit secret references when appendIndexToSecretRefs: true is set, enabling unique cloud-init configurations for each VM in the pool.  | [#15868](https://github.com/kubevirt/kubevirt/pull/15868) | [frank-gen](https://github.com/frank-gen) |
| The `EnableVirtioFsConfigVolumes` feature has graduated to GA and no longer requires the associated feature gate to be enabled.  | [#15913](https://github.com/kubevirt/kubevirt/pull/15913) | [germag](https://github.com/germag) |
| Test Fix: make Alpine ISO mount checks architecture-agnostic  | [#15863](https://github.com/kubevirt/kubevirt/pull/15863) | [HarshithaMS005](https://github.com/HarshithaMS005) |
| Document allowed values for `spec.runStrategy`.  | [#16122](https://github.com/kubevirt/kubevirt/pull/16122) | [dasionov](https://github.com/dasionov) |
| Don't use attachment pods marked for deletion for hotplug volume status updates.  | [#16159](https://github.com/kubevirt/kubevirt/pull/16159) | [Dsanatar](https://github.com/Dsanatar) |
| Allow VMExport with PVCs from Completed Pods  | [#15442](https://github.com/kubevirt/kubevirt/pull/15442) | [Dsanatar](https://github.com/Dsanatar) |
| Migration is using dedicated certificate for mTLS.  | [#15949](https://github.com/kubevirt/kubevirt/pull/15949) | [xpivarc](https://github.com/xpivarc) |
| fix: KSM is enabled in case of node pressure within 3 minutes  | [#16049](https://github.com/kubevirt/kubevirt/pull/16049) | [fossedihelm](https://github.com/fossedihelm) |
| Introduce new API - UtilityVolumes - direct virt-launcher attachment mechanism  | [#15922](https://github.com/kubevirt/kubevirt/pull/15922) | [ShellyKa13](https://github.com/ShellyKa13) |
| kubevirt.io/cpumanager label is advertised for nodes capable of running dedicated VMs.  | [#14892](https://github.com/kubevirt/kubevirt/pull/14892) | [xpivarc](https://github.com/xpivarc) |
| Allow migration when host model changes after libvirt upgrade.  | [#15694](https://github.com/kubevirt/kubevirt/pull/15694) | [Barakmor1](https://github.com/Barakmor1) |
| Add RestartRequired when detaching CD-ROMs from a running VM  | [#15969](https://github.com/kubevirt/kubevirt/pull/15969) | [Dsanatar](https://github.com/Dsanatar) |
| Add GuestFilesystemAlmostOutOfSpace alerts  | [#15714](https://github.com/kubevirt/kubevirt/pull/15714) | [machadovilaca](https://github.com/machadovilaca) |
| Introduce a new subresource `/evacuate/cancel` and `virtctl evacuate-cancel` command to allow users to cancel the evacuation process for a VirtualMachineInstance (VMI). This clears the `evacuationNodeName` field in the VMI's status, stopping the automatic creation of migration resources and fully aborting the eviction cycle.  | [#15957](https://github.com/kubevirt/kubevirt/pull/15957) | [xpivarc](https://github.com/xpivarc) |
| The `MultiArchitecture` feature gate has been deprecated and is no longer used to determine if VirtualMachines with a differing architecture to the control plane should be rejected by the admission webhooks  | [#16023](https://github.com/kubevirt/kubevirt/pull/16023) | [lyarwood](https://github.com/lyarwood) |
| Reject stop requests for paused VMIs.  A paused VMI must be unpaused before it can be stopped.  | [#15405](https://github.com/kubevirt/kubevirt/pull/15405) | [dasionov](https://github.com/dasionov) |
| A decentralized live migration failure is properly propagates between source and target  | [#15716](https://github.com/kubevirt/kubevirt/pull/15716) | [awels](https://github.com/awels) |
| NodeRestriction: Source of node update is now verified  | [#15374](https://github.com/kubevirt/kubevirt/pull/15374) | [xpivarc](https://github.com/xpivarc) |
| Bug fix: KubeVirt.spec.imagetag installation is working again  | [#16050](https://github.com/kubevirt/kubevirt/pull/16050) | [xpivarc](https://github.com/xpivarc) |
| Recording rule kubevirt_vmi_vcpu_count name changes to vmi:kubevirt_vmi_vcpu:count  | [#15968](https://github.com/kubevirt/kubevirt/pull/15968) | [sradco](https://github.com/sradco) |
| Introduce pool.kubevirt.io/v1beta1  | [#15166](https://github.com/kubevirt/kubevirt/pull/15166) | [Sreeja1725](https://github.com/Sreeja1725) |
| VMSnapshot: add SourceIndications status field to list snapshot indications with descriptions for clearer meaning.  | [#15409](https://github.com/kubevirt/kubevirt/pull/15409) | [noamasu](https://github.com/noamasu) |
| Promote IBM Secure Execution Feature to Beta stage.  | [#15934](https://github.com/kubevirt/kubevirt/pull/15934) | [jschintag](https://github.com/jschintag) |
| BugFix: The migration limit was not accurately being used with decentralized live migrations  | [#15767](https://github.com/kubevirt/kubevirt/pull/15767) | [awels](https://github.com/awels) |
| The KubevirtSeccompProfile feature is now in Beta  | [#15970](https://github.com/kubevirt/kubevirt/pull/15970) | [jean-edouard](https://github.com/jean-edouard) |
| promote ImageVolume FG to Beta  | [#15960](https://github.com/kubevirt/kubevirt/pull/15960) | [Barakmor1](https://github.com/Barakmor1) |
| VMPool: Add support for auto-healing startegy  | [#15638](https://github.com/kubevirt/kubevirt/pull/15638) | [Sreeja1725](https://github.com/Sreeja1725) |
| VMpool: Add Scale-in strategy support with Proactive, opportunistic modes and statePreservation  | [#15604](https://github.com/kubevirt/kubevirt/pull/15604) | [Sreeja1725](https://github.com/Sreeja1725) |
| support v0.32.5 code generator  | [#15529](https://github.com/kubevirt/kubevirt/pull/15529) | [Yu-Jack](https://github.com/Yu-Jack) |


_This page is updated daily._
