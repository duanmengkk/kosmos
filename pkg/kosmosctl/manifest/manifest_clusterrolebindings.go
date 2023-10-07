package manifest

const (
	ClusterlinkNetworkManagerClusterRoleBinding = `
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: clusterlink-network-manager
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: clusterlink-network-manager
subjects:
  - kind: ServiceAccount
    name: clusterlink-network-manager
    namespace: {{ .Namespace }}
`

	ClusterlinkFloaterClusterRoleBinding = `
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: clusterlink-floater
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: clusterlink-floater
subjects:
  - kind: ServiceAccount
    name: clusterlink-floater
    namespace: {{ .Namespace }}
`
	ClusterlinkClusterRoleBinding = `
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: clusterlink
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: clusterlink
subjects:
  - kind: ServiceAccount
    name: clusterlink-controller-manager
    namespace: {{ .Namespace }}
  - kind: ServiceAccount
    name: clusterlink-operator
    namespace: {{ .Namespace }}
`

	ClusterTreeKnodeManagerClusterRoleBinding = `
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: clustertree-knode-manager
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: clustertree-knode-manager
subjects:
  - kind: ServiceAccount
    name: clustertree-knode-manager
    namespace: {{ .Namespace }}
`
)

type ClusterRoleBindingReplace struct {
	Namespace string
}
