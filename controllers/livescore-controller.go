package controllers

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
)

type response_livescore struct {
	Record interface{} `json:"scores"`
}

func Livescorehome(c *fiber.Ctx) error {
	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(response_livescore{}).
		SetError(responseerror{}).
		Get("https://www.goalserve.com/getfeed/95fb9eaa689e427a67eb08db2ea27bf2/soccernew/home?json=1")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()
	result := resp.Result().(*response_livescore)
	return c.JSON(fiber.Map{
		"status": 200,
		"record": result.Record,
		"time":   time.Since(render_page).String(),
	})
}
