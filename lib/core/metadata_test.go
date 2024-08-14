package core

import (
	"github.com/iancoleman/strcase"
	"github.com/ichaly/go-next/lib/base"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"reflect"
	"strings"
	"testing"
)

type _MetadataSuite struct {
	suite.Suite
	d *gorm.DB
	v *viper.Viper
}

func TestMetadata(t *testing.T) {
	suite.Run(t, new(_MetadataSuite))
}

func (my *_MetadataSuite) SetupSuite() {
	var err error
	my.v, err = base.NewViper("../../cfg/dev.yml")
	my.Require().NoError(err)
	my.d, err = base.NewConnect(my.v, []gorm.Plugin{base.NewSonyFlake()}, []interface{}{})
	my.Require().NoError(err)
}

func (my *_MetadataSuite) TestMetadata() {
	metadata, err := NewMetadata(my.v, my.d)
	my.Require().NoError(err)
	str, err := metadata.MarshalSchema()
	my.Require().NoError(err)
	my.T().Log(str)
}

func (my *_MetadataSuite) TestDecoder() {
	var c Column
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Result:           &c,
		WeaklyTypedInput: true,
		MatchName: func(mapKey, fieldName string) bool {
			mapKey = strcase.ToSnake(mapKey)
			return strings.EqualFold(mapKey, fieldName)
		},
		DecodeHook: func(f reflect.Value, t reflect.Value) (interface{}, error) {
			data := f.Interface()
			println(data)
			return data, nil
		},
	})
	my.Require().NoError(err)
	err = decoder.Decode(struct {
		DataType          string `gorm:"column:data_type;"`
		IsPrimary         bool   `gorm:"column:is_primary;"`
		IsForeign         bool   `gorm:"column:is_foreign;"`
		IsNullable        bool   `gorm:"column:is_nullable;"`
		TableName         string `gorm:"column:table_name;"`
		ColumnName        string `gorm:"column:column_name;"`
		TableRelation     string `gorm:"column:table_relation;"`
		ColumnRelation    string `gorm:"column:column_relation;"`
		TableDescription  string `gorm:"column:table_description;"`
		ColumnDescription string `gorm:"column:column_description;"`
	}{
		IsPrimary:         true,
		IsForeign:         true,
		IsNullable:        true,
		DataType:          "varchar",
		TableName:         "users",
		ColumnName:        "id",
		TableRelation:     "users",
		ColumnRelation:    "id",
		TableDescription:  "用户表",
		ColumnDescription: "用户id",
	})
	my.Require().NoError(err)
}
