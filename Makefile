VERSION=0.0.1
TARGETS_NOVENDOR=$(shell glide novendor)

build: cmd/redis-cluster/main.go cluster/*.go config/*.go
	go build -ldflags "-X main.version=${VERSION}" -o bin/redis-cluster cmd/redis-cluster/main.go

clean:
	rm -rf bin/*

bundle:
	glide install

test:
	go test -v $(TARGETS_NOVENDOR)

build-all: build-mac build-linux

build-mac: cmd/redis-cluster/main.go cluster/*.go config/*.go
	GOOS=darwin GOARCH=amd64 go build -ldflags "-X main.version=${VERSION}" -o bin/darwin/amd64/redis-cluster-${VERSION}/redis-cluster cmd/redis-cluster/main.go

build-linux: cmd/redis-cluster/main.go cluster/*.go config/*.go
	GOOS=linux GOARCH=amd64 go build -ldflags "-X main.version=${VERSION}" -o bin/linux/amd64/redis-cluster-${VERSION}/redis-cluster cmd/redis-cluster/main.go

dist: build-all
	cd bin/darwin/amd64 && tar zcvf redis-cluster-darwin-amd64-${VERSION}.tar.gz redis-cluster-${VERSION}
	cd bin/linux/amd64 && tar zcvf redis-cluster-linux-amd64-${VERSION}.tar.gz redis-cluster-${VERSION}
