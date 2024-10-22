package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type Place struct {
	DisplayName string `json:"display_name"`
}

type model struct {
	query   string
	results []Place
}

// Custom message for debounced search
type searchMsg struct {
	places []Place
}

// Delay constant for debouncing
const debounceDuration = 500 * time.Millisecond

var debounceTimer *time.Timer

// Function to search places asynchronously
func searchPlacesDebounced(query string) tea.Cmd {
	return func() tea.Msg {
		if len(query) == 0 {
			return searchMsg{places: []Place{}}
		}

		// Call the Nominatim API
		baseURL := "https://nominatim.openstreetmap.org/search"
		params := url.Values{}
		params.Add("q", query)
		params.Add("format", "json")
		params.Add("addressdetails", "1")
		params.Add("limit", "5")

		resp, err := http.Get(baseURL + "?" + params.Encode())
		if err != nil {
			fmt.Println("Error:", err)
			return searchMsg{places: []Place{}}
		}
		defer resp.Body.Close()

		// Parse the response
		var places []Place
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response:", err)
			return searchMsg{places: []Place{}}
		}

		err = json.Unmarshal(body, &places)
		if err != nil {
			fmt.Println("Error parsing JSON:", err)
			return searchMsg{places: []Place{}}
		}

		return searchMsg{places: places}
	}
}

// Update handles keypress events and updates the state
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return m, tea.Quit
		case "backspace":
			if len(m.query) > 0 {
				m.query = m.query[:len(m.query)-1]
			}
		default:
			m.query += msg.String()
		}

		// Reset the debounce timer
		if debounceTimer != nil {
			debounceTimer.Stop()
		}
		debounceTimer = time.NewTimer(debounceDuration)

		// Trigger search after the debounce duration
		go func(query string) {
			<-debounceTimer.C
			m.results = fetchPlaces(m.query)
		}(m.query)

	case searchMsg:
		// Update the results after the debounced search is done
		m.results = msg.places
	}

	return m, nil
}

// View renders the output in the terminal
func (m model) View() string {
	var b strings.Builder
	b.WriteString("Type to search locations (esc to quit):\n")
	b.WriteString(m.query + "\n\n")
	b.WriteString("Results:\n")
	for _, place := range m.results {
		b.WriteString(place.DisplayName + "\n")
	}
	return b.String()
}

func main() {
	m := model{}
	p := tea.NewProgram(m)
	if err := p.Start(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
