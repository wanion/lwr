version := $(shell date +"%y.%m.%d")
.DEFAULT_GOAL := build

clean:
	go clean
.PHONY:clean

fmt:
	go fmt ./...
.PHONY:fmt

lint: fmt
	golint ./...
.PHONY:lint

vet: fmt
	go vet ./...
.PHONY:vet

rpmdev:
	rpmdev-setuptree
.PHONY:rpmdev

archive: rpmdev
	git archive --format=tar.gz --prefix=lwr-$(version)/ -o ~/rpmbuild/SOURCES/lwr-$(version).tar.gz HEAD
.PHONY:archive

rpm: archive
	rpmbuild -ba lwr.spec
	ls -1 ~/rpmbuild/RPMS/x86_64/
.PHONY:rpm

build: vet
	CGO_ENABLED=0 go build -ldflags "-X 'main.appVersion=$(version)'"
.PHONY:build
