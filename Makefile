#!make

TEST?=$$(go list ./... | grep -v 'vendor')
HOSTNAME=supply.com
NAMESPACE=devops
NAME=k8s_client_go
BINARY=${NAME}
VERSION=1.4
OS_ARCH=darwin_amd64

go.mod:
	go mod init k8s_client_go

go.sum:
	go get k8s.io/client-go@v0.19.0

init: go.mod go.sum

build:
	go build -o ${BINARY}

install: build
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

release:
	GOOS=darwin GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_darwin_amd64
	GOOS=freebsd GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_freebsd_386
	GOOS=freebsd GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_freebsd_amd64
	GOOS=freebsd GOARCH=arm go build -o ./bin/${BINARY}_${VERSION}_freebsd_arm
	GOOS=linux GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_linux_386
	GOOS=linux GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_linux_amd64
	GOOS=linux GOARCH=arm go build -o ./bin/${BINARY}_${VERSION}_linux_arm
	GOOS=openbsd GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_openbsd_386
	GOOS=openbsd GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_openbsd_amd64
	GOOS=solaris GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_solaris_amd64
	GOOS=windows GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_windows_386
	GOOS=windows GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_windows_amd64

test:
	go test -i $(TEST) || exit 1
	echo $(TEST) | xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4

testacc:
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m

clean:
	go clean

az-login:
	az login

connect-dev:
	az account set --subscription="b54f955a-67c4-4680-888e-17bf609fe9c2"
	az aks get-credentials --name "feidsupplyaks001" --resource-group "feideu2-supplyaks-rg-001" --overwrite-existing
	kubectl config use-context "feidsupplyaks001"
	kubectl get namespaces

connect-sandbox:
	az account set --subscription="8d01f77a-4a6f-4548-b5f3-743769a1b178"
	az aks get-credentials --name "tfprovtest" --resource-group "feiseu2-supply-rg-001" --overwrite-existing
	kubectl config use-context "tfprovtest"
	kubectl get namespaces
