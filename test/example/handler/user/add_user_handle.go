package user

import (
	"github.com/MasterJoyHunan/gengin/test/example/internal/response"
	"github.com/MasterJoyHunan/gengin/test/example/logic/user"
	userType "github.com/MasterJoyHunan/gengin/test/example/types/user"

	"github.com/gin-gonic/gin"
)

func AddUserHandle(c *gin.Context) {
	var req userType.UserRequest
	if err := c.ShouldBindUri(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	if err := c.ShouldBind(&req); err != nil {
		response.HandleResponse(c, nil, err)
		return
	}

	resp, err := user.AddUser(&req)
	response.HandleResponse(c, resp, err)
}
