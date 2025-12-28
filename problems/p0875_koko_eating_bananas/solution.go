package p0875_koko_eating_bananas

func minEatingSpeed(piles []int, h int) int {
	left := 1
	right := piles[0]

	for _, pile := range piles {
		left = min(pile, left)
		right = max(pile, right)
	}

	for left <= right {
		mid := (left + right) / 2

		estH := calH(piles, mid)

		if estH > h {
			left = mid + 1
		} else if estH <= h {
			right = mid - 1
		}
	}

	return right + 1
}

func calH(piles []int, eSpeed int) int {
	esth := 0

	for _, pile := range piles {
		esth += (pile / eSpeed)
		if (pile % eSpeed) > 0 {
			esth += 1
		}
	}

	return esth
}
