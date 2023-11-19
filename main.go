package main

import (
	"kubernetes-operator-sample/controller"
	"os"
	"os/signal"
	"syscall"

	"k8s.io/klog/v2"
)

func main() {
	defer klog.Flush()

	klog.InfoS("main start")
	defer klog.InfoS("main end")

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	controller := controller.Controller{}
	if err := controller.Start(); err != nil {
		klog.ErrorS(err, "controller start error")
	}

	<-signals

	controller.Stop()
}
