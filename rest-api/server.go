package main

import (
	"fmt"
	"net/http"

	"./controller"
	router "./http"
	"./repository"
	"./services"
)

var (
	userRepo       repository.UserRepo       = repository.NewUserRepository()
	userService    services.UserService      = services.NewUserService(userRepo)
	userController controller.UserController = controller.NewUserController(userService)
	httpRouter     router.Router             = router.NewMuxRouter()
)

func main() {

	const port string = ":8080"
	httpRouter.GET("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Println(res, "up and running")
	})
	httpRouter.POST("/addUser", userController.CreateUser)
	httpRouter.GET("/getUser", userController.GetUser)
	//httpRouter.POST("/posts", addPost)
	httpRouter.SERVE(port)
}
