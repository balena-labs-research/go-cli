BINARY_NAME=go-cli
LDFLAGS=-w -s
ENVIRONMENT=CGO_ENABLED=1

test:
	${ENVIRONMENT} go test -ldflags '${LDFLAGS}'  ./...
 
build:
	${ENVIRONMENT} go build -ldflags '${LDFLAGS}' ./cmd/${BINARY_NAME}/

dist:
	${ENVIRONMENT} go build -o dist/ -ldflags '${LDFLAGS}' ./cmd/${BINARY_NAME}/

run:
	${ENVIRONMENT} go build -ldflags '${LDFLAGS}' ./cmd/${BINARY_NAME}/
	./${BINARY_NAME}
