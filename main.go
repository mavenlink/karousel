package main

import (
	"fmt"
	"log"
	// "os"
	"strconv"
	"time"

	"k8s.io/client-go/1.4/kubernetes"
	"k8s.io/client-go/1.4/pkg/api"
	"k8s.io/client-go/1.4/rest"
)

var (
	Version = "No Version Provided"
)

func deletePods(kubeClient *kubernetes.Clientset) {
	list, err := kubeClient.Core().Pods("").List(api.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	for _, resource := range list.Items {
		startTime := resource.GetCreationTimestamp()
		ttl, err := strconv.ParseFloat(resource.Labels["ttl"], 64)
		if err != nil {
			continue
		}

		resourceAge := time.Now().Sub(startTime.Time)
		resourceAgeHours := resourceAge.Hours()
		if resourceAgeHours <= ttl {
			log.Println(resource.Name, "is", int(resourceAge.Hours()), "hours old, it will be deleted in", ttl, "hours")
			continue
		}

		// Delete Resource
		fmt.Println("Attempting to kill", resource.Name, "it is older then", ttl, "hours")
		err = kubeClient.Pods(resource.Namespace).Delete(resource.Name, &api.DeleteOptions{})
		if err != nil {
			log.Printf("%v was deleted\n", resource.Name)
		}
	}
}

func deleteServices(kubeClient *kubernetes.Clientset) {
	list, err := kubeClient.Core().Services("").List(api.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	for _, resource := range list.Items {
		startTime := resource.GetCreationTimestamp()
		ttl, err := strconv.ParseFloat(resource.Labels["ttl"], 64)
		if err != nil {
			continue
		}

		resourceAge := time.Now().Sub(startTime.Time)
		resourceAgeHours := resourceAge.Hours()
		if resourceAgeHours <= ttl {
			log.Println(resource.Name, "is", int(resourceAge.Hours()), "hours old, it will be deleted in", ttl, "hours")
			continue
		}

		// Delete Resource
		fmt.Println("Attempting to kill", resource.Name, "it is older then", ttl, "hours")
		err = kubeClient.Services(resource.Namespace).Delete(resource.Name, &api.DeleteOptions{})
		if err != nil {
			log.Printf("%v was deleted\n", resource.Name)
		}
	}
}
func deleteDeployments(kubeClient *kubernetes.Clientset) {
	list, err := kubeClient.Deployments("").List(api.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	for _, resource := range list.Items {
		startTime := resource.GetCreationTimestamp()
		ttl, err := strconv.ParseFloat(resource.Labels["ttl"], 64)
		if err != nil {
			continue
		}

		resourceAge := time.Now().Sub(startTime.Time)
		resourceAgeHours := resourceAge.Hours()
		if resourceAgeHours <= ttl {
			log.Println(resource.Name, "is", int(resourceAge.Hours()), "hours old, it will be deleted in", ttl, "hours")
			continue
		}

		// Delete Resource
		fmt.Println("Attempting to kill", resource.Name, "it is older then", ttl, "hours")
		err = kubeClient.Deployments(resource.Namespace).Delete(resource.Name, &api.DeleteOptions{})
		if err != nil {
			log.Printf("%v was deleted\n", resource.Name)
		}
	}
}

func deleteIngresses(kubeClient *kubernetes.Clientset) {
	list, err := kubeClient.Deployments("").List(api.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	for _, resource := range list.Items {
		startTime := resource.GetCreationTimestamp()
		ttl, err := strconv.ParseFloat(resource.Labels["ttl"], 64)
		if err != nil {
			continue
		}

		resourceAge := time.Now().Sub(startTime.Time)
		resourceAgeHours := resourceAge.Hours()
		if resourceAgeHours <= ttl {
			log.Println(resource.Name, "is", int(resourceAge.Hours()), "hours old, it will be deleted in", ttl, "hours")
			continue
		}

		// Delete Resource
		fmt.Println("Attempting to kill", resource.Name, "it is older then", ttl, "hours")
		err = kubeClient.Ingresses(resource.Namespace).Delete(resource.Name, &api.DeleteOptions{})
		if err != nil {
			log.Printf("%v was deleted\n", resource.Name)
		}
	}
}

func main() {
	fmt.Printf("Karousel(%s) Started...Please stand by for ascension\n", Version)
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	kubeClient, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	fmt.Printf("Value: %v Error: %v", kubeClient, err)
	for {
		deleteDeployments(kubeClient)
		deletePods(kubeClient)
		deleteServices(kubeClient)
		deleteIngresses(kubeClient)
	}
}
