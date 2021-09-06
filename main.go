package main

import (
	progress "github.com/mr-joshcrane/progressbar/progress"
)

func main() {
	bar := progress.New(13)
	bar.Start()
}