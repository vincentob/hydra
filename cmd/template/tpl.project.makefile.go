package template

func init() {
	CommonProjectFiles["Makefile"] = `# go build output binary filename
NAME={{ .ProjectName }}
BINARY={{ .ProjectName }}
GO111MODULE=on
CGO=0

# check for runtime os
ifeq ($(shell uname),Darwin)
	GOOS := darwin
else
	GOOS := linux
endif

# build binary for current os
build:
	@echo "Build ${NAME} for ${GOOS}:"
	GO111MODULE=${GO111MODULE} GOPROXY=${GOPROXY} CGO_ENABLED=${CGO} GOOS=${GOOS} go build -a -o ${BINARY} cmd/*.go
	@echo "Done"

# build docker image
build_docker:
	@echo "Build ${NAME} docker image:"
	docker build -t ${BINARY} .
	docker tag ${BINARY} ${BINARY}:latest
	@echo "Done"

# build binary for all os, include linux darwin windows and docker.
release: release_linux release_darwin release_windows release_docker
release_linux: prepare_linux build dist
release_darwin: prepare_darwin build dist
release_windows: prepare_windows build dist
release_docker docker: build_docker

# specify os for building.
linux: prepare_linux build
darwin: prepare_darwin build
windows: prepare_windows build

prepare_linux:
	$(eval GOOS := linux)
	$(eval BINARY := ${NAME})
	@rm -f ${BINARY}

prepare_darwin:
	$(eval GOOS := darwin)
	$(eval BINARY := ${NAME})
	@rm -f ${BINARY}

prepare_windows:
	$(eval GOOS := windows)
	$(eval BINARY := ${NAME}.exe)
	@rm -f ${BINARY}


dist:
	@mkdir -p dist
	@rm -rf dist/*
	@mv $(BINARY) dist/
	@echo "The building output is: dist/${BINARY}"

.PHONY: build* release* windows linux darwin docker dist help

`
}
