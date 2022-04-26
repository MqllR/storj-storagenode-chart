## STORJ Storagenode

![Release Helm Charts](https://github.com/MqllR/storj-storagenode-chart/workflows/Release%20Helm%20Charts/badge.svg)

Helm chart to install a [storj](https://storj.io/) storagenode.

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

#### Sysctl Configuration for UDP

If you have `service.quic` enabled (the default), you will need to update the `net.core.rmem_max` sysctl value, or the storj storagenode will complain in the logs. Doing this on your worker node(s) will pass the setting through to the storj pod. Note that this may impact other pods as well, since it is being done at the node level. Please see https://docs.storj.io/node/dependencies/quic-requirements/linux-configuration-for-udp/ for full details and instructions.

### Configuration

Parameter | Description | Default | Required
--- | --- | --- | ---
`config.wallet` | ERC20 address for payment  | `nil` | yes
`config.email` | Email address used by Storj  | `nil` | yes
`config.address` | (domain\|ip):port for external communication | `nil` | yes
`config.storage` | Storage size allocated  | `1TB` | no
`service.storagenode.type` | Service type for the storagenode | `NodePort` | no
`service.storagenode.annotations` | Service annotations | `{}` | no
`service.storagenode.loadBalancerIP` | Secify a Load balancer IP if the provider allow you | null | no
`service.storagenode.port` | Service port for the storagenode | `28967` | no
`service.storagenode.nodePort` | Node port to expose for the storagenode | "" | no
`service.stats.enabled` | Expose the node's Dashboard | `true` | no
`service.stats.type` | Service type for the dashboard | `ClusterIP` | no
`service.stats.port` | Service port for the dashboard | `14002` | no
`service.quic.enabled` | Expose the storagenode's UDP port for quic | `true` | no
`service.quic.type` | Service type for the storagenode's UDP port | `NodePort` | no
`service.quic.loadBalancerIP` | Secify a Load balancer IP if the provider allow you | null | no
`service.quic.port` | Service port for the storagenode's UDP port | `28967` | no
`service.quic.nodePort` | Node port to expose for the storagenode's UDP port | "" | no
`replicaCount` | Number of replica | `1` | no
`podAnnotations` | Annotations for the pod | `{}` | no
`podSecurityContext` | Custom security context | `{}` | no
`nodeSelector` | Node labels for pod assignment	 | `{}` | no
`tolerations` | Node taints to tolerate | `{}` | no
`affinity` | Pod affinity | `{}` | no
`imagePullSecrets` |  | `{}` | no
`storagenode.image.repository` | Image name | `storjlabs/storagenode:latest` | no
`storagenode.image.pullPolicy` | Container pull policy | `Always` | no
`storagenode.securtyContext` | Custom security context for container | `{}` | no
`storagenode.resources` | Resources request and limit YAML | `{}` | no
`identity.externalSecret.secretName` | Specify the secretName | `""` | yes
`persistence.enabled` | Create a persistence volume | `true` | no
`persistence.annotations` | Persistent volume claim annotation | `{}` | no
`persistence.volumeClaimTemplate` | Persistent volume claim YAML | `{}` | no
`metrics.enabled` | Start the container storj-exporter | `true` | no
`metrics.image.repository` | Image name | `anclrii/storj-exporter:1.0.10` | no
`metrics.image.pullPolicy` | Container pull policy | `IfNotPresent` | no
`metrics.securtyContext` | Custom security context for container | `{}` | no
`metrics.resources` | Resources request and limit YAML | `{}` | no
`ingress.enabled` | If true, an ingress object will be created to expose the __stat dashboard only__ | `false` | no
`ingress.annotations` | Annotations for the ingress | `{}` | no
`ingress.hosts` | Ingress hostname | `[]` | no
`ingress.tls` | Ingress TLS | `[]` | no

### Installing the Chart

1. Add the helm repo

```
helm repo add storj-storagenode-chart https://mqllr.github.io/storj-storagenode-chart
```

2. Create your identity

Create a secret in kubernetes with your node identity

```
curl -L https://github.com/MqllR/storj-storagenode-chart/releases/download/identity/identity-to-kube-secret-amd64 -o identity-to-kube-secret
./identity-to-kube-secret -secret-name "storj-identity-node1" | kubectl apply -f -
```

3. Install the chart

```
helm install node storj-storagenode-chart/storj-storagenode --set config.email=mymail@domain.com,config.wallet=0xdfca4035b9f16c40b558218d1bedc08590fe28d4,config.address=mydomain.net:28967,identity.externalSecret.secretName="storj-identity-node1"
```
