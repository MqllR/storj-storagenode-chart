## STORJ Storagenode

This helm chart install a storj storagenode https://documentation.storj.io/setup/cli/storage-node

NOTE: Will soon implement its prometheus exporter.

### Prerequisites

#### Identity

The storage node needs an identity previously generated https://documentation.storj.io/dependencies/identity. After that, copy the identity from `$HOME/.local/share/storj/identity/storagenode/` to `<chart_path>/secrets/<release_name>/`

#### Storage

This helm chart use a StatefulSet object with a volumeClaimTemplates. If the storagenode run on a baremetal machine without a PersistentVolume implemented (https://kubernetes.io/docs/concepts/storage/persistent-volumes/#types-of-persistent-volumes), it can use the no-provisioner:

```
---
kind: StorageClass
apiVersion: storage.k8s.io/v1
metadata:
  name: local-storage
provisioner: kubernetes.io/no-provisioner
volumeBindingMode: WaitForFirstConsumer

---
apiVersion: v1
kind: PersistentVolume
metadata:
  name: local-pv-storagenode1
spec:
  capacity:
    storage: 2Ti
  accessModes:
  - ReadWriteOnce
  persistentVolumeReclaimPolicy: Retain
  storageClassName: local-storage
  local:
    path: /data/disk1    # ->> CHANGEME
  nodeAffinity:
    required:
      nodeSelectorTerms:
      - matchExpressions:
        - key: kubernetes.io/hostname
          operator: In
          values:
            - mysweetmachine
```

### Configuration

Parameter | Description | Default | Required
--- | --- | --- | ---
`configWallet` | ERC20 address for payment  | `nil` | yes
`configEmail` | Email address used by Storj  | `nil` | yes
`configStorage` | Storage size allocated  | `nil` | yes
`configBandwidth` | Bandwidth allocated  | `nil` | yes
`configBandwidth` | Bandwidth allocated  | `nil` | yes
`identityLocalPath` | Relative path to the chart. Must be in secrets/ | `nil` | yes
`volumeClaimTemplate.storageClassName` | storageClass used for this node | `nil` | no
