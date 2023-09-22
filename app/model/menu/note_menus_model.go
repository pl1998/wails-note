package menu

import "changeme/app/model/note"

type NoteMenus struct {
	MenuId     uint64       `gorm:"column:menu_id;primaryKey;autoIncrement;" json:"menu_id"`
	Name       string       `gorm:"column:name" json:"name"`
	PId        uint64       `gorm:"column:p_id" json:"p_id"`
	AddTime    int64        `gorm:"column:add_time" json:"add_time"`
	UpdateTime int64        `gorm:"column:add_time" json:"update_time"`
	IsDeleted  int          `gorm:"column:is_deleted" json:"is_deleted"`
	IsDir      int          `gorm:"column:is_dir" json:"is_dir"`
	Sort       int          `gorm:"column:sort" json:"sort"`
	SortTime   int          `gorm:"column:sort_time" json:"sort_time"`
	Children   any          `gorm:"-" json:"children,omitempty"`
	NoteList   []note.Notes `gorm:"foreignKey:NoteId;references:MenuId" json:"users"`
}
