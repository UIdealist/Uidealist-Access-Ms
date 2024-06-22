package routes

import (
	"uidealist/app/controllers"
	"uidealist/pkg/utils"

	"github.com/casbin/casbin/v2"
	"github.com/gofiber/fiber/v2"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(e *casbin.Enforcer, a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	// Routes for GET method:
	// TODO: Complete this controller to get permissions of user over a resource.
	// route.Get("/permissions/user/:usrId/resource/:resId", utils.WrapWithCasbinEnforcer(controllers.GetPermissions))

	// Routes for POST method:
	route.Post("/check", utils.WrapWithCasbinEnforcer(e, controllers.CheckAccess))
	route.Post("/grant", utils.WrapWithCasbinEnforcer(e, controllers.GrantAccess))
	route.Post("/revoke", utils.WrapWithCasbinEnforcer(e, controllers.RevokeAccess))

	// Routes for DELETE method:
	route.Post("/user/:usrId/resource/:resId/action/:actId", utils.WrapWithCasbinEnforcer(e, controllers.GrantAccess))
}
