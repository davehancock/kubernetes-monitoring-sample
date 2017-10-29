# kubernetes-monitoring-sample

A sample project demonstrating how to provision a simple service on Kubernetes, monitor metrics via Prometheus and visualize with Grafana.

# Prerequisites 

Install minikube - see instructions at: https://github.com/kubernetes/minikube

# TL;DR

Start Minikube
``` shell
minikube start 
```

Apply all manifests (sample-app, prometheus-operator, grafana)
``` shell
./manifests/apply.sh 
```

Navigate to Grafana local address (obtained via minikube cli)
``` shell
minikube service grafana --url
```

Login (foo:foo) to view the selection of infrastructural and business metric dashboards:

![Imgur](https://i.imgur.com/hPJCayh.png)


# Work In Progress / TODO

- Plot business metrics properly
- Define a smooth dashboard workflow. I.e create -> update -> version -> provision
- All RBAC
- Consistent labels / expression for service monitor selectors
- Proper namespacing
- Alert Manager / Prometheus Rules
- Extra Monitoring Pieces i.e kube-controller-manager, kube-scheduler, apiserver, etcd


# Detailed Summary

## Sample App

This project contains a simple Golang app for metric demonstration purposes. This will be deployed as a service with 3 replicated pods. The container image itself is publicly available so no need to build from source locally. 

If you do wish to build the sample app locally from the command line:

``` shell
env GOOS=linux GOARCH=386 go build && docker build -t test . && docker run -p 8081:8080 test
```

minikube service sample-app --url


## Prometheus / Prometheus Operator

Chosen to deploy prometheus using the operator to ease the complex configuration of maintaining a server manually see:

https://coreos.com/operators/prometheus/docs/latest/user-guides/getting-started.html


minikube service prometheus --url

http://192.168.99.100:31899


## Grafana

Modelled on example configuration found here:
https://github.com/coreos/prometheus-operator/tree/master/contrib/kube-prometheus/manifests/grafana

Adding the json for the dashboards by hand is pretty painful, discussions and tools below:
https://github.com/coreos/prometheus-operator/issues/272 
https://github.com/coreos/prometheus-operator/blob/master/contrib/kube-prometheus/docs/developing-alerts-and-dashboards.md
https://github.com/coreos/prometheus-operator/tree/master/contrib/kube-prometheus/hack/grafana-dashboards-configmap-generator 

minikube service grafana --url

http://192.168.99.100:31900


## Node Exporter / Kube State Metrics / Kubelet etc

Modelled on example configuration found here:
https://github.com/coreos/prometheus-operator/tree/master/contrib/kube-prometheus
