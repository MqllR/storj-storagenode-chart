## STORJ Storagenode

Helm chart install a [storj](https://storj.io/) storagenode.

### Prerequisites

#### Identity

The storage node needs an identity previously generated https://documentation.storj.io/dependencies/identity.

#### Storage

This helm chart use a StatefulSet object with a volumeClaimTemplates. If the storagenode run on a baremetal machine without a PersistentVolume controller (https://kubernetes.io/docs/concepts/storage/persistent-volumes/#types-of-persistent-volumes), you can use the no-provisioner:

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
    path: /data/disk1    # -> Where your disk is locally mounted
  nodeAffinity:
    required:
      nodeSelectorTerms:
      - matchExpressions:
        - key: kubernetes.io/hostname
          operator: In
          values:
            - mysweetmachine   # -> Set your node hostname
```

### Configuration

Parameter | Description | Default | Required
--- | --- | --- | ---
`config.wallet` | ERC20 address for payment  | `nil` | yes
`config.email` | Email address used by Storj  | `nil` | yes
`config.address` | (domain\|ip):port for external communication | `nil` | yes
`config.storage` | Storage size allocated  | `1TB` | no
`config.bandwidth` | Bandwidth allocated  | `2TiB` | no
`service.type` | Service type for the storagenode | `NodePort` | no
`service.annotations` |  | `{}` | no
`service.port` | Service port for the storagenode | `28967` | no
`service.nodePort` | Node port to expose for the storagenode | "" | no
`replicaCount` | Number of replica | `1` | no
`podAnnotations` | Annotations for the pod | `{}` | no
`podSecurityContext` | Custom security context | `{}` | no
`nodeSelector` | Node labels for pod assignment	 | `{}` | no
`tolerations` | Node taints to tolerate | `{}` | no
`affinity` | Pod affinity | `{}` | no
`imagePullSecrets` |  | `{}` | no
`storagenode.image.repository` | Image name | `storjlabs/storagenode:beta` | no
`storagenode.image.pullPolicy` | Container pull policy | `Always` | no
`storagenode.securtyContext` | Custom security context for container | `{}` | no
`storagenode.resources` | Resources request and limit YAML | `{}` | no
`nodeStats.enabled` | Expose the node's Dashboard | `true` | no
`nodeStats.service.type` | Service type for the dashboard | `ClusterIP` | no
`nodeStats.service.port` | Service port for the dashboard | `14002` | no
`identity.localPath` | Relative path to the chart. Must be in the chart directory | `secrets/*` | no
`identity.externalSecret.enabled` | Use an external secret | `false` | no
`identity.externalSecret.secretName` | Specify the secretName | `""` | no
`persistence.enabled` | Create a persistence volume | `true` | no
`persistence.annotations` | Persistent volume claim annotation | `{}` | no
`persistence.volumeClaimTemplate` | Persistent volume claim YAML | `{}` | no
`metrics.enabled` | Start the container storj-exporter | `true` | no
`metrics.image.repository` | Image name | `anclrii/storj-exporter:0.2.4` | no
`metrics.image.pullPolicy` | Container pull policy | `IfNotPresent` | no
`metrics.securtyContext` | Custom security context for container | `{}` | no
`metrics.resources` | Resources request and limit YAML | `{}` | no
`ingress.enabled` | If true, an ingress object will be created | `false` | no
`ingress.annotations` | Annotations for the ingress | `{}` | no
`ingress.hosts` | Ingress hostname | `[]` | no
`ingress.tls` | Ingress TLS | `[]` | no

### Installing the Chart

1. Clone the repo

```
git clone https://github.com/MqllR/storj-storagenode-chart
```

2. Create your identity

As explain in this [documentation](https://helm.sh/docs/chart_template_guide/accessing_files/), helm can't access files outside of the chart directory. You can create the identity by yourself or copy the storj identity path to the local chart directory:

```
mkdir -p secrets/identity-storagenode/
cp ~/.local/share/storj/identity/storagenode/ secrets/identity-storagenode/
```

Install the chart

```
helm install node1 . --set config.email=mymail@domain.com,config.wallet=0xdfca4035b9f16c40b558218d1bedc08590fe28d4,config.address=mydomain.net:28967,identity.localPath=secrets/identity-storagenode/\*
```
