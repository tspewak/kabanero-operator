IMAGE ?= kabanero/kabanero-operator:latest

.PHONY: build deploy

build: 
	go install ./cmd/manager

build-image: dependencies
	operator-sdk build ${IMAGE}

generate:
	operator-sdk generate k8s

install:
	kubectl apply -f deploy/crds/kabanero_v1alpha1_kabanero_crd.yaml
	
deploy: 
	kubectl create namespace kabanero || true
	kubectl -n kabanero apply -f deploy/

dependencies:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure

# Requires https://github.com/pmezard/licenses
dependency-report: dependencies
	go get -u github.com/pmezard/licenses
	licenses ./pkg/... | cut -c49- > 3RD_PARTY