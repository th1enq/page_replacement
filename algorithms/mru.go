package algorithms

func MRU(pages []int, frameCount int) SimulationResult {
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
			lastUsed[page] = t
		} else {
			result.PageFaults++
			if len(frames) < frameCount {
				frames = append(frames, page)
			} else {
				// Tìm trang được dùng gần nhất (MRU)
				mruPage := frames[0]
				mruTime := lastUsed[mruPage]
				for _, p := range frames {
					if lastUsed[p] > mruTime {
						mruPage = p
						mruTime = lastUsed[p]
					}
				}
				// Thay thế mruPage bằng page mới
				for i, p := range frames {
					if p == mruPage {
						frames[i] = page
						break
					}
				}
				delete(lastUsed, mruPage)
			}
			lastUsed[page] = t
		}

		step := createStep(page, frames, !found, append([]int{}, frames...), make(map[int]int))
		result.Steps = append(result.Steps, step)
	}

	return result
}
