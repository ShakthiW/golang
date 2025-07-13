package main

func main() {
	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	cmdFlags := NewCommandFlags()

	storage.Load(&todos)
	cmdFlags.Execute(&todos)
	storage.Save(todos)
}