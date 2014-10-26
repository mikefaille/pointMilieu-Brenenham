package main

// See https://en.wikipedia.org/wiki/Bresenham%27s_line_algorithm

func main() {
	brenhenham(0,0,10,6)

}

func brenhenham(x1 int, y1 int, x2 int, y2 int) {
	println("x1 =", x1, ", y1=", y1, ", x2 = ", x2, ", y2 =", y2)
	dx := x2 - x1
	dy := y2 - y1
	incCE := 2 * dy
	incNE := 2 * dy - 2 * dx
	d := 2 * dy - dx
	y := y1
	println("dx =", dx, ", dy=", dy, ", incCE = ", incCE, "incNE = ", incNE, ", d =",  d, ", y =", y, "\n" )


	for x := x1; x <= x2; x++ {
		println("x = ", x, ", y = ", y)
		if d <= 0 {
			println("if \n d =", d, " , incCE =", incCE, "\n")
			d += incCE
		} else {
			println("else \n d =", d, ",  incCE =", incNE, "\n")
			d += incNE
			y++
		}

	}

}
