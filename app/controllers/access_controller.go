package controllers

import (
	"uidealist/app/crud"
	"uidealist/pkg/repository"

	"github.com/casbin/casbin/v2"
	"github.com/gofiber/fiber/v2"
)

// CheckAccess to check if user is allowed to access a resource.
// @Description Checks against a RBAC policy if a user is allowed to access a resource.
// @Summary Check if resource is accessible by user
// @Tags Access
// @Accept json
// @Produce json
// @Param policies body crud.AccessList true "Policies to check"
// @Success 200 {string} status "ok"
// @Router /v1/check [post]
func CheckAccess(e *casbin.Enforcer, c *fiber.Ctx) error {

	// Parse policies from request body.
	policies := new(crud.AccessList[interface{}])
	if err := c.BodyParser(policies); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(crud.BaseResponse{
			Error:   true,
			Code:    repository.INVALID_REQUEST,
			Message: err.Error(),
		})
	}

	// Check if user is allowed to access resources.
	allowed, err := e.BatchEnforce(policies.ParseToArray())
	if err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(crud.BaseResponse{
			Error:   true,
			Code:    repository.ACCESS_ERROR,
			Message: err.Error(),
		})
	}

	// Return status 200 and success message.
	return c.Status(fiber.StatusOK).JSON(crud.CheckResponse{
		BaseResponse: crud.BaseResponse{
			Error:   false,
			Code:    repository.AUTHORIZED,
			Message: "Access granted",
		},
		Data: allowed,
	})
}

// GrantAccess to grant to a resource.
// @Description Alters a RBAC policy to grant access to a resource.
// @Summary Grant access to resource
// @Tags Access
// @Accept json
// @Produce json
// @Param policies body crud.AccessList true "Policies to grant"
// @Success 200 {string} status "ok"
// @Router /v1/grant [post]
func GrantAccess(e *casbin.Enforcer, c *fiber.Ctx) error {
	// Parse policies from request body.
	policies := new(crud.AccessList[string])
	if err := c.BodyParser(policies); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(crud.BaseResponse{
			Error:   true,
			Code:    repository.INVALID_REQUEST,
			Message: err.Error(),
		})
	}

	// Grant access to user.
	_, err := e.AddPolicies(policies.ParseToArray())
	if err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(crud.BaseResponse{
			Error:   true,
			Code:    repository.ACCESS_ERROR,
			Message: err.Error(),
		})
	}

	// Return status 200 and success message.
	return c.Status(fiber.StatusOK).JSON(crud.BaseResponse{
		Error:   false,
		Code:    repository.SUCCESS,
		Message: "Access granted",
	})
}

// RevokeAccess to grant to a resource.
// @Description Alters a RBAC policy to revoke access to a resource.
// @Summary Revoke access to resource
// @Tags Access
// @Accept json
// @Produce json
// @Param policies body crud.AccessList true "Policies to revoke"
// @Success 200 {string} status "ok"
// @Router /v1/revoke [post]
func RevokeAccess(e *casbin.Enforcer, c *fiber.Ctx) error {
	// Parse policies from request body.
	policies := new(crud.AccessList[string])
	if err := c.BodyParser(policies); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(crud.BaseResponse{
			Error:   true,
			Code:    repository.INVALID_REQUEST,
			Message: err.Error(),
		})
	}

	// Grant access to user.
	_, err := e.RemovePolicies(policies.ParseToArray())
	if err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(crud.BaseResponse{
			Error:   true,
			Code:    repository.ACCESS_ERROR,
			Message: err.Error(),
		})
	}

	// Return status 200 and success message.
	return c.Status(fiber.StatusOK).JSON(crud.BaseResponse{
		Error:   false,
		Code:    repository.SUCCESS,
		Message: "Access revoked",
	})
}

// UpdateAccess to alter access to a resource.
// @Description Alters a RBAC policy to update access to a resource.
// @Summary Update access to resource
// @Tags Access
// @Accept json
// @Produce json
// @Param policies body crud.AccessList true "Policies to update"
// @Success 200 {string} status "ok"
// @Router /v1/update [post]
func UpdateAccess(e *casbin.Enforcer, c *fiber.Ctx) error {
	// Parse policies from request body.
	policies := new(crud.AccessListUpdate)
	if err := c.BodyParser(policies); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(crud.BaseResponse{
			Error:   true,
			Code:    repository.INVALID_REQUEST,
			Message: err.Error(),
		})
	}

	// Grant access to user.
	_, err := e.UpdatePolicies(policies.ParseToArray())
	if err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(crud.BaseResponse{
			Error:   true,
			Code:    repository.ACCESS_ERROR,
			Message: err.Error(),
		})
	}

	// Return status 200 and success message.
	return c.Status(fiber.StatusOK).JSON(crud.BaseResponse{
		Error:   false,
		Code:    repository.SUCCESS,
		Message: "Access revoked",
	})
}
