#/bin/bash

cd $(dirname $0)

kubectl apply -f operator.yml
kubectl apply -f prometheus.yml
kubectl apply -f service.yml
