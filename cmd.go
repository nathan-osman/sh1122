package sh1122

const (
	cmdSET_DISP     = 0xae
	cmdSET_CONTRAST = 0x81
)

// SetDisplay turns the screen on or off depending on the provided parameter.
func (s *SH1122) SetDisplay(on bool) error {
	if on {
		return s.send([]byte{cmdSET_DISP | 0x01}, true)
	} else {
		return s.send([]byte{cmdSET_DISP}, true)
	}
}

// SetContrast sets the display's contrast to the provided value.
func (s *SH1122) SetContrast(v byte) error {
	return s.send([]byte{
		cmdSET_CONTRAST, v,
	}, true)
}
