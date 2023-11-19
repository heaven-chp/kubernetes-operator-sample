package v1alpha1

import (
	"context"
	fruitV1alpha1 "kubernetes-operator-sample/api/fruit/v1alpha1"
	"kubernetes-operator-sample/client"

	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/klog/v2"
)

type Apple struct {
}

func (this *Apple) WatchInterface(clientset *client.Clientset) (watch.Interface, error) {
	return clientset.FruitV1alpha1().Apples("").Watch(context.TODO(), metaV1.ListOptions{})
}

func (this *Apple) Job(event watch.Event) {
	switch event.Object.(type) {
	case *fruitV1alpha1.Apple:
		apple := event.Object.(*fruitV1alpha1.Apple)

		klog.InfoS("event", "name", apple.Name, "type", event.Type)
	}
}
