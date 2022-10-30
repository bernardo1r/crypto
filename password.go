package password

import (
	"fmt"
	"golang.org/x/term"
	"os"
	"os/signal"
)

func ReadStdin() (pass []byte, err error) {
	stdin := int(os.Stdin.Fd())

	state, err := term.GetState(stdin)
	if err != nil {
		return pass, fmt.Errorf("could not get terminal state: %w", err)
	}

	c := make(chan os.Signal, 1)
	defer close(c)
	signal.Notify(c, os.Interrupt)
	go func() {
		_, ok := <-c
		if ok {
			term.Restore(stdin, state)
			os.Exit(1)
		}
	}()

	pass, err = term.ReadPassword(stdin)
	return pass, err
}
