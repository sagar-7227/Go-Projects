package lead

import "gorm.io/gorm"
import "github.com/sagar-7227/go-fiber-crm-basic/database"
import "github.com/gofiber/fiber/v2"


type Lead struct {
	gorm.Model
	Name  string `json:"name"`
	Company string `json:"company"`
	Email string `json:"email"`
	Phone int `json:"phone"`
}

func GetLeads(c *fiber.Ctx) error {
	db := database.GetDB()
	var leads []Lead
	db.Find(&leads)
	return c.JSON(leads)
}

func NewLead(c *fiber.Ctx) error {
	db := database.GetDB()
	lead := new(Lead)
	if err := c.BodyParser(lead); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	db.Create(&lead)
	return c.JSON(lead)
}

func GetLead(c *fiber.Ctx) error {
	id := c.Params("id")
	var lead Lead
	db := database.GetDB()
	result := db.First(&lead, id)
	if result.Error != nil {
		return c.Status(404).SendString("No lead found with ID")
	}
	return c.JSON(lead)
}

func DeleteLead(c *fiber.Ctx) error {
	id := c.Params("id")
	var lead Lead
	db := database.GetDB()
	result := db.First(&lead, id)
	if result.Error != nil {
		return c.Status(404).SendString("No lead found with ID")
	}
	db.Delete(&lead)
	return c.SendString("Lead successfully deleted")
}