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

CILIUM_VERSION ?= 1.16.12

.PHONY: build-bird 
build-bird:
	docker build -t bird:latest --platform $(PLATFORM) bird

.PHONY: deploy
deploy: build-bird cilium-manifests
	sudo containerlab -t topology.yaml deploy
	kind export kubeconfig --name kind0
	kubectl taint nodes kind0-control-plane node-role.kubernetes.io/control-plane:NoSchedule-
	# install cilium
	kubectl apply -f cilium/upstream.yaml
	cilium status --wait --namespace kube-system
	# install bgpcp
	kubectl apply -f cilium/bgpcp.yaml
	# deploy test setup
	$(MAKE) -C ../test/ deploy

.PHONY: destroy
destroy:
	sudo containerlab -t topology.yaml destroy -c
	rm -f .topo.yaml.bak

.PHONY: cilium-manifests
cilium-manifests: cilium/upstream.yaml
cilium/upstream.yaml: cilium/values.yaml
	rm -rf /tmp/work-cilium
	mkdir -p /tmp/work-cilium
	git clone --depth 1 --branch v$(shell echo $(CILIUM_VERSION) | cut -d \. -f 1,2,3) https://github.com/cilium/cilium /tmp/work-cilium
	helm template /tmp/work-cilium/install/kubernetes/cilium/ \
		--namespace=kube-system \
		--values cilium/values.yaml > cilium/upstream.yaml
	rm -rf /tmp/work-cilium

.PHONY: reload
reload:
	$(MAKE) destroy
	$(MAKE) deploy

.PHONY: test
test:
	$(MAKE) -C ../test/ test-l4lb
