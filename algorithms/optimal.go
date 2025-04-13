package algorithms

import "sort"

// Optimal implements the optimal page replacement algorithm
func Optimal(pages []int, frameCount int) SimulationResult {
	frames := make([]int, frameCount)
	for i := range frames {
		frames[i] = -1
	}

	result := SimulationResult{}

	// Tiền xử lý vị trí xuất hiện của từng trang
	occurrences := make(map[int][]int)
	for idx, page := range pages {
		occurrences[page] = append(occurrences[page], idx)
	}

	for i, page := range pages {
		found := false
		for j := range frames {
			if frames[j] == page {
				found = true
				break
			}
		}

		if !found {
			result.PageFaults++
			// Chưa đầy khung, thêm luôn
			if i < frameCount {
				frames[i] = page
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

				frames[replaceIdx] = page
			}
		}

		step := createStep(page, frames, !found, append([]int{}, frames...), make(map[int]int))
		result.Steps = append(result.Steps, step)
	}

	return result
}
