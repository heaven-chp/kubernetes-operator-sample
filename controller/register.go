package controller

import (
	"kubernetes-operator-sample/controller/fruit/v1alpha1"

	"k8s.io/apimachinery/pkg/watch"
)

func appendItems(items map[Interface]watch.Interface) {
	items[&v1alpha1.Apple{}] = nil
	items[&v1alpha1.Banana{}] = nil
}
