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
`config.wallet` | ERC20 address for payment  | `nil` | yes
`config.email` | Email address used by Storj  | `nil` | yes
`config.address` | (domain\|ip):port for external communication | `nil` | yes
`config.storage` | Storage size allocated  | 1TB | no
`configbandwidth` | Bandwidth allocated  | 2TiB | no
`nodeStats.enabled` | Expose the node Dashboard | true | no
`identityLocalPath` | Relative path to the chart. Must be in secrets/ | `secrets/*` | no
`volumeClaimTemplate.storageClassName` | storageClass used for this node | `nil` | no
`metrics.enabled` | Start the container storj exporter | `true` | no

### Installing the Chart

Clone the repo

```
git clone https://github.com/MqllR/storj-storagenode-chart
```

Create your identity

```
mkdir -p secrets/identity-storagenode/
cp ~/.local/share/storj/identity/storagenode/ secrets/identity-storagenode/
```

Install the chart

```
helm install node . --set config.email=mymail@domain.com,config.wallet=0xdfca4035b9f16c40b558218d1bedc08590fe28d4,config.address=mydomain.net:28967,identityLocalPath=secrets/identity-storagenode/\*
```
