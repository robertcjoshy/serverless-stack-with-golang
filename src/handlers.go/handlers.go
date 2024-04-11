package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	r "github.com/robert/serverless-stack-with-golang/src/models"
)

type createemail struct {
	MAIL string `json:"mail"`
}

type Response struct {
	Res string `json:"res"`
}

func getuser(c *gin.Context) {

	email := c.Param("email")

	email_result, err := r.GetEmail(email)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(201, email_result)

}

func createuser(c *gin.Context) {
	var req createemail

	err := c.BindJSON(&req)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	mail := req.MAIL

	created_email, err := r.SaveEmail(mail)

	c.JSON(201, created_email)

}

func updateuser() {

}

func deleteuser() {

}
