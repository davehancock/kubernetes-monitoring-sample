#/bin/bash

cd $(dirname $0)

./monitoring/prometheus/apply.sh
./monitoring/grafana/apply.sh
./services/sample/apply.sh
