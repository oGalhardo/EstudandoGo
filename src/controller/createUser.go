package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/oGalhardo/EstudandoGo/src/configuration/rest_err"
	"github.com/oGalhardo/EstudandoGo/src/model/request"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError(fmt.Sprintf("There are some incorrect filds, err=%s\n", err.Error()))
		c.JSON(restErr.Code, restErr)
		return
	}
	fmt.Println(userRequest)

}
