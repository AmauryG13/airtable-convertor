# Golang executable
GOCMD=go
GOBUILD=${GOCMD} build
GOTEST=${GOCMD} test
GOFMT=${GOCMD} fmt
GOCLEAN=${GOCMD} clean

# Project parameters 
BINARY = airtable-convertor
GOARCH = amd64

VERSION=$(shell git describe --tags)
COMMIT=$(shell git rev-parse HEAD)
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)

# Symlink into GOPATH
GITHUB_USERNAME=amauryg13
PROJECT_DIR=${GOPATH}/src/github.com/${GITHUB_USERNAME}/${BINARY}

BUILD_DIR = ${PROJECT_DIR}/build
MAIN = ${PROJECT_DIR}/cmd/${BINARY}/main.go

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS = -ldflags "-X main.Version=${VERSION} -X main.Commit=${COMMIT} -X main.Branch=${BRANCH}"

# Build the project
all: link clean test linux darwin windows

link:
	mkdir -p ${BUILD_DIR}

linux: 
	cd ${BUILD_DIR} \
	GOOS=linux GOARCH=${GOARCH} ${GOBUILD} ${LDFLAGS} -o ${BUILD_DIR}/${BINARY}-linux-${GOARCH} ${MAIN}

darwin:
	cd ${BUILD_DIR} \
	GOOS=darwin GOARCH=${GOARCH} ${GOBUILD} ${LDFLAGS} -o ${BUILD_DIR}/${BINARY}-darwin-${GOARCH} ${MAIN}

windows:
	cd ${BUILD_DIR} \
	GOOS=windows GOARCH=${GOARCH} ${GOBUILD} ${LDFLAGS} -o ${BUILD_DIR}/${BINARY}-windows-${GOARCH}.exe ${MAIN}
test:
	${GOTEST} ./...

clean:
	${GOCLEAN} -r

.PHONY: link linux darwin windows test clean