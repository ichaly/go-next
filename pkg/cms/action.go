package cms

import (
	"github.com/ichaly/go-next/lib/base"
)

type Action struct {
	base.Primary `mapstructure:",squash"`
	ObjectId     int64 `gorm:"comment:内容ID"`
	View         int   `gorm:"comment:阅读量"`
	Like         int   `gorm:"comment:点赞量"`
	Diss         int   `gorm:"comment:差评量"`
	Share        int   `gorm:"comment:分享量"`
	Follow       int   `gorm:"comment:关注量"`
	Comment      int   `gorm:"comment:评论量"`
	Favorite     int   `gorm:"comment:收藏量"`
}

func (Action) TableName() string {
	return "cms_action"
}
