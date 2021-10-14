GOPATH := $(shell go env GOPATH)
VERSION := v0.0.1
WORK_PATH := $(shell pwd)
BIN_PATH := $(subst /,\/,$(WORK_PATH))

depend:
	go get ../...

build: depend
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -installsuffix cgo -o ./login ../main.go

install:
	\cp -rf ../conf/conf.yaml ./helm_loginExample/conf
	# helm repo add stable https://charts.helm.sh/stable
	# helm install mysql stable/mysql
	helm install loginexample ./helm_loginExample
test:
	go test -v ../... -cover

docker: build
	docker build -f ./Dockerfile -t loginexample:$(VERSION) ../
	rm -rf login
clean:
	docker rmi loginexample:$(VERSION)

deploy: docker
	kubectl apply -f ns-rbac.yaml
	kubectl apply -f project.yaml
k8s:
	kubectl set image deployment/project project=registry.cn-beijing.aliyuncs.com/julive/project:$(VERSION) -n micro
	kubectl replace --force -f project.yaml

.PHONY:  depend build test docker clean deploy k8s
