package requests

type NoteMenuStoreForm struct {
	Name  string `json:"name" validate:"required,min=2,max=20"`
	Pid   uint64 `json:"p_id" validate:"required,number"`
	IsDir int    `json:"is_dir" validate:"required,number"`
}

type NoteMenuUpdateForm struct {
	Name   string `json:"name" validate:"required,min=2,max=20"`
	Pid    uint64 `json:"p_id" validate:"required,number"`
	IsDir  int    `json:"is_dir" validate:"required,number"`
	Sort   int    `json:"sort" validate:"number"`
	MenuId uint64 `json:"menu_id"  validate:"required,number"`
}
