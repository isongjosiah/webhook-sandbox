package main

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName: "ec-webhook-sandbox",
	})

	webhookGroup := app.Group("/webhook")
	webhookGroup.Post("/ec-notification", func(ctx *fiber.Ctx) error {
		r := ctx.Request()
		res := ctx.Response()

		var payload map[string]interface{}
		unMarshalErr := json.Unmarshal(r.Body(), &payload)
		if unMarshalErr != nil {
			return unMarshalErr
		}

		fmt.Println("webhook payload", payload)
		res.SetStatusCode(http.StatusOK)
		return nil
	})
}
