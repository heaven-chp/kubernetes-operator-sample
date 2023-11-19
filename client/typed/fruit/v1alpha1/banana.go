package v1alpha1

import (
	"context"
	v1alpha1 "kubernetes-operator-sample/api/fruit/v1alpha1"
	"kubernetes-operator-sample/client/scheme"
	"time"

	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/rest"
)

type BananasGetter interface {
	Bananas(namespace string) BananaInterface
}

type BananaInterface interface {
	Create(ctx context.Context, banana *v1alpha1.Banana, opts metaV1.CreateOptions) (*v1alpha1.Banana, error)
	Update(ctx context.Context, banana *v1alpha1.Banana, opts metaV1.UpdateOptions) (*v1alpha1.Banana, error)
	UpdateStatus(ctx context.Context, banana *v1alpha1.Banana, opts metaV1.UpdateOptions) (*v1alpha1.Banana, error)
	Delete(ctx context.Context, name string, opts metaV1.DeleteOptions) error
	Get(ctx context.Context, name string, opts metaV1.GetOptions) (*v1alpha1.Banana, error)
	List(ctx context.Context, opts metaV1.ListOptions) (*v1alpha1.BananaList, error)
	Watch(ctx context.Context, opts metaV1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metaV1.PatchOptions, subresources ...string) (result *v1alpha1.Banana, err error)
}

type bananas struct {
	client    rest.Interface
	namespace string
}

func newBananas(client *FruitV1alpha1Client, namespace string) *bananas {
	return &bananas{
		client:    client.RESTClient(),
		namespace: namespace,
	}
}

func (this *bananas) Get(ctx context.Context, name string, options metaV1.GetOptions) (result *v1alpha1.Banana, err error) {
	result = &v1alpha1.Banana{}
	err = this.client.Get().
		Namespace(this.namespace).
		Resource("bananas").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

func (this *bananas) List(ctx context.Context, opts metaV1.ListOptions) (result *v1alpha1.BananaList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.BananaList{}
	err = this.client.Get().
		Namespace(this.namespace).
		Resource("bananas").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

func (this *bananas) Watch(ctx context.Context, opts metaV1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return this.client.Get().
		Namespace(this.namespace).
		Resource("bananas").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

func (this *bananas) Create(ctx context.Context, banana *v1alpha1.Banana, opts metaV1.CreateOptions) (result *v1alpha1.Banana, err error) {
	result = &v1alpha1.Banana{}
	err = this.client.Post().
		Namespace(this.namespace).
		Resource("bananas").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(banana).
		Do(ctx).
		Into(result)
	return
}

func (this *bananas) Update(ctx context.Context, banana *v1alpha1.Banana, opts metaV1.UpdateOptions) (result *v1alpha1.Banana, err error) {
	result = &v1alpha1.Banana{}
	err = this.client.Put().
		Namespace(this.namespace).
		Resource("bananas").
		Name(banana.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(banana).
		Do(ctx).
		Into(result)
	return
}

func (this *bananas) UpdateStatus(ctx context.Context, banana *v1alpha1.Banana, opts metaV1.UpdateOptions) (result *v1alpha1.Banana, err error) {
	result = &v1alpha1.Banana{}
	err = this.client.Put().
		Namespace(this.namespace).
		Resource("bananas").
		Name(banana.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(banana).
		Do(ctx).
		Into(result)
	return
}

func (this *bananas) Delete(ctx context.Context, name string, opts metaV1.DeleteOptions) error {
	return this.client.Delete().
		Namespace(this.namespace).
		Resource("bananas").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

func (this *bananas) DeleteCollection(ctx context.Context, opts metaV1.DeleteOptions, listOpts metaV1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return this.client.Delete().
		Namespace(this.namespace).
		Resource("bananas").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

func (this *bananas) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metaV1.PatchOptions, subresources ...string) (result *v1alpha1.Banana, err error) {
	result = &v1alpha1.Banana{}
	err = this.client.Patch(pt).
		Namespace(this.namespace).
		Resource("bananas").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
