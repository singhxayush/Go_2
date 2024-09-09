package main

func main() {
	todos := Todos{}

	storage := NewStorage[Todos]("todos.json")
	storage.Load(&todos)

	CmdFlags := NewCmdflags()
	CmdFlags.Execute(&todos)

	storage.Save(todos)
}
