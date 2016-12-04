package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/oskca/gopherjs-vue"

	"github.com/robvdl/experiments/ex1/components/login"
)

// Model is the application model.
type Model struct {
	*js.Object

	IntValue int    `js:"integer"`
	Str      string `js:"str"`
}

// Inc is a method that simply increases a value.
func (m *Model) Inc() {
	m.IntValue++
	println("Inc called")
}

// Repeat is a method that duplicates a string on the app model.
func (m *Model) Repeat() {
	m.Str = m.Str + m.Str
}

// Reset will reset the string on the app model.
func (m *Model) Reset() {
	m.Str = "a string "
}

func main() {
	// register components.
	login.Register()

	// m is the application model.
	m := &Model{
		Object: js.Global.Get("Object").New(),
	}

	// field assignment is required in this way to make data passing works.
	m.IntValue = 100
	m.Str = "a string "

	// create the VueJS viewModel using a struct pointer.
	vue.New("#app", m)
}
