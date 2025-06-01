package sh1122

import (
	"image"
	"image/color"

	"periph.io/x/conn/v3"
)

// Image data sent to the display is 4-bit

var palette = []color.Color{
	color.RGBA{0x00, 0x00, 0x00, 0xff},
	color.RGBA{0x11, 0x11, 0x11, 0xff},
	color.RGBA{0x22, 0x22, 0x22, 0xff},
	color.RGBA{0x33, 0x33, 0x33, 0xff},
	color.RGBA{0x44, 0x44, 0x44, 0xff},
	color.RGBA{0x55, 0x55, 0x55, 0xff},
	color.RGBA{0x66, 0x66, 0x66, 0xff},
	color.RGBA{0x77, 0x77, 0x77, 0xff},
	color.RGBA{0x88, 0x88, 0x88, 0xff},
	color.RGBA{0x99, 0x99, 0x99, 0xff},
	color.RGBA{0xaa, 0xaa, 0xaa, 0xff},
	color.RGBA{0xbb, 0xbb, 0xbb, 0xff},
	color.RGBA{0xcc, 0xcc, 0xcc, 0xff},
	color.RGBA{0xdd, 0xdd, 0xdd, 0xff},
	color.RGBA{0xee, 0xee, 0xee, 0xff},
	color.RGBA{0xff, 0xff, 0xff, 0xff},
}

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
	for y := 0; y < Height; y++ {
		for x := 0; x < Width; x += 2 {
			p1 := s.img.ColorIndexAt(x, y)
			p2 := s.img.ColorIndexAt(x+1, y)
			b[(y*Width+x)/2] = (p1 << 4) | p2
		}
	}
	max := s.conn.(conn.Limits).MaxTxSize()
	for len(b) > max {
		s.send(b[:max], false)
		b = b[max:]
	}
	return s.send(b, false)
}
