package note

type Notes struct {
	NoteId     uint64 `gorm:"column:note_id;primaryKey;autoIncrement;" json:"note_id"`
	Name       string `gorm:"column:name" json:"name"`
	Content    string `gorm:"column:content" json:"content"`
	AddTime    int64  `gorm:"column:add_time" json:"add_time"`
	UpdateTime int64  `gorm:"column:update_time" json:"update_time"`
	IsDeleted  int    `gorm:"column:is_deleted" json:"is_deleted"`
	MenuId     int    `json:"menu_id" validate:"required,number"`
}
