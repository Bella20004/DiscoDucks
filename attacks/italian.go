package attacks

import (
	"context"
	"fmt"
	"hue-bridge-attacker/colors"
	"image/color"
	"time"

	"github.com/amimof/huego"
)

// TODO: Stop when context is cancelled
func Italian(ctx context.Context, bridge *huego.Bridge) error {
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

	var i int = 0

	for {
		for _, light := range lights {
			var col color.RGBA
			if i == 0 {
				col = colors.Green
			} else if i == 1 {
				col = colors.White
			} else { // i == 2
				col = colors.Red
				i = -1 // as after this i++
			}

			err = light.Col(col)
			if err != nil {
				return err
			}
		}
		i++

		time.Sleep(500 * time.Millisecond)
	}
}
