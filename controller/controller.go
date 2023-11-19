package controller

import (
	"kubernetes-operator-sample/client"

	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/rest"
	"k8s.io/klog/v2"
)

type Interface interface {
	WatchInterface(clientset *client.Clientset) (watch.Interface, error)

	Job(event watch.Event)
}

type Controller struct {
	items map[Interface]watch.Interface
}

func (this *Controller) Start() error {
	if err := this.Stop(); err != nil {
		return err
	}

	appendItems(this.items)

	for resource, _ := range this.items {
		if config, err := rest.InClusterConfig(); err != nil {
			return err
		} else if clientset, err := client.NewForConfig(config); err != nil {
			return err
		} else if watchInterface, err := resource.WatchInterface(clientset); err != nil {
			klog.ErrorS(err, "resource Watch error")
			continue
		} else {
			this.items[resource] = watchInterface
		}

		go func(resource Interface, watchInterface watch.Interface) {
			for event := range watchInterface.ResultChan() {
				resource.Job(event)
			}
		}(resource, this.items[resource])
	}

	return nil
}

func (this *Controller) Stop() error {
	for _, watchInterface := range this.items {
		if watchInterface != nil {
			watchInterface.Stop()
		}
	}

	this.items = make(map[Interface]watch.Interface)

	return nil
}
