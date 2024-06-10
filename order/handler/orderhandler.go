package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"order-service/database"
	"order-service/model/entities"

	"github.com/gofiber/fiber/v2"
)

func CreateOrder(c *fiber.Ctx) error {
	order := new(entities.Order)
	if err := c.BodyParser(order); err != nil {
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := checkProductAvailability(order.CatalogId); err != nil {
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&order)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data":    order,
	})
}

func FindOrders(c *fiber.Ctx) error {
	var order []entities.Order
	database.DB.Db.Find(&order)
	return c.JSON(order)
}

func checkProductAvailability(CatalogId uint) error {
	resp, err := http.Get(fmt.Sprintf("http://catalog-service:5001/products/%d", CatalogId))
	if err != nil {
		return errors.New("failed to connect to catalog service")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("product not available")
	}

	var product entities.Product
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &product)

	if product.ID == 0 {
		return errors.New("product not found")
	}

	return nil
}

func CheckCustomerAvailability(CustomerId uint) error {
	resp, err := http.Get(fmt.Sprintf("http://customer-service:5002/customers/%d", CustomerId))
	if err != nil {
		return errors.New("Failed to connect to Customer Service")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("Customer not logged in")
	}

	var customer entities.Customer
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &customer)

	if customer.ID == 0 {
		return errors.New("Customer not found")
	}

	return nil
}
