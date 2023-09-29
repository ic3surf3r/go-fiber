package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	// Load templates
	engine := html.New("./views", ".tmpl")

	// Create fibre app
    app := fiber.New(fiber.Config{
        Views: engine,
    })


	// Serve Static
	app.Static("/", "./public") 

	// Middleware
	m := (func(c *fiber.Ctx) error {
		fmt.Println("I'm an amendobobo yeah")

		return c.Next()
	})

	// Add routes
	pages := app.Group("/pages", m)

    pages.Get("/", func(c *fiber.Ctx) error {
		// return fiber.NewError(fiber.StatusServiceUnavailable, "on vacation")
		return c.JSON(struct {
			Name string
			Age int
		}{
			Name: "Bob",
			Age: 42,
		})

        // return c.Render("index", fiber.Map{
        //     "Name": "Bob",
        // })
    })
    pages.Get("/names/:name/age/:age", func(c *fiber.Ctx) error {
		name := c.Params("name")
		age := c.Params("age")
        return c.Render("index", fiber.Map{
            "Name": name,
			"Age": age,
        })
    })

	app.Post("/", func(c *fiber.Ctx) error {
		var body struct {
			Message string
		}

        if err := c.BodyParser(&body); err != nil {
            return err
        }

		return c.Render("index", fiber.Map{
            "Name": "Bob",
			"Message" : body.Message,
        })
	})

	// Start app
    app.Listen(":3000")
}