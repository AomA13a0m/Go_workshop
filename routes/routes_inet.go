package routes

import (
	c "Gofiber_test/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func InetRoute(app *fiber.App) {

	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"gofiber": "15112565",
		},
	}))

	api := app.Group("/api") // /api
	v1 := api.Group("/v1")   // /api/v1
	v1.Get("/", c.HelloWorld)

	v1.Get("/test", c.HelloAom)

	v1.Post("/", c.TestBodyParser)

	v1.Get("/user/:name", c.TestParams)

	v1.Post("/inet", c.TestQuery)

	v1.Post("/fact/:number", c.Factorial)

	//##########################################

	dog := v1.Group("/dog")

	dog.Get("/", c.GetDogs)

	dog.Get("/filter", c.GetDog)

	dog.Post("", c.AddDog)

	dog.Put("/:id", c.UpdateDog)

	dog.Delete("/:id", c.RemoveDog)

	dog.Get("/json", c.GetDogsJson)

	dog.Get("/rgb", c.GetDogsJson_Ex)

	dog.Get("/amount", c.GetAmount)

	//##########################################

	v2 := api.Group("/v2")
	v2.Get("/", c.HelloV2)

	v2.Post("/user", c.TestValidate)

	v3 := api.Group("/v3")
	v3.Get("/asc2/:tex_id", c.Asc2)

}
