package sharding

import (
	"errors"
	"fmt"
	"github.com/ichaly/go-next/pkg/util"
	"gorm.io/gorm"
	"strconv"
	"sync"
)

var (
	ErrInvalidID                               = errors.New("invalid id format")
	ErrMissingShardingKey                      = errors.New("sharding key or id required, and use operator =")
	ErrInsertDiffSuffix                        = errors.New("can not insert different suffix table in one query ")
	ErrInvalidShardingAlgorithm                = errors.New("specify NumberOfShards or ShardingAlgorithm")
	ErrInvalidValueForDefaultShardingAlgorithm = errors.New("default algorithm only support integer and string column," +
		"if you use other type, specify you own ShardingAlgorithm")
)

func Register(config Config, tables ...any) *Sharding {
	return &Sharding{_config: config, _tables: tables}
}

type Sharding struct {
	*gorm.DB
	ConnPool *Connection

	mutex   sync.RWMutex
	configs map[string]Config

	_config Config
	_tables []any
}

type Config struct {
	DoubleWrite       bool
	NumberOfShards    uint
	ShardingKey       string
	ShardingAlgorithm func(columnValue any) (suffix string, err error)

	tableFormat string
}

func (my *Sharding) Name() string {
	return "gorm:sharding"
}

func (my *Sharding) Initialize(db *gorm.DB) error {
	db.Dialector = NewShardingDialector(db.Dialector, my)
	my.DB = db

	_ = my.Callback().Create().Before("*").Register("gorm:sharding", my.decoration)
	_ = my.Callback().Query().Before("*").Register("gorm:sharding", my.decoration)
	_ = my.Callback().Update().Before("*").Register("gorm:sharding", my.decoration)
	_ = my.Callback().Delete().Before("*").Register("gorm:sharding", my.decoration)
	_ = my.Callback().Row().Before("*").Register("gorm:sharding", my.decoration)
	_ = my.Callback().Raw().Before("*").Register("gorm:sharding", my.decoration)

	if my._config.ShardingAlgorithm == nil {
		if my._config.NumberOfShards == 0 {
			return ErrInvalidShardingAlgorithm
		}
		if my._config.NumberOfShards < 10 {
			my._config.tableFormat = "_%01d"
		} else if my._config.NumberOfShards < 100 {
			my._config.tableFormat = "_%02d"
		} else if my._config.NumberOfShards < 1000 {
			my._config.tableFormat = "_%03d"
		} else if my._config.NumberOfShards < 10000 {
			my._config.tableFormat = "_%04d"
		}
		my._config.ShardingAlgorithm = func(value any) (suffix string, err error) {
			num := 0
			switch value := value.(type) {
			case int:
				num = value
			case int64:
				num = int(value)
			case string:
				num, err = strconv.Atoi(value)
				if err != nil {
					num = util.Hash(value)
				}
			default:
				return "", ErrInvalidValueForDefaultShardingAlgorithm
			}
			return fmt.Sprintf(my._config.tableFormat, num%int(my._config.NumberOfShards)), nil
		}
	}

	if my.configs == nil {
		my.configs = make(map[string]Config)
	}

	for _, table := range my._tables {
		if t, ok := table.(string); ok {
			my.configs[t] = my._config
		} else {
			stmt := &gorm.Statement{DB: my.DB}
			if err := stmt.Parse(table); err == nil {
				my.configs[stmt.Table] = my._config
			} else {
				return err
			}
		}
	}

	return nil
}

func (my *Sharding) decoration(db *gorm.DB) {
	my.mutex.Lock()
	if db.Statement.ConnPool != nil {
		my.ConnPool = &Connection{ConnPool: db.Statement.ConnPool, sharding: my}
		db.Statement.ConnPool = my.ConnPool
	}
	my.mutex.Unlock()
}

func (my *Sharding) resolve(query string, args ...any) (ftQuery, stQuery, tableName string, err error) {
	return
}
