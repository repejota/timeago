PACKAGES = $(shell go list ./...)
VERSION=`cat VERSION`
BUILD=`git symbolic-ref HEAD 2> /dev/null | cut -b 12-`-`git log --pretty=format:%h -1`
DIST_INCLUDE_FILES=README.md ROADMAP.md LICENSE VERSION

# Setup the -ldflags option for go build here, interpolate the variable
# values
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"

# Build & Install

install:		## Build and install package on your system
	go install $(LDFLAGS) -v $(PACKAGES)

.PHONY: version
version:		## Show version information
	@echo $(VERSION)-$(BUILD)

# Testing

.PHONY: test
test:			## Execute package tests 
	go test -v $(PACKAGES)

.PHONY: cover-profile
cover-profile:
	echo "mode: count" > coverage-all.out
	$(foreach pkg,$(PACKAGES),\
		go test -coverprofile=coverage.out -covermode=count $(pkg);\
		tail -n +2 coverage.out >> coverage-all.out;)
	rm -rf coverage.out

.PHONY: cover
cover: cover-profile	
cover: 			## Generate test coverage data
	go tool cover -func=coverage-all.out

.PHONY: cover-html
cover-html: cover-profile
cover-html:		## Generate coverage report
	go tool cover -html=coverage-all.out

.PHONY: coveralls
coveralls:
	goveralls -service circle-ci -repotoken iB20py5PBrpiugL8JnXlxFacGvk8UbBTU

# Lint

lint:			## Lint source code
	gometalinter --disable-all --enable=errcheck --enable=vet --enable=vetshadow

# Dependencies

deps:			## Install package dependencies

	
dev-deps:		## Install dev dependencies
	go get -u github.com/smartystreets/goconvey
	go get -u github.com/mattn/goveralls
	go get -u github.com/inconshreveable/mousetrap
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install

.PHONY: clean
clean:			## Delete generated development environment
	go clean
	rm -rf coverage-all.out

# Docs

godoc-serve:		## Serve documentation (godoc format) for this package at port HTTP 9090
	godoc -http=":9090"

include Makefile.help.mk
