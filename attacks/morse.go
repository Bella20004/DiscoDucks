package attacks

import (
	"context"
	"fmt"
	"github.com/gSpera/morse"
	"hue-bridge-attacker/colors"
	"time"

	"github.com/amimof/huego"
)

// TODO: Stop when context is cancelled
func Morse(ctx context.Context, bridge *huego.Bridge) error {
	lights, err := bridge.GetLights()
	if err != nil {
		return fmt.Errorf("error getting lights: %w", err)
	}
	fmt.Printf("Found %d lights\n", len(lights))

	var messageMorse string

	if message, ok := ctx.Value("message").(string); ok {
		fmt.Println("Retrieved message:", message)
		messageMorse = morse.ToMorse(message)
	} else {
		fmt.Println("No message found in context")
	}

	// Turn on lights and white and set transition time to shortest possible.
	for _, light := range lights {
		err = light.On()
		if err != nil {
			return err
		}

		err = light.TransitionTime(0)
		if err != nil {
			return err
		}

		err = light.Col(colors.White)
		if err != nil {
			return err
		}
	}
	for {
		for _, char := range messageMorse {
			if char == ' ' {
				time.Sleep(1000 * time.Millisecond) // Space between words
			}

			// Turn lights on and white
			for _, light := range lights {
				light.On()
				if err != nil {
					return err
				}

				err = light.Col(colors.White)
				if err != nil {
					return err
				}
			}

			// Wait for the duration of the Morse character
			if char == '.' {
				time.Sleep(200 * time.Millisecond)
			} else if char == '-' {
				time.Sleep(600 * time.Millisecond)
			}

			// Turn lights off
			for _, light := range lights {
				err = light.Off()
				if err != nil {
					return err
				}
			}

			time.Sleep(500 * time.Millisecond)
		}

		// Wait before repeating the message
		time.Sleep(2000 * time.Millisecond)
	}
}
