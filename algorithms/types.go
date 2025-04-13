package algorithms

// SimulationResult represents the result of a page replacement simulation
type SimulationResult struct {
	Steps      []Step
	PageFaults int
}

// Step represents a single step in the simulation
type Step struct {
	PageNumber int
	Frames     []int
	IsFault    bool
	Queue      []int
	Frequency  map[int]int
}

// createStep creates a new step with the given parameters
func createStep(pageNum int, frames []int, isFault bool, queue []int, freq map[int]int) Step {
	return Step{
		PageNumber: pageNum,
		Frames:     append([]int{}, frames...),
		IsFault:    isFault,
		Queue:      append([]int{}, queue...),
		Frequency:  freq,
	}
}
