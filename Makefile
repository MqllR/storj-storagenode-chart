GOOS=linux
GOBUILD=go build
BINARY=identity-to-kube-secret

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
