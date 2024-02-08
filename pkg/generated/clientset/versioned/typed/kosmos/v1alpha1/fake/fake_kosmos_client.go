// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/kosmos.io/kosmos/pkg/generated/clientset/versioned/typed/kosmos/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeKosmosV1alpha1 struct {
	*testing.Fake
}

func (c *FakeKosmosV1alpha1) Clusters() v1alpha1.ClusterInterface {
	return &FakeClusters{c}
}

func (c *FakeKosmosV1alpha1) ClusterNodes() v1alpha1.ClusterNodeInterface {
	return &FakeClusterNodes{c}
}

func (c *FakeKosmosV1alpha1) ClusterPodConvertPolicies() v1alpha1.ClusterPodConvertPolicyInterface {
	return &FakeClusterPodConvertPolicies{c}
}

func (c *FakeKosmosV1alpha1) DaemonSets(namespace string) v1alpha1.DaemonSetInterface {
	return &FakeDaemonSets{c, namespace}
}

func (c *FakeKosmosV1alpha1) Knodes() v1alpha1.KnodeInterface {
	return &FakeKnodes{c}
}

func (c *FakeKosmosV1alpha1) NodeConfigs() v1alpha1.NodeConfigInterface {
	return &FakeNodeConfigs{c}
}

func (c *FakeKosmosV1alpha1) PodConvertPolicies(namespace string) v1alpha1.PodConvertPolicyInterface {
	return &FakePodConvertPolicies{c, namespace}
}

func (c *FakeKosmosV1alpha1) ShadowDaemonSets(namespace string) v1alpha1.ShadowDaemonSetInterface {
	return &FakeShadowDaemonSets{c, namespace}
}

func (c *FakeKosmosV1alpha1) VirtualClusters() v1alpha1.VirtualClusterInterface {
	return &FakeVirtualClusters{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeKosmosV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
