package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	// "github.com/golang/glog"
	"github.com/spf13/pflag"

	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/client/unversioned"
	kubectl_util "k8s.io/kubernetes/pkg/kubectl/cmd/util"
)

var (
	flags = pflag.NewFlagSet("", pflag.ExitOnError)
)

func main() {
	flags.AddGoFlagSet(flag.CommandLine)
	flags.Parse(os.Args)
	clientConfig := kubectl_util.DefaultClientConfig(flags)

	config, err := clientConfig.ClientConfig()
	if err != nil {
		log.Fatalf("error connecting to the client: %v", err)
	}

	kubeClient, err := unversioned.New(config)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	for {
		podlist, err := kubeClient.Pods(api.NamespaceDefault).List(api.ListOptions{})
		if err != nil {
			log.Fatalf("failed list the pods: %v", err)
		}

		for _, pod := range podlist.Items {
			currentTime := time.Now()
			startTime := pod.Status.StartTime
			ttl, err := strconv.ParseFloat(pod.Labels["ttl"], 64)
			if err != nil {
				// log.Printf("failed parse label ttl: %v", err)
				continue
			}

			podAge := currentTime.Sub(startTime.Time)
			podAgeHours := podAge.Hours()
			if podAgeHours <= ttl {
				log.Println("Pod", pod.Name, "is younger then", ttl, "hours")
				continue
			}

			fmt.Println("Attempting to kill pod", pod.Name, "it is older then", ttl, "hours")
			err = kubeClient.Pods(pod.Namespace).Delete(pod.Name, &api.DeleteOptions{})
			if err != nil {
				log.Printf("Pod %v was deleted\n", pod)
			}
		}
	}
}
