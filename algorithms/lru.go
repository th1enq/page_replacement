package algorithms

func LRU(pages []int, frameCount int) SimulationResult {
	frames := make([]int, frameCount)
	for i := range frames {
		frames[i] = -1
	}

	result := SimulationResult{}
	inPage := make(map[int]int) // page -> frame index

	index := make(map[int]int)

	lastUsed := 0

	for i, page := range pages {
		for lastUsed < i {
			frontPage := pages[lastUsed]
			_, exist := inPage[frontPage]
			if !exist || index[frontPage] != lastUsed {
				lastUsed++
			} else {
				break
			}
		}

		_, found := inPage[page]

		if !found {
			// Page fault
			result.PageFaults++

			if i < frameCount {
				frames[i] = page
				inPage[page] = i
			} else {
				oldPage := pages[lastUsed]
				lruIdx := inPage[oldPage]
				delete(inPage, oldPage)

				frames[lruIdx] = page
				inPage[page] = lruIdx
			}
		}
		index[page] = i

		step := createStep(page, frames, !found, append([]int{}, frames...), make(map[int]int))
		result.Steps = append(result.Steps, step)
	}

	return result
}
