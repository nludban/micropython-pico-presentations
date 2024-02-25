package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/ajstarks/svgo"
)


func TextStyle(anchor string) string {
	// "text-anchor:left;font-size:20px;fill:black"
	return strings.Join(
		[]string{
			"text-anchor:", anchor,
			";font-size:20px;fill:black;stroke-width:1",
			";alignment-baseline:middle",
		},
		"",
	)
}


func Power(c *svg.SVG, title string, x, y int) {
	w := 18

	c.Line(x - w, y + 40, x + w, y + 40)
	c.Line(x, y + 40, x, y + 72)

	c.Text(x, y + 30, title, TextStyle("middle"))
	return
}


func LED(c *svg.SVG, title string, x, y int) {
	w := 18

	c.Line(x, y, x, y + 20)				// |
	c.Polygon(
		[]int{ x - w, x + w, x },		// --
		[]int{ y + 20, y + 20, y + 52 },	// \/
	)
	c.Line(x - w, y + 52, x + w, y + 52)		// --
	c.Line(x, y + 52, x, y + 72)			// |

	c.Line(x + w, y + 30, x + 25, y + 25)
	c.Line(x + w, y + 35, x + 25, y + 35)
	c.Line(x + w, y + 40, x + 25, y + 45)

	c.Text(x + 30, y + 20, title, TextStyle("left"))
	return
}


func Ground(c *svg.SVG, x, y int) {
	c.Line(x, y, x, y + 20)
	c.Line(x - 20, y + 20, x + 20, y + 20)
	c.Line(x - 14, y + 30, x + 14, y + 30)
	c.Line(x - 8, y + 40, x + 8, y + 40)
}


func Ohm(c *svg.SVG, title string, x, y int) {
	w := 18
	h := 6

	c.Polyline(
		// XXX length mismatch => silently outputs empty.
		[]int{ x, x, x + w, x - w, x + w, x - w, x + w, x - w, x, x },
		[]int{ y, y + 12,
			y + 12 + 1 * h,
			y + 12 + 2 * h,
			y + 12 + 3 * h,
			y + 12 + 4 * h,
			y + 12 + 5 * h,
			y + 12 + 6 * h,
			y + 12 + 7 * h,
			y + 72 },
	)

	c.Text(x + 30, y + 50, title, TextStyle("left"))
	return
}


func Inverter(c *svg.SVG, x, y int) {
	c.Line(x, y, x + 15, y)
	c.Polygon(
		[]int{ x + 15, x + 47, x + 15 },
		[]int{ y + 18, y, y - 18 },
	)
	c.Line(x + 57, y, x + 72, y)
	c.Circle(x + 52, y, 5)

	c.Text(x - 5, y + 40, "74ahct14", TextStyle("center"))
}


func GPOut(c *svg.SVG, title string, x, y int) {
	c.Polygon(
		[]int{ x + 10, x + 62, x + 72, x + 62, x + 10 },
		[]int{ y + 15, y + 15, y, y - 15, y - 15 },
	)
	c.Text(x + 15, y, title, TextStyle("start"))
}


func GPIn(c *svg.SVG, title string, x, y int) {
	h := 15
	c.Polygon(
		[]int{ x + 72, x + 20, x + 10, x + 20, x + 72 },
		[]int{ y + h, y + h, y, y - h, y - h },
	)
	c.Text(x + 20, y, title, TextStyle("start"))
}


func Potentiometer(c *svg.SVG, x, y int) {
	h := 6
	w := 18
	px := []int{ x, x }
	py := []int{ y, y + 20 - h }
	for i := 0; i < 16; i += 2 {
		px = append(px, x + w)
		py = append(py, y + 20 + (i + 0) * h)

		px = append(px, x - w)
		py = append(py, y + 20 + (i + 1) * h)
	}
	px = append(px, x)
	py = append(py, py[len(py) - 1] + h)
	px = append(px, x)
	py = append(py, y + 144)

	fmt.Println("px=", px)
	fmt.Println("py=", py)
	c.Polyline(px, py)

	c.Line(x - 72, y + 72, x - 18, y + 72)
	c.Line(x - 24, y + 72 - 10, x - 18, y + 72)
	c.Line(x - 24, y + 72 + 10, x - 18, y + 72)

}

//--------------------------------------------------

func P(p int) int {
	return 20 + p * 72
}


func ExternalLED() {
	fout, _ := os.Create("pin-18-led.svg")

	var c *svg.SVG = svg.New(fout)

	width := 300
	height := 300
	c.Start(width, height)
	c.Roundrect(0, 0, width, height, 15, 15,
		"fill:snow;stroke:black;stroke-width:3")

	c.Gstyle("fill:none;stroke:black;stroke-width:3")

	Power(c, "+5V",		P(2), P(0))
	Ohm(c, "1k0",		P(2), P(1))
	LED(c, "Red",		P(2), P(2))

	GPOut(c, "GP18",	P(0), P(3))
	Inverter(c,		P(1), P(3))

	c.Gend()
	c.End()
	
	fout.Close()
}


func ThreeExternalLEDs() {
	fout, _ := os.Create("pin-18-20-leds.svg")

	var c *svg.SVG = svg.New(fout)

	width := 550
	height := 450
	c.Start(width, height)
	c.Roundrect(0, 0, width, height, 15, 15,
		"fill:snow;stroke:black;stroke-width:3")

	c.Gstyle("fill:none;stroke:black;stroke-width:3")

	Power(c, "+5v",		P(1), P(0))
	c.Line(P(1), P(1), P(6), P(1))

	Ohm(c, "1k0",		P(2), P(1))
	LED(c, "Red",		P(2), P(2))
	GPOut(c, "GP18",	P(0), P(3))
	Inverter(c,		P(1), P(3))

	Ohm(c, "1k0",		P(4), P(1))
	LED(c, "Green",		P(4), P(2))
	GPOut(c, "GP19",	P(0), P(4))
	Inverter(c,		P(1), P(4))
	c.Line(P(2), P(4), P(4), P(4))
	c.Line(P(4), P(4), P(4), P(3))

	Ohm(c, "1k0",		P(6), P(1))
	LED(c, "Blue",		P(6), P(2))
	GPOut(c, "GP20",	P(0), P(5))
	Inverter(c,		P(1), P(5))
	c.Line(P(2), P(5), P(6), P(5))
	c.Line(P(6), P(5), P(6), P(3))

	c.Gend()
	c.End()
	
	fout.Close()
}


func AnalogInput() {
	fout, _ := os.Create("pin-28-analog.svg")

	var c *svg.SVG = svg.New(fout)

	width := 250
	height := 350
	c.Start(width, height)
	c.Roundrect(0, 0, width, height, 15, 15,
		"fill:snow;stroke:black;stroke-width:3")

	c.Gstyle("fill:none;stroke:black;stroke-width:3")

	Power(c, "+3v3",	P(2), P(0))
	Potentiometer(c,	P(2), P(1))
	Ground(c,		P(2), P(3))

	GPIn(c, "GP28",		P(0), P(2))

	c.Gend()
	c.End()
	
	fout.Close()
}


func main() {
	ExternalLED()
	ThreeExternalLEDs()
	AnalogInput()
}
