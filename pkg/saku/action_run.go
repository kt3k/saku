package saku

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func actionRun(titles []string, tasks *TaskCollection, l *logger, runOpts *runOptions) error {
	if runOpts.isSerialAndParallel() {
		return fmt.Errorf("both --serial and --parallel options are specified")
	}

	for _, title := range titles {
		if tasks.findByTitle(title) == nil {
			return fmt.Errorf("Task not defined: %s", title)
		}
	}

	done := make(chan error, 1)
	sigs := make(chan os.Signal, 1)

	runTasks := tasks.filterByTitles(titles)
	runTasks.SetRunMode(runOpts.runMode())

	stack := newTaskStack()
	channels := newTaskChannels()

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	go func() {
		sig := <-sigs

		runTasks.abort()

		done <- fmt.Errorf("aborted: signal=%s", sig)
	}()

	go func() {
		for {
			l.println("+" + <-channels.onCommand)
		}
	}()

	go func() {
		done <- runTasks.Run(runOpts, channels, stack, l)
	}()

	return <-done
}
