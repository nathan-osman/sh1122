package sh1122

// Config stores configuration for initializing a SH1122 device.
type Config struct {

	// Port is the name of the port to use.
	Port string

	// RSTPin is the GPIO pin used for the reset line.
	RSTPin string

	// DCPin is the GPIO pin used for setting data / command.
	DCPin string

	// CSPin is the GPIO pin used for chip select.
	CSPin string
}
