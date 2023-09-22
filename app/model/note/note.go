package note

type Notes struct {
	NoteId     uint64 `gorm:"column:note_id;primaryKey;autoIncrement;" json:"note_id"`
	Name       string `gorm:"column:name" json:"name"`
	Desc       string `gorm:"column:desc" json:"desc"`
	AddTime    int64  `gorm:"column:add_time" json:"add_time"`
	UpdateTime int64  `gorm:"column:update_time" json:"update_time"`
	IsDeleted  int    `gorm:"column:is_deleted" json:"is_deleted"`
}
