package ppm

import (
	"fmt"
	"io"
	"os"

	p "github.com/avkapustin/grt/internal/primitives"
)

func Save(canvas *p.Canvas, to io.Writer) error {
	_, err := fmt.Fprintf(to, "P6\n%d %d\n%d\n", canvas.Width(), canvas.Height(), 255)
	if err != nil {
		return err
	}

	buf := make([]byte, canvas.Width()*canvas.Height()*3) // r,g,b per pixel
	position := 0
	// TODO check reverse write order (from width-1, heigth-1 downto zero) - maybe bounds checking will be eliminated on each buf access
	for i := 0; i < canvas.Height(); i++ {
		for j := 0; j < canvas.Width(); j++ {
			color := canvas.GetScreenPixel(j, i)
			buf[position] = color.R
			buf[position+1] = color.G
			buf[position+2] = color.B
			position += 3
		}
	}
	_, err = to.Write(buf)
	return err
}

func SaveToFile(canvas *p.Canvas, to string) error {
	fd, err := os.Create(to)
	if err != nil {
		return err
	}
	defer fd.Close()

	return Save(canvas, fd)
}
