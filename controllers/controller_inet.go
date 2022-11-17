package controllers

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"Gofiber_test/database"
	m "Gofiber_test/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func HelloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func HelloV2(c *fiber.Ctx) error {
	return c.SendString("Hello, World! this is v2")
}

func HelloAom(c *fiber.Ctx) error {
	return c.SendString("Hello, Aomsin!")
}

func TestBodyParser(c *fiber.Ctx) error {
	p := new(m.Person)
	// var p m.Person

	if err := c.BodyParser(p); err != nil {
		return err
	}

	log.Println(p.Name) // john
	log.Println(p.Pass) // doe

	str := p.Name + " " + p.Pass
	return c.SendString(str)
}

func TestParams(c *fiber.Ctx) error {

	str := "Hello ---> " + c.Params("name")
	return c.SendString(str)

}

func TestQuery(c *fiber.Ctx) error {
	a := c.Query("search")
	name := c.Query("name")
	str := "my search is  " + a + " my name is " + name
	return c.JSON(str)
}

func TestValidate(c *fiber.Ctx) error {
	//Connect to database

	user := new(m.User)
	// var p m.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	validate := validator.New()
	errors := validate.Struct(user)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors.Error())
	}
	return c.JSON(user)
}

func Factorial(c *fiber.Ctx) error {
	// return c.SendString("Hello Factorail")
	num := c.Params("number")
	str1, err := strconv.Atoi(num)
	if err != nil {
		panic(err)
	}
	sum := 1
	for i := str1; i >= 1; i-- {
		sum = sum * i

	}
	strsum := strconv.Itoa(sum)
	return c.SendString(strsum)

}

func Asc2(c *fiber.Ctx) error {
	c.Query("tax_id") // "fenny"
	a := c.Query("tax_id")
	runes := []rune(a)
	text := ""

	var result []int
	result1 := ""

	for i := 0; i < len(runes); i++ {
		result = append(result, int(runes[i]))
		result1 = strconv.Itoa(result[i])
		text = text + " " + result1
		fmt.Println(result1)
	}
	text1 := text
	fmt.Println(text)
	return c.JSON(text1)
}

func GetDogs(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Find(&dogs)
	return c.Status(200).JSON(dogs)
}

func GetDog(c *fiber.Ctx) error {
	db := database.DBConn
	search := strings.TrimSpace(c.Query("search"))
	var dog []m.Dogs

	result := db.Find(&dog, "dog_id = ?", search)

	// returns found records count, equals `len(users)
	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON(&dog)
}

