package controllers

import (
	"frankie/universal/sdk/model"
	"frankie/universal/sdk/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type validation struct {
}

//NewValidation init validation
func NewValidation() *validation {
	return &validation{}
}

//Validate used as controller for validate data
func (v *validation) Validate(c *gin.Context) {
	var paramInput []model.InputValidation
	if err := c.ShouldBindJSON(&paramInput); err != nil {
		msg := utils.SetError(utils.IncorrectParam, err)
		c.AbortWithStatusJSON(msg.StatusCode, msg)
		return
	}

	//validate input
	for i := range paramInput {
		if err := paramInput[i].Validate(); err != nil {
			c.AbortWithStatusJSON(err.StatusCode, err)
			return
		}
	}

	data := map[string]interface{}{
		"puppy": true,
	}
	c.JSON(http.StatusOK, data)
}
