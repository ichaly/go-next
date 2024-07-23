package core

import "fmt"

type Language struct {
	m *Metadata
}

func NewLanguage(m *Metadata) (*Language, error) {
	return &Language{m: m}, nil
}

func (my *Language) build() string {
	for s := range my.m.Nodes {
		fmt.Printf("%s\n", s)
	}
	return ""
}
