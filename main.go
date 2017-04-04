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

func deletePod(clientset *kubernetes.Clientset) {
	// Get pods to delete
	podlist, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		log.Fatalf("(2)failed list pods: %v", err)
	}

	for _, pod := range podlist.Items {
		startTime := pod.GetCreationTimestamp()
		if startTime.IsZero() {
			continue
		}
		ttl, err := strconv.ParseFloat(pod.Labels["ttl"], 64)
		if err != nil {
			continue
		}

		podAge := time.Now().Sub(startTime.Time)
		podAgeHours := podAge.Hours()
		if podAgeHours <= ttl {
			log.Println("Pod", pod.Name, "is", int(podAge.Hours()), "hours old, it will be deleted in", ttl, "hours")
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
		startTime := service.GetCreationTimestamp()
		if startTime.IsZero() {
			continue
		}
		ttl, err := strconv.ParseFloat(service.Labels["ttl"], 64)
		if err != nil {
			continue
		}

		serviceAge := time.Now().Sub(startTime.Time)
		serviceAgeHours := serviceAge.Hours()
		if serviceAgeHours <= ttl {
			log.Println("Service", service.Name, "is", int(serviceAge.Hours()), "hours old, it will be deleted in", ttl, "hours")
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
		startTime := replicaset.GetCreationTimestamp()
		if startTime.IsZero() {
			continue
		}
		ttl, err := strconv.ParseFloat(replicaset.Labels["ttl"], 64)
		if err != nil {
			continue
		}

		replicasetAge := time.Now().Sub(startTime.Time)
		replicasetAgeHours := replicasetAge.Hours()
		if replicasetAgeHours <= ttl {
			log.Println("Replicaset", replicaset.Name, "is", int(replicasetAge.Hours()), "hours old, it will be deleted in", ttl, "hours")
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
		startTime := ingress.GetCreationTimestamp()
		if startTime.IsZero() {
			continue
		}
		ttl, err := strconv.ParseFloat(ingress.Labels["ttl"], 64)
		if err != nil {
			continue
		}

		ingressAge := time.Now().Sub(startTime.Time)
		ingressAgeHours := ingressAge.Hours()
		if ingressAgeHours <= ttl {
			log.Println("Ingress", ingress.Name, "is", int(ingressAge.Hours()), "hours old, it will be deleted in", ttl, "hours")
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
		startTime := deployment.GetCreationTimestamp()
		if startTime.IsZero() {
			continue
		}
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

		// err = kubeClient.Deployments(deployment.Namespace).Delete(deployment.Name, &api.DeleteOptions{})
		err = clientset.ExtensionsV1beta1().Deployments(deployment.Namespace).Delete(deployment.Name, &metav1.DeleteOptions{})

		if err != nil {
			log.Printf("Deployment %v was deleted\n", deployment.Name)
		}
	}
}

func main() {
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
