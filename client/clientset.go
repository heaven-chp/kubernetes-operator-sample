package client

import (
	"fmt"
	fruitV1alpha1 "kubernetes-operator-sample/client/typed/fruit/v1alpha1"
	"net/http"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/util/flowcontrol"
)

type Clientset struct {
	fruitV1alpha1 *fruitV1alpha1.FruitV1alpha1Client
}

func (this *Clientset) FruitV1alpha1() fruitV1alpha1.FruitV1alpha1Interface {
	return this.fruitV1alpha1
}

func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c

	if configShallowCopy.UserAgent == "" {
		configShallowCopy.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	httpClient, err := rest.HTTPClientFor(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	return NewForConfigAndClient(&configShallowCopy, httpClient)
}

func NewForConfigAndClient(c *rest.Config, httpClient *http.Client) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		if configShallowCopy.Burst <= 0 {
			return nil, fmt.Errorf("burst is required to be greater than 0 when RateLimiter is not set and QPS is set to greater than 0")
		}
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}

	var cs Clientset
	var err error
	if cs.fruitV1alpha1, err = fruitV1alpha1.NewForConfigAndClient(&configShallowCopy, httpClient); err != nil {
		return nil, err
	}

	return &cs, nil
}
