package main

import (
	"log"
	"math/rand"
	"os"
	"os/signal"
	"time"
)

func main() {
	setupSignals()
}

func setupSignals() {
	sigs := make(chan os.Signal, 1)

	// catch all signals since not explicitly listing
	signal.Notify(sigs)
	//signal.Notify(sigs,syscall.SIGQUIT)

	// method invoked upon seeing signal
	go func() {
		s := <-sigs
		log.Printf("RECEIVED SIGNAL: %s", s)
		appCleanup()
		os.Exit(1)
	}()

	for {
		log.Printf("hello, world!")

		// wait random number of milliseconds
		Nsecs := rand.Intn(3000)
		log.Printf("About to sleep %dms before looping again", Nsecs)
		time.Sleep(time.Millisecond * time.Duration(Nsecs))
	}

}

func appCleanup() {
}
