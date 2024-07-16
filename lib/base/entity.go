package base

import (
	"context"
	"github.com/ichaly/go-next/lib/gql"
	"github.com/ichaly/go-next/lib/util"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type userContextKeyType struct{}

var UserContextKey = userContextKeyType{}

type Id uint64

func (my Id) ID() {}

func (my Id) String() string {
	str, err := gql.ShortId.Encode([]uint64{uint64(my)})
	if err != nil {
		return util.FormatLong(int64(my))
	}
	return str
}

type Primary struct {
	Id Id `gorm:"primary_key;comment:主键;next:sonyflake;" json:",omitempty"`
}

type General struct {
	State     int8              `gorm:"index;comment:状态;" `
	Remark    datatypes.JSONMap `gorm:"comment:备注" json:",omitempty"`
	CreatedAt *time.Time        `gorm:"comment:创建时间;" json:",omitempty"`
	UpdatedAt *time.Time        `gorm:"comment:更新时间;" json:",omitempty"`
}

type Entity struct {
	Primary `mapstructure:",squash"`
	General `mapstructure:",squash"`
}

type Deleted struct {
	DeletedAt *gorm.DeletedAt `gorm:"index;comment:逻辑删除;" json:",omitempty"`
}

type DeleteEntity struct {
	AuditorEntity `mapstructure:",squash"`
	Deleted       `mapstructure:",squash"`
}

type AuditorEntity struct {
	Entity    `mapstructure:",squash"`
	CreatedBy *Id `gorm:"comment:创建人;" json:",omitempty"`
	UpdatedBy *Id `gorm:"comment:更新人;" json:",omitempty"`
	DeletedBy *Id `gorm:"comment:删除人;" json:",omitempty"`
}

func GetUserFromContext(ctx context.Context) (interface{}, bool) {
	val, ok := ctx.Value(UserContextKey).(string)
	if !ok || val == "" {
		return nil, false
	}
	num, err := strconv.ParseUint(val, 10, 64)
	if err != nil {
		return nil, false
	}
	return &num, ok
}