# Docker

## Build locally

Statically compile app with all libraries built in.

Disabling cgo produces a static binary. Set the OS to Linux. The -a flag means to rebuild 
all the packages which means all the imports will be rebuilt with cgo disabled. 

Creates a static binary.

```console
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
```

## Build Docker image
```console
docker build -t erraticmotion/kubia:alpine .
docker image prune --filter label=stage=builder -f
docker run --name kubia-container -p 8080:8080 -d erraticmotion/kubia:alpine

curl localhost:8080
docker exec -it kubia-container bash
:/# ps aux
exit

docker stop kubia-container
docker rm kubia-container
docker push erraticmotion/kubia:alpine
```