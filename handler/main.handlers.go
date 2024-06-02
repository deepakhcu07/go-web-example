package handler

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
	"time"
)

func HandleViewHome(c *fiber.Ctx) error {

	return c.Render("index", fiber.Map{
		"Page": "Dev Tools",
	})
}

func HandleViewEpochConverter(c *fiber.Ctx) error {
	now := time.Now().UTC()
	return c.Render("epoch-converter", fiber.Map{
		"Page":              "Epoch Converter",
		"EpochInSeconds":    strconv.FormatInt(now.Unix(), 10),
		"EpochInMillis":     strconv.FormatInt(now.UnixMilli(), 10),
		"EpochInMicro":      strconv.FormatInt(now.UnixMicro(), 10),
		"EpochInNanoSecond": strconv.FormatInt(now.UnixNano(), 10),
	})
}

func HandleTimeSSE(c *fiber.Ctx) error {
	now := time.Now().UTC()
	return c.Render("epoch-converter", fiber.Map{
		"EpochInSeconds": strconv.FormatInt(now.Unix(), 10),
	})
}
