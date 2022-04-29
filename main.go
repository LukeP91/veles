package main

import (
	"os"

    "veles/mocking"
)

func main() {
	sleeper := &mocking.DefaultSleeper{}
	mocking.Countdown(os.Stdout, sleeper)
}
