package controllers

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"time"

	"bitbucket.org/isbtotogroup/isbpanel_frontend_backend/helpers"
	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func Crmhome(c *fiber.Ctx) error {
	type payload_crmhome struct {
		Crm_status string `json:"crm_status"`
		Crm_search string `json:"crm_search"`
		Crm_page   int    `json:"crm_page"`
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_crmhome)
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
			"client_hostname": hostname,
			"crm_status":      client.Crm_status,
			"crm_search":      client.Crm_search,
			"crm_page":        client.Crm_page,
		}).
		Post(PATH + "api/crm")
	if err != nil {
		log.Println(err.Error())
	}

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
func Crmsaleshome(c *fiber.Ctx) error {
	type payload_crmhome struct {
		Crmsales_phone  string `json:"crmsales_phone"`
		Crmsales_status string `json:"crmsales_status"`
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_crmhome)
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
			"client_hostname": hostname,
			"crmsales_phone":  client.Crmsales_phone,
			"crmsales_status": client.Crmsales_status,
		}).
		Post(PATH + "api/crmsales")
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
func Crmsalesperform(c *fiber.Ctx) error {
	type payload_crmhome struct {
		Page               string `json:"page"`
		Employee_iddepart  string `json:"employee_iddepart"`
		Employee_username  string `json:"employee_username"`
		Employee_startdate string `json:"employee_startdate"`
		Employee_enddate   string `json:"employee_enddate"`
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_crmhome)
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
			"page":               client.Page,
			"employee_iddepart":  client.Employee_iddepart,
			"employee_username":  client.Employee_username,
			"employee_startdate": client.Employee_startdate,
			"employee_enddate":   client.Employee_enddate,
		}).
		Post(PATH + "api/employeebysalesperformance")
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
func Crmdeposithome(c *fiber.Ctx) error {
	type payload_crmhome struct {
		Crmsales_idcrmsales int `json:"crmsales_idcrmsales"`
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_crmhome)
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
			"client_hostname":     hostname,
			"crmsales_idcrmsales": client.Crmsales_idcrmsales,
		}).
		Post(PATH + "api/crmdeposit")
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
func Crmisbtvhome(c *fiber.Ctx) error {
	type payload_crmisbtvhome struct {
		Crmisbtv_search string `json:"crmisbtv_search"`
		Crmisbtv_page   int    `json:"crmisbtv_page"`
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_crmisbtvhome)
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
			"client_hostname": hostname,
			"crmisbtv_page":   client.Crmisbtv_page,
			"crmisbtv_search": client.Crmisbtv_search,
		}).
		Post(PATH + "api/crmisbtv")
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
func Crmduniafilm(c *fiber.Ctx) error {
	type payload_crmisbtvhome struct {
		Crmisbtv_search string `json:"crmisbtv_search"`
		Crmisbtv_page   int    `json:"crmisbtv_page"`
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_crmisbtvhome)
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
			"client_hostname": hostname,
			"crmisbtv_page":   client.Crmisbtv_page,
			"crmisbtv_search": client.Crmisbtv_search,
		}).
		Post(PATH + "api/crmduniafilm")
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

