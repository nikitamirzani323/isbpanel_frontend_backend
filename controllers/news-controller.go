package controllers

import (
	"log"
	"strings"
	"time"

	"bitbucket.org/isbtotogroup/isbpanel_frontend_backend/entities"
	"bitbucket.org/isbtotogroup/isbpanel_frontend_backend/helpers"
	"bitbucket.org/isbtotogroup/isbpanel_frontend_backend/models"
	"github.com/buger/jsonparser"
	"github.com/go-playground/validator/v10"
	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

const Fieldnews_home_redis = "LISTNEWS_BACKEND_ISBPANEL"
const Fieldcategory_home_redis = "LISTCATEGORY_BACKEND_ISBPANEL"
const Fieldnews_client_home_redis = "LISTNEWS_FRONTEND_ISBPANEL"
const Fieldnewsmovie_client_home_redis = "LISTNEWSMOVIES_FRONTEND_ISBPANEL"

func Newshome(c *fiber.Ctx) error {
	type payload_newshome struct {
		News_search string `json:"news_search"`
		News_page   int    `json:"news_page"`
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_newshome)
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
			"news_page":       client.News_page,
			"news_search":     client.News_search,
		}).
		Post(PATH + "api/news")
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
func Newssave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_newssave)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}
	user := c.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	temp_decp := helpers.Decryption(name)
	client_admin, _ := helpers.Parsing_Decry(temp_decp, "==")

	result, err := models.Save_news(
		client_admin, client.Sdata,
		client.News_title, client.News_descp, client.News_url, client.News_image, client.News_idrecord, client.News_category)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	val_news := helpers.DeleteRedis(Fieldnews_home_redis + "_")
	log.Printf("Redis Delete BACKEND NEWS : %d", val_news)
	val_client_news := helpers.DeleteRedis(Fieldnews_client_home_redis)
	log.Printf("Redis Delete CLIENT NEWS : %d", val_client_news)
	val_client_newsmovie := helpers.DeleteRedis(Fieldnewsmovie_client_home_redis)
	log.Printf("Redis Delete CLIENT NEWS MOVIE : %d", val_client_newsmovie)
	return c.JSON(result)
}
func Newsdelete(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_newsdelete)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}
	user := c.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	temp_decp := helpers.Decryption(name)
	client_admin, _ := helpers.Parsing_Decry(temp_decp, "==")

	result, err := models.Delete_news(client_admin, client.News_id)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	val_news := helpers.DeleteRedis(Fieldnews_home_redis + "_")
	log.Printf("Redis Delete BACKEND NEWS : %d", val_news)
	val_client_news := helpers.DeleteRedis(Fieldnews_client_home_redis)
	log.Printf("Redis Delete CLIENT NEWS : %d", val_client_news)
	val_client_newsmovie := helpers.DeleteRedis(Fieldnewsmovie_client_home_redis)
	log.Printf("Redis Delete CLIENT NEWS MOVIE : %d", val_client_newsmovie)
	return c.JSON(result)
}
func Categoryhome(c *fiber.Ctx) error {
	var obj entities.Model_category
	var arraobj []entities.Model_category
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldcategory_home_redis)
	jsonredis := []byte(resultredis)
	message_RD, _ := jsonparser.GetString(jsonredis, "message")
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		category_id, _ := jsonparser.GetInt(value, "category_id")
		category_name, _ := jsonparser.GetString(value, "category_name")
		category_totalnews, _ := jsonparser.GetInt(value, "category_totalnews")
		category_display, _ := jsonparser.GetInt(value, "category_display")
		category_status, _ := jsonparser.GetString(value, "category_status")
		category_statuscss, _ := jsonparser.GetString(value, "category_statuscss")
		category_create, _ := jsonparser.GetString(value, "category_create")
		category_update, _ := jsonparser.GetString(value, "category_update")

		obj.Category_id = int(category_id)
		obj.Category_name = category_name
		obj.Category_totalnews = int(category_totalnews)
		obj.Category_display = int(category_display)
		obj.Category_status = category_status
		obj.Category_statuscss = category_statuscss
		obj.Category_create = category_create
		obj.Category_update = category_update
		arraobj = append(arraobj, obj)
	})
	if !flag {
		result, err := models.Fetch_category()
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldcategory_home_redis, result, 60*time.Minute)
		log.Println("CATEGORY NEWS MYSQL")
		return c.JSON(result)
	} else {
		log.Println("CATEGORY NEWS CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": message_RD,
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Categorysave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_categorysave)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}
	user := c.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	temp_decp := helpers.Decryption(name)
	client_admin, _ := helpers.Parsing_Decry(temp_decp, "==")

	result, err := models.Save_category(
		client_admin,
		client.Category_name, client.Category_status, client.Sdata, client.Category_idrecord, client.Category_display)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	val_category := helpers.DeleteRedis(Fieldcategory_home_redis)
	log.Printf("Redis Delete BACKEND NEWS CATEGORY : %d", val_category)
	return c.JSON(result)
}
func Categorydelete(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_categorydelete)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}
	user := c.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	temp_decp := helpers.Decryption(name)
	client_admin, _ := helpers.Parsing_Decry(temp_decp, "==")

	result, err := models.Delete_category(client_admin, client.Category_idrecord)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	val_category := helpers.DeleteRedis(Fieldcategory_home_redis)
	log.Printf("Redis Delete BACKEND NEWS CATEGORY : %d", val_category)
	return c.JSON(result)
}
