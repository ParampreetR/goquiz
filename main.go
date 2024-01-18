package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/template/handlebars/v2"
)

func main() {
	engine := handlebars.New("./views", ".hbs")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/static", "./public")
	app.Use(encryptcookie.New(encryptcookie.Config{
		Key: encryptcookie.GenerateKey(),
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/login")
	})

	app.Get("/login", func(c *fiber.Ctx) error {
		return c.Render("login", fiber.Map{})
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		user := User{Name: ""}
		if c.BodyParser(&user) != nil {
			return c.SendString("Error")
		}

		c.Cookie(&fiber.Cookie{
			Name:  "name",
			Value: user.Name,
		})
		return c.Redirect("/quiz")
	})

	app.Get("/quiz", func(c *fiber.Ctx) error {
		if c.Cookies("name", "") == "" {
			c.Redirect("/login")
		}

		return c.Render("quiz", fiber.Map{
			"Questions": GetQuestionsFromDatabase(),
		})
	})

	app.Post("/submit", func(c *fiber.Ctx) error {
		body := c.Body()
		var quiz_answers_raw []string
		var quiz_answers map[string]string = make(map[string]string)
		quiz_answers_raw = strings.Split(string(body[:]), "&")
		fmt.Println(string(body[:]))

		for _, val := range quiz_answers_raw {
			fmt.Println(strings.Split(val, "="))
			if strings.Split(val, "=")[1] != "" {
				quiz_answers[strings.Split(val, "=")[0]] = strings.Split(val, "=")[1]
			}
		}

		for key, val := range quiz_answers {
			fmt.Println(key + "\t" + val)
		}

		questions := GetQuestionsFromDatabase()
		result := CheckAnswers(quiz_answers, questions)

		fmt.Println(result)

		return c.Redirect("/result")
	})

	app.Get("/add_question", func(c *fiber.Ctx) error {
		return c.Render("add_question", fiber.Map{})
	})

	app.Post("/add_question", func(c *fiber.Ctx) error {
		var question_form FormQuestion = FormQuestion{}
		err := c.BodyParser(&question_form)
		if err != nil {
			fmt.Println(err)
			log.Fatalf("Error in /add_question")
			return c.SendString("Error")
		}

		fmt.Println(question_form)
		SaveQuestionToDatabase(question_form)
		return c.Redirect("/list_questions")
	})

	app.Get("/list_questions", func(c *fiber.Ctx) error {
		questions := GetQuestionsFromDatabase()
		return c.Render("list_questions", fiber.Map{
			"Questions": questions,
		})
	})

	app.Post("/delete_question/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))

		if err != nil {
			log.Fatalf("Error: %s", err.Error())
			return c.SendString("Error")
		}
		DeleteQuestionFromDatabase(id)
		return c.Redirect("/list_questions")
	})

	app.Post("/logout", func(c *fiber.Ctx) error {
		c.ClearCookie("name")
		return c.Redirect("/")
	})

	log.Fatal(app.Listen(":3000"))
}
