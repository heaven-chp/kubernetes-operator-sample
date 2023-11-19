package v1alpha1

import (
	"context"
	fruitV1alpha1 "kubernetes-operator-sample/api/fruit/v1alpha1"
	"kubernetes-operator-sample/client"

	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/klog/v2"
)

type Banana struct {
}

func (this *Banana) WatchInterface(clientset *client.Clientset) (watch.Interface, error) {
	return clientset.FruitV1alpha1().Bananas("").Watch(context.TODO(), metaV1.ListOptions{})
}

func (this *Banana) Job(event watch.Event) {
	switch event.Object.(type) {
	case *fruitV1alpha1.Banana:
		banana := event.Object.(*fruitV1alpha1.Banana)

		klog.InfoS("event", "name", banana.Name, "type", event.Type)
	}
}
