# hkpeople

A HomeKit sensor to activate when people (i.e. their phones) are home.
This is super useful if you want to turn off the lights after everyone left, or turn on if anyone comes home.

Thanks [@brutella](https://github.com/brutella) for the awesome [HAP Go library](https://github.com/brutella/hc) which made this possible.

## Usage

### On your machine

```sh
# Requires valid go install
$ go install github.com/bahlo/hkpeople/cmd/hkpeople

# You can find out the hostnames of your phones on iOS in Settings -> General -> Info -> Name
$ TARGETS=my-iphone.local,so-iphone.local hkpeople
```

### On your iPhone
Then open the Home app and add the sensor.
The pin is `45736753`, you can type the numbers that contain the respective letters for `hkpeople`.

## Precompile (for Raspberry Pi)

```sh
$ git clone github.com/bahlo/hkpeople

# Requires github.com/golang/dep to be installed
$ dep ensure

# Requires valid go install
$ GOOS=linux GOARCH=arm GOARM=6 go build -o hkpeople-rpi cmd/hkpeople/main.go

# Copy hkpeople-rpi to your raspberry and start it like describe above
```
