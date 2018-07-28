package saku

type taskChannels struct {
	onCommand chan string
}

func newTaskChannels() *taskChannels {
	return &taskChannels{
		onCommand: make(chan string),
	}
}
