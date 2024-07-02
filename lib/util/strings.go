package util

import (
	"errors"
	"strings"
)

func JoinString(elem ...string) string {
	b := strings.Builder{}
	for _, e := range elem {
		b.WriteString(e)
	}
	return b.String()
}

func StartWithAny(s string, prefixes ...string) (error, string) {
	for _, p := range prefixes {
		if strings.HasPrefix(s, p) {
			return nil, p
		}
	}
	return errors.New("not found"), ""
}
