package base

import (
	"github.com/ichaly/go-next/lib/util"
	"github.com/spf13/viper"
	"path/filepath"
	"strings"
)

func NewViper(file string) (error, *viper.Viper) {
	//解析文件路径和名称
	path := filepath.Dir(file)
	name := strings.TrimSuffix(filepath.Base(file), filepath.Ext(file))

	//初始化配置
	v := viper.New()
	v.SetConfigName(name)
	v.AddConfigPath(path)

	//支持环境变量自动替换
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	v.AutomaticEnv()

	//读取跟配置文件
	if err := v.ReadInConfig(); err != nil {
		return err, nil
	}

	//合并其他配置文件
	profiles := strings.Split(v.GetString("profiles.active"), ",")
	for _, p := range profiles {
		if len(p) == 0 {
			continue
		}
		file = util.JoinString(name, "-", p)
		v.SetConfigName(file)
		_ = v.MergeInConfig()
	}

	return nil, v
}
