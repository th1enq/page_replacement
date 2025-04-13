package algorithms

// SecondChance implements Second Chance (Clock) page replacement
func SecondChance(pages []int, frameCount int) SimulationResult {
	frames := make([]int, frameCount)
	for i := range frames {
		frames[i] = -1
	}

	result := SimulationResult{}
	refBits := make([]bool, frameCount)
	nextFrame := 0

	inPage := make(map[int]int) // map page -> index in frames

	for _, page := range pages {
		idx, found := inPage[page]

		if found {
			refBits[idx] = true // cấp cơ hội thứ hai
		} else {
			result.PageFaults++

			for {
				currentPage := frames[nextFrame]

				// Nếu khung trống hoặc không được cấp cơ hội
				if currentPage == -1 || !refBits[nextFrame] {
					// Nếu trang trước đó có tồn tại, xoá khỏi map
					if currentPage != -1 {
						delete(inPage, currentPage)
					}
					frames[nextFrame] = page
					refBits[nextFrame] = true
					inPage[page] = nextFrame
					nextFrame = (nextFrame + 1) % frameCount
					break
				} else {
					// Đã cấp cơ hội, reset và tiếp tục
					refBits[nextFrame] = false
					nextFrame = (nextFrame + 1) % frameCount
				}
			}
		}

		step := createStep(page, frames, !found, append([]int{}, frames...), make(map[int]int))
		result.Steps = append(result.Steps, step)
	}

	return result
}
