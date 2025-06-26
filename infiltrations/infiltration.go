package infiltrations

import (
	"context"
	"github.com/amimof/huego"
)

type InfiltrationFunc func(ctx context.Context, bridge *huego.Bridge) (string, error)
