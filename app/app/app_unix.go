package app

import (
	"os/signal"
	"syscall"
)

func (a *App) mainRoutine()  {
	signal.Notify(a.signal, syscall.SIGUSR1, syscall.SIGUSR2, syscall.SIGTERM, syscall.SIGINT)

	a.Stop(<-a.signal)
}