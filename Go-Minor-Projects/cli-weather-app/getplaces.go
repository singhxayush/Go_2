// package main
//
// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"net/url"
// )
//
// // Structure to hold Nominatim response
// type Place struct {
// 	DisplayName string `json:"display_name"`
// }
//
// // Function to search locations using Nominatim
// func searchPlaces(query string) ([]Place, error) {
// 	baseURL := "https://nominatim.openstreetmap.org/search"
// 	params := url.Values{}
// 	params.Add("q", query)
// 	params.Add("format", "json")
// 	params.Add("addressdetails", "1")
// 	params.Add("limit", "5")
//
// 	// Make the API request
// 	resp, err := http.Get(baseURL + "?" + params.Encode())
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()
//
// 	// Parse the response
// 	var places []Place
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	err = json.Unmarshal(body, &places)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	return places, nil
// }
//
// func main() {
// 	query := "Sydney"
// 	places, err := searchPlaces(query)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
//
// 	for _, place := range places {
// 		fmt.Println("Location:", place.DisplayName)
// 	}
// }

// package main
//
// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"net/url"
// 	"os"
// 	"strings"
// 	"time"
//
// 	tea "github.com/charmbracelet/bubbletea"
// )
//
// // Structure to hold Nominatim response
// type Place struct {
// 	DisplayName string `json:"display_name"`
// }
//
// // Model for Bubble Tea, which holds the query and results
// type model struct {
// 	query   string
// 	results []Place
// }
//
// // Initialize the program
// func (m model) Init() tea.Cmd {
// 	return tea.Batch(tick(), nil)
// }
//
// // tick triggers the search every 500ms or so
// func tick() tea.Cmd {
// 	return tea.Tick(time.Millisecond*500, func(t time.Time) tea.Msg {
// 		return tickMsg{}
// 	})
// }
//
// // Custom message to re-trigger search periodically
// type tickMsg struct{}
//
// // Update handles keypress events and updates the state
// func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
// 	switch msg := msg.(type) {
//
// 	// Handle keyboard input
// 	case tea.KeyMsg:
// 		switch msg.String() {
// 		case "esc":
// 			return m, tea.Quit
// 		case "backspace":
// 			if len(m.query) > 0 {
// 				m.query = m.query[:len(m.query)-1]
// 			}
// 		default:
// 			m.query += msg.String()
// 		}
//
// 		// Fetch updated search results asynchronously
// 		m.results = fetchPlaces(m.query)
//
// 	// Periodic tick message to update the view
// 	case tickMsg:
// 		// No specific action, but re-tick the UI to refresh results
// 	}
//
// 	return m, tick()
// }
//
// // View renders the output in the terminal
// func (m model) View() string {
// 	var b strings.Builder
// 	b.WriteString("Type to search locations (esc to quit):\n")
// 	b.WriteString(m.query + "\n\n")
// 	b.WriteString("Results:\n")
// 	for _, place := range m.results {
// 		b.WriteString(place.DisplayName + "\n")
// 	}
// 	return b.String()
// }
//
// // Function to search locations using Nominatim
// func fetchPlaces(query string) []Place {
// 	if len(query) == 0 {
// 		return []Place{}
// 	}
//
// 	// Call the Nominatim API
// 	baseURL := "https://nominatim.openstreetmap.org/search"
// 	params := url.Values{}
// 	params.Add("q", query)
// 	params.Add("format", "json")
// 	params.Add("addressdetails", "1")
// 	params.Add("limit", "5")
//
// 	// Make the API request
// 	resp, err := http.Get(baseURL + "?" + params.Encode())
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return []Place{}
// 	}
// 	defer resp.Body.Close()
//
// 	// Parse the response
// 	var places []Place
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println("Error reading response:", err)
// 		return []Place{}
// 	}
//
// 	err = json.Unmarshal(body, &places)
// 	if err != nil {
// 		fmt.Println("Error parsing JSON:", err)
// 		return []Place{}
// 	}
//
// 	return places
// }
//
// // Main function to run the Bubble Tea program
// func main() {
// 	m := model{}
// 	p := tea.NewProgram(m)
// 	if err := p.Start(); err != nil {
// 		fmt.Println("Error:", err)
// 		os.Exit(1)
// 	}
// }
