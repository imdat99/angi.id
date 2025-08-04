package middlewares

import (
	"strings"

	"angi.account/config"
	"angi.account/modules/common"
	"angi.account/types"
	"github.com/gofiber/fiber/v2"
)

func Auth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var authHeader string
		authHeader = c.Get("Authorization")
		if authHeader == "" {
			authHeader = c.Cookies("Authorization")
		}
		token := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))

		if token == "" {
			return fiber.NewError(fiber.StatusUnauthorized, "Please authenticate")
		}

		payload, err := common.VerifyToken(token, config.Acfg.JWTSecret, types.TokenTypeAccessToken)
		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, "Please authenticate")
		}

		// user, err := userService.GetUserByID(c, payload.UserID)
		// if err != nil || user == nil {
		// 	return fiber.NewError(fiber.StatusUnauthorized, "Please authenticate")
		// }

		c.Locals("TokenPayload", payload)

		// Không Phân Quyền
		// if len(requiredRights) > 0 {
		// 	userRights, hasRights := config.RoleRights[user.Role]
		// 	if (!hasRights || !hasAllRights(userRights, requiredRights)) && c.Params("userId") != userID {
		// 		return fiber.NewError(fiber.StatusForbidden, "You don't have permission to access this resource")
		// 	}
		// }

		return c.Next()
	}
}
