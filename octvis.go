package main

import (
	"math"
	"math/rand"
	"bytes"
	"strconv"
	"image"
	"image/color"
	"image/png"
	"net/http"
)

func hexcol(col uint32) color.RGBA {
	b := uint8(col)
	col >>= 8
	g := uint8(col)
	col >>= 8
	r := uint8(col)
	return color.RGBA{r, g, b, 255}
}

func point(img *image.RGBA, x int, y int, c color.Color) {
	b := img.Bounds()
	if x < b.Min.X || x >= b.Max.X { return }
	if y < b.Min.Y || y >= b.Max.Y { return }
	img.Set(x, y, c)
}

func genPNG(col []color.RGBA, rnd *rand.Rand) []byte {
	width := 1024
	height := 1024
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	var o, p, q, r Octonion

	for i := 0; i < 8; i++ { o[i] = float32(rnd.Intn(100)) }
	o[0] = 100000
	o = o.Normalized()

	for i := 0; i < 8; i++ { p[i] = float32(rnd.Intn(100)) }
	p = p.Normalized()

	for i := 0; i < 8; i++ { q[i] = float32(rnd.Intn(100))/1000 }
	q[0] = 1000000
	q = q.Normalized()

	for i := 0; i < 8; i++ { r[i] = float32(rnd.Intn(100))/100000 }
	r[0] = 10000000
	r = r.Normalized()

	for y := 0; y < 1024; y++ {
		for x := 0; x < 1024; x++ {
			img.Set(x, y, color.RGBA{0, 0, 0, 255})
		}
	}
	sc := float64(0.5)
	for i := 0; i < 100000; i++ {
		w := float32(math.Floor(float64(width)*sc))
		h := float32(math.Floor(float64(height)*sc))

		point(img, width/2+int(p[1]*w), height/2+int(p[7]*h), col[0])
		point(img, width/2+int(p[2]*w), height/2+int(p[7]*h), col[1])
		point(img, width/2+int(p[3]*w), height/2+int(p[7]*h), col[2])
		point(img, width/2+int(p[4]*w), height/2+int(p[7]*h), col[3])
		point(img, width/2+int(p[5]*w), height/2+int(p[7]*h), col[4])
		point(img, width/2+int(p[6]*w), height/2+int(p[7]*h), col[5])
		point(img, width/2+int(p[0]*w), height/2+int(p[7]*h), col[6])

		p = p.Mul(o)
		o = o.Mul(q)
		q = q.Mul(r)
	}

	buf := new(bytes.Buffer)
	_ = png.Encode(buf, img)
	return buf.Bytes()
}

func handleimgOLD(w http.ResponseWriter, req *http.Request) {
	var col [7]color.RGBA
	col[0] = color.RGBA{255, 0, 0, 255}
	col[1] = color.RGBA{0, 255, 0, 255}
	col[2] = color.RGBA{0, 0, 255, 255}
	col[3] = color.RGBA{0, 255, 255, 255}
	col[4] = color.RGBA{255, 0, 255, 255}
	col[5] = color.RGBA{255, 255, 0, 255}
	col[6] = color.RGBA{255, 255, 255, 255}

	rnd := rand.New(rand.NewSource(rand.Int63()))
	png := genPNG(col[:], rnd)

	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Lnegth", string(len(png)))

	_, _ = w.Write(png)
}

func handleimg(w http.ResponseWriter, req *http.Request) {
	// init RNG
	rnd := rand.New(rand.NewSource(rand.Int63()))
	// overwrite with explicit seed
	seed := req.URL.Query()["seed"]
	if len(seed) > 0 {
		n, err := strconv.Atoi(seed[0])
		if err == nil {
			rnd = rand.New(rand.NewSource(int64(n)))
		}
	}

	// init colors from RNG
	var col [7]color.RGBA
	for i := 0; i < 7; i++ {
		col[i] = color.RGBA{uint8(rnd.Int()), uint8(rnd.Int()),
			uint8(rnd.Int()), 255}
	}
	// overwrite with explicit colors
	cols := req.URL.Query()["c"]
	j := 0
	for i := 0; i < len(cols); i++ {
		if j >= len(col) {
			break
		}
		c, err := strconv.ParseUint(cols[i], 16, 32)
		if err == nil {
			col[j] = hexcol(uint32(c))
			j++
		}
	}

	png := genPNG(col[:], rnd)

	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Lnegth", string(len(png)))

	_, _ = w.Write(png)
}

func main() {
	http.HandleFunc("/image", handleimg)
	http.ListenAndServe(":8090", nil)
}
