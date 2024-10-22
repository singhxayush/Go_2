// package main
//
// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
//
// 	tea "github.com/charmbracelet/bubbletea"
// )
//
// type model struct {
// 	cursor    int      // Which list item our cursor is on
// 	locations []string // All locations loaded from file
// 	filtered  []string // Filtered locations based on input
// 	query     string   // Current search input
// }
//
// func main() {
// 	// Load the locations from the file
// 	locations, err := loadLocations("locations.txt")
// 	if err != nil {
// 		fmt.Println("Error loading locations:", err)
// 		return
// 	}
//
// 	// Initialize the model
// 	initialModel := model{
// 		locations: locations,
// 		filtered:  locations, // Start with all locations
// 		query:     "",
// 	}
//
// 	// Run the Bubble Tea program
// 	p := tea.NewProgram(initialModel)
// 	if err := p.Start(); err != nil {
// 		fmt.Println("Error:", err)
// 		os.Exit(1)
// 	}
// }
//
// // Load locations from a file into a slice
// func loadLocations(fileName string) ([]string, error) {
// 	file, err := os.Open(fileName)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer file.Close()
//
// 	var locations []string
// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		locations = append(locations, scanner.Text())
// 	}
//
// 	if err := scanner.Err(); err != nil {
// 		return nil, err
// 	}
// 	return locations, nil
// }
//
// // Bubble Tea's Init function (initialization)
// func (m model) Init() tea.Cmd {
// 	return nil
// }
//
// // Update function is called when an event (e.g., key press) occurs
// func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
// 	switch msg := msg.(type) {
//
// 	// Handle key presses
// 	case tea.KeyMsg:
// 		switch msg.String() {
//
// 		// Move the cursor up
// 		case "up":
// 			if m.cursor > 0 {
// 				m.cursor--
// 			}
//
// 		// Move the cursor down
// 		case "down":
// 			if m.cursor < len(m.filtered)-1 {
// 				m.cursor++
// 			}
//
// 		// Enter key selects the current item
// 		case "enter":
// 			fmt.Printf("\nYou selected: %s\n", m.filtered[m.cursor])
// 			return m, tea.Quit
//
// 		// Handle backspace (delete last character in query)
// 		case "backspace":
// 			if len(m.query) > 0 {
// 				m.query = m.query[:len(m.query)-1]
// 				m.filtered = searchLocations(m.query, m.locations)
// 			}
//
// 		// Handle space (add space to the query)
// 		case " ":
// 			m.query += " "
// 			m.filtered = searchLocations(m.query, m.locations)
//
// 		// ESC key to quit
// 		case "esc":
// 			return m, tea.Quit
//
// 		// Handle any other character (add to query)
// 		default:
// 			if msg.String() != "" && msg.String() != " " {
// 				m.query += msg.String()
// 				m.filtered = searchLocations(m.query, m.locations)
// 			}
// 		}
// 	}
//
// 	return m, nil
// }
//
// // View function renders the UI
// func (m model) View() string {
// 	var s strings.Builder
//
// 	// Show the current query
// 	s.WriteString(fmt.Sprintf("Search: %s\n\n", m.query))
//
// 	// Display the list of matching locations
// 	for i, location := range m.filtered {
// 		// Highlight the cursor
// 		cursor := " "
// 		if m.cursor == i {
// 			cursor = ">"
// 		}
// 		s.WriteString(fmt.Sprintf("%s %s\n", cursor, location))
// 	}
//
// 	// Return the UI as a string
// 	return s.String()
// }
//
// // Search function filters the locations based on the query
// func searchLocations(query string, locations []string) []string {
// 	var matches []string
// 	for _, location := range locations {
// 		if strings.Contains(strings.ToLower(location), strings.ToLower(query)) {
// 			matches = append(matches, location)
// 		}
// 	}
// 	return matches
// }

package main

import "net/http"

func main() {

	url := "https://api.tomorrow.io/v4/weather/realtime?location=JamaDobA&apikey=0YDnpp2PKiSCvj1rRW9SZl4ac2oyG4pY"

	// req, _ := http.NewRequest("GET", url, nil)
	//
	// req.Header.Add("accept", "application/json")
	//
	// res, _ := http.DefaultClient.Do(req)
	//
	// defer res.Body.Close()
	// body, _ := io.ReadAll(res.Body)
	//
	// fmt.Println(string(body))

	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}

}

// 0YDnpp2PKiSCvj1rRW9SZl4ac2oyG4pY
