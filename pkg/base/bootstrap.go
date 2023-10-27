package base

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"net/http"
	"regexp"
	"strings"
	"time"
)

var (
	Version   = "V0.0.0"
	GitCommit = "Unknown"
	BuildTime = ""

	routers = make(map[string]gin.IRouter)
	reg     = regexp.MustCompile(`/+`)
)

func Bootstrap(l fx.Lifecycle, c *Config, e *gin.Engine, g PluginGroup) {
	if BuildTime == "" {
		BuildTime = time.Now().Format("2006-01-02 15:04:05")
	}
	routers["/"] = e
	all := append(g.Middlewares, g.Plugins...)
	for _, p := range all {
		base := fmt.Sprintf("%s/", strings.TrimRight(reg.ReplaceAllString(p.Base(), "/"), "/"))
		r, ok := routers[base]
		if !ok {
			r = e.Group(base)
			routers[base] = r
		}
		p.Init(r)
	}
	srv := &http.Server{Addr: fmt.Sprintf(":%v", c.App.Port), Handler: e}
	l.Append(fx.StartStopHook(func(ctx context.Context) {
		go startServer(srv, c)
	}, func(ctx context.Context) error {
		return stopServer(srv, c)
	}))
	fmt.Printf("当前版本:%s-%s 发布日期:%s\n", Version, GitCommit, BuildTime)
}

func startServer(srv *http.Server, c *Config) {
	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("%v failed to start: %v\n", c.App.Name, err)
	}
}

func stopServer(srv *http.Server, c *Config) error {
	fmt.Printf("%v shutdown complete\n", c.App.Name)
	return srv.Shutdown(context.Background())
}
