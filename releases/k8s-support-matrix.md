# KubeVirt to Kubernetes version support matrix

| KubeVirt version | 1.36               | 1.35               | 1.34               | 1.33<sup>EOL</sup> | 1.32<sup>EOL</sup> | 1.31<sup>EOL</sup> |
|------------------|--------------------|--------------------|--------------------|--------------------|--------------------|--------------------|
|              1.9 | ✓                  | ✓                  | ✓                  | -                  | -                  | -                  |
|              1.8 | -                  | ✓                  | ✓                  | EOL                | -                  | -                  |
|              1.7 | -                  | -                  | ✓                  | EOL                | EOL                | -                  |
|              1.6 | -                  | -                  | -                  | EOL                | EOL                | EOL                |
|              1.5 | -                  | -                  | -                  | -                  | EOL                | EOL                |
|              1.4 | -                  | -                  | -                  | -                  | -                  | EOL                |


Note: _EOL_ means that the Kubernetes version was supported by KubeVirt but has reached end of life. See [Kubernetes releases](https://kubernetes.io/releases/) for more details

## Support Exceptions

The KubeVirt community maintains the ability to extend support for a KubeVirt
release to older Kubernetes versions that may fall outside of the
latest 3 Kubernetes releases present at the time of the KubeVirt release.

For more details on Kubernetes version compatibility, see the [Kubernetes Version Compatibility](https://github.com/kubevirt/kubevirt/blob/main/docs/kubernetes-compatibility.md#support-exceptions)
documentation in the kubevirt/kubevirt repository.
