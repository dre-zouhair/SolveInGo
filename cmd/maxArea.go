package cmd

func MaxArea(height []int) int {
	var surface, l, r = 0, 0, len(height) - 1

	for l <= r {
		if height[l] < height[r] {
			if newSurface := (r - l) * height[l]; newSurface > surface {
				surface = newSurface
			}
			l++
		} else {
			if newSurface := (r - l) * height[r]; newSurface > surface {
				surface = newSurface
			}
			r--
		}
	}

	return surface
}
