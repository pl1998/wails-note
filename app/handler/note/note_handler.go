package note

import (
	"changeme/app/handler"
	"changeme/app/model/note"
	"changeme/app/requests"
	"changeme/pkg/helpers"
	"changeme/pkg/model"
	response "changeme/pkg/respone"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"time"
)

type NoteHandler struct {
}

// Index
func (*NoteHandler) Index(cxt *gin.Context) {
	id := handler.GetId(cxt)
	var notes note.Notes
	if result := model.DB.
		Model(&note.Notes{}).
		Where("note_id=? and is_deleted = ?", id, 0).
		First(&notes); result.RowsAffected == 0 {
		response.FailResponse(http.StatusInternalServerError, "数据不存在！").WriteTo(cxt)
		return
	}
	response.SuccessResponse(notes).WriteTo(cxt)
	return

}

// Store
func (*NoteHandler) Store(cxt *gin.Context) {
	params := requests.NoteStoreForm{
		Name:    cxt.PostForm("name"),
		Content: cxt.PostForm("content"),
		MenuId:  helpers.StringToInt(cxt.PostForm("menu_id")),
	}

	errs := validator.New().Struct(params)
	if errs != nil {
		response.ErrorResponse(http.StatusInternalServerError, errs.Error()).ToJson(cxt)
		return
	}
	var notes note.Notes
	notes.Content = params.Content
	notes.Name = params.Name
	notes.MenuId = params.MenuId
	notes.AddTime = time.Now().Unix()
	notes.UpdateTime = time.Now().Unix()
	if model.DB.Model(&note.Notes{}).Create(&notes).Error != nil {
		response.ErrorResponse(http.StatusInternalServerError, "添加失败").ToJson(cxt)
		return
	}
	response.
		SuccessResponse(notes).
		ToJson(cxt)
	return
}

// Update
func (*NoteHandler) Update(cxt *gin.Context) {
	id := handler.GetId(cxt)
	params := requests.NoteStoreForm{
		Name:    cxt.PostForm("name"),
		Content: cxt.PostForm("desc"),
	}
	errs := validator.New().Struct(params)
	if errs != nil {
		response.ErrorResponse(http.StatusInternalServerError, errs.Error()).ToJson(cxt)
		return
	}
	var notes note.Notes
	notes.Name = params.Name
	notes.Content = params.Content
	notes.UpdateTime = time.Now().Unix()
	if model.DB.Model(&note.Notes{}).Where("note_id =? ", id).Updates(&notes).Error != nil {
		response.ErrorResponse(http.StatusInternalServerError, "更新失败").ToJson(cxt)
		return
	}
	response.
		SuccessResponse(notes).
		ToJson(cxt)
	return
}

// Delete
func (*NoteHandler) Delete(cxt *gin.Context) {
	id := handler.GetId(cxt)
	var notes note.Notes
	model.DB.Model(&note.Notes{}).Where("note_id =? ", id).Delete(&notes)
	response.
		SuccessResponse(notes).
		ToJson(cxt)
	return
}
