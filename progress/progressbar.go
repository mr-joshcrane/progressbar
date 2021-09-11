package progressbar

import (
	"fmt"
	"io"
	"bytes"
)

type ProgressBar struct {
	percentDone int
	buf 		bytes.Buffer
}

func New() ProgressBar {
	var buf bytes.Buffer
	return ProgressBar{0, buf}
}


type Work func() int

func (bar *ProgressBar ) Start(work Work) {
	for bar.percentDone < 100 {
		bar.percentDone += work()
		fmt.Fprint(&bar.buf, fmt.Sprintf("\r%d percent done", bar.percentDone))
	}
	bar.percentDone = 100
}

func (bar *ProgressBar) Render() error {
	for bar.percentDone != 100 {
		str, err := io.ReadAll(&bar.buf)
		if err != nil {
			return err
		}
		fmt.Print(string(str))
	}
	fmt.Println("\r100 percent done")
	return nil
}

// func percentage(elapsed, seconds int) int {
// 	return int(float64(elapsed)/float64(seconds) * 100)
// }