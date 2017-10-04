# kubernetes-monitoring-sample

To build the app locally from the command line:

```shell
env GOOS=linux GOARCH=386 go build && docker build -t test . && docker run -p 8081:8080 test
```
