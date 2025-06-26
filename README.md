# Introduction 

The Hue-bridge attacker command line support tool to hijack lights from a Philips Hue network with bridge integration. The tool finds a bridge IP-address that it wants to connect with if it is in the proximity of the device the tool is running on and on the same network.

# Installation

Direct to the hue-bridge-attacker folder and open a terminal.

1. Clone this repository

2. Build the `hue-attack` binary:

   ```sh
   go build -o hue-attack main.go
   ```

3. You can now execute the `hue-attack` binary.

## Notes on installation

- This program was developed using go version 1.24.4

# Usage

The usage is described in the command help text. This can be displayed by running the binary without arguments (`hue-attack`) or with the help flag set. (`hue-attack --help`)

Below is the output of the help text.

```
NAME:
   hue-attack - Attacks Philips Hue lights

USAGE:
   hue-attack [global options] [infiltration strategy] [attack strategy]

DESCRIPTION:
   This program attacks a Philips Hue Bridge. The attack consists of 2 stages. Strategies for these stages are configured using arguments.

   Infiltration stage: The goal of this stage is to gain access to the Philips Hue Bridge. Available infiltration strategies are described below.

   Attack stage: The goal of this stage is to perform some kind of attack. Available attack strategies are described below.

   This program was made for the course 2IC80 Lab on Offensive Computer Security at the Eindhoven University of Technology. Only use this program in environments where you are allowed to.

GLOBAL OPTIONS:
   --ip string                  Don't discover the Philips Hue Bridge, use this IP by instead
   --message string, -m string  Only available with "morse" attack strategy. Sets the message to communicate via Morse code.
   --help, -h                   show help

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
```

## Notes on usage

- The `dos` infiltration strategy uses the `hping3` command and therefore only works on systems that have this installed.
- `sudo` might be needed to use the `dos` infiltration strategy.
- The `disconnect` attack strategy is best used together with another Zigbee coordinator with `permit-join` enabled such that other coordinator can connect with the lights.

# Disclaimer
The hue-bridge-attacker tool is intended solely for educational and ethical use. Any misuse, such as use with malicious intent, is strictly prohibited. The developers disclaim all responsibility and liability for any harm, damages, or legal consequences resulting from the improper use of this tool. Users take full responsibility for their own usage.
