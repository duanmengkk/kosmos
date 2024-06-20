/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

@CHANGELOG
KubeEdge Authors: To make a bridge between kubeclient and metaclient,
This file is derived from K8S client-go code with reduced set of methods
Changes done are
1. Package corev1 got some functions from "k8s.io/client-go/kubernetes/typed/core/corev1/fake/fake_core_client.go"
and made some variant
*/

package v1

import (
	v1 "k8s.io/api/core/v1"
	fakekube "k8s.io/client-go/kubernetes/fake"
	"k8s.io/client-go/kubernetes/scheme"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	fakecorev1 "k8s.io/client-go/kubernetes/typed/core/v1/fake"
	"k8s.io/client-go/rest"
	"net/http"
)

// UnifyCoreV1 is a coreV1 bridge
type UnifyCoreV1 struct {
	fakecorev1.FakeCoreV1
	RestClient rest.Interface
}

// NewForConfigAndClient creates a new CoreV1Client for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client, fake *fakekube.Clientset) (*UnifyCoreV1, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientForConfigAndClient(&config, h)
	if err != nil {
		return nil, err
	}
	return &UnifyCoreV1{
		FakeCoreV1: fakecorev1.FakeCoreV1{Fake: &fake.Fake},
		RestClient: client,
	}, nil
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/api"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

func (c *UnifyCoreV1) Pods(namespace string) corev1.PodInterface {
	return &UnifyPods{fakecorev1.FakePods{Fake: &c.FakeCoreV1}, c.RestClient, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *UnifyCoreV1) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.RestClient
}
