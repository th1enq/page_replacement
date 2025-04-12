# Page Replacement Algorithms Simulation

This is a web application that simulates various page replacement algorithms used in operating systems. The application implements the following algorithms:

- FIFO (First-In-First-Out)
- Optimal
- LRU (Least Recently Used)
- LFU (Least Frequently Used)
- MFU (Most Frequently Used)
- Second Chance (Clock)

## Requirements

- Go 1.21 or later
- Docker

## Running the Application

1. Clone this repository
2. Navigate to the project directory
3. Run the application:
   ```bash
   docker compose up --build
   ```
4. Open your web browser and visit `http://localhost:8080`

## Usage

1. Enter a sequence of page numbers in the input field (e.g., "7 0 1 2 0 3 0 4 2 3")
2. Specify the number of frames available
3. Click "Simulate" to see how each algorithm handles the page replacement

The results will show the state of the frames after each page request for all algorithms simultaneously, allowing you to compare their behavior.

## Algorithm Descriptions

- **FIFO**: Replaces the oldest page in memory
- **Optimal**: Replaces the page that will not be used for the longest time in the future
- **LRU**: Replaces the page that hasn't been used for the longest time
- **LFU**: Replaces the page that has been used least frequently
- **MFU**: Replaces the page that has been used most frequently
- **Second Chance**: Similar to FIFO but gives pages a second chance before replacement 