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

type ApplesGetter interface {
	Apples(namespace string) AppleInterface
}

type AppleInterface interface {
	Create(ctx context.Context, apple *v1alpha1.Apple, opts metaV1.CreateOptions) (*v1alpha1.Apple, error)
	Update(ctx context.Context, apple *v1alpha1.Apple, opts metaV1.UpdateOptions) (*v1alpha1.Apple, error)
	UpdateStatus(ctx context.Context, apple *v1alpha1.Apple, opts metaV1.UpdateOptions) (*v1alpha1.Apple, error)
	Delete(ctx context.Context, name string, opts metaV1.DeleteOptions) error
	Get(ctx context.Context, name string, opts metaV1.GetOptions) (*v1alpha1.Apple, error)
	List(ctx context.Context, opts metaV1.ListOptions) (*v1alpha1.AppleList, error)
	Watch(ctx context.Context, opts metaV1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metaV1.PatchOptions, subresources ...string) (result *v1alpha1.Apple, err error)
}

type apples struct {
	client    rest.Interface
	namespace string
}

func newApples(client *FruitV1alpha1Client, namespace string) *apples {
	return &apples{
		client:    client.RESTClient(),
		namespace: namespace,
	}
}

func (this *apples) Get(ctx context.Context, name string, options metaV1.GetOptions) (result *v1alpha1.Apple, err error) {
	result = &v1alpha1.Apple{}
	err = this.client.Get().
		Namespace(this.namespace).
		Resource("apples").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

func (this *apples) List(ctx context.Context, opts metaV1.ListOptions) (result *v1alpha1.AppleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AppleList{}
	err = this.client.Get().
		Namespace(this.namespace).
		Resource("apples").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

func (this *apples) Watch(ctx context.Context, opts metaV1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return this.client.Get().
		Namespace(this.namespace).
		Resource("apples").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

func (this *apples) Create(ctx context.Context, apple *v1alpha1.Apple, opts metaV1.CreateOptions) (result *v1alpha1.Apple, err error) {
	result = &v1alpha1.Apple{}
	err = this.client.Post().
		Namespace(this.namespace).
		Resource("apples").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(apple).
		Do(ctx).
		Into(result)
	return
}

func (this *apples) Update(ctx context.Context, apple *v1alpha1.Apple, opts metaV1.UpdateOptions) (result *v1alpha1.Apple, err error) {
	result = &v1alpha1.Apple{}
	err = this.client.Put().
		Namespace(this.namespace).
		Resource("apples").
		Name(apple.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(apple).
		Do(ctx).
		Into(result)
	return
}

func (this *apples) UpdateStatus(ctx context.Context, apple *v1alpha1.Apple, opts metaV1.UpdateOptions) (result *v1alpha1.Apple, err error) {
	result = &v1alpha1.Apple{}
	err = this.client.Put().
		Namespace(this.namespace).
		Resource("apples").
		Name(apple.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(apple).
		Do(ctx).
		Into(result)
	return
}

func (this *apples) Delete(ctx context.Context, name string, opts metaV1.DeleteOptions) error {
	return this.client.Delete().
		Namespace(this.namespace).
		Resource("apples").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

func (this *apples) DeleteCollection(ctx context.Context, opts metaV1.DeleteOptions, listOpts metaV1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return this.client.Delete().
		Namespace(this.namespace).
		Resource("apples").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

func (this *apples) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metaV1.PatchOptions, subresources ...string) (result *v1alpha1.Apple, err error) {
	result = &v1alpha1.Apple{}
	err = this.client.Patch(pt).
		Namespace(this.namespace).
		Resource("apples").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
