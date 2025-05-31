## sh1122

This package provides an easy way to control SH1122 devices in Go using SPI.

### Usage

The following example assumes an SSH1122 is connected to a Raspberry Pi Zero via SPI with the GPIO pins as described in the config:

```golang
package main

import (
    "github.com/nathan-osman/sh1122"
)

func main() {

    // Initialize the connection
    s, err := sh1122.New(&sh1122.Config{
        Port:   "/dev/spidev0.0",
        RSTPin: "GPIO23",
        DCPin:  "GPIO24",
        CSPin:  "GPIO25",
    })
    if err != nil {
        panic(err)
    }
    defer s.Close()

    // Turn the display on
    if err := s.SetDisplay(true) {
        panic(err)
    }

    // Flip what we've drawn to the display
    if err := s.Flip(); err != nil {
        panic(err)
    }
}
```
