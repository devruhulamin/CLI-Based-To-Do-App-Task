package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type todoItem struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type List []todoItem

// implements my custom string likes [to string] in dart

func (l *List) String() string {
	formatter := ""
	for i, v := range *l {
		prefix := " "
		if v.Done {
			prefix = "X "
		}
		formatter += fmt.Sprintf("%s%d: %s \n", prefix, i+1, v.Task)
	}
	return formatter
}

// details print
func (l *List) DetailsShow() string {
	formatter := ""
	for i, v := range *l {
		prefix := " "
		if v.Done {
			prefix = "X "
		}
		formatter += fmt.Sprintf("%s ->  %s%d: %s \n", v.CreatedAt.Format("02 Jan 2006 15:04"), prefix, i+1, v.Task)
	}

	return formatter
}

func (l *List) ShowPending() string {
	formatter := ""
	for i, v := range *l {
		if v.Done {
			continue
		}
		formatter += fmt.Sprintf("%s ->  %d: %s \n", v.CreatedAt.Format("02 Jan 2006 15:04"), i+1, v.Task)
	}

	return formatter
}

// add item to the list
func (l *List) Add(task string) {
	t := todoItem{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}
	*l = append(*l, t)
}

// marks task as completed

func (l *List) Complted(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("task item %d not exits", i)
	}
	ls[i-1].Done = true
	ls[i-1].CompletedAt = time.Now()

	return nil
}

// delete a todo item

func (l *List) Delete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("task item %d not exits", i)
	}
	*l = append(ls[:i-1], ls[i:]...)
	return nil
}

// get single task
func (l *List) GetOne(i int) (todoItem, error) {
	ls := *l
	if i <= 0 || i > len(ls) {
		return todoItem{}, fmt.Errorf("task item %d not exits", i)
	}
	item := *l
	return item[i], nil
}

// save list of todo as json
func (l *List) Save(filename string) error {
	js, err := json.Marshal(l)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, js, 0644)
}

// read the file from the disk and return

func (l *List) Get(filename string) error {
	file, err := os.ReadFile(filename)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	if len(file) == 0 {
		return nil
	}

	return json.Unmarshal(file, l)
}

func (l *List) Erash(filename string) {
	os.Remove(filename)
}
