package sh1122

import (
	"image"
	"time"

	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
	"periph.io/x/conn/v3/physic"
	"periph.io/x/conn/v3/spi"
	"periph.io/x/conn/v3/spi/spireg"
	"periph.io/x/host/v3"
)

const (
	// Width indicates the number of horizontal pixels.
	Width = 256

	// Height indicates the number of vertical pixels.
	Height = 64
)

// SH1122 provides an interface for controlling SH1122 devices connected via SPI.
type SH1122 struct {
	port   spi.PortCloser
	conn   spi.Conn
	rstPin gpio.PinIO
	dcPin  gpio.PinIO
	csPin  gpio.PinIO
	img    *image.Gray
}

// New creates a new SH1122 instance.
func New(cfg *Config) (*SH1122, error) {
	_, err := host.Init()
	if err != nil {
		return nil, err
	}
	p, err := spireg.Open(cfg.Port)
	if err != nil {
		return nil, err
	}
	c, err := p.Connect(1*physic.MegaHertz, spi.Mode0, 8)
	if err != nil {
		p.Close()
		return nil, err
	}
	s := &SH1122{
		port:   p,
		conn:   c,
		rstPin: gpioreg.ByName(cfg.RSTPin),
		dcPin:  gpioreg.ByName(cfg.DCPin),
		csPin:  gpioreg.ByName(cfg.CSPin),
		img:    image.NewGray(image.Rect(0, 0, Width, Height)),
	}
	if err := s.init(); err != nil {
		p.Close()
		return nil, err
	}
	return s, nil
}

func (s *SH1122) init() error {
	if err := s.rstPin.Out(gpio.High); err != nil {
		return err
	}
	time.Sleep(10 * time.Millisecond)
	if err := s.rstPin.Out(gpio.Low); err != nil {
		return err
	}
	time.Sleep(10 * time.Millisecond)
	if err := s.rstPin.Out(gpio.High); err != nil {
		return err
	}
	return s.csPin.Out(gpio.High)
}

// Close shuts down the SH1122 device.
func (s *SH1122) Close() {
	s.port.Close()
}
