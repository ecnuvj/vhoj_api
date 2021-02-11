package handler

import (
	"github.com/ecnuvj/vhoj_api/model/contract"
	"github.com/ecnuvj/vhoj_api/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SubmitCode(c *gin.Context) {
	request := &contract.SubmitCodeRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, &contract.SubmitCodeResponse{
			BaseResponse: util.NewFailureResponse("request param error, err: %v", err),
		})
		return
	}

}

func ReSubmitCode(c *gin.Context) {

}

func ListSubmissions(c *gin.Context) {

}
