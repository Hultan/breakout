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

func removeEntity(e gameObject) {
	for i := range entities {
		if entities[i] == e {
			entities[i] = entities[len(entities)-1]
			entities = entities[:len(entities)-1]
		}
	}
}
