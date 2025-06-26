package infiltrations

import (
	"context"
	"fmt"
	"github.com/amimof/huego"
	"time"
)

const (
	HueBridgeAppName        = "disco-ducks"
	UserCreationRetryPeriod = 10 * time.Second
)

func Passive(ctx context.Context, bridge *huego.Bridge) (string, error) {
	attempts := 0

	for {
		attempts++
		fmt.Printf("User creation attempt #%d - Starting...\n", attempts)

		user, err := bridge.CreateUser(HueBridgeAppName)
		if err != nil {
			fmt.Printf("User createion attempt #%d - Failed, trying again in %d seconds.\n", attempts, UserCreationRetryPeriod/time.Second)
			time.Sleep(UserCreationRetryPeriod)
			continue
		}

		fmt.Printf("User createion attempt #%d - Success!\n", attempts)
		return user, nil
	}
}
