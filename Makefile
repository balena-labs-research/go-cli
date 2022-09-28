BINARY_NAME=go-cli
LDFLAGS=-w -s
ENVIRONMENT=CGO_ENABLED=1

build:
	${ENVIRONMENT} go build -ldflags '${LDFLAGS}' ./

test:
	${ENVIRONMENT} go test -ldflags '${LDFLAGS}'  ./...

dist:
	${ENVIRONMENT} go build -o dist/ -ldflags '${LDFLAGS}' ./

run:
	${ENVIRONMENT} go build -ldflags '${LDFLAGS}' ./
	./${BINARY_NAME}
