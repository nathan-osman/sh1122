package sh1122

import (
	"image"
	"image/color"

	"periph.io/x/conn/v3"
)

// draw.Image methods:

func (s *SH1122) ColorModel() color.Model {
	return s.img.ColorModel()
}

func (s *SH1122) Bounds() image.Rectangle {
	return s.img.Bounds()
}

func (s *SH1122) At(x, y int) color.Color {
	return s.img.At(x, y)
}

func (s *SH1122) Set(x, y int, c color.Color) {
	s.img.Set(x, y, c)
}

// Flip blits the content of the internal buffer to the display. This is done
// using multiple SPI transfers if needed (due to internal limits).
func (s *SH1122) Flip() error {
	b := make([]byte, (Width*Height)/2)
	for i := 0; i < len(b); i++ {
		var (
			p1 = s.img.Pix[i*2] / 0x10
			p2 = s.img.Pix[i*2+1] / 0x10
		)
		b[i] = (p1 << 4) | p2
	}
	max := s.conn.(conn.Limits).MaxTxSize()
	for len(b) > max {
		s.send(b[:max], false)
		b = b[max:]
	}
	return s.send(b, false)
}
