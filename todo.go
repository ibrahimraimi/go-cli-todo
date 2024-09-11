package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

/* methods: --add */
func (todos *Todos) add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CompletedAt: nil,
		CreatedAt:   time.Now(),
	}

	*todos = append(*todos, todo)
}

func (todos *Todos) validateIdx(idx int) error {
	if idx < 0 || idx >= len(*todos) {
		err := errors.New("Invalid index")
		fmt.Println(err)
		return err
	}

	return nil
}

/* methods: --del */
func (todos *Todos) delete(idx int) error {
	t := *todos

	if err := t.validateIdx(idx); err != nil {
		return err
	}

	*todos = append(t[:idx], t[idx+1:]...)

	return nil
}

func (todos *Todos) toggle(idx int) error {
	t := *todos
	if err := t.validateIdx((idx)); err != nil {
		return err
	}

	isCompleted := t[idx].Completed

	if !isCompleted {
		completionTime := time.Now()
		t[idx].CompletedAt = &completionTime
	}

	t[idx].Completed = !isCompleted

	return nil
}

/* --edit */
func (todos *Todos) edit(idx int, title string) error {
	t := *todos

	if err := t.validateIdx(idx); err != nil {
		return err
	}

	t[idx].Title = title

	return nil
}

func (todos *Todos) print() {
	tb := table.New(os.Stdout)

	tb.SetPadding(5)
	tb.SetHeaderStyle(table.StyleBold)
	tb.SetLineStyle(table.StyleBrightYellow)
	tb.SetDividers(table.UnicodeRoundedDividers)

	tb.SetHeaders("ID", "Title", "Completed", "Created At", "Completed At")

	for index, t := range *todos {
		completed := "❌"
		completedAt := ""

		if t.Completed {
			completed = "✅"
			if t.CompletedAt != nil {
				completedAt = t.CompletedAt.Format(time.RFC1123)
			}
		}

		tb.AddRow(strconv.Itoa(index), t.Title, completed, t.CreatedAt.Format(time.RFC1123), completedAt)
	}

	tb.Render()
}
