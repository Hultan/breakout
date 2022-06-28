package game

import (
	"image/color"
)

func getColor(color color.Color) (float64, float64, float64, float64) {
	r, g, b, a := color.RGBA()
	return float64(r) / 65535, float64(g) / 65535, float64(b) / 65535, float64(a) / 65535
}

func clamp(v, minv, maxv float64) float64 {
	if v < minv {
		return minv
	} else if v > maxv {
		return maxv
	}

	return v
}
