ARCH := $(shell uname -m)
ifeq ($(ARCH),x86_64)
    PLATFORM := linux/amd64
else ifeq ($(ARCH),aarch64)
    PLATFORM := linux/arm64
else ifeq ($(ARCH),arm64)
    PLATFORM := linux/arm64
else
    $(warning Unknown architecture $(ARCH), defaulting to linux/amd64)
    PLATFORM := linux/amd64
endif
APP_NAME := go-http-echo
IMAGE := $(APP_NAME):latest
KIND_CLUSTER := kind0
DEPLOYMENT_FILES := k8s
TEST_LB_NAME := default-lb
TEST_LB_IP := 10.0.10.0
ROUTER_CONTAINER_NAME := clab-bgp-on-kind-router0
CURL_POD := $(shell yq '.metadata.name' k8s/curl.yaml)

.PHONY: build load deploy clean

## Build the container image
build:
	docker build -t $(IMAGE) --platform $(PLATFORM) .

## Load image into Kind cluster
load:
	kind load docker-image $(IMAGE) --name $(KIND_CLUSTER)

## Apply Kubernetes manifests
deploy: build load
	kubectl apply -f $(DEPLOYMENT_FILES)
	kubectl wait --for=condition=Ready -n default pod -l app=curl
	kubectl wait --for=condition=Ready -n default pod -l app=go-http-echo

## Remove deployment
clean:
	kubectl delete -f $(DEPLOYMENT_FILES) || true

test: deploy
	kubectl exec $(CURL_POD) -- curl -s $(TEST_LB_NAME)
	docker exec $(ROUTER_CONTAINER_NAME) curl -s $(TEST_LB_IP)
