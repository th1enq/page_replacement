package algorithms

func LRU(pages []int, frameCount int) SimulationResult {
	frames := make([]int, 0, frameCount)
	lastUsed := make(map[int]int)
	result := SimulationResult{}

	for t, page := range pages {
		found := false
		for _, p := range frames {
			if p == page {
				found = true
				break
			}
		}

		if found {
			// Cập nhật thời gian sử dụng
			lastUsed[page] = t
		} else {
			// Page fault
			result.PageFaults++
			if len(frames) < frameCount {
				frames = append(frames, page)
			} else {
				// Tìm LRU page
				lruPage := frames[0]
				lruTime := lastUsed[lruPage]
				for _, p := range frames {
					if lastUsed[p] < lruTime {
						lruPage = p
						lruTime = lastUsed[p]
					}
				}
				// Thay thế
				for i, p := range frames {
					if p == lruPage {
						frames[i] = page
						break
					}
				}
				delete(lastUsed, lruPage)
			}
			lastUsed[page] = t
		}

		step := createStep(page, frames, !found, append([]int{}, frames...), make(map[int]int))
		result.Steps = append(result.Steps, step)
	}

	return result
}
