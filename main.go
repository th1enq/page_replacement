package main

import (
	"html/template"
	"log"
	"net/http"
	"page_replacement/algorithms"
	"strconv"
	"strings"
)

type PageData struct {
	Input     string
	Frames    int
	Algorithm string
	Results   map[string]algorithms.SimulationResult
	Message   string
	Step      int
	IsRunning bool
}

func main() {
	http.HandleFunc("/", handleSimulation)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleSimulation(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("index.html").Funcs(template.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
	}).ParseFiles("templates/index.html"))

	data := PageData{
		Results:   make(map[string]algorithms.SimulationResult),
		Step:      -1,
		IsRunning: false,
	}

	if r.Method == http.MethodPost {
		input := r.FormValue("input")
		framesStr := r.FormValue("frames")
		algorithm := r.FormValue("algorithm")
		action := r.FormValue("action") // "step", "clear", "end" or empty for initial simulation

		// Handle clear action
		if action == "clear" {
			tmpl.Execute(w, data)
			return
		}

		frames, err := strconv.Atoi(framesStr)
		if err != nil {
			data.Message = "Invalid number of frames"
			tmpl.Execute(w, data)
			return
		}

		pages := parseInput(input)
		if len(pages) == 0 {
			data.Message = "Invalid input sequence"
			tmpl.Execute(w, data)
			return
		}

		data.Input = input
		data.Frames = frames
		data.Algorithm = algorithm

		// Get current step from form
		stepStr := r.FormValue("current_step")
		currentStep := -1
		if stepStr != "" {
			currentStep, _ = strconv.Atoi(stepStr)
		}

		// Handle different actions
		switch action {
		case "step":
			if currentStep < len(pages)-1 {
				data.Step = currentStep + 1
				data.IsRunning = true
			} else {
				data.Step = currentStep
				data.IsRunning = false
			}
		case "end":
			data.Step = len(pages) - 1
			data.IsRunning = false
		default:
			// Initial simulation
			data.Step = 0
			data.IsRunning = true
		}

		// Run simulation for selected algorithm
		switch algorithm {
		case "FIFO":
			data.Results["FIFO"] = algorithms.FIFO(pages, frames)
		case "Optimal":
			data.Results["Optimal"] = algorithms.Optimal(pages, frames)
		case "LRU":
			data.Results["LRU"] = algorithms.LRU(pages, frames)
		case "MRU":
			data.Results["LRU"] = algorithms.MRU(pages, frames)
		case "LFU":
			data.Results["LFU"] = algorithms.LFU(pages, frames)
		case "MFU":
			data.Results["MFU"] = algorithms.MFU(pages, frames)
		case "Second Chance":
			data.Results["Second Chance"] = algorithms.SecondChance(pages, frames)
		default:
			data.Message = "Please select an algorithm"
			tmpl.Execute(w, data)
			return
		}
	}

	tmpl.Execute(w, data)
}

func parseInput(input string) []int {
	parts := strings.Fields(input)
	result := make([]int, 0, len(parts))

	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			continue
		}
		result = append(result, num)
	}

	return result
}
