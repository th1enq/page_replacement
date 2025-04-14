package algorithms

import "sort"

// Optimal implements the optimal page replacement algorithm
func Optimal(pages []int, frameCount int) SimulationResult {
	frames := make([]int, frameCount)
	for i := range frames {
		frames[i] = -1
	}

	result := SimulationResult{}
	inPage := make(map[int]int)
	currentIndex := 0

	// Tiền xử lý vị trí xuất hiện của từng trang
	occurrences := make(map[int][]int)
	for idx, page := range pages {
		occurrences[page] = append(occurrences[page], idx)
	}

	for i, page := range pages {
		_, found := inPage[page]

		if !found {
			result.PageFaults++
			// Chưa đầy khung, thêm luôn
			if currentIndex < frameCount {
				frames[currentIndex] = page
				inPage[page] = 1
				currentIndex++
			} else {
				farthest := -1
				replaceIdx := 0

				for j := range frames {
					nextUseList := occurrences[frames[j]]
					pos := sort.Search(len(nextUseList), func(k int) bool {
						return nextUseList[k] > i
					})

					if pos == len(nextUseList) {
						replaceIdx = j
						break
					}

					if nextUseList[pos] > farthest {
						farthest = nextUseList[pos]
						replaceIdx = j
					}
				}

				delete(inPage, frames[replaceIdx])

				frames[replaceIdx] = page
				inPage[page] = 1
			}
		}

		step := createStep(page, frames, !found, append([]int{}, frames...), make(map[int]int))
		result.Steps = append(result.Steps, step)
	}

	return result
}
