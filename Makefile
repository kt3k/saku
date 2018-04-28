.PHONY: clean build test fmt cov chglog

integration-test:
	go build -o saku
	./saku

test:
	go test -race -v .

clean:
	rm -rf vendor
	rm -rf $(GOPATH)/bin/saku

build:
	go build -o saku

fmt:
	go fmt

cov:
	go test -race -coverprofile=coverage.txt -covermode=atomic .

chglog:
	git tag v0.1.0
	git-chglog v0.1.0
	git tag -d v0.1.0
