package validation

import (
	"errors"
	"frankie/universal/sdk/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type validation struct{}

//NewValidation init validation
func NewValidation() *validation {
	return &validation{}
}

/**
 * @api {GET} /web/:id Get Sample Data
 * @apiDescription Get Sample Data
 * @apiGroup Sample
 * @apiVersion 0.1.0
 * @apiPermission Authorize
 * @apiUse AuthorizeHeader
 * @apiExample {curl} Example usage:
 *   curl -i https://api.sample.com/web/:id
 * @apiSuccess {Int} id Sample ID
 * @apiSuccess {String} field_a Sample Field A
 * @apiSuccess {Float} field_b Sample Field B
 * @apiSuccess {Bool} field_c Sample Field C
 * @apiSuccess {Timestamp} created_on Time When Data Created
 * @apiSuccess {Timestamp} updated_on Time When Data Updated
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 Ok
 * {
 * 		"id": 3,
 *		"field_a": "value of field a",
 * 		"field_b": 88.88,
 * 		"field_c": true,
 * 		"created_on": "2018-01-09 12:00:00",
 * 		"updated_on": "2018-01-09 12:00:00"
 * }
 * @apiUse Error400
 * @apiUse Error401
 * @apiUse Error500
 */

//Validate used as controller for validate data
func (v *validation) Validate(c *gin.Context) {
	err := utils.SetError(utils.CallServicesFailed, errors.New("asd"))

	if err != nil {
		c.AbortWithStatusJSON(err.StatusCode, err)
		return
	}
	data := map[string]interface{}{
		"ping": "pong",
	}
	c.JSON(http.StatusCreated, data)
}
