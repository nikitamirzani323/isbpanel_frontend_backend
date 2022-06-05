package controllers

import (
	"log"
	"strings"
	"time"

	"bitbucket.org/isbtotogroup/isbpanel_frontend_backend/entities"
	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
)

func Providerslothome(c *fiber.Ctx) error {
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(entities.Home)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	log.Println("Hostname: ", hostname)
	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname": hostname,
			"page":            client.Page,
		}).
		Post(PATH + "api/slotprovider")
	if err != nil {
		log.Println(err.Error())
	}
	// log.Println("Response Info:")
	// log.Println("  Error      :", err)
	// log.Println("  Status Code:", resp.StatusCode())
	// log.Println("  Status     :", resp.Status())
	// log.Println("  Proto      :", resp.Proto())
	// log.Println("  Time       :", resp.Time())
	// log.Println("  Received At:", resp.ReceivedAt())
	// log.Println("  Body       :\n", resp)
	// log.Println()
	result := resp.Result().(*responsedefault)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":  result.Status,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func ProviderslotSave(c *fiber.Ctx) error {
	type payload_save struct {
		Page                 string `json:"page"`
		Sdata                string `json:"sdata" `
		Providerslot_id      int    `json:"providerslot_id"`
		Providerslot_name    string `json:"providerslot_name" `
		Providerslot_display int    `json:"providerslot_display" `
		Providerslot_image   string `json:"providerslot_image" `
		Providerslot_slug    string `json:"providerslot_slug" `
		Providerslot_title   string `json:"providerslot_title" `
		Providerslot_descp   string `json:"providerslot_descp" `
		Providerslot_status  string `json:"providerslot_status" `
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_save)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	log.Println("Hostname: ", hostname)
	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname":      hostname,
			"page":                 client.Page,
			"sdata":                client.Sdata,
			"providerslot_id":      client.Providerslot_id,
			"providerslot_name":    client.Providerslot_name,
			"providerslot_display": client.Providerslot_display,
			"providerslot_image":   client.Providerslot_image,
			"providerslot_slug":    client.Providerslot_slug,
			"providerslot_title":   client.Providerslot_title,
			"providerslot_descp":   client.Providerslot_descp,
			"providerslot_status":  client.Providerslot_status,
		}).
		Post(PATH + "api/slotprovidersave")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Response Info:")
	log.Println("  Error      :", err)
	log.Println("  Status Code:", resp.StatusCode())
	log.Println("  Status     :", resp.Status())
	log.Println("  Proto      :", resp.Proto())
	log.Println("  Time       :", resp.Time())
	log.Println("  Received At:", resp.ReceivedAt())
	log.Println("  Body       :\n", resp)
	log.Println()
	result := resp.Result().(*responsedefault)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":  result.Status,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