func CrmSave(c *fiber.Ctx) error {
	type payload_crmsave struct {
		Page       string `json:"page"`
		Sdata      string `json:"sdata" `
		Crm_page   int    `json:"crm_page"`
		Crm_id     int    `json:"crm_id"`
		Crm_phone  string `json:"crm_phone" `
		Crm_name   string `json:"crm_name" `
		Crm_status string `json:"crm_status" `
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_crmsave)
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
			"sdata":           client.Sdata,
			"crm_page":        client.Crm_page,
			"crm_id":          client.Crm_id,
			"crm_phone":       client.Crm_phone,
			"crm_name":        client.Crm_name,
			"crm_status":      client.Crm_status,
		}).
		Post(PATH + "api/crmsave")
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
func CrmSaveStatus(c *fiber.Ctx) error {
	type payload_crmsave struct {
		Page       string `json:"page"`
		Sdata      string `json:"sdata" `
		Crm_page   int    `json:"crm_page"`
		Crm_id     int    `json:"crm_id"`
		Crm_status string `json:"crm_status" `
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_crmsave)
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
			"sdata":           client.Sdata,
			"crm_page":        client.Crm_page,
			"crm_id":          client.Crm_id,
			"crm_status":      client.Crm_status,
		}).
		Post(PATH + "api/crmsavestatus")
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
func CrmSalesSave(c *fiber.Ctx) error {
	type payload_crud struct {
		Page              string `json:"page"`
		Search            string `json:"search" `
		Crm_page          int    `json:"crm_page"`
		Crmsales_phone    string `json:"crmsales_phone"`
		Crmsales_username string `json:"crmsales_username"`
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_crud)
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
			"search":            client.Search,
			"crm_page":          client.Crm_page,
			"crmsales_phone":    client.Crmsales_phone,
			"crmsales_username": client.Crmsales_username,
		}).
		Post(PATH + "api/crmsalessave")
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
func CrmSalesDelete(c *fiber.Ctx) error {
	type payload_crud struct {
		Page           string `json:"page"`
		Search         string `json:"search" `
		Crm_page       int    `json:"crm_page"`
		Crmsales_id    int    `json:"crmsales_id" `
		Crmsales_phone string `json:"crmsales_phone" `
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_crud)
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
			"search":          client.Search,
			"crm_page":        client.Crm_page,
			"crmsales_id":     client.Crmsales_id,
			"crmsales_phone":  client.Crmsales_phone,
		}).
		Post(PATH + "api/crmsalesdelete")
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
func CrmSavesource(c *fiber.Ctx) error {
	type payload_crmsavesource struct {
		Page       string          `json:"page"`
		Sdata      string          `json:"sdata" `
		Crm_page   int             `json:"crm_page"`
		Crm_source string          `json:"crm_source" `
		Crm_data   json.RawMessage `json:"crm_data" `
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_crmsavesource)
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
			"sdata":           client.Sdata,
			"crm_page":        client.Crm_page,
			"crm_source":      client.Crm_source,
			"crm_data":        client.Crm_data,
		}).
		Post(PATH + "api/crmsavesource")
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
func CrmSaveDatabase(c *fiber.Ctx) error {
	type payload_crmsavesource struct {
		Page       string          `json:"page"`
		Sdata      string          `json:"sdata" `
		Crm_page   int             `json:"crm_page"`
		Crm_source string          `json:"crm_source" `
		Crm_data   json.RawMessage `json:"crm_data" `
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_crmsavesource)
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
			"sdata":           client.Sdata,
			"crm_page":        client.Crm_page,
			"crm_source":      client.Crm_source,
			"crm_data":        client.Crm_data,
		}).
		Post(PATH + "api/crmsavedatabase")
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
func CrmSaveMaintenance(c *fiber.Ctx) error {
	type payload_crmsavesource struct {
		Page     string          `json:"page"`
		Sdata    string          `json:"sdata" `
		Crm_page int             `json:"crm_page"`
		Crm_data json.RawMessage `json:"crm_data" `
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_crmsavesource)
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
			"sdata":           client.Sdata,
			"crm_page":        client.Crm_page,
			"crm_data":        client.Crm_data,
		}).
		Post(PATH + "api/crmsavemaintenance")
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

type DatabaseRecord struct {
	Phone string
	Nama  string
}

func Crmdatabase(c *fiber.Ctx) error {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	dir_img := "/frontend/public/csv/"

	if runtime.GOOS == "windows" {
		dir_img = strings.Replace(dir_img, "/", "\\", -1)
	}

	file, err := c.FormFile("file")
	if err != nil {
		log.Println("csv upload error --> ", err)
		return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})

	}
	log.Println("File : ", file.Filename)
	uniqueId := uuid.New()
	filename := strings.Replace(uniqueId.String(), "-", "", -1)
	fileExt := strings.Split(file.Filename, ".")[1]
	image := fmt.Sprintf("%s.%s", filename, fileExt)
	log.Println(image)
	// path_imagelocal := `F:\ISBPROJECT\ISBPANEL\isbpanel_backend\frontend\public\images\` + image
	path_imageserver := dir + dir_img
	err = c.SaveFile(file, fmt.Sprintf(`%s/%s`, path_imageserver, image))
	if err != nil {
		log.Println("csv save error --> ", err)
		return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
	}

	// path_imagelocal = `F:\ISBPROJECT\ISBPANEL\isbpanel_backend\frontend\public\images\` + image
	path_imageserver = dir + dir_img + image
	log.Println(path_imageserver)

	f, err := os.Open(path_imageserver)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// convert records to array of structs
	databaseList := createDatabaseList(data)
	fmt.Printf("%+v\n", databaseList)
	return c.JSON(fiber.Map{
		"status": 200,
		"record": databaseList,
	})
}
func createDatabaseList(data [][]string) []DatabaseRecord {
	var databaseList []DatabaseRecord
	for i, line := range data {
		if i > 0 { // omit header line
			var rec DatabaseRecord
			for j, field := range line {
				if j == 0 {
					rec.Phone = field
				} else if j == 1 {
					rec.Nama = field
				}
			}
			databaseList = append(databaseList, rec)
		}
	}
	return databaseList
}
