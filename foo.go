package main

import (
  client "k8s.io/kubernetes/pkg/client/unversioned"
  "k8s.io/kubernetes/pkg/api"
)

func main() {
  config := &client.Config{
    Host:     "http://localhost:8080",
    Username: "test",
    Password: "password",
  }
  client, err := client.New(config)
  if err != nil {
    // handle error
  }
  pods, err := client.Pods(api.NamespaceDefault).List(api.ListOptions{})
  if err != nil {
    // handle error
  }
}
