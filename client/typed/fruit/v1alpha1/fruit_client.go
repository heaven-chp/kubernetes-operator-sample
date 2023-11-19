package v1alpha1

import (
	v1alpha1 "kubernetes-operator-sample/api/fruit/v1alpha1"
	"kubernetes-operator-sample/client/scheme"
	"net/http"

	"k8s.io/client-go/rest"
)

type FruitV1alpha1Interface interface {
	RESTClient() rest.Interface
	ApplesGetter
	BananasGetter
}

type FruitV1alpha1Client struct {
	restClient rest.Interface
}

func (this *FruitV1alpha1Client) Apples(namespace string) AppleInterface {
	return newApples(this, namespace)
}

func (this *FruitV1alpha1Client) Bananas(namespace string) BananaInterface {
	return newBananas(this, namespace)
}

func NewForConfig(c *rest.Config) (*FruitV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	httpClient, err := rest.HTTPClientFor(&config)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(&config, httpClient)
}

func NewForConfigAndClient(c *rest.Config, h *http.Client) (*FruitV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientForConfigAndClient(&config, h)
	if err != nil {
		return nil, err
	}
	return &FruitV1alpha1Client{client}, nil
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

func (this *FruitV1alpha1Client) RESTClient() rest.Interface {
	if this == nil {
		return nil
	}
	return this.restClient
}
