MAC = GOOS=darwin GOARCH=amd64
LINUX = GOOS=linux GOARCH=amd64
PACKAGE = github.com/alphagov/stscreds
COMMAND = ${PACKAGE}/cmd/stscreds
SOURCES = $(shell find ${GOPATH}/src/${PACKAGE} -name \*.go)
FLAGS = -ldflags "-X main.versionNumber=${VERSION}"
VERSION ?= DEVELOPMENT

RELEASE_TARBALL = ${GOPATH}/release/stscreds-${VERSION}.tar.gz

default: build

${RELEASE_TARBALL}: ${GOPATH}/release/mac/stscreds ${GOPATH}/release/linux/stscreds
	mkdir -p ${GOPATH}/release/
	tar -zcf ${RELEASE_TARBALL} -C ${GOPATH}/release/ mac/stscreds linux/stscreds

${GOPATH}/release/mac/stscreds: ${SOURCES}
	${MAC} go build ${FLAGS} -o ${GOPATH}/release/mac/stscreds ${COMMAND}
	chmod +x ${GOPATH}/release/mac/stscreds

${GOPATH}/release/linux/stscreds: ${SOURCES}
	${LINUX} go build ${FLAGS} -o ${GOPATH}/release/linux/stscreds ${COMMAND}
	chmod +x ${GOPATH}/release/linux/stscreds

${GOPATH}/bin/stscreds: ${SOURCES}
	go build ${FLAGS} -o ${GOPATH}/bin/stscreds ${COMMAND}
	chmod +x ${GOPATH}/bin/stscreds

release: ${RELEASE_TARBALL}

build: ${GOPATH}/bin/stscreds

clean:
	rm -f ${GOPATH}/release/linux/stscreds
	rm -f ${GOPATH}/release/mac/stscreds
	rm -f ${GOPATH}/bin/stscreds
	rm -f ${RELEASE_TARBALL}
