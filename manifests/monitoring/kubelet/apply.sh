#/bin/bash

cd $(dirname $0)

until kubectl get servicemonitor; do sleep 1; done
kubectl apply -f service-monitor.yml
