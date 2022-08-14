package su

import (
	"fmt"
	"web-api/config/db"
	"web-api/model/su"

	"github.com/gofiber/fiber/v2"
)

func CreatePage(c *fiber.Ctx) error {
	req := new(su.SuPage)
	if err := c.BodyParser(req); err != nil {
		fmt.Println(err)
	}
	db.Create(req, &req)
	return c.Status(fiber.StatusOK).JSON(req)
}

func ListPage(c *fiber.Ctx) error {
	v := new([]su.SuPage)
	db.Read(v, &v, v)
	return c.Status(fiber.StatusOK).JSON(v)
}

func UpdatePage(c *fiber.Ctx) error {
	req := new(su.SuPage)
	if err := c.BodyParser(req); err != nil {
		fmt.Println(err)
	}
	db.Update(req, &req, su.SuPage{PageCode: req.PageCode})
	return c.Status(fiber.StatusOK).JSON(req)
}

func DeletePage(c *fiber.Ctx) error {
	req := c.Query("code")
	v := new(su.SuPage)
	db.Delete(v, &v, su.SuPage{PageCode: req})
	return c.Status(fiber.StatusOK).JSON(v)
}

func GetParameter(c *fiber.Ctx) error {
	req := c.Query("code")
	v := new(su.SuParam)
	db.Read(v, &v, su.SuParam{ParamCode: req})
	return c.Status(fiber.StatusOK).JSON(v)
}
