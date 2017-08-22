package main

import (
	"fmt"

	"github.com/fatih/structs"
)

func main() {
	ExampleNew()
	ExampleMap()
	ExampleMaps()
}

func ExampleNew() {
	type Server struct {
		Name    string
		ID      int32
		Enabled bool
	}

	server := &Server{
		Name:    "Arslan",
		ID:      123456,
		Enabled: true,
	}

	s := structs.New(server)

	fmt.Printf("Name        : %v\n", s.Name())
	fmt.Printf("Values      : %v\n", s.Values())
	fmt.Printf("Value of ID : %v\n", s.Field("ID").Value())
	// Output:
	// Name        : Server
	// Values      : [Arslan 123456 true]
	// Value of ID : 123456

}

func ExampleMap() {
	type Server struct {
		Name    string
		ID      int32
		Enabled bool
	}

	s := &Server{
		Name:    "Arslan",
		ID:      123456,
		Enabled: true,
	}

	m := structs.Map(s)

	fmt.Printf("%#v\n", m["Name"])
	fmt.Printf("%#v\n", m["ID"])
	fmt.Printf("%#v\n", m["Enabled"])
	// Output:
	// "Arslan"
	// 123456
	// true

}

// Map of slice
func ExampleMaps0() {
	type Server struct {
		Name    string
		ID      int32
		Enabled bool
	}

	s := []Server{
		{
			Name:    "Arslan",
			ID:      123456,
			Enabled: true,
		},
		{
			Name:    "Arslan",
			ID:      123456,
			Enabled: true,
		},
	}

	// XX: m := structs.Map(s)
	// panic: not struct
	fmt.Printf("%#v\n", s)
}

// Map of slice
func ExampleMaps() {
	type Server struct {
		Name    string
		ID      int32
		Enabled bool
	}

	s := struct {
		Servers []Server
	}{[]Server{
		{
			Name:    "Arslan",
			ID:      123456,
			Enabled: true,
		},
		{
			Name:    "Arslan",
			ID:      123456,
			Enabled: true,
		},
	},
	}

	m := structs.Map(s)
	fmt.Printf("%#v\n", m)
}