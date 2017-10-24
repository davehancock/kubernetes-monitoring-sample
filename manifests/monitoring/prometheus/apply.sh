#/bin/bash

cd $(dirname $0)

kubectl apply -f operator.yml
kubectl apply -f service.yml

# Wait until the TPR (third party resource) for Prometheus is registered before applying dependent manifests
until kubectl get prometheus; do sleep 1; done
kubectl apply -f prometheus.yml
