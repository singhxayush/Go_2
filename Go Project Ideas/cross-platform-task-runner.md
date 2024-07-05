Okay, here's a conceptual outline and code snippets to get you started on building a cross-platform task runner in Go:

**Core Features:**

1. **Command Execution:**

   * Use the `os/exec` package to run shell commands.
   * Handle command arguments and input/output redirection.
   * Provide options for error handling and logging.

2. **Template Rendering:**

   * Integrate a templating engine like `text/template` or a more feature-rich option like `gotemplate`.
   * Allow users to define templates in a separate file or embed them within the task runner's configuration.
   * Support variables and data passing to templates.

3. **Configuration:**

   * Use a configuration file format like YAML, JSON, or TOML to define tasks.
   * Allow users to specify:
      * Task names
      * Commands to execute
      * Template files and data
      * Dependencies between tasks

4. **Cross-Platform Compatibility:**

   * Rely on Go's built-in cross-compilation capabilities.
   * Consider using platform-specific code blocks (if necessary) for handling OS-specific differences.

**Code Structure (Example):**

```go
package main

import (
	"fmt"
	"os"
	"os/exec"
	"text/template"
)

// Task represents a single task in the runner.
type Task struct {
	Name    string
	Command string
	Template string // Path to a template file
	Data    interface{}
}

func main() {
	// 1. Load tasks from a configuration file (e.g., tasks.yaml)

	// 2. Process each task
	for _, task := range tasks {
		fmt.Printf("Running task: %s\n", task.Name)

		// 3. Execute the command
		cmd := exec.Command("sh", "-c", task.Command)
		err := cmd.Run()
		if err != nil {
			fmt.Printf("Error executing command: %v\n", err)
			continue // Skip to the next task
		}

		// 4. Render the template (if provided)
		if task.Template != "" {
			tmpl, err := template.ParseFiles(task.Template)
			if err != nil {
				fmt.Printf("Error parsing template: %v\n", err)
				continue
			}

			err = tmpl.Execute(os.Stdout, task.Data)
			if err != nil {
				fmt.Printf("Error executing template: %v\n", err)
			}
		}
	}
}
```

**Additional Considerations:**

* **Error Handling:** Implement robust error handling and logging to provide informative feedback to users.
* **Parallel Execution:** Consider adding support for parallel task execution to improve performance.
* **Dependency Management:** Allow users to define dependencies between tasks so that tasks are executed in the correct order.
* **User Interface:** You could create a command-line interface (CLI) or a graphical user interface (GUI) for interacting with the task runner.

**Libraries to Explore:**

* **`os/exec`:** For executing shell commands.
* **`text/template` or `gotemplate`:** For template rendering.
* **`yaml`, `json`, or `toml`:** For configuration file parsing.
* **`flag`:** For parsing command-line arguments.

Remember that this is a basic outline. Building a full-featured task runner will involve more complex logic and error handling.