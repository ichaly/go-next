package core

import (
	"fmt"
	"golang.org/x/exp/maps"
	"sort"
	"strings"
)

type Language struct {
	m *Metadata
}

func NewLanguage(m *Metadata) (*Language, error) {
	return &Language{m: m}, nil
}

func (my *Language) build() string {
	keys := maps.Keys(my.m.Nodes)
	sort.Strings(keys)

	var w strings.Builder
	for _, k := range keys {
		_, _ = fmt.Fprintf(&w, "%s\n", k)
	}

	return w.String()
}
