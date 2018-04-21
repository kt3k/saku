.PHONY: clean build test fmt cov chglog

test:
	go build -o saku
	./saku
	# go test -v .

clean:
	rm -rf vendor
	rm -rf $(GOPATH)/bin/saku

build:
	go build -o saku

fmt:
	go fmt

cov:
	echo goverage -coverprofile=cover.out `go list ./... | grep -v /vendor/`
	echo go tool cover -func=cover.out
	echo rm -rf cover.out

chglog:
	git tag v0.1.0
	git-chglog v0.1.0
	git tag -d v0.1.0
