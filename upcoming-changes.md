---
title: Noteworthy changes for the next KubeVirt release
---

# Noteworthy changes for the next KubeVirt release

This list contains the noteworthy changes made after the latest KubeVirt release. The community expects these changes to be included in the next Kubevirt release.

> [!WARNING]
> **Please be aware that any of these might be excluded from the next release.**

| Upcoming changes | PR                                                                   | Author                                          |
|------------------|----------------------------------------------------------------------|-------------------------------------------------|
| Graduate the clone API to v1beta1 and deprecate v1alpha1  | [#13520](https://github.com/kubevirt/kubevirt/pull/13520) | [iholder101](https://github.com/iholder101) |
| Drop `ExperimentalVirtiofsSupport` feature gate in favor of `EnableVirtioFsConfigVolumes` for sharing ConfigMaps, Secrets, DownwardAPI and ServiceAccounts and `EnableVirtioFsPVC` for sharing PVCs.  | [#11997](https://github.com/kubevirt/kubevirt/pull/11997) | [jcanocan](https://github.com/jcanocan) |
| This version of KubeVirt includes upgraded virtualization technology based on libvirt 10.10.0 and QEMU 9.1.0.<br>Each new release of libvirt and QEMU contains numerous improvements and bug fixes.  | [#13641](https://github.com/kubevirt/kubevirt/pull/13641) | [andreabolognani](https://github.com/andreabolognani) |
| Bugfix: Support online snapshot of VMs with backend storage  | [#13682](https://github.com/kubevirt/kubevirt/pull/13682) | [alromeros](https://github.com/alromeros) |
| Bugfix: Support offline snapshot of VMs with backend storage  | [#13207](https://github.com/kubevirt/kubevirt/pull/13207) | [alromeros](https://github.com/alromeros) |
| Alert KubevirtVmHighMemoryUsage has been deprecated.  | [#13587](https://github.com/kubevirt/kubevirt/pull/13587) | [sradco](https://github.com/sradco) |
| Test suite: 3 new labels are available to filter tests: HostDiskGate, requireHugepages1Gi, blockrwo  | [#13109](https://github.com/kubevirt/kubevirt/pull/13109) | [xpivarc](https://github.com/xpivarc) |
| Add the iothreads option to specify number of iothreads to be used  | [#13110](https://github.com/kubevirt/kubevirt/pull/13110) | [alicefr](https://github.com/alicefr) |
| storage tests: assemble storage-oriented conformance test suite  | [#13586](https://github.com/kubevirt/kubevirt/pull/13586) | [akalenyu](https://github.com/akalenyu) |
| add support for virtio video device for amd64  | [#13606](https://github.com/kubevirt/kubevirt/pull/13606) | [dasionov](https://github.com/dasionov) |
| Storage tests: eliminate runtime skips  | [#13603](https://github.com/kubevirt/kubevirt/pull/13603) | [akalenyu](https://github.com/akalenyu) |
| BugFix: Volume hotplug broken with crun >= 1.18  | [#13546](https://github.com/kubevirt/kubevirt/pull/13546) | [akalenyu](https://github.com/akalenyu) |
| Ensure virt-tail and virt-monitor have the same timeout, preventing early termination of virt-tail while virt-monitor is still starting  | [#13588](https://github.com/kubevirt/kubevirt/pull/13588) | [Yu-Jack](https://github.com/Yu-Jack) |
| Upgrade of virt stack  | [#13545](https://github.com/kubevirt/kubevirt/pull/13545) | [alicefr](https://github.com/alicefr) |
| VMExport: exported DV uses the storage API  | [#13152](https://github.com/kubevirt/kubevirt/pull/13152) | [akalenyu](https://github.com/akalenyu) |
| Updated common-instancetypes bundles to v1.2.1  | [#13562](https://github.com/kubevirt/kubevirt/pull/13562) | [kubevirt-bot](https://github.com/kubevirt-bot) |
| virtctl expose now uses the unique `vm.kubevirt.io/name` label found on every virt-launcher Pod as a service selector.  | [#13496](https://github.com/kubevirt/kubevirt/pull/13496) | [0xFelix](https://github.com/0xFelix) |
| virtctl create vm validates disk names and prevents disk names that will lead to rejection of a VM upon creation.  | [#13547](https://github.com/kubevirt/kubevirt/pull/13547) | [0xFelix](https://github.com/0xFelix) |
| Fixed bug where VMs may not get the persistent EFI they requested  | [#13544](https://github.com/kubevirt/kubevirt/pull/13544) | [jean-edouard](https://github.com/jean-edouard) |
| Add kubevirt_vm_create_date_timestamp_seconds metric  | [#13431](https://github.com/kubevirt/kubevirt/pull/13431) | [avlitman](https://github.com/avlitman) |
| Bugfix: Support exporting backend PVC  | [#13460](https://github.com/kubevirt/kubevirt/pull/13460) | [alromeros](https://github.com/alromeros) |
| Build KubeVirt with go v1.22.10  | [#13495](https://github.com/kubevirt/kubevirt/pull/13495) | [brianmcarey](https://github.com/brianmcarey) |
| Remove deprecated DataVolume garbage collection tests  | [#13437](https://github.com/kubevirt/kubevirt/pull/13437) | [arnongilboa](https://github.com/arnongilboa) |
| Ensure IP not empty in kubevirt_vmi_status_addresses metric  | [#13386](https://github.com/kubevirt/kubevirt/pull/13386) | [machadovilaca](https://github.com/machadovilaca) |
| Bugfix: fix possible virt-handler race condition and stuck situation during shutdown  | [#13424](https://github.com/kubevirt/kubevirt/pull/13424) | [fossedihelm](https://github.com/fossedihelm) |
| Adjust managedTap binding to work with VMs with Address Conflict Detection enabled  | [#13458](https://github.com/kubevirt/kubevirt/pull/13458) | [orelmisan](https://github.com/orelmisan) |
| Add virt-handler cpu and memory usage metrics  | [#13250](https://github.com/kubevirt/kubevirt/pull/13250) | [Sreeja1725](https://github.com/Sreeja1725) |
| /var/lib/kubelet on the nodes can now be a symlink  | [#13263](https://github.com/kubevirt/kubevirt/pull/13263) | [jean-edouard](https://github.com/jean-edouard) |
| Auto-configured parallel QEMU-level migration threads (a.k.a. multifd)  | [#12705](https://github.com/kubevirt/kubevirt/pull/12705) | [iholder101](https://github.com/iholder101) |
| bug-fix: prevent status update for old migrations  | [#13426](https://github.com/kubevirt/kubevirt/pull/13426) | [dasionov](https://github.com/dasionov) |
| Unconditionally disable libvirt's VMPort feature which is relevant for VMWare only  | [#13252](https://github.com/kubevirt/kubevirt/pull/13252) | [iholder101](https://github.com/iholder101) |
| VMRestore: remove VMSnapshot logic from vmrestore webhook  | [#13305](https://github.com/kubevirt/kubevirt/pull/13305) | [ShellyKa13](https://github.com/ShellyKa13) |
| Bug-fix: Reduced probability of false "failed to detect socket for containerDisk disk0: ... connection refused" warnings  | [#13367](https://github.com/kubevirt/kubevirt/pull/13367) | [xpivarc](https://github.com/xpivarc) |
| Network Binding Plugin feature is declared GA  | [#13314](https://github.com/kubevirt/kubevirt/pull/13314) | [EdDev](https://github.com/EdDev) |
| Add node label to migration metrics  | [#13325](https://github.com/kubevirt/kubevirt/pull/13325) | [machadovilaca](https://github.com/machadovilaca) |
| Add Guest and Hugepages memory to kubevirt_vm_resource_requests  | [#13294](https://github.com/kubevirt/kubevirt/pull/13294) | [machadovilaca](https://github.com/machadovilaca) |
| Vmrestore - add options to handle cases when target is not ready  | [#13195](https://github.com/kubevirt/kubevirt/pull/13195) | [ShellyKa13](https://github.com/ShellyKa13) |
| Avoid NPE when getting filesystem overhead  | [#13138](https://github.com/kubevirt/kubevirt/pull/13138) | [mhenriks](https://github.com/mhenriks) |
| VMSnapshot: propagate freeze error failure  | [#13270](https://github.com/kubevirt/kubevirt/pull/13270) | [ShellyKa13](https://github.com/ShellyKa13) |
| added a new label to kubevirt_vmi_info metric named vmi_pod and contain the current pod name that runs the VMI.  | [#13148](https://github.com/kubevirt/kubevirt/pull/13148) | [avlitman](https://github.com/avlitman) |
| Enable volume migration for hotplugged volumes  | [#12800](https://github.com/kubevirt/kubevirt/pull/12800) | [alicefr](https://github.com/alicefr) |
| virtctl: Image uploads are retried up to 15 times  | [#12925](https://github.com/kubevirt/kubevirt/pull/12925) | [0xFelix](https://github.com/0xFelix) |
| BugFix: VMSnapshot 'InProgress' and Failing for a VM with InstanceType and Preference  | [#13260](https://github.com/kubevirt/kubevirt/pull/13260) | [akalenyu](https://github.com/akalenyu) |
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
