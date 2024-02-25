package main

import (
	"fmt"
	"math"
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


func sincos(r, a int) (x, y int) {
	phi := float64(-a) * math.Pi / 180.0
	x = int(math.Round(float64(r) * math.Cos(phi)))
	y = int(math.Round(float64(r) * math.Sin(phi)))
	return
}


func Sector(c *svg.SVG, cx, cy, r, a1, a2 int) {
	sx, sy := sincos(72, a1)
	ex, ey := sincos(72, a2)
	c.Arc(cx + sx, cy + sy,		// start
		72, 72, 0,		// axis lengths and rotation
		false, false,		// large? sweep?
		cx + ex, cy + ey,	// end
		"fill:gray;stroke:none")
	c.Polygon(
		[]int{ cx, cx + sx, cx + ex },
		[]int{ cy, cy + sy, cy + ey },
		"fill:gray;stroke:none")
	return
}


func PolarText(c *svg.SVG, cx, cy, r, a int, s string) {
	tx, ty := sincos(r, a)
	c.Text(cx + tx, cy + ty, s, TextStyle("middle"))
	return
}


func PolarSensor(c *svg.SVG, cx, cy, r, a int, name string) {
	p1x, p1y := sincos(r, a)
	p2x, p2y := sincos(r + 15, a + 5)
	p3x, p3y := sincos(r + 15, a - 5)
	c.Polygon(
		[]int{ cx + p1x, cx + p2x, cx + p3x },
		[]int{ cy + p1y, cy + p2y, cy + p3y },
		"stroke:blue;fill:blue",
	)
	PolarText(c, cx, cy, r + 25, a + 15, name)
	return
}


func PolarArrow(c *svg.SVG, cx, cy, r, as, ae int, label string) {
	ao := ae - 5
	at := ae - 20
	ah := ae
	if ae < as {
		ao = ae + 5
		at = ae + 20
		as, ae = ae, as
	}
	h1x, h1y := sincos(r + 10, ao)
	h2x, h2y := sincos(r, ah)
	h3x, h3y := sincos(r - 10, ao)
	sx, sy := sincos(r, as)
	ex, ey := sincos(r, ae)
	c.Arc(cx + sx, cy + sy,
		r, r, 0,
		false, false,
		cx + ex, cy + ey,
	)
	c.Polyline(
		[]int{ cx + h1x, cx + h2x, cx + h3x },
		[]int{ cy + h1y, cy + h2y, cy + h3y })

	PolarText(c, cx, cy, r + 20, at, label)
}


func RotaryStepDir() {
	fout, _ := os.Create("rotary-step-dir.svg")

	var c *svg.SVG = svg.New(fout)

	width := 500
	height := 250
	c.Start(width, height)
	c.Roundrect(0, 0, width, height, 15, 15,
		"fill:snow;stroke:black;stroke-width:3")

	c.Gstyle("fill:none;stroke:black;stroke-width:3")

	cx := 125
	cy := 150
	Sector(c, cx, cy, 72, 90 + 10, 180 + 10)
	Sector(c, cx, cy, 72, 270 + 10, 0 + 10)
	c.Circle(cx, cy, 72)
	c.Circle(cx, cy, 10, "fill:black")
	PolarText(c, cx, cy, 40, 15 + 45 * 1, "1")
	PolarText(c, cx, cy, 40, 15 + 45 * 3, "0")
	PolarText(c, cx, cy, 40, 15 + 45 * 5, "1")
	PolarText(c, cx, cy, 40, 15 + 45 * 7, "0")
	PolarSensor(c, cx, cy, 75, 0, "Step")
	PolarSensor(c, cx, cy, 75, 45, "Dir")
	PolarArrow(c, cx, cy, 90, 120, 90, "(+1)")

	cx = 350
	cy = 150
	Sector(c, cx, cy, 72, 0 - 10, 90 - 10)
	Sector(c, cx, cy, 72, 180 - 10, 270 - 10)
	c.Circle(cx, cy, 72)
	c.Circle(cx, cy, 10, "fill:black")
	PolarText(c, cx, cy, 40, -15 + 45 * 1, "0")
	PolarText(c, cx, cy, 40, -15 + 45 * 3, "1")
	PolarText(c, cx, cy, 40, -15 + 45 * 5, "0")
	PolarText(c, cx, cy, 40, -15 + 45 * 7, "1")
	PolarSensor(c, cx, cy, 75, 0, "Step")
	PolarSensor(c, cx, cy, 75, 45, "Dir")
	PolarArrow(c, cx, cy, 90, 90, 120, "(-1)")


	c.Gend()
	c.End()
	
	fout.Close()

}


func main() {
	ExternalLED()
	ThreeExternalLEDs()
	AnalogInput()
	RotaryStepDir()
}

/**/