func AddDog(c *fiber.Ctx) error {
	//twst3
	db := database.DBConn
	var dog m.Dogs

	if err := c.BodyParser(&dog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Create(&dog)
	return c.Status(201).JSON(dog)
}

func UpdateDog(c *fiber.Ctx) error {
	db := database.DBConn
	var dog m.Dogs
	id := c.Params("id")

	if err := c.BodyParser(&dog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Where("id = ?", id).Updates(&dog)
	return c.Status(200).JSON(dog)
}

func RemoveDog(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var dog m.Dogs

	result := db.Delete(&dog, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}

func GetDogsJson(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Find(&dogs)

	var dataResults []m.DogsRes
	for _, v := range dogs {

		typeStr := ""
		if v.DogID == 111 {
			typeStr = "red"
		} else if v.DogID == 113 {
			typeStr = "green"
		} else if v.DogID == 999 {
			typeStr = "pink"
		} else {
			typeStr = "no color"
		}

		d := m.DogsRes{
			Name:  v.Name,  //coco
			DogID: v.DogID, //199
			Type:  typeStr, //no color
		}
		dataResults = append(dataResults, d)
		// sumAmount += v.Amount
	}

	r := m.ResultData{
		Data:  dataResults,
		Name:  "golang-test",
		Count: len(dogs), //หาผลรวม,
	}
	return c.Status(200).JSON(r)
}

//################################### 7.2 ##################################################

func GetDogsJson_Ex(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	sred := 0
	sgreen := 0
	spink := 0
	nocolor := 0

	db.Find(&dogs)

	var dataResults []m.DogsRes
	for _, v := range dogs {

		typeStr := ""
		if v.DogID >= 10 && v.DogID <= 50 {
			typeStr = "red"
			sred += 1

		} else if v.DogID >= 100 && v.DogID <= 150 {
			typeStr = "green"
			sgreen += 1

		} else if v.DogID >= 200 && v.DogID <= 250 {
			typeStr = "pink"
			spink += 1

		} else {
			typeStr = "no color"
			nocolor += 1
		}

		d := m.DogsRes{
			Name:  v.Name,  //coco
			DogID: v.DogID, //199
			Type:  typeStr, //no color
		}
		dataResults = append(dataResults, d)
		// sumAmount += v.Amount
	}

	r := m.ResultData{
		Data:        dataResults,
		Name:        "golang-test",
		Count:       len(dogs), //หาผลรวม,
		Sum_red:     sred,
		Sum_green:   sgreen,
		Sum_pink:    spink,
		Sum_nocolor: nocolor,
	}
	return c.Status(200).JSON(r)
}

//############################################# 7.1 ####################################

func AmountGreaterThan1000(db *gorm.DB) *gorm.DB {
	return db.Where("dog_id >= ? && dog_id <= ?", 50, 100)
}

func GetAmount(c *fiber.Ctx) error {

	db := database.DBConn
	var dogs []m.Dogs

	db.Scopes(AmountGreaterThan1000).Find(&dogs)
	return c.Status(200).JSON(dogs)
}

//############################################ add user ####################################

func AddUser(c *fiber.Ctx) error {
	//twst3
	db := database.DBConn
	var user m.Users

	if err := c.BodyParser(&user); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Create(&user)
	return c.Status(201).JSON(user)
}

//########################################## get user ######################################

func GetUser(c *fiber.Ctx) error {
	db := database.DBConn
	var users []m.Users

	db.Find(&users)
	return c.Status(200).JSON(users)
}

//########################################## update user ######################################

func UpdateUser(c *fiber.Ctx) error {
	db := database.DBConn
	var user m.Users
	id := c.Params("id")

	if err := c.BodyParser(&user); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Where("employee_id = ?", id).Updates(&user)
	return c.Status(200).JSON(user)
}

//########################################## delete user ######################################

func DeleteUser(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var user m.Users

	result := db.Delete(&user, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)

}

//########################################## get user filter ######################################

func GetUserFilter(c *fiber.Ctx) error {
	db := database.DBConn
	search := strings.TrimSpace(c.Query("search"))
	var user []m.Users

	result := db.Find(&user, "employee_id = ? || name = ? || lastname = ? ", search)

	// returns found records count, equals `len(users)
	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON(&user)
}

//########################################## get user filter ######################################

func GetGENage(c *fiber.Ctx) error {
	db := database.DBConn
	var user []m.Users

	genZ := 0
	genY := 0
	genX := 0
	baby_boomer := 0
	gi_generation := 0

	db.Find(&user)

	var dataResults []m.UsersRes
	for _, v := range user {

		genStr := ""
		if v.Age <= 23 {
			genStr = "GenZ"
			genZ += 1

		} else if v.Age >= 24 && v.Age <= 41 {
			genStr = "GenY"
			genY += 1

		} else if v.Age >= 42 && v.Age <= 56 {
			genStr = "GenX"
			genX += 1

		} else if v.Age >= 57 && v.Age <= 75 {
			genStr = "Baby Boomer"
			baby_boomer += 1
		} else if v.Age > 76 {
			genStr = "G.I. Generation"
			gi_generation += 1
		}

		d := m.UsersRes{
			Employee_id: v.Employee_id,
			Name:        v.Name, //coco
			Age:         v.Age,  //199
			Generation:  genStr, //no color
		}
		dataResults = append(dataResults, d)
		// sumAmount += v.Amount
	}

	r := m.User_Generation{
		Data:          dataResults,
		Count:         len(user), //หาผลรวม,
		GenZ:          genZ,
		GenY:          genY,
		GenX:          genX,
		Baby_Boomer:   baby_boomer,
		GI_Generation: gi_generation,
	}
	return c.Status(200).JSON(r)
}
