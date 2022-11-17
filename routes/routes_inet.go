package routes

import (
	c "Gofiber_test/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func InetRoute(app *fiber.App) {

	api := app.Group("/api")  // /api
	v1 := api.Group("/v1")    // /api/v1
	user := v1.Group("/user") // api/v1/user

	user.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"testgo": "17112565",
		},
	}))

	v1.Get("/", c.HelloWorld)

	v1.Get("/test", c.HelloAom)

	v1.Post("/", c.TestBodyParser)

	v1.Get("/users/:name", c.TestParams)

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

	user.Post("", c.AddUser) //adduser
	// user.Get("/", c.GetUser)             //getuser/Read
	user.Put("/:id", c.UpdateUser)    //updateuser
	user.Delete("/:id", c.DeleteUser) //deleteuser
	user.Get("/sum", c.GetGENage)
	user.Get("/search", c.GetUserFilter) //getuser filter/Read

	//##########################################

	v2 := api.Group("/v2")
	v2.Get("/", c.GetUser) // api/v2/

	//##########################################

	// v2.Get("/", c.HelloV2)
	// v2.Post("/user", c.TestValidate)

	// v3 := api.Group("/v3")
	// v3.Get("/asc2/:tex_id", c.Asc2)

}
