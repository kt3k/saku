# clean
> Cleans the workspace

    rm -rf vendor

# build
> Builds the binary

    go build -o saku

# test
> Runs the tests

    go test -v .

# fix
> Format source code

    go fmt

# cov
> Takes the coverage data

> Foo
> Bar

    echo goverage -coverprofile=cover.out `go list ./... | grep -v /vendor/`

> Baz


    echo go tool cover -func=cover.out


    echo rm -rf cover.out

# changelog
> Creates the change log

    git tag v0.1.0
    git-chglog v0.1.0
    git tag -d v0.1.0
