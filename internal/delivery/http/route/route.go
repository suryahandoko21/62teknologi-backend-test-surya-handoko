package route

import (
	"be/internal/delivery/http"

	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App                 *fiber.App
	UserController      *http.UserController
	BusinesseController *http.BusinessController
	AuthMiddleware      fiber.Handler
}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoute()
	c.SetupAuthRoute()
}

/*public route can be access without login*/
func (c *RouteConfig) SetupGuestRoute() {
	c.App.Post("/api/users", c.UserController.Register)
	c.App.Post("/api/users/login", c.UserController.Login)

}

/*for access this route must login and attach token in header by name: Authorization*/
func (c *RouteConfig) SetupAuthRoute() {
	c.App.Use(c.AuthMiddleware)
	c.App.Get("/api/businesses/list", c.BusinesseController.List)
	c.App.Post("/api/businesses/create", c.BusinesseController.Create)
	c.App.Put("/api/businesses/update/:uuid", c.BusinesseController.Update)
	c.App.Delete("/api/businesses/delete/:uuid", c.BusinesseController.Delete)
}
