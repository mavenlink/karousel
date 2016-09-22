![bespoked](logo.png)

# karousel

A utility to delete "expired" resources from a kubernetes cluster.

# Usage

Karousel uses a meta data field "ttl" (Time to Live) to determine the correct time to delete a resource.  See examples directory for sample yaml. 

# Local Testing

Make sure to remove the pflags dir from your $GOPATH/src/k8s/...

We recommend running minikube https://github.com/kubernetes/minikube

