package handler

import (
	"UAS-micro/database"
	"UAS-micro/model"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/streadway/amqp"
)

func GetAllCatalog(c *fiber.Ctx) error {
	var catalogues []model.Catalog
	database.DB.Db.Find(&catalogues)
	return c.Status(fiber.StatusOK).JSON(catalogues)
}

func GetCatalogById(c *fiber.Ctx) error {
	catalog := &model.Catalog{}
	id := c.Params("id")
	if err := database.DB.Db.First(catalog, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(catalog)
}

func CreateCatalog(c *fiber.Ctx) error {
	catalog := new(model.Catalog)
	if err := c.BodyParser(catalog); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	database.DB.Db.Create(&catalog)

	err := PublishCatalog(catalog)
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(catalog)
}

func PublishCatalog(catalog *model.Catalog) error {
	body, err := json.Marshal(catalog)
	if err != nil {
		return err
	}

	var ch *amqp.Channel
	err = ch.Publish(
		"",              // exchange
		"catalog_queue", // routing key
		false,           // mandatory
		false,           // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	return err
}
