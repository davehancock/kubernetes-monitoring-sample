#/bin/bash

cd $(dirname $0)

kubectl apply -f secret.yml
kubectl apply -f dashboards.yml
kubectl apply -f deployment.yml
kubectl apply -f service.yml
