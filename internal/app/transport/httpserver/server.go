package httpserver

import "movie-project/internal/app/services"

type AppServer struct {
	userService services.UserService
}

func NewAppServer(userService services.UserService) *AppServer {
	return &AppServer{
		userService: userService,
	}
}
