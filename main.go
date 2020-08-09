// Package todo contains a basic CLI todo app
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"time"
)

var todoItem = flag.String("add", "Something worth doing", "Item to add to todo list.\n\t{\"task\": task to do, \"due\": date due (YYYY-MM-DD)}")

// TodoItem is the internal type used to store the JSON data that is
// deserialised by the app.
type TodoItem struct {
	Todo string `json:"todo"`
	Due  string `json:"due,omitempty"`
}

// ParsedTodoItem is the same as the TodoItem type, albeit with the `Due` field
// parsed to due the `time.Time` struct.
type ParsedTodoItem struct {
	Todo string
	Due  time.Time
}

func parseDuedate(dueDate string) (parseddueDate time.Time) {
	if dueDate == "" {
		return time.Time{}
	}

	const dueDataFormat = "2006-01-02"

	parsedDueDate, parseErr := time.Parse(dueDataFormat, dueDate)

	if parseErr != nil {
		log.Fatal("Badly formed due date.")
	}

	return parsedDueDate
}

// ParseInput parses the input into something more usable.
// This includes checking for empty input and parsing the
// `due` field.
func ParseInput(input *string) ParsedTodoItem {
	var todoItem TodoItem

	err := json.Unmarshal([]byte(*input), &todoItem)
	if err != nil {
		log.Fatal("Invalid JSON passed to ./todo-app")
	}

	parsedDueDate := parseDuedate(todoItem.Due)

	parsedItem := ParsedTodoItem{Todo: todoItem.Todo, Due: parsedDueDate}

	return parsedItem

}

// PrettyPrintItem echoes back the parsed command line input.
func PrettyPrintItem(item ParsedTodoItem) (n int, err error) {
	formattedItem, err := json.MarshalIndent(item, "	", "	")
	if err != nil {
		return 0, err
	}
	return fmt.Printf("You entered:\n\n\t%s", string(formattedItem))
}

func main() {
	flag.Parse()

	PrettyPrintItem(ParseInput(todoItem))
}
