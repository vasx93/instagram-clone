package user

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/vasx93/instagram-clone/internal/tools"
)

// controller > service > repository
type UserController struct {
	service *UserService
}

func NewUserController(s *UserService, g *gin.Engine) *UserController {
	controller := &UserController{
		service: s,
	}

	g.GET("/users", controller.GetAll())

	return controller
}

func (uc *UserController) GetAll() gin.HandlerFunc {

	return func(c *gin.Context) {

		//do auth
		// unmarshal marshal if neded
		users, err := uc.service.GetAll()

		if err != nil {
			fmt.Println(err)
			tools.WriteError(c, 400)
			return

		}

		tools.WriteJSON(c, 200, users)
		// c.IndentedJSON(400, gin.H{"users": &users})
	}
}
