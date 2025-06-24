# Introduction 

The Hue-bridge attacker command line support tool to hijack lights from a Philips Hue network with bridge integration. The tool finds
a bridge IP-address that it wants to connect with if it is in the proximity of the device the tool is running on and on the same network.

# Technical setup
The tool was developed solely for a Philips Hue network in combination with a bridge system. The tool is designed for Linux,
but the passive infiltration strategy also works on a Windows device. The disconnect attack strategy solely works if you have another
active Zigbee coordinator running in a permit-join mode.
The device it is being run on needs to have go installed. As it is developed with go1.24.4, it is only for this version, we can guarantee it works.

# Using the tool 

Direct to the hue-bridge-attacker folder and open a terminal.

1.

```
go build -o hue-attack main.go
```
2. Fill in commands with the following structure.

```
./hue-attack [global options] [infiltration strategy] [attack strategy]
```
3. for an overview of the options fill in the following command

```
./hue-attack -help
```
Please note that for the infiltration strategy dos, you need to add sudo to the command on a Linux device.

# Disclaimer
The hue-bridge-attacker tool is intended solely for educational and ethical use. Any misuse, such as use with malicious intent,
is strictly prohibited. The developers disclaim all responsibility and liability for any 
harm, damages, or legal consequences resulting from the improper use of this tool. Users take full responsibility for 
their own usage.
