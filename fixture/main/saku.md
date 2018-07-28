# hello
> Say hello
> To the world

    echo hello
    echo world

# foo
> This command fails

    foo

# build-and-test
> Builds and tests

    saku build test

# nodesc

    echo nodesc

# fail-without-extra-args

    go run ../../fixture/main/fail-without-extra-args.go

# parent-task

    echo has-child

## child-task-a

    echo child-task-a

## child-task-b

    echo child-task-b

### grand-child-task-a

    echo child-task-a

### grand-child-task-b

    echo child-task-b
