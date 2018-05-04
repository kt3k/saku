# clean
> Cleans the workspace

    rm -rf $GOPATH/bin/saku
    rm -f coverage.txt
    rm saku

# build
> Builds the binary

    go build -o saku

# test
> Runs the tests

    go test -v .

# fmt
> Format source code

    go fmt

# cov
> Creates the coverage data

    go test -race -coverprofile=coverage.txt -covermode=atomic .

# cov-html
> Creates the html coverage report

    saku cov
    go tool cover -html=coverage.txt

# changelog
> Creates the change log

    git tag v0.1.0
    git-chglog v0.1.0
    git tag -d v0.1.0

# hello
> Say hello
> To the world

    echo hello
    echo world

# foo
> This command fails

    go foo

# hello
> Say hello

    echo hello

# build-and-test
> Builds and tests

    saku build test

# p9090
> Runs a static server

    static -p 9090
