package main

import (
	// "encoding/json"
	"fmt"
	flag "github.com/spf13/pflag"
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/client/restclient"
	client "k8s.io/kubernetes/pkg/client/unversioned"
	kubectl_util "k8s.io/kubernetes/pkg/kubectl/cmd/util"
	// "log"
	// "net/http"
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

type itemdata [][]string

func getPods() {
	clientConfig := kubectl_util.DefaultClientConfig(flags)

	// flags.Parse(os.Args)

	config := &restclient.Config{
		Host:     *host,
		Insecure: *insecure,
	}

	var err error
	// var datas itemdata

	config, err = clientConfig.ClientConfig()
	check(err)
	kubeClient, err := client.New(config)
	check(err)
	podlist, err := kubeClient.Pods(api.NamespaceDefault).List(api.ListOptions{})
	check(err)
	for _, pod := range podlist.Items {
		fmt.Println(pod.Name)
		fmt.Println(pod.Status)
	}
	// fmt.Fprintf("Pods: %s\n", podlist.Items)
	// fmt.Fprintf(w, "Connection: %s\n", html.EscapeString(r.kubeClient))
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	// http.HandleFunc("/", getPods) // set route to get pods
	// err := http.ListenAndServe(":9090", nil)
	// check(err)
	getPods()
}
