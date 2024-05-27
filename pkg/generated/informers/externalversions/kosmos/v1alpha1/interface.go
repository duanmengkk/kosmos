// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/kosmos.io/kosmos/pkg/generated/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Clusters returns a ClusterInformer.
	Clusters() ClusterInformer
	// ClusterDistributionPolicies returns a ClusterDistributionPolicyInformer.
	ClusterDistributionPolicies() ClusterDistributionPolicyInformer
	// ClusterNodes returns a ClusterNodeInformer.
	ClusterNodes() ClusterNodeInformer
	// ClusterPodConvertPolicies returns a ClusterPodConvertPolicyInformer.
	ClusterPodConvertPolicies() ClusterPodConvertPolicyInformer
	// DaemonSets returns a DaemonSetInformer.
	DaemonSets() DaemonSetInformer
	// DistributionPolicies returns a DistributionPolicyInformer.
	DistributionPolicies() DistributionPolicyInformer
	// GlobalNodes returns a GlobalNodeInformer.
	GlobalNodes() GlobalNodeInformer
	// NodeConfigs returns a NodeConfigInformer.
	NodeConfigs() NodeConfigInformer
	// PodConvertPolicies returns a PodConvertPolicyInformer.
	PodConvertPolicies() PodConvertPolicyInformer
	// ShadowDaemonSets returns a ShadowDaemonSetInformer.
	ShadowDaemonSets() ShadowDaemonSetInformer
	// VirtualClusters returns a VirtualClusterInformer.
	VirtualClusters() VirtualClusterInformer
	// VirtualClusterPlugins returns a VirtualClusterPluginInformer.
	VirtualClusterPlugins() VirtualClusterPluginInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// Clusters returns a ClusterInformer.
func (v *version) Clusters() ClusterInformer {
	return &clusterInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ClusterDistributionPolicies returns a ClusterDistributionPolicyInformer.
func (v *version) ClusterDistributionPolicies() ClusterDistributionPolicyInformer {
	return &clusterDistributionPolicyInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ClusterNodes returns a ClusterNodeInformer.
func (v *version) ClusterNodes() ClusterNodeInformer {
	return &clusterNodeInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ClusterPodConvertPolicies returns a ClusterPodConvertPolicyInformer.
func (v *version) ClusterPodConvertPolicies() ClusterPodConvertPolicyInformer {
	return &clusterPodConvertPolicyInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// DaemonSets returns a DaemonSetInformer.
func (v *version) DaemonSets() DaemonSetInformer {
	return &daemonSetInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// DistributionPolicies returns a DistributionPolicyInformer.
func (v *version) DistributionPolicies() DistributionPolicyInformer {
	return &distributionPolicyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// GlobalNodes returns a GlobalNodeInformer.
func (v *version) GlobalNodes() GlobalNodeInformer {
	return &globalNodeInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// NodeConfigs returns a NodeConfigInformer.
func (v *version) NodeConfigs() NodeConfigInformer {
	return &nodeConfigInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// PodConvertPolicies returns a PodConvertPolicyInformer.
func (v *version) PodConvertPolicies() PodConvertPolicyInformer {
	return &podConvertPolicyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ShadowDaemonSets returns a ShadowDaemonSetInformer.
func (v *version) ShadowDaemonSets() ShadowDaemonSetInformer {
	return &shadowDaemonSetInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// VirtualClusters returns a VirtualClusterInformer.
func (v *version) VirtualClusters() VirtualClusterInformer {
	return &virtualClusterInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// VirtualClusterPlugins returns a VirtualClusterPluginInformer.
func (v *version) VirtualClusterPlugins() VirtualClusterPluginInformer {
	return &virtualClusterPluginInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
