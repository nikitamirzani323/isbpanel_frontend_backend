package controllers

import (
	"log"
	"strings"
	"time"

	"bitbucket.org/isbtotogroup/isbpanel_frontend_backend/helpers"
	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
)

const Fieldwebsiteagen_home_redis = "LISTWEBSITEAGEN_BACKEND_ISBPANEL"

func Websiteagenhome(c *fiber.Ctx) error {
	type payload_websiteagenhome struct {
		Websiteagen_search string `json:"websiteagen_search"`
		Websiteagen_page   int    `json:"websiteagen_page"`
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_websiteagenhome)
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
			"websiteagen_search": client.Websiteagen_search,
			"websiteagen_page":   client.Websiteagen_page,
		}).
		Post(PATH + "api/webagen")
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
func Websiteagensave(c *fiber.Ctx) error {
	type payload_websiteagensave struct {
		Page                 string `json:"page"`
		Sdata                string `json:"sdata" `
		Websiteagen_search   string `json:"websiteagen_search"`
		Websiteagen_page     int    `json:"websiteagen_page"`
		Websiteagen_idrecord int    `json:"websiteagen_id"`
		Websiteagen_name     string `json:"websiteagen_name" `
		Websiteagen_status   string `json:"websiteagen_status" `
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_websiteagensave)
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
			"client_hostname":    hostname,
			"page":               client.Page,
			"sdata":              client.Sdata,
			"websiteagen_search": client.Websiteagen_search,
			"websiteagen_page":   client.Websiteagen_page,
			"websiteagen_id":     client.Websiteagen_idrecord,
			"websiteagen_name":   client.Websiteagen_name,
			"websiteagen_status": client.Websiteagen_status,
		}).
		Post(PATH + "api/webagensave")
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
