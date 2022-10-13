package password

import (
	"golang.org/x/term"
	"os"
	"os/signal"
	"syscall"
)

func ReadStdin() (pass []byte, err error) {
	stdin := int(os.Stdin.Fd())

	state, err := term.GetState(stdin)
	if err != nil {
		return
	}

	c := make(chan os.Signal, 1)
	defer close(c)
	signal.Notify(c, os.Interrupt, os.Kill)
	go func() {
		_, ok := <-c
		if ok {
			term.Restore(stdin, state)
			syscall.Kill(syscall.Getpid(), syscall.SIGKILL)
		}
	}()

	pass, err = term.ReadPassword(stdin)
	return
}
