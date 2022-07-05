package controllers

import (
	"log"
	"strings"
	"time"

	"bitbucket.org/isbtotogroup/isbpanel_frontend_backend/helpers"
	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
)

const Field_tafsirmimpihome_redis = "LISTTAFSIRMIMPI_BACKEND_ISBPANEL"

func Tafsirmimpihome(c *fiber.Ctx) error {
	type payload_tafsirmimpihome struct {
		Tafsirmimpi_search string `json:"tafsirmimpi_search"`
		Tafsirmimpi_page   int    `json:"tafsirmimpi_page"`
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_tafsirmimpihome)
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
		SetResult(helpers.Responsepaging{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname":    hostname,
			"tafsirmimpi_search": client.Tafsirmimpi_search,
			"tafsirmimpi_page":   client.Tafsirmimpi_page,
		}).
		Post(PATH + "api/tafsirmimpi")
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
	result := resp.Result().(*helpers.Responsepaging)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":      result.Status,
			"perpage":     result.Perpage,
			"totalrecord": result.Totalrecord,
			"message":     result.Message,
			"record":      result.Record,
			"time":        time.Since(render_page).String(),
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
func Tafsirmimpisave(c *fiber.Ctx) error {
	type payload_tafsirmimpisave struct {
		Page                  string `json:"page"`
		Sdata                 string `json:"sdata" `
		Tafsirmimpi_page      int    `json:"tafsirmimpi_page"`
		Tafsirmimpi_search    string `json:"tafsirmimpi_search"`
		Tafsirmimpi_id        int    `json:"tafsirmimpi_id"`
		Tafsirmimpi_mimpi     string `json:"tafsirmimpi_mimpi" `
		Tafsirmimpi_artimimpi string `json:"tafsirmimpi_artimimpi" `
		Tafsirmimpi_angka2d   string `json:"tafsirmimpi_angka2d" `
		Tafsirmimpi_angka3d   string `json:"tafsirmimpi_angka3d" `
		Tafsirmimpi_angka4d   string `json:"tafsirmimpi_angka4d" `
		Tafsirmimpi_status    string `json:"tafsirmimpi_status"`
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_tafsirmimpisave)
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
			"tafsirmimpi_page":      client.Tafsirmimpi_page,
			"tafsirmimpi_search":    client.Tafsirmimpi_search,
			"tafsirmimpi_id":        client.Tafsirmimpi_id,
			"tafsirmimpi_mimpi":     client.Tafsirmimpi_mimpi,
			"tafsirmimpi_artimimpi": client.Tafsirmimpi_artimimpi,
			"tafsirmimpi_angka2d":   client.Tafsirmimpi_angka2d,
			"tafsirmimpi_angka3d":   client.Tafsirmimpi_angka3d,
			"tafsirmimpi_angka4d":   client.Tafsirmimpi_angka4d,
			"tafsirmimpi_status":    client.Tafsirmimpi_status,
		}).
		Post(PATH + "api/tafsirmimpisave")
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
