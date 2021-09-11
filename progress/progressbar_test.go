package progressbar_test

import (
	// "fmt"
	pb "github.com/mr-joshcrane/progressbar/progress"
	"math/rand"
	"testing"
	"time"
)

func DoWork() int {
	time.Sleep(1 * time.Second)
	return rand.Intn(20)
}

func TestProgress(t *testing.T) {
	bar := pb.New()
	go bar.Start(DoWork)
	err := bar.Render()
	if err != nil {
		t.Error(err)
	}
}
