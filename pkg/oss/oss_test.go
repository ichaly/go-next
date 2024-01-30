package oss

import (
	"path"
	"strings"
	"testing"
)

func TestPath(t *testing.T) {
	name := "test.png"
	t.Log(path.Join("", name))
	t.Log(path.Join("/", name))
	t.Log(path.Join("/key", name))
	t.Log(path.Join("/key//", name))
	t.Log(path.Join("key//", name))
	t.Log(path.Join("///key/demo.jpg", name))
	t.Log(strings.TrimPrefix(path.Join("///key/demo.jpg", name), "/"))
}
