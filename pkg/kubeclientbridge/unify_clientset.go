package kubeclientbridge

import (
	"fmt"
	kosmosv1 "github.com/kosmos.io/kosmos/pkg/kubeclientbridge/typed/core/v1"
	fakekube "k8s.io/client-go/kubernetes/fake"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/util/flowcontrol"
	"net/http"
)

// UnifyClientset extends Clientset
type UnifyClientset struct {
	*fakekube.Clientset
	coreV1 *kosmosv1.UnifyCoreV1
}

func NewUnifyClientset(c *rest.Config) (*UnifyClientset, error) {
	configShallowCopy := *c

	if configShallowCopy.UserAgent == "" {
		configShallowCopy.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	// share the transport between all clients
	httpClient, err := rest.HTTPClientFor(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(&configShallowCopy, httpClient)
}

func NewForConfigAndClient(c *rest.Config, client *http.Client) (*UnifyClientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		if configShallowCopy.Burst <= 0 {
			return nil, fmt.Errorf("burst is required to be greater than 0 when RateLimiter is not set and QPS is set to greater than 0")
		}
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}

	var cs UnifyClientset
	var err error
	cs.Clientset = fakekube.NewSimpleClientset()
	cs.coreV1, err = kosmosv1.NewForConfigAndClient(&configShallowCopy, client, cs.Clientset)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// CoreV1 retrieves the CoreV1Client
func (c *UnifyClientset) CoreV1() corev1.CoreV1Interface {
	return c.coreV1
}
