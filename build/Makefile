GOPATH := $(shell go env GOPATH)
VERSION := 0.0.1.0
WORK_PATH := $(shell pwd)
BIN_PATH := $(subst /,\/,$(WORK_PATH))

depend:
	go get ../...

build: depend
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -installsuffix cgo -o ./login ../main.go

docker: start
	\cp -rf ../conf/conf.yaml ./helm_loginExample/conf

install:
   helm repo add stable https://charts.helm.sh/stable 
   helm install mysql stable/mysql
test:
	go test -v ../... -cover

docker: build
	docker build -f ./Dockerfile -t registry.cn-beijing.aliyuncs.com/julive/project:$(VERSION) ../
	rm -rf project

buildhealth:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o grpc-health-probe ../pkg/grpc-health-probe/main.go

clean:
	docker rmi registry.cn-beijing.aliyuncs.com/julive/project:$(VERSION)

deploy: docker
	kubectl apply -f ns-rbac.yaml
	kubectl apply -f project.yaml

text:
	kubectl apply -f project-dev.yaml

dev:build
	docker build -f ./DockerDevfile -t registry.cn-beijing.aliyuncs.com/julive/project:$(VERSION) ../
	kubectl apply -f ns-rbac.yaml
	kubectl apply -f project-dev.yaml
	kubectl get deploy project -o yaml -n micro | sed "s/sed_bin_hostpath_target/$(BIN_PATH)\/project/g" | kubectl replace -f -

reload: docker k8s

k8s:
	kubectl set image deployment/project project=registry.cn-beijing.aliyuncs.com/julive/project:$(VERSION) -n micro
	kubectl replace --force -f project.yaml

.PHONY:  depend build test docker clean deploy k8s
