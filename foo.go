package main

import (
	"fmt"
	"k8s.io/kubernetes/pkg/api"
	restclient "k8s.io/kubernetes/pkg/client/restclient"
	client "k8s.io/kubernetes/pkg/client/unversioned"
	"log"
)

func main() {
	config := &restclient.Config{
		Host:     "http://192.168.99.100:8081",
		Username: "admin",
		Password: "merge4justice",
		Insecure: true,
	}
	kubeclient, err := client.New(config)
	if err != nil {
		log.Fatalln("Can't connect to Kubernetes API:", err)
	}
	podlist, err := kubeclient.Pods(api.NamespaceDefault).List(api.ListOptions{})
	if err != nil {
		log.Fatalln("Can't get pods:", err)
	}
	fmt.Printf("Pods: %s\n", podlist.Items)
	// fmt.Printf("Connection: %s\n", kubeclient)
	// fmt.Printf("%d\n", len(podlist.Items))
}
