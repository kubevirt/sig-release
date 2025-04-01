---
title: Noteworthy changes for the next KubeVirt release
---

# Noteworthy changes for the next KubeVirt release

This list contains the noteworthy changes made after the latest KubeVirt release. The community expects these changes to be included in the next Kubevirt release.

> [!WARNING]
> **Please be aware that any of these might be excluded from the next release.**

| Upcoming changes | PR                                                                   | Author                                          |
|------------------|----------------------------------------------------------------------|-------------------------------------------------|
| Updated common-instancetypes bundles to v1.3.1  | [#14374](https://github.com/kubevirt/kubevirt/pull/14374) | [kubevirt-bot](https://github.com/kubevirt-bot) |
| A request to create a VirtualMachines that references a non-existent  instance type or preference are no longer rejected. The VirtualMachine will be created but will fail to start until the missing resources are created in the cluster.  | [#14219](https://github.com/kubevirt/kubevirt/pull/14219) | [lyarwood](https://github.com/lyarwood) |
| Don't expose as VMI status the implicit qemu domain pause at the end of live migration  | [#14288](https://github.com/kubevirt/kubevirt/pull/14288) | [qinqon](https://github.com/qinqon) |
| Fixed persistent reservation support for multipathd by improving socket access and multipath files in pr-helper  | [#14309](https://github.com/kubevirt/kubevirt/pull/14309) | [alicefr](https://github.com/alicefr) |
| fix: disks-images-provider is pointing to wrong alpine image for s390x.  | [#14325](https://github.com/kubevirt/kubevirt/pull/14325) | [vamsikrishna-siddu](https://github.com/vamsikrishna-siddu) |
| The `v1alpha{1,2}` versions of the `instancetype.kubevirt.io` API group are no longer served or supported.  | [#14048](https://github.com/kubevirt/kubevirt/pull/14048) | [lyarwood](https://github.com/lyarwood) |
| A new `Enabled` attribute has been added to the `TPM` device allowing users to explicitly disable the device regardless of any referenced preference.  | [#14316](https://github.com/kubevirt/kubevirt/pull/14316) | [lyarwood](https://github.com/lyarwood) |
| Cleanup: Fix unit tests on a sane, non-host-cgroup-sharing development setup  | [#14328](https://github.com/kubevirt/kubevirt/pull/14328) | [akalenyu](https://github.com/akalenyu) |
| Add interface name label to kubevirt_vmi_status_addresses  | [#14108](https://github.com/kubevirt/kubevirt/pull/14108) | [machadovilaca](https://github.com/machadovilaca) |
| The `InstancetypeReferencePolicy` feature has graduated to GA and no longer requires the associated feature gate to be enabled.  | [#14050](https://github.com/kubevirt/kubevirt/pull/14050) | [lyarwood](https://github.com/lyarwood) |
| Register k8s client-go latency metrics on init  | [#14286](https://github.com/kubevirt/kubevirt/pull/14286) | [machadovilaca](https://github.com/machadovilaca) |
| Update module github.com/containers/common to v0.60.4  | [#14304](https://github.com/kubevirt/kubevirt/pull/14304) | [jean-edouard](https://github.com/jean-edouard) |
| VM Persistent State GA  | [#14065](https://github.com/kubevirt/kubevirt/pull/14065) | [jean-edouard](https://github.com/jean-edouard) |
| VMSnapshot: add QuiesceFailed indication to snapshot if freeze failed  | [#14096](https://github.com/kubevirt/kubevirt/pull/14096) | [ShellyKa13](https://github.com/ShellyKa13) |
| Update module golang.org/x/oauth2 to v0.27.0  | [#14215](https://github.com/kubevirt/kubevirt/pull/14215) | [dominikholler](https://github.com/dominikholler) |
| Default VM Rollout Strategy is now LiveUpdate. Important: to preserve previous behavior, rolloutStrategy needs to be set to "Stage" in the KubeVirt CR.  | [#14068](https://github.com/kubevirt/kubevirt/pull/14068) | [jean-edouard](https://github.com/jean-edouard) |
| Update module golang.org/x/net to v0.36.0  | [#14222](https://github.com/kubevirt/kubevirt/pull/14222) | [dominikholler](https://github.com/dominikholler) |
| Update golang.org/x/crypto to v0.35.0  | [#14218](https://github.com/kubevirt/kubevirt/pull/14218) | [dominikholler](https://github.com/dominikholler) |
| Update module github.com/opencontainers/runc to v1.1.14  | [#14217](https://github.com/kubevirt/kubevirt/pull/14217) | [dominikholler](https://github.com/dominikholler) |
| Large number of migrations should no longer lead to active migrations timing out  | [#14141](https://github.com/kubevirt/kubevirt/pull/14141) | [jean-edouard](https://github.com/jean-edouard) |
| Ensure launcher pods are finalized and deleted before removing the VMI finalizer when the VMI is marked for deletion.  | [#13870](https://github.com/kubevirt/kubevirt/pull/13870) | [dasionov](https://github.com/dasionov) |
| libvirt: 10.10.0-7, qemu: 9.1.0-15  | [#14101](https://github.com/kubevirt/kubevirt/pull/14101) | [qinqon](https://github.com/qinqon) |
| Add entrypoint to the pr-helper for creating the symlink to the multipath socket  | [#14071](https://github.com/kubevirt/kubevirt/pull/14071) | [alicefr](https://github.com/alicefr) |
| Support live migration to a named node  | [#12725](https://github.com/kubevirt/kubevirt/pull/12725) | [tiraboschi](https://github.com/tiraboschi) |
| Add v1.5.0 perf and scale benchmarks data  | [#13888](https://github.com/kubevirt/kubevirt/pull/13888) | [Sreeja1725](https://github.com/Sreeja1725) |
| The virtctl port-forward/ssh/scp syntax was changed to type/name[/namespace]. It now supports resources with dots in their name properly.  | [#13939](https://github.com/kubevirt/kubevirt/pull/13939) | [0xFelix](https://github.com/0xFelix) |
| virt-launcher now uses bash to retrieve disk info and verify container-disk files, requiring bash to be included in the launcher image  | [#13807](https://github.com/kubevirt/kubevirt/pull/13807) | [Barakmor1](https://github.com/Barakmor1) |
| Network interfaces state can be set to `down` or `up` in order to set the link state accordingly when VM is running. Hot plugging of interface in these states is also supported.  | [#13744](https://github.com/kubevirt/kubevirt/pull/13744) | [nirdothan](https://github.com/nirdothan) |
| Interrupted migrations will now be reconciled on next VM start.  | [#13536](https://github.com/kubevirt/kubevirt/pull/13536) | [jean-edouard](https://github.com/jean-edouard) |
| bug-fix: add machine type to `NodeSelector` to prevent breaking changes on unsupported nodes  | [#13690](https://github.com/kubevirt/kubevirt/pull/13690) | [dasionov](https://github.com/dasionov) |
| The node-restriction Validating Admission Policy will return consistent reasons on failures  | [#13940](https://github.com/kubevirt/kubevirt/pull/13940) | [tiraboschi](https://github.com/tiraboschi) |
| Instance type and preference runtime data is now stored under `Status.{Instancetype,Preference}Ref` and is no longer mutated into the core VirtualMachine` `Spec`.  | [#13916](https://github.com/kubevirt/kubevirt/pull/13916) | [lyarwood](https://github.com/lyarwood) |
| VMClone: Remove webhook that checks Snapshot Source  | [#13831](https://github.com/kubevirt/kubevirt/pull/13831) | [ShellyKa13](https://github.com/ShellyKa13) |
| GA ClusterProfiler FG and add a config to enable it  | [#13815](https://github.com/kubevirt/kubevirt/pull/13815) | [acardace](https://github.com/acardace) |
| Updated common-instancetypes bundles to v1.3.0  | [#13928](https://github.com/kubevirt/kubevirt/pull/13928) | [kubevirt-bot](https://github.com/kubevirt-bot) |
| Fetch non-cluster instance type and preferences with namespace key  | [#13805](https://github.com/kubevirt/kubevirt/pull/13805) | [machadovilaca](https://github.com/machadovilaca) |


_This page is updated daily._
