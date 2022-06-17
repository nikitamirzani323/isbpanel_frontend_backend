package controllers

import (
	"log"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
)

func Gameslothome(c *fiber.Ctx) error {
	type payload_fetch struct {
		Page            string `json:"page" `
		Providerslot_id int    `json:"providerslot_id" `
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_fetch)
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
			"providerslot_id": client.Providerslot_id,
		}).
		Post(PATH + "api/prediksislot")
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
func GameslotSave(c *fiber.Ctx) error {
	type payload_save struct {
		Page                  string `json:"page"`
		Sdata                 string `json:"sdata" `
		Prediksislot_id       int    `json:"prediksislot_id"`
		Providerslot_id       int    `json:"providerslot_id" `
		Providerslot_slug     string `json:"providerslot_slug" `
		Prediksislot_name     string `json:"prediksislot_name" `
		Prediksislot_prediksi int    `json:"prediksislot_prediksi"`
		Prediksislot_image    string `json:"prediksislot_image" `
		Prediksislot_status   string `json:"prediksislot_status" `
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
			"client_hostname":       hostname,
			"page":                  client.Page,
			"sdata":                 client.Sdata,
			"prediksislot_id":       client.Prediksislot_id,
			"providerslot_id":       client.Providerslot_id,
			"providerslot_slug":     client.Providerslot_slug,
			"prediksislot_name":     client.Prediksislot_name,
			"prediksislot_prediksi": client.Prediksislot_prediksi,
			"prediksislot_image":    client.Prediksislot_image,
			"prediksislot_status":   client.Prediksislot_status,
		}).
		Post(PATH + "api/prediksislotsave")
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
func GameslotDelete(c *fiber.Ctx) error {
	type payload_save struct {
		Page              string `json:"page"`
		Prediksislot_id   int    `json:"prediksislot_id" `
		Providerslot_id   int    `json:"providerslot_id" `
		Providerslot_slug string `json:"providerslot_slug" `
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
			"client_hostname":   hostname,
			"page":              client.Page,
			"prediksislot_id":   client.Prediksislot_id,
			"providerslot_id":   client.Providerslot_id,
			"providerslot_slug": client.Providerslot_slug,
		}).
		Post(PATH + "api/prediksislotdelete")
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
func GameslotGenerator(c *fiber.Ctx) error {
	type payload_save struct {
		Page              string `json:"page"`
		Providerslot_id   int    `json:"providerslot_id" `
		Providerslot_slug string `json:"providerslot_slug" `
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
			"client_hostname":   hostname,
			"page":              client.Page,
			"providerslot_id":   client.Providerslot_id,
			"providerslot_slug": client.Providerslot_slug,
		}).
		Post(PATH + "api/prediksislotgenerator")
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
