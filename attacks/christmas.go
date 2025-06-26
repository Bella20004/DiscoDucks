package attacks

import (
	"context"
	"fmt"
	"github.com/amimof/huego"
	"image/color"
	"time"
	"hue-bridge-attacker/colors"
)

// TODO: Stop when context is cancelled
func Christmas(ctx context.Context, bridge *huego.Bridge) error {
	lights, err := bridge.GetLights()
	if err != nil {
		return fmt.Errorf("error getting lights: %w", err)
	}
	fmt.Printf("Found %d lights\n", len(lights))

	// Turn on lights and set transition time to shortest possible.
	for _, light := range lights {
		err = light.On()
		if err != nil {
			return err
		}

		err = light.TransitionTime(0)
		if err != nil {
			return err
		}
	}

	isGreen := false

	for {
		isGreen = !isGreen
		for _, light := range lights {
			var col color.RGBA
			if isGreen {
				col = colors.Green
			} else {
				col = colors.Red
			}

			err = light.Col(col)
			if err != nil {
				return err
			}
		}

		time.Sleep(500 * time.Millisecond)
	}
}
