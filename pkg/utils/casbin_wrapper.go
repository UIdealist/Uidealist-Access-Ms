package utils

import (
	"github.com/casbin/casbin/v2"
	"github.com/gofiber/fiber/v2"
)

// WrapWithCasbinEnforcer wraps the fiber handler with Casbin enforcer.
func WrapWithCasbinEnforcer(e *casbin.Enforcer, f func(*casbin.Enforcer, *fiber.Ctx) error) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return f(e, c)
	}
}
