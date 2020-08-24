package main

import (
	"bytes"
	"crypto/tls"
	"log"

	"github.com/Terisback/subkers"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"golang.org/x/crypto/acme/autocert"
)

const TLS = false

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

	if TLS {
		m := &autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			HostPolicy: autocert.HostWhitelist("subkers.terisback.ru"),
			Cache:      autocert.DirCache("./certs"),
		}

		tls := &tls.Config{
			GetCertificate: m.GetCertificate,
			NextProtos: []string{
				"http/1.1", "acme-tls/1",
			},
		}

		log.Fatal(app.Listen(443, tls))
	} else {
		log.Fatal(app.Listen(3000))
	}
}

func convertHandler(c *fiber.Ctx) {
	file, err := c.FormFile("subtitle")
	if err != nil {
		log.Println(err)
		c.SendStatus(fiber.StatusBadRequest)
		return
	}
	ext := c.FormValue("extension")
	f, _ := file.Open()

	subType, err := subkers.SubtitlesType(ext)
	if err != nil {
		c.SendStatus(fiber.StatusBadRequest)
		return
	}

	markers, err := subkers.ProcessSpecific(subType, f)
	if err != nil {
		c.SendStatus(fiber.StatusBadRequest)
		return
	}

	var buf bytes.Buffer
	err = subkers.WriteAll(markers, &buf)
	if err != nil {
		c.SendStatus(fiber.StatusBadRequest)
		return
	}

	c.SendStream(&buf)
}
