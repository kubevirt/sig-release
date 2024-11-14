---
title: Noteworthy changes for the next KubeVirt release
---

# Noteworthy changes for the next KubeVirt release

This list contains the noteworthy changes made after the latest KubeVirt release. The community expects these changes to be included in the next Kubevirt release.

> [!WARNING]
> **Please be aware that any of these might be excluded from the next release.**

| Upcoming changes | PR                                                                   | Author                                          |
|------------------|----------------------------------------------------------------------|-------------------------------------------------|
| Fix issue starting Virtual Machine Export when succeed/failed VMI exists for that VM  | [#13240](https://github.com/kubevirt/kubevirt/pull/13240) | [awels](https://github.com/awels) |
| The inflexible `PreferredUseEFi` and `PreferredUseSecureBoot` preference fields have been deprecated ahead of removal in a future version of the `instancetype.kubevirt.io` API. Users should instead use `PreferredEfi` to provide a preferred `EFI` configuration for their `VirtualMachine`.  | [#12750](https://github.com/kubevirt/kubevirt/pull/12750) | [lyarwood](https://github.com/lyarwood) |
| backend-storage will now correctly use the default virtualization storage class  | [#13219](https://github.com/kubevirt/kubevirt/pull/13219) | [jean-edouard](https://github.com/jean-edouard) |
| Add release v1.4.0 perf and scale benchmarks data  | [#13204](https://github.com/kubevirt/kubevirt/pull/13204) | [Sreeja1725](https://github.com/Sreeja1725) |
| BugFix: VMSnapshots broken on OpenShift  | [#13197](https://github.com/kubevirt/kubevirt/pull/13197) | [akalenyu](https://github.com/akalenyu) |
| kubevirt_vm_disk_allocated_size_bytes metric added in order to monitor vm sizes  | [#12765](https://github.com/kubevirt/kubevirt/pull/12765) | [avlitman](https://github.com/avlitman) |
| Update promql query of cpu and memory metrics for sig-performance tests  | [#12546](https://github.com/kubevirt/kubevirt/pull/12546) | [Sreeja1725](https://github.com/Sreeja1725) |
| Enable virt-exportproxy and virt-exportserver image for s390x  | [#12844](https://github.com/kubevirt/kubevirt/pull/12844) | [jschintag](https://github.com/jschintag) |
| VMs admitter: remove validation of vm clone volume from the webhook  | [#12628](https://github.com/kubevirt/kubevirt/pull/12628) | [ShellyKa13](https://github.com/ShellyKa13) |
| Added labels, annotations to VM Export resources and configurable pod readiness timeout  | [#13006](https://github.com/kubevirt/kubevirt/pull/13006) | [chomatdam](https://github.com/chomatdam) |
| GA the `VMLiveUpdateFeatures` feature-gate.  | [#13091](https://github.com/kubevirt/kubevirt/pull/13091) | [acardace](https://github.com/acardace) |


_This page is updated daily._
