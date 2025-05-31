package sh1122

import (
	"periph.io/x/conn/v3/gpio"
)

func (s *SH1122) send(c []byte, cmd bool) error {
	var l gpio.Level
	if cmd {
		l = gpio.Low
	} else {
		l = gpio.High
	}
	if err := s.dcPin.Out(l); err != nil {
		return err
	}
	if err := s.csPin.Out(gpio.Low); err != nil {
		return err
	}
	if err := s.conn.Tx(c, make([]byte, len(c))); err != nil {
		return err
	}
	return s.csPin.Out(gpio.High)
}
