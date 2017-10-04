FROM scratch

COPY ./kubernetes-monitoring-sample /go/bin/kubernetes-monitoring-sample

EXPOSE 8080

ENTRYPOINT ["/go/bin/kubernetes-monitoring-sample"]
