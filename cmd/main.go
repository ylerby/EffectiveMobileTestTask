package main

import (
	"EffectiveMobileTask/internal/app"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	application := app.NewApplication()
	go application.Run()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	application.Stop()
}
