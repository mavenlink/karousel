package main

import (
	"flag"
	"fmt"
	"github.com/spf13/pflag"
	"k8s.io/kubernetes/pkg/api"
	kubeClient "k8s.io/kubernetes/pkg/client/restclient"
	client "k8s.io/kubernetes/pkg/client/unversioned"
	kubectl_util "k8s.io/kubernetes/pkg/kubectl/cmd/util"
	"log"
)

var (
	flags = pflag.NewFlagSet("", pflag.ExitOnError)
)

func main() {
	// config := &restclient.Config{
	// 	Host:     "https://192.168.99.100:8443",
	// 	Username: "admin",
	// 	Password: "merge4justice",
	// 	Insecure: true,
	// }
	clientConfig := kubectl_util.DefaultClientConfig()

	var err error
	if *inCluster {
		kubeClient, err = client.NewInCluster()
	} else {
		config, connErr := clientConfig.ClientConfig()
		if connErr != nil {
			glog.Fatalf("error connecting to the client: %v", err)
		}
		kubeClient, err = client.New(config)
	}

	podlist, err := kubeClient.Pods(api.NamespaceDefault).List(api.ListOptions{})
	if err != nil {
		log.Fatalln("Can't get pods:", err)
	}
	fmt.Printf("Pods: %s\n", podlist.Items)
	// fmt.Printf("Connection: %s\n", kubeclient)
	// fmt.Printf("%d\n", len(podlist.Items))
}
