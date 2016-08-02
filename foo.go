package main

import (
  "fmt"
  client "k8s.io/kubernetes/pkg/client/unversioned"
  restclient "k8s.io/kubernetes/pkg/client/restclient"
  "k8s.io/kubernetes/pkg/api"
)

func main() {
  config := &restclient.Config{
    Host:     "https://192.168.99.112:8443",
  }
  kubeclient, err := client.New(config)
  if err != nil {
    fmt.Printf("%s\n",err)
  }
  podlist, err := kubeclient.Pods(api.NamespaceDefault).List(api.ListOptions{})
  if err != nil {
    fmt.Printf("%s\n",err)
  }
  fmt.Printf("%s\n",podlist.Items)
  fmt.Printf("%d\n",len(podlist.Items))
}
