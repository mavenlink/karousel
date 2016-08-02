package main

import (
    "fmt"
    "log"
    client "k8s.io/kubernetes/pkg/client/unversioned"
    restclient "k8s.io/kubernetes/pkg/client/restclient"
    "k8s.io/kubernetes/pkg/api"
)

func main() {

    config := restclient.Config {
        Host: "https://192.168.99.100:8443",
    }
    c, err := client.New(&config)
    if err != nil {
        log.Fatalln("Can't connect to Kubernetes API:", err)
    }

    s, err := c.Services(api.NamespaceDefault).Get("some-service-name")
    if err != nil {
        log.Fatalln("Can't get service:", err)
    }
    fmt.Println("Name:", s.Name)
    for p, _ := range s.Spec.Ports {
        fmt.Println("Port:", s.Spec.Ports[p].Port)
        fmt.Println("NodePort:", s.Spec.Ports[p].NodePort)
    }
  }
