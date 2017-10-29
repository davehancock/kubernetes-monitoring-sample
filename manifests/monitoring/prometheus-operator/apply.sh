#/bin/bash

cd $(dirname $0)

kubectl apply -f ./rbac.yml
kubectl apply -f ./deployment.yml
kubectl apply -f ./service.yml

until kubectl get servicemonitor; do sleep 1; done
kubectl apply -f ./service-monitor.yml
