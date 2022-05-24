GOOS=linux
GOBUILD=go build
BINARY=identity-to-kube-secret
CHART_NAME=storj-storagenode
CHART_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))/charts/$(CHART_NAME)

.PHONY: build
build: build-arm build-amd64

.PHONY: build-arm
build-arm:
	GOARCH=arm ${GOBUILD} -o ${BINARY}-arm script/${BINARY}.go

.PHONY: build-amd64
build-amd64:
	GOARCH=amd64 ${GOBUILD} -o ${BINARY}-amd64 script/${BINARY}.go

.PHONY: cleanup
cleanup:
	rm -f ${BINARY}-*
	rm -f ${CHART_NAME}-*

.PHONE: helm_lint
helm_lint:
	helm lint ${CHART_DIR} --strict -f linter_values.yaml
