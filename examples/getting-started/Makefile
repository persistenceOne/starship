NAME = starship-getting-started
FILE = configs/starship.yaml
TINY_FILE = configs/tiny-starship.yaml

HELM_REPO = starship
HELM_CHART = devnet
HELM_VERSION = v0.1.36

###############################################################################
###                              All commands                               ###
###############################################################################

.PHONY: setup
setup: setup-deps setup-helm setup-kind

.PHONY: start
start: install port-forward

.PHONY: start-tiny
start-tiny: install-tiny port-forward

.PHONY: stop
stop: stop-forward delete

.PHONY: clean
clean: stop clean-kind

###############################################################################
###                          Dependency check                               ###
###############################################################################

.PHONY: check
setup-deps:
	bash $(CURDIR)/scripts/dev-setup.sh

###############################################################################
###                              Helm Charts                                ###
###############################################################################

setup-helm:
	helm repo add $(HELM_REPO) https://cosmology-tech.github.io/starship/
	helm repo update
	helm search repo $(HELM_REPO)/$(HELM_CHART) --version $(HELM_VERSION)

install:
	helm install -f $(FILE) $(NAME) $(HELM_REPO)/$(HELM_CHART) --version $(HELM_VERSION) --wait --debug --timeout 20m

install-tiny:
	$(MAKE) install FILE=$(TINY_FILE)

delete:
	-helm delete $(NAME)

###############################################################################
###                                 Port forward                            ###
###############################################################################

.PHONY: port-forward
port-forward:
	bash $(CURDIR)/scripts/port-forward.sh --config=$(FILE)

.PHONY: stop-forward
stop-forward:
	-pkill -f "port-forward"

###############################################################################
###                          Local Kind Setup                               ###
###############################################################################
KIND_CLUSTER=starship

.PHONY: setup-kind
setup-kind:
	kind create cluster --name $(KIND_CLUSTER)

.PHONY: clean-kind
clean-kind:
	kind delete cluster --name $(KIND_CLUSTER)
