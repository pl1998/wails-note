package menu

import (
	"changeme/app/handler"
	"changeme/app/model/menu"
	"changeme/app/requests"
	"changeme/pkg/helpers"
	"changeme/pkg/model"
	response "changeme/pkg/respone"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"time"
)

// MenuHandler menu handler struct
type MenuHandler struct {
}

// Index menu list func
func (*MenuHandler) Index(cxt *gin.Context) {
	var menuList []menu.NoteMenus
	model.DB.
		Model(&menu.NoteMenus{}).
		//Preload("NoteList").
		Where("is_deleted =?", 0).
		Order("sort desc").
		Order("sort_time desc").
		Find(&menuList)
	response.
		SuccessResponse(GetMenuTree(menuList, 0)).
		ToJson(cxt)
	return
}

// Store create note menu
func (*MenuHandler) Store(cxt *gin.Context) {
	idDir := helpers.StringToInt(cxt.DefaultPostForm("is_dir", "0"))
	params := requests.NoteMenuStoreForm{
		Name:    cxt.PostForm("name"),
		IsDir:   idDir,
		Pid:     uint64(helpers.StringToInt(cxt.DefaultPostForm("p_id", "0"))),
		Content: cxt.DefaultPostForm("content", ""),
	}
	errs := validator.New().Struct(params)
	if errs != nil {
		response.ErrorResponse(http.StatusInternalServerError, errs.Error()).ToJson(cxt)
		return
	}
	var note_menu menu.NoteMenus
	note_menu.PId = params.Pid
	note_menu.Name = params.Name
	note_menu.IsDir = params.IsDir
	note_menu.Content = params.Content
	note_menu.AddTime = time.Now().Unix()
	if model.DB.Model(&menu.NoteMenus{}).Create(&note_menu).Error != nil {
		response.ErrorResponse(http.StatusInternalServerError, "更新失败").ToJson(cxt)
		return
	}
	response.
		SuccessResponse(note_menu).
		ToJson(cxt)
}

// Delete menu
func (*MenuHandler) Delete(cxt *gin.Context) {
	id := handler.GetId(cxt)
	var note_menu menu.NoteMenus
	if result := model.DB.Model(&menu.NoteMenus{}).Where("menu_id=?", id).Find(&note_menu); result.RowsAffected == 0 {
		response.FailResponse(http.StatusInternalServerError, "删除失败！").WriteTo(cxt)
		return
	}
	response.SuccessResponse().WriteTo(cxt)
	return
}

// Update menu
func (*MenuHandler) Update(cxt *gin.Context) {
	params := requests.NoteMenuUpdateForm{
		Name:   cxt.PostForm("name"),
		IsDir:  helpers.StringToInt(cxt.PostForm("is_dir")),
		Pid:    uint64(helpers.StringToInt(cxt.PostForm("p_id"))),
		MenuId: uint64(helpers.StringToInt(cxt.PostForm("menu_id"))),
	}
	errs := validator.New().Struct(params)
	if errs != nil {
		response.ErrorResponse(http.StatusInternalServerError, errs.Error()).ToJson(cxt)
		return
	}
	var note_menu menu.NoteMenus
	note_menu.PId = params.Pid
	note_menu.Name = params.Name
	note_menu.IsDir = params.IsDir
	note_menu.UpdateTime = time.Now().Unix()
	if model.DB.Model(&menu.NoteMenus{}).Where("menu_id =? ", params.MenuId).Updates(&note_menu).Error != nil {
		response.ErrorResponse(http.StatusInternalServerError, "更新失败").ToJson(cxt)
		return
	}
	response.
		SuccessResponse(note_menu).
		ToJson(cxt)
}

// GetMenuTree get tree data structure
func GetMenuTree(menus []menu.NoteMenus, pid uint64) any {
	var list []menu.NoteMenus
	for _, v := range menus {
		if v.PId == pid {
			v.Children = GetMenuTree(menus, v.MenuId)
			list = append(list, v)
		}
	}
	return list
}
