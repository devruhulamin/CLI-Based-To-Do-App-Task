package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"ruhultodo/todo"
	"strings"
)

var todoFileName = "todo.json"

func main() {
	envVal := os.Getenv("TODO_FILENAME")
	if len(envVal) > 0 {
		todoFileName = envVal
	}
	add := flag.Bool("add", false, "add a new task")
	del := flag.Int("del", 0, "delete a task")
	listAll := flag.Bool("list", false, "List all task")
	completed := flag.Int("complete", 0, "Marks Task As Complete")
	details := flag.Bool("det", false, "show details info")
	pending := flag.Bool("pending", false, "show uncompleted task")
	getOne := flag.Int("i", 0, "show a specifics")

	flag.Parse()

	l := &todo.List{}

	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case *listAll:
		fmt.Print(l)
	case *completed > 0:
		if err := l.Complted(*completed); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *add:
		task, err := getTask(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		l.Add(task)
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *del > 0:
		if err := l.Delete(*del); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	case *details:
		fmt.Print(l.DetailsShow())
	case *pending:
		fmt.Print(l.ShowPending())
	case *getOne > 0:
		fmt.Print(l.GetOne(*getOne))

	default:
		fmt.Fprintln(os.Stderr, "Invalid Option")
		os.Exit(1)

	}
}

func getTask(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, ""), nil
	}
	s := bufio.NewScanner(r)
	s.Scan()
	if err := s.Err(); err != nil {
		return "", err
	}
	text := s.Text()
	fmt.Println(text)
	if len(text) == 0 {
		return "", fmt.Errorf("task Can Not Be emptly")
	}
	return text, nil
}
