GOOS=linux
GOBUILD=go build
BINARY=identity-to-kube-secret
CHART_NAME=storj-storagenode
ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

.PHONY: go_build
build: go_build-arm go_build-amd64

.PHONY: go_build-arm
build-arm:
	GOARCH=arm ${GOBUILD} -o ${BINARY}-arm script/${BINARY}.go

.PHONY: go_build-amd64
build-amd64:
	GOARCH=amd64 ${GOBUILD} -o ${BINARY}-amd64 script/${BINARY}.go

.PHONY: cleanup
cleanup:
	rm -f ${BINARY}-*
	rm -f ${CHART_NAME}-*

.PHONE: helm_lint
helm_lint:
	helm lint ${ROOT_DIR} --strict -f linter_values.yaml

.PHONE: helm_package
helm_package:
	helm package ${ROOT_DIR}
