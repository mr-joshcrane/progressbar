package progressbar

import (
	"fmt"
	"time"
)

type ProgressBar struct {
	seconds int
}

func New(time int) ProgressBar {
	return ProgressBar{time}
}

func (bar *ProgressBar) Start() {
	for elapsed := 1; elapsed <= bar.seconds; elapsed++ {
		time.Sleep(1 * time.Second)
		done := percentage(elapsed, bar.seconds)
		fmt.Printf(fmt.Sprintf("\r%d percent done", done))
	}
}

func percentage(elapsed, seconds int) int {
	return int(float64(elapsed)/float64(seconds) * 100)
}