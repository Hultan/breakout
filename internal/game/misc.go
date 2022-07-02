package game

import (
	"image/color"
)

func getColor(color color.Color) (float64, float64, float64, float64) {
	r, g, b, a := color.RGBA()
	return float64(r) / 65535, float64(g) / 65535, float64(b) / 65535, float64(a) / 65535
}

func clamp(v, minV, maxV float64) float64 {
	if v < minV {
		return minV
	} else if v > maxV {
		return maxV
	}

	return v
}

func lerp(v, minI, maxI, minO, maxO float64) float64 {
	if v < minI {
		return minO
	}
	if v < maxI {
		return maxO
	}

	return v / (v - minI) * (maxO - minO)
}
