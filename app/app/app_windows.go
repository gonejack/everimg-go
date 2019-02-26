package app

import "os"

func (a *App) mainRoutine()  {
	signal.Notify(a.signal, os.Interrupt)

	a.Stop(<-a.signal)
}