---
title: Noteworthy changes for the next KubeVirt release
---

# Noteworthy changes for the next KubeVirt release

This list contains the noteworthy changes made after the latest KubeVirt release. The community expects these changes to be included in the next Kubevirt release.

> [!WARNING]
> **Please be aware that any of these might be excluded from the next release.**

| Upcoming changes | PR                                                                   | Author                                          |
|------------------|----------------------------------------------------------------------|-------------------------------------------------|
| bugfix: snapshot and restore now works correctly for VMs after a storage volume migration  | [#15305](https://github.com/kubevirt/kubevirt/pull/15305) | [Acedus](https://github.com/Acedus) |
| Common Names are now enforce for aggregated API  | [#15314](https://github.com/kubevirt/kubevirt/pull/15314) | [xpivarc](https://github.com/xpivarc) |
| Bumped the bundled common-instancetypes to v1.4.0 which add new preferences.  | [#15253](https://github.com/kubevirt/kubevirt/pull/15253) | [0xFelix](https://github.com/0xFelix) |
| BugFix: export fails when VMExport has dots in secret  | [#15182](https://github.com/kubevirt/kubevirt/pull/15182) | [akalenyu](https://github.com/akalenyu) |
| Support for all `*_SHASUM` environment variables has been removed from the `virt-operator` component. Users should instead use the remaining `*_IMAGE` environment variables to request a specific image version using a tag, digest or both.  | [#15061](https://github.com/kubevirt/kubevirt/pull/15061) | [lyarwood](https://github.com/lyarwood) |
| virt-operator won't schedule on worker nodes  | [#15157](https://github.com/kubevirt/kubevirt/pull/15157) | [jean-edouard](https://github.com/jean-edouard) |
| Drop an arbitrary limitation on VM's domain.firmaware.serial. Any string is passed verbatim to smbios. Illegal may be tweaked or ignored based on qemu/smbios version.  | [#15118](https://github.com/kubevirt/kubevirt/pull/15118) | [dankenigsberg](https://github.com/dankenigsberg) |
| Update dependecy golang.org/x/oauth2 to v0.27.0  | [#15098](https://github.com/kubevirt/kubevirt/pull/15098) | [dominikholler](https://github.com/dominikholler) |
| Fix postcopy multifd compatibility during upgrade  | [#15016](https://github.com/kubevirt/kubevirt/pull/15016) | [fossedihelm](https://github.com/fossedihelm) |
| Update dependecy golang.org/x/net to v0.38.0  | [#15100](https://github.com/kubevirt/kubevirt/pull/15100) | [dominikholler](https://github.com/dominikholler) |
| BugFix: export fails when VMExport has dots in name  | [#15099](https://github.com/kubevirt/kubevirt/pull/15099) | [akalenyu](https://github.com/akalenyu) |
| allows virtual machine instances with an instance type to specify memory fields that do not conflict with the instance type  | [#14685](https://github.com/kubevirt/kubevirt/pull/14685) | [seanbanko](https://github.com/seanbanko) |
| Cleanup: libvmi: add consistently named cpu/mem setters  | [#14888](https://github.com/kubevirt/kubevirt/pull/14888) | [akalenyu](https://github.com/akalenyu) |
| Bugfix: Label upload PVCs to support CDI WebhookPvcRendering  | [#15067](https://github.com/kubevirt/kubevirt/pull/15067) | [alromeros](https://github.com/alromeros) |
| HostDisk: KubeVirt no longer performs chown/chmod to compensate for storage that doesn't support fsGroup  | [#15037](https://github.com/kubevirt/kubevirt/pull/15037) | [jean-edouard](https://github.com/jean-edouard) |
| Added support for architecture-specific configuration of `s390x` (IBM Z) in KubeVirt cluster config.  | [#15017](https://github.com/kubevirt/kubevirt/pull/15017) | [nekkunti](https://github.com/nekkunti) |
| The synchronization controller migration network IP address is advertised by the KubeVirt CR  | [#15022](https://github.com/kubevirt/kubevirt/pull/15022) | [awels](https://github.com/awels) |
| Decentralized migration resource now shows the synchronization address  | [#15021](https://github.com/kubevirt/kubevirt/pull/15021) | [awels](https://github.com/awels) |
| Add support for DRA devices such as GPUs and HostDevices.  | [#14365](https://github.com/kubevirt/kubevirt/pull/14365) | [alaypatel07](https://github.com/alaypatel07) |
| Decentralized live migration is available to allow migration across namespaces and clusters  | [#14882](https://github.com/kubevirt/kubevirt/pull/14882) | [awels](https://github.com/awels) |
| Beta: NodeRestriction  | [#14964](https://github.com/kubevirt/kubevirt/pull/14964) | [xpivarc](https://github.com/xpivarc) |
| Possible to trust additional CAs for verifying kubevirt infra structure components  | [#14986](https://github.com/kubevirt/kubevirt/pull/14986) | [awels](https://github.com/awels) |
| Support seamless TCP migration with passt (alpha)  | [#14875](https://github.com/kubevirt/kubevirt/pull/14875) | [nirdothan](https://github.com/nirdothan) |


_This page is updated daily._
