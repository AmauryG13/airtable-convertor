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
CURRENT_DIR=$(shell pwd)
PROJECT_DIR_LINK=$(shell readlink ${BUILD_DIR})

BUILD_DIR = ${PROJECT_DIR}/build
MAIN = ${PROJECT_DIR}/cmd/${BINARY}/main.go

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS = -ldflags "-X main.VERSION=${VERSION} -X main.COMMIT=${COMMIT} -X main.BRANCH=${BRANCH}"

# Build the project
all: link clean test linux darwin windows

link:
	PROJECT_DIR=${PROJECT_DIR}; \
	PROJECT_DIR_LINK=${PROJECT_DIR_LINK}; \
	CURRENT_DIR=${CURRENT_DIR}; \
	
	if [ "$${PROJECT_DIR_LINK}" != "$${CURRENT_DIR}" ]; then \
	    echo "Fixing symlinks for PROJECT dir"; \
	    rm -f $${PROJECT_DIR}; \
	    ln -s $${CURRENT_DIR} $${PROJECT_DIR}; \
	fi

	mkdir -p ${BUILD_DIR}

linux: 
	cd ${BUILD_DIR}; \
	GOOS=linux GOARCH=${GOARCH} ${GOBUILD} ${LDFLAGS} -o ${BUILD_DIR}/${BINARY}-linux-${GOARCH} ${MAIN} ; \
	cd - >/dev/null

darwin:
	cd ${BUILD_DIR}; \
	GOOS=darwin GOARCH=${GOARCH} ${GOBUILD} ${LDFLAGS} -o ${BUILD_DIR}/${BINARY}-darwin-${GOARCH} ${MAIN} ; \
	cd - >/dev/null

windows:
	cd ${BUILD_DIR}; \
	GOOS=windows GOARCH=${GOARCH} ${GOBUILD} ${LDFLAGS} -o ${BUILD_DIR}/${BINARY}-windows-${GOARCH}.exe ${MAIN} ; \
	cd - >/dev/null

test:
	${GOTEST} ./...

clean:
	${GOCLEAN} -r

.PHONY: link linux darwin windows test clean