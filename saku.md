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

    cd pkg/saku; go test -v .

# fmt
> Format source code

    go fmt

# cov
> Creates the coverage data

    cd pkg/saku; go test -coverprofile=coverage.txt -covermode=atomic .

# cov-html
> Creates the html coverage report

    saku cov
    go tool cover -html=pkg/saku/coverage.txt

# changelog
> Creates the change log

    git tag v0.1.0
    git-chglog v0.1.0
    git tag -d v0.1.0
