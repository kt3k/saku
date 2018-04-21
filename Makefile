.PHONY: clean build test fmt cov chglog

clean:
	rm -rf vendor

build:
	go build -o saku

test:
	go test -v .

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
