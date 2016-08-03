package main

import (
	"fmt"
	flag "github.com/spf13/pflag"
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/client/restclient"
	client "k8s.io/kubernetes/pkg/client/unversioned"
	kubectl_util "k8s.io/kubernetes/pkg/kubectl/cmd/util"
	"log"
	"os"
)

var (
	flags     = flag.NewFlagSet("", flag.ContinueOnError)
	inCluster = flags.Bool("use-kubernetes-cluster-service", false, `If true, use the built in kubernetes
        cluster for creating the client`)
	host = flags.String("host", "",
		`K8s host`)
	insecure = flags.Bool("insecure", false,
		`Enforce CA check for cert`)
)

func main() {
	clientConfig := kubectl_util.DefaultClientConfig(flags)

	flags.Parse(os.Args)

	config := &restclient.Config{
		Host:     *host,
		Insecure: *insecure,
	}
	var err error
	// if *inCluster {
	// 	kubeClient, err := client.NewInCluster()
	// } else {
	// 	config, connErr := clientConfig.ClientConfig()
	// 	if connErr != nil {
	// 		log.Fatalf("error connecting to the client: %v", err)
	// 	}
	// 	kubeClient, err := client.New(config)
	// }
	config, err = clientConfig.ClientConfig()

	kubeClient, err := client.New(config)
	podlist, err := kubeClient.Pods(api.NamespaceDefault).List(api.ListOptions{})
	if err != nil {
		log.Fatalln("Can't get pods:", err)
	}
	fmt.Printf("Pods: %s\n", podlist.Items)
	fmt.Printf("Connection: %s\n", kubeClient)
	fmt.Printf("%d\n", len(podlist.Items))
}
