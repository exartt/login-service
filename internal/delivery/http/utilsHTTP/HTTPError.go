package utilsHTTP

import "github.com/gofiber/fiber/v2"

func sendStatusMessage(c *fiber.Ctx, statusCode int, message string) error {
	return c.Status(statusCode).SendString(message)
}

func BadRequest(c *fiber.Ctx) error {
	return sendStatusMessage(c, 400, "Requisição inválida")
}

func Unauthorized(c *fiber.Ctx) error {
	return sendStatusMessage(c, 401, "Não autorizado")
}

func Forbidden(c *fiber.Ctx) error {
	return sendStatusMessage(c, 403, "Acesso proibido")
}

func NotFound(c *fiber.Ctx) error {
	return sendStatusMessage(c, 404, "Página não encontrada")
}

func InternalServerError(c *fiber.Ctx) error {
	return sendStatusMessage(c, 500, "Erro interno do servidor")
}

func Ok(c *fiber.Ctx) error {
	return sendStatusMessage(c, 200, "Tudo certo")
}

func Accepted(c *fiber.Ctx) error {
	return sendStatusMessage(c, 202, "Requisição aceita")
}

func NoContent(c *fiber.Ctx) error {
	return sendStatusMessage(c, 204, "Sem conteúdo")
}

func NewError(c *fiber.Ctx, status int, err error) error {
	httpError := HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	return c.Status(status).JSON(httpError)
}

type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}
