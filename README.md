<img width="150" src="https://kt3k.github.io/saku/media/saku-logo.svg" />

# saku (作) v0.1.4 WIP

[![CircleCI](https://circleci.com/gh/kt3k/saku.svg?style=svg)](https://circleci.com/gh/kt3k/saku)
[![Build status](https://ci.appveyor.com/api/projects/status/56lo7khd1an83lh3?svg=true)](https://ci.appveyor.com/project/kt3k/saku-elt31)
[![codecov](https://codecov.io/gh/kt3k/saku/branch/master/graph/badge.svg)](https://codecov.io/gh/kt3k/saku)

> Markdown-based task runner

`saku` is a simple task runner based on markdown syntax. You can define and describe your tasks in markdown file `saku.md` and execute them with `saku` command.

(You can optionally define tasks in `README.md`. See [The below](#embed-sakumd-in-readmemd) for details.)

# :cd: Install

**Go users**

    go get -u github.com/kt3k/saku

**Mac OS users**

    brew install kt3k/tap/saku

Or download binary from the [release page][].

# :leaves: Usage

First, create a markdown file `saku.md`:

```md
# build
> Build the go binary.

    go build -v -i main.go

# test
> Run all the go tests.

    go test -race ./...

# js

    minify -o public/script.js src/js

# css

    minify -o public/style.css src/css
```

The above defines 4 tasks `build` `test` `js` `css`. (A heading (#) is a task title!)

If you hit `saku` (without arguments) it shows the list of the descriptions of the all tasks.

If you hit the command `saku build`, it invokes `build` task, `go build -v -i main.go` in the above example.

**Note**: 4-space or tab indent makes code block in markdown syntax. See [here](https://daringfireball.net/projects/markdown/syntax#precode)

```md
    echo hello
    echo world
```

The above is a code block of `echo hello` for the first line and `echo world` for the second line.

# `saku.md` Rules

- Heading (# title) starts the task definition.
- Code blocks are commands.
  - Code blocks can have multiple commands. They will be executed sequentially.
- Blockquotes are description of the task.
- Anything else is ignored.
- Anything before the first heading is ignored.

For example:

```
# build
> Build the go binary.

    echo Starting build go binary
    go build -v -i main.go
```

The above defines the task `build`, which has the description `Build the go binary.`. It has two commands `echo Starting build go binary` and `go build -v -i main.go` and they run in sequence.

## Parallel execution

With `-p, --parallel` option, you can run tasks in parallel like the below:

```
saku -p watch-scripts run-server
```

## Race execution

With `-r, --race` option, you can run tasks in parallel and terminate tasks when the first task finished. This is useful for testing servers.

This option takes effect only when `-p` option is specified.

```
saku -p -r run-server test-server
```

## Nesting tasks (Dependency of tasks)

You can use `saku` in `saku.md` like below:

```md
# dist

    saku js minify

# js

    browserify src/main.js > build/app.js

# minify

    uglify-js < build/app.js > build/app.min.js
```

In this way, you can express the dependencies of tasks.

If you need to invoke tasks in parallel from another task, use `saku -p`.

```md
# start

    saku -p serve watch

# watch

    my-watcher

# serve

    my-server
```

# CLI Usage

```
Usage: saku [options] <task, ...> [-- extra-options]

Options:
  -v, --version   - - - Shows the version number and exits.
  -h, --help  - - - - - Shows the help message and exits.
  -i, --info  - - - - - Shows the task information and exits.
  -p, --parallel  - - - Runs tasks in parallel. Default false.
  -s, --sequential  - - Runs tasks in serial. Default true.
  -c, --config <path> - Specifies the config file. Default is 'saku.md'.
  -r, --race  - - - - - Set the flag to kill all tasks when a task
                        finished with zero. This option is valid only
                        with 'parallel' option.
  -q, --quiet   - - - - Stops the logging.

The extra options after '--' are passed to each task command.
```

# Notes

## Embed `saku.md` in `README.md`

You can optionally embed `saku.md` in `README.md`. See the below for details.

README.md

```md
# Development

These are the commands for development. You can invoke them with `saku` command.

<!-- saku start -->

## build

    go build -v -i main.go

## test

    go test -race ./...

<!-- saku end -->
```

The contents between `<!-- saku start -->` and `<!-- saku end -->` are used as `saku.md`. You can write them as if they are the part of your `README.md`.

## Example use cases

- [moneybit-app](https://github.com/kt3k/moneybit-app)'s [saku.md](https://github.com/kt3k/moneybit-app/blob/master/saku.md)
  - A project of accounting app for mobile, written in JavaScript.

**Note**: Please add yours if you use saku in your OSS project!

## The origin of the name

Saku is the Japanese name for the Chinese character "作", which means "make". Saku is intended to be an alternative of `make` command (of a task runner use case).

# Prior Art

- [node-saku][]
  - the original implementation of saku in JavaScript
- [make][]
- [npm-run-all][] by @mysticatea
- yaml-based tools
  - [robo][]
  - [go-task][]

# History

- 2018-01-31   v0.11.0   Pass options after `--`.

# License

MIT

[make]: https://en.wikipedia.org/wiki/Make_(software)
[npm-run-all]: https://github.com/mysticatea/npm-run-all
[robo]: https://github.com/tj/robo
[go-task]: https://github.com/go-task/task
[node-saku]: https://github.com/kt3k/node-saku
[release page]: https://github.com/kt3k/saku/releases/tag/v0.1.4
