BINARY_NAME=balena-go
LDFLAGS=-w -s
ENVIRONMENT=CGO_ENABLED=1

build:
	${ENVIRONMENT} go build -o ${BINARY_NAME} -ldflags '${LDFLAGS}' ./

test:
	${ENVIRONMENT} go test -ldflags '${LDFLAGS}' ./...

dist:
	${ENVIRONMENT} go build -o dist/${BINARY_NAME} -ldflags '${LDFLAGS}' ./

run:
	${ENVIRONMENT} go build -o ${BINARY_NAME} -ldflags '${LDFLAGS}' ./
	./${BINARY_NAME}
