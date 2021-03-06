/*
Copyright 2016 The Kubernetes Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Note: the example only works with the code within the same release/branch.
package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var (
	Version = "No Version Provided"
)

func deletePod(clientset *kubernetes.Clientset) {
	// Get pods to delete
	podlist, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		log.Println("(2)failed list pods: %v", err)
	}

	for _, pod := range podlist.Items {
		dt, err := strconv.ParseFloat(pod.Annotations["deployTime"], 64)
		if err != nil {
			log.Printf("(2)failed to parse deployTime for pod %v: %v", pod.Name, err)
			continue
		}
		var i int64 = int64(dt)
		startTime := time.Unix(i, 64)
		ttl, err := strconv.ParseFloat(pod.Annotations["ttl"], 64)
		if err != nil {
			log.Printf("(2)failed to parse ttl for pod %v:  %v", pod.Name, err)
			continue
		}

		podAge := time.Now().Sub(startTime)
		podAgeHours := podAge.Hours()
		if podAgeHours <= ttl {
			continue
		}

		// Delete Pod
		fmt.Println("Attempting to kill pod", pod.Name, "it is older then", ttl, "hours")

		err = clientset.CoreV1().Pods(pod.Namespace).Delete(pod.Name, &metav1.DeleteOptions{})

		if err != nil {
			log.Printf("Pod %v was deleted\n", pod.Name)
		}
	}
}

func deleteService(clientset *kubernetes.Clientset) {
	// Get service to delete
	servicelist, err := clientset.CoreV1().Services("").List(metav1.ListOptions{})
	if err != nil {
		log.Fatalf("(2)failed list services: %v", err)
	}

	for _, service := range servicelist.Items {
		dt, err := strconv.ParseFloat(service.Annotations["deployTime"], 64)
		if err != nil {
			log.Printf("(2)failed to parse deployTime for service %v: %v", service.Name, err)
			continue
		}
		var i int64 = int64(dt)
		startTime := time.Unix(i, 64)
		ttl, err := strconv.ParseFloat(service.Annotations["ttl"], 64)
		if err != nil {
			log.Printf("(2)failed to parse ttl for service %v: %v", service.Name, err)
			continue
		}

		serviceAge := time.Now().Sub(startTime)
		serviceAgeHours := serviceAge.Hours()
		if serviceAgeHours <= ttl {
			continue
		}

		// Delete service
		fmt.Println("Attempting to kill service", service.Name, "it is older then", ttl, "hours")

		err = clientset.CoreV1().Services(service.Namespace).Delete(service.Name, &metav1.DeleteOptions{})

		if err != nil {
			log.Printf("Service %v was deleted\n", service.Name)
		}
	}
}

func deleteReplicaSet(clientset *kubernetes.Clientset) {
	// Get replicaset to delete
	replicasetlist, err := clientset.ExtensionsV1beta1().ReplicaSets("").List(metav1.ListOptions{})
	if err != nil {
		log.Fatalf("(2)failed list replicasets: %v", err)
	}

	for _, replicaset := range replicasetlist.Items {
		dt, err := strconv.ParseFloat(replicaset.Annotations["deployTime"], 64)
		if err != nil {
			log.Printf("(2)failed to parse deployTime for service %v: %v", replicaset.Name, err)
			continue
		}
		var i int64 = int64(dt)
		startTime := time.Unix(i, 64)
		ttl, err := strconv.ParseFloat(replicaset.Annotations["ttl"], 64)
		if err != nil {
			log.Printf("(2)failed to parse ttl for service %v: %v", replicaset.Name, err)
			continue
		}

		replicasetAge := time.Now().Sub(startTime)
		replicasetAgeHours := replicasetAge.Hours()
		if replicasetAgeHours <= ttl {
			continue
		}

		// Delete replicaset
		fmt.Println("Attempting to kill replicaset", replicaset.Name, "it is older then", ttl, "hours")

		err = clientset.ExtensionsV1beta1().ReplicaSets(replicaset.Namespace).Delete(replicaset.Name, &metav1.DeleteOptions{})

		if err != nil {
			log.Printf("Replciaset %v was deleted\n", replicaset.Name)
		}
	}
}

func deleteIngress(clientset *kubernetes.Clientset) {
	// Get ingress to delete
	ingresslist, err := clientset.ExtensionsV1beta1().Ingresses("").List(metav1.ListOptions{})
	if err != nil {
		log.Fatalf("(2)failed list ingresses: %v", err)
	}

	for _, ingress := range ingresslist.Items {
		dt, err := strconv.ParseFloat(ingress.Annotations["deployTime"], 64)
		if err != nil {
			log.Printf("(2)failed to parse deployTime for ingress %v: %v", ingress.Name, err)
			continue
		}
		var i int64 = int64(dt)
		startTime := time.Unix(i, 64)
		ttl, err := strconv.ParseFloat(ingress.Annotations["ttl"], 64)
		if err != nil {
			log.Printf("(2)failed to parse ttl for ingress %v: %v", ingress.Name, err)
			continue
		}

		ingressAge := time.Now().Sub(startTime)
		ingressAgeHours := ingressAge.Hours()
		if ingressAgeHours <= ttl {
			continue
		}

		// Delete Ingress
		fmt.Println("Attempting to kill ingress", ingress.Name, "it is older then", ttl, "hours")

		err = clientset.ExtensionsV1beta1().Ingresses(ingress.Namespace).Delete(ingress.Name, &metav1.DeleteOptions{})

		if err != nil {
			log.Printf("Ingress %v was deleted\n", ingress.Name)
		}
	}
}

func deleteDeployment(clientset *kubernetes.Clientset) {
	// Get deployments to delete
	deploymentlist, err := clientset.ExtensionsV1beta1().Deployments("").List(metav1.ListOptions{})
	if err != nil {
		log.Fatalf("(2)failed list deployments: %v", err)
	}

	for _, deployment := range deploymentlist.Items {
		dt, err := strconv.ParseFloat(deployment.Annotations["deployTime"], 64)
		if err != nil {
			log.Printf("(2)failed to parse deployTime for deployment %v: %v", deployment.Name, err)
			continue
		}
		var i int64 = int64(dt)
		startTime := time.Unix(i, 64)
		ttl, err := strconv.ParseFloat(deployment.Annotations["ttl"], 64)
		if err != nil {
			log.Printf("(2)failed to parse ttl for deployment %v: %v", deployment.Name, err)
			continue
		}

		deploymentAge := time.Now().Sub(startTime)
		deploymentAgeHours := deploymentAge.Hours()
		if deploymentAgeHours <= ttl {
			continue
		}

		// Delete Deployment
		fmt.Println("Attempting to kill deployment", deployment.Name, "it is older then", ttl, "hours")

		// err = kubeClient.Deployments(deployment.Namespace).Delete(deployment.Name, &api.DeleteOptions{})
		err = clientset.ExtensionsV1beta1().Deployments(deployment.Namespace).Delete(deployment.Name, &metav1.DeleteOptions{})

		if err != nil {
			log.Printf("Deployment %v was deleted\n", deployment.Name)
		}
	}
}

func main() {
	fmt.Printf("Karousel(%s) Started...Please stand by for ascension\n", Version)
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	for {
		deleteDeployment(clientset)
		deleteReplicaSet(clientset)
		deletePod(clientset)
		deleteService(clientset)
		deleteIngress(clientset)
		time.Sleep(300 * time.Second)
	}
}
