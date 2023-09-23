package handler

import (
	response "changeme/pkg/respone"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Person struct {
	ID string `uri:"id" binding:"required"`
}

func GetPersonId(cxt *gin.Context) (error, Person) {
	var person Person
	if err := cxt.ShouldBindUri(&person); err != nil {
		return err, person
	}
	return nil, person
}

func GetId(cxt *gin.Context) string {
	err, person := GetPersonId(cxt)
	if err != nil {
		response.FailResponse(http.StatusInternalServerError, "参数错误！").WriteTo(cxt)
	}
	return person.ID
}
