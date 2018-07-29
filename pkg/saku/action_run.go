package saku

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func actionRun(titles []string, tasks *TaskCollection, runCtx *runContext) error {
	for _, title := range titles {
		if tasks.findByTitle(title) == nil {
			return fmt.Errorf("Task not defined: %s", title)
		}
	}

	done := make(chan error, 1)
	sigs := make(chan os.Signal, 1)

	runTasks := tasks.filterByTitles(titles)
	runTasks.SetRunMode(runCtx.mode)

	stack := newTaskStack()

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	go func() {
		sig := <-sigs

		runTasks.abort()

		done <- fmt.Errorf("aborted: signal=%s", sig)
	}()

	go func() {
		done <- runTasks.Run(runCtx, stack)
	}()

	return <-done
}
