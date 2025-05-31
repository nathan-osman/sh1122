package sh1122

import (
	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
	"periph.io/x/conn/v3/spi"
	"periph.io/x/conn/v3/spi/spireg"
)

// SH1122 provides an interface for controlling SH1122 devices connected via SPI.
type SH1122 struct {
	port   spi.PortCloser
	conn   spi.Conn
	rstPin gpio.PinIO
	dcPin  gpio.PinIO
	csPin  gpio.PinIO
}

// New creates a new SH1122 instance.
func New(cfg *Config) (*SH1122, error) {
	p, err := spireg.Open(cfg.Port)
	if err != nil {
		return nil, err
	}
	c, err := p.Connect(0, spi.Mode0, 8)
	if err != nil {
		p.Close()
		return nil, err
	}
	return &SH1122{
		port:   p,
		conn:   c,
		rstPin: gpioreg.ByName(cfg.RSTPin),
		dcPin:  gpioreg.ByName(cfg.DCPin),
		csPin:  gpioreg.ByName(cfg.CSPin),
	}, nil
}

// Close shuts down the SH1122 device.
func (s *SH1122) Close() {
	s.port.Close()
}
