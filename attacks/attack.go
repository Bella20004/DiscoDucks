package attacks

import (
	"context"
	"github.com/amimof/huego"
)

type AttackFunc func(ctx context.Context, bridge *huego.Bridge) error
