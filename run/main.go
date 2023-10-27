package main

import (
	"fmt"
	"github.com/ichaly/go-next/main/cmd"
	"github.com/ichaly/go-next/pkg/zlog"
	"github.com/rs/zerolog"
	"os"
	"strings"
	"time"
)

func main() {
	c := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.DateTime}
	c.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	c.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s=", i)
	}
	e := zlog.NewRotate(
		zlog.WithFilename("logs/error.log"),
		zlog.WithLevelFunc(func(l zlog.Level) bool { return l > zlog.WarnLevel }),
	)
	t := zlog.NewRotate(
		zlog.WithFilename("logs/trace.log"),
		zlog.WithLevelFunc(func(l zlog.Level) bool { return l <= zlog.WarnLevel }),
	)
	o := zerolog.MultiLevelWriter(c, e, t)
	l := zlog.New(zlog.WithOut(o), zlog.WithLevel(zlog.DebugLevel))
	zlog.SetDefault(l)

	cmd.Execute()
}
