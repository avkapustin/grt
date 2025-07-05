package primitives

import "fmt"

var MAX_CANVAS_WIDTH = 16384
var MAX_CANVAS_HEIGHT = 16384

type Canvas struct {
	width, height int
	data          []Color
}

func MakeCanvas(width, height int) (*Canvas, error) {
	if width <= 0 || width > MAX_CANVAS_WIDTH {
		return nil, fmt.Errorf("Canvas: width should be between 1 and %d, actually %d\n", MAX_CANVAS_WIDTH, width)
	}
	if height <= 0 || height > MAX_CANVAS_HEIGHT {
		return nil, fmt.Errorf("Canvas: height should be between 1 and %d, actually %d\n", MAX_CANVAS_HEIGHT, height)
	}

	return &Canvas{
		width,
		height,
		make([]Color, width*height),
	}, nil
}

func (c *Canvas) Width() int {
	return c.width
}

func (c *Canvas) Height() int {
	return c.height
}

func (c *Canvas) GetScreenPixel(x, y int) Color {
	if x < 0 || x >= c.width {
		return Color{}
	}
	if y < 0 || y >= c.height {
		return Color{}
	}
	return c.data[x+y*c.width]
}

func (c *Canvas) SetScreenPixel(x, y int, color Color) {
	if x < 0 || x >= c.width {
		return
	}
	if y < 0 || y >= c.height {
		return
	}
	c.data[x+y*c.width] = color
}

func (c *Canvas) SetViewportPixel(x, y int, color Color) {
	if x < -c.width/2 || x >= c.width/2 {
		return
	}
	if y <= -c.height/2 || y > c.height/2 {
		return
	}
	c.SetScreenPixel(x+c.width/2, -y+c.height/2, color)
}
