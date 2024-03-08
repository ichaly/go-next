package cms

import (
	"github.com/ichaly/go-next/lib/base"
)

type Interact struct {
	base.Primary `mapstructure:",squash"`
	ContentId    int64 `gorm:"comment:内容ID"`
	View         int   `gorm:"comment:阅读量"`
	Like         int   `gorm:"comment:点赞量"`
	Share        int   `gorm:"comment:分享量"`
	Comment      int   `gorm:"comment:评论量"`
}

func (Interact) TableName() string {
	return "cms_interact"
}
