package main

import (
	"bytes"
	"errors"
	"log"
	"os"

	"github.com/Terisback/subkers"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
)

const MaxSizeOfFile = 5000000

func main() {
	app := fiber.New(&fiber.Settings{
		ServerHeader:          "Subkers",
		DisableStartupMessage: true,
	})

	app.Use(middleware.Logger())
	app.Use(middleware.Compress(middleware.CompressLevelBestSpeed))

	app.Post("/api/convert", convertHandler)

	app.Static("/", "./client/dist/")
	app.Get("*", func(c *fiber.Ctx) {
		c.Redirect("/")
	})

	port := os.Getenv("SUBKERS_PORT")
	if port == "" {
		port = "80"
	}

	log.Fatal(app.Listen(port))
}

func convertHandler(c *fiber.Ctx) {
	file, err := c.FormFile("subtitle")
	if err != nil {
		log.Println(err)
		sendErrorViaJSON(c, err)
		return
	}

	if file.Size > MaxSizeOfFile {
		sendErrorViaJSON(c, errors.New("File is too big (Max size of file is 5 MB)"))
		return
	}

	ext := c.FormValue("extension")

	subType, err := subkers.SubtitlesType(ext)
	if err != nil {
		sendErrorViaJSON(c, err)
		return
	}

	f, _ := file.Open()
	defer f.Close()

	markers, err := subkers.ProcessSpecific(subType, f)
	if err != nil {
		sendErrorViaJSON(c, err)
		return
	}

	var buf bytes.Buffer
	err = subkers.WriteAll(markers, &buf)
	if err != nil {
		sendErrorViaJSON(c, err)
		return
	}

	c.SendStream(&buf)
}

func sendErrorViaJSON(c *fiber.Ctx, err error) {
	c.JSON(struct {
		Err string `json:"error"`
	}{
		err.Error(),
	})
}
