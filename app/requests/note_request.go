package requests

type NoteStoreForm struct {
	Name    string `json:"name" validate:"required,min=2,max=40"`
	Content string `json:"content" validate:"required"`
	MenuId  int    `json:"menu_id" validate:"required,number"`
}
