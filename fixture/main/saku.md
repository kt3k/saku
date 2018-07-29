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

# parallel-parent-task

<!-- saku parallel -->

## parallel-child-a

    sleep 0.1
    echo parallel-child-a

## parallel-child-b
    sleep 0.2
    echo parallel-child-b

# race-parent-task

<!-- saku parallel race -->

## race-child-a

    sleep 0.2
    echo race-child-a

## race-child-b

    sleep 0.1
    echo race-child-b

### race-grand-child-a

    sleep 1
    echo race-grand-child-a

### race-grand-child-b

    sleep 2
    echo race-grand-child-b
