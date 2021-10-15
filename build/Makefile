GOPATH := $(shell go env GOPATH)
VERSION := v0.0.2
WORK_PATH := $(shell pwd)
BIN_PATH := $(subst /,\/,$(WORK_PATH))

depend:
	go get ../...

build: depend
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -installsuffix cgo -o ./login ../main.go

install:
	# \cp -rf ../conf/conf.yaml ./helm_loginExample/conf //configmap模式可以使用
	helm install loginexample ./helm_loginExample
test:
	go test -v ../... -cover

docker: build
	docker build -f ./Dockerfile -t loginexample:$(VERSION) ../
	rm -rf login
clean:
	docker rmi loginexample:$(VERSION)
installtraefik:
	helm repo add traefik https://helm.traefik.io/traefik
	helm repo update
	helm install traefik traefik/traefik
.PHONY:  depend build install test docker clean installtraefik
