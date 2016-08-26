package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/spf13/pflag"

	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/client/unversioned"
	kubectl_util "k8s.io/kubernetes/pkg/kubectl/cmd/util"
)

var (
	flags   = pflag.NewFlagSet("", pflag.ExitOnError)
	Version = "No Version Provided"
)

func deleteDeployment(kubeClient *unversioned.Client) {
	// Get deployments to delete
	deploymentlist, err := kubeClient.Deployments(api.NamespaceDefault).List(api.ListOptions{})
	if err != nil {
		log.Fatalf("(2)failed list deployments: %v", err)
	}

	for _, deployment := range deploymentlist.Items {
		startTime := deployment.GetCreationTimestamp()
		ttl, err := strconv.ParseFloat(deployment.Labels["ttl"], 64)
		if err != nil {
			continue
		}

		deploymentAge := time.Now().Sub(startTime.Time)
		deploymentAgeHours := deploymentAge.Hours()
		if deploymentAgeHours <= ttl {
			log.Println("Deployment", deployment.Name, "is", int(deploymentAge.Hours()), "hours old, it will be deleted in", ttl, "hours")
			continue
		}

		// Delete Deployment
		fmt.Println("Attempting to kill deployment", deployment.Name, "it is older then", ttl, "hours")
		err = kubeClient.Deployments(deployment.Namespace).Delete(deployment.Name, &api.DeleteOptions{})
		if err != nil {
			log.Printf("Deployment %v was deleted\n", deployment.Name)
		}
	}
}

// func deletePod(kubeClient *unversioned.Client) {
// 	// Get pods to delete
// 	podlist, err := kubeClient.Pods(api.NamespaceDefault).List(api.ListOptions{})
// 	if err != nil {
// 		log.Fatalf("(2)failed list pods: %v", err)
// 	}

// 	for _, pod := range podlist.Items {
// 		startTime := pod.Status.StartTime
// 		ttl, err := strconv.ParseFloat(pod.Labels["ttl"], 64)
// 		if err != nil {
// 			continue
// 		}

// 		podAge := time.Now().Sub(startTime.Time)
// 		podAgeHours := podAge.Hours()
// 		if podAgeHours <= ttl {
// 			log.Println("Pod", pod.Name, "is", int(podAge.Hours()), "hours old, it will be deleted in", ttl, "hours")
// 			continue
// 		}

// 		// Delete Pod
// 		fmt.Println("Attempting to kill pod", pod.Name, "it is older then", ttl, "hours")
// 		err = kubeClient.Pods(pod.Namespace).Delete(pod.Name, &api.DeleteOptions{})
// 		if err != nil {
// 			log.Printf("Pod %v was deleted\n", pod)
// 		}
// 	}
// }

func deleteService(kubeClient *unversioned.Client) {
	// Get service to delete
	servicelist, err := kubeClient.Services(api.NamespaceDefault).List(api.ListOptions{})
	if err != nil {
		log.Fatalf("(2)failed list service: %v", err)
	}

	for _, service := range servicelist.Items {
		ttl, err := strconv.ParseFloat(service.Labels["ttl"], 64)
		if err != nil {
			continue
		}
		serviceCreation := service.GetCreationTimestamp().Time
		serviceAge := time.Now().Sub(serviceCreation)
		if serviceAge.Hours() <= ttl {
			log.Println("Service", service.Name, "is", int(serviceAge.Hours()), "hours old, it will be deleted in", ttl, "hours")
			continue
		}

		// Delete Service
		fmt.Println("Attempting to kill service", service.Name, "it is older then", ttl, "hours")
		err = kubeClient.Services(service.Namespace).Delete(service.Name)
		if err != nil {
			log.Printf("Service %v was deleted\n", service.Name)
		}
	}
}

func deleteIngress(kubeClient *unversioned.Client) {
	// Get ingress to delete
	ingresslist, err := kubeClient.Ingress(api.NamespaceDefault).List(api.ListOptions{})
	if err != nil {
		log.Fatalf("(2)failed list ingress: %v", err)
	}

	for _, ingress := range ingresslist.Items {
		ttl, err := strconv.ParseFloat(ingress.Labels["ttl"], 64)
		if err != nil {
			continue
		}
		ingressCreation := ingress.GetCreationTimestamp().Time
		ingressAge := time.Now().Sub(ingressCreation)
		if ingressAge.Hours() <= ttl {
			log.Println("Ingress", ingress.Name, "is", int(ingressAge.Hours()), "hours old, it will be deleted in", ttl, "hours")
			continue
		}

		// Delete Ingress
		fmt.Println("Attempting to kill ingress", ingress.Name, "it is older then", ttl, "hours")
		err = kubeClient.Ingress(ingress.Namespace).Delete(ingress.Name, &api.DeleteOptions{})
		if err != nil {
			log.Printf("Ingress %v was deleted\n", ingress.Name)
		}
	}
}

func main() {
	fmt.Printf("Karousel(%s) Started...Please stand by for ascension\n", Version)
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
		deleteDeployment(kubeClient)
		// deletePod(kubeClient)
		deleteService(kubeClient)
		deleteIngress(kubeClient)
		time.Sleep(300 * time.Second)
	}
}
