<img width="150" src="https://kt3k.github.io/saku/media/saku-logo.svg" />

# clean
> Cleans the workspace

    rm -rf dist
    rm -rf homebrew-tap
    rm -f coverage.txt

# install
> builds and installs it locally

    go install

# release
> Releases saku

    saku clean
    goreleaser
    git clone https://github.com/kt3k/homebrew-tap.git
    cd homebrew-tap; ./update_saku.sh 1.0.0 ; git commit -a -m "update saku (1.0.0)" ; git push origin head

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

# chglog
> Creates the change log

    git-chglog --next-tag v1.0.0 -o CHANGELOG.md
