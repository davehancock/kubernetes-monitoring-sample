#/bin/bash

cd $(dirname $0)

./monitoring/prometheus-operator/apply.sh
./monitoring/prometheus/apply.sh
./monitoring/kubelet/apply.sh
./monitoring/kube-state-metrics/apply.sh
./monitoring/node-exporter/apply.sh
./monitoring/grafana/apply.sh

./services/sample/apply.sh
