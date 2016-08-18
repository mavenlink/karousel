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

func deletePod() {
	// Get pods to delete
	podlist, err := KubeClient.Pods(api.NamespaceDefault).List(api.ListOptions{})
	if err != nil {
		log.Fatalf("(2)failed list pods: %v", err)
	}

	for _, pod := range podlist.Items {
		startTime := pod.Status.StartTime
		ttl, err := strconv.ParseFloat(pod.Labels["ttl"], 64)
		if err != nil {
			// log.Printf("failed parse label ttl: %v", err)
			continue
		}

		podAge := CurrentTime.Sub(startTime.Time)
		podAgeHours := podAge.Hours()
		if podAgeHours <= ttl {
			log.Println("Pod", pod.Name, "is younger then", ttl, "hours")
			continue
		}

		// Delete Pod
		fmt.Println("Attempting to kill pod", pod.Name, "it is older then", ttl, "hours")
		err = KubeClient.Pods(pod.Namespace).Delete(pod.Name, &api.DeleteOptions{})
		if err != nil {
			log.Printf("Pod %v was deleted\n", pod)
		}
	}
}

func deleteService() {
	// Get service to delete
	servicelist, err := KubeClient.Services(api.NamespaceDefault).List(api.ListOptions{})
	if err != nil {
		log.Fatalf("(2)failed list service: %v", err)
	}

	for _, service := range servicelist.Items {
		ttl, err := strconv.ParseFloat(service.Labels["ttl"], 64)
		if err != nil {
			// log.Printf("failed parse label ttl: %v", err)
			continue
		}
		serviceCreation := service.GetCreationTimestamp().Time
		serviceAge := CurrentTime.Sub(serviceCreation)
		// serviceAge := CurrentTime.Sub(service.GetCreationTimestamp)
		if serviceAge.Hours() <= ttl {
			log.Println("Service", service.Name, "is younger then", ttl, "hours")
			continue
		}

		// Delete Service
		fmt.Println("Attempting to kill service", service.Name, "it is older then", ttl, "hours")
		err = KubeClient.Services(service.Namespace).Delete(service.Name)
		if err != nil {
			log.Printf("Service %v was deleted\n", service.Name)
		}
	}
}

func main() {
	fmt.Println("Instastage Started...Please stand by for ascension")
	flags.AddGoFlagSet(flag.CommandLine)
	flags.Parse(os.Args)
	clientConfig := kubectl_util.DefaultClientConfig(flags)

	config, err := clientConfig.ClientConfig()
	if err != nil {
		log.Fatalf("error connecting to the client: %v", err)
	}

	KubeClient, err := unversioned.New(config)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	CurrentTime := time.Now()

	for {
		deleteService()
		deletePod()

		// // Get ingress to delete
		// ingresslist, err := KubeClient.Ingress(api.NamespaceDefault).List(api.ListOptions{})
		// if err != nil {
		// 	log.Fatalf("(2)failed list ingress: %v", err)
		// }

		// for _, ingress := range ingresslist.Items {
		// 	ttl, err := strconv.ParseFloat(ingress.Labels["ttl"], 64)
		// 	if err != nil {
		// 		// log.Printf("failed parse label ttl: %v", err)
		// 		continue
		// 	}

		// 	ingressAge := CurrentTime.Sub(ingress.Status.StartTime)
		// 	if ingressAge.Hours() <= ttl {
		// 		log.Println("Ingress", ingress.Name, "is younger then", ttl, "hours")
		// 		continue
		// 	}

		// 	// Delete Ingress
		// 	fmt.Println("Attempting to kill ingress", ingress.Name, "it is older then", ttl, "hours")
		// 	err = KubeClient.Ingress(ingress.Namespace).Delete(ingress.Name, &api.DeleteOptions{})
		// 	if err != nil {
		// 		log.Printf("Ingress %v was deleted\n", ingress.Name)
		// 	}
		// }

	}
}
