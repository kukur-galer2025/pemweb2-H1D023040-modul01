package main

import (
    "log"
    "github.com/gofiber/fiber/v3"
)

func main() {
    // Membuat aplikasi Fiber baru
    app := fiber.New()

    // Endpoint bawaan dari Modul
    app.Get("/", func(c fiber.Ctx) error {
        return c.SendString("Halo Pemrograman Web II")
    })

    // Endpoint API Info dari Modul
    app.Get("/api/info", func(c fiber.Ctx) error {
        return c.JSON(fiber.Map{
            "aplikasi": "Latihan Fiber",
            "versi":    "1.0.0",
            "status":   "berjalan",
        })
    })

    // TUGAS PRAKTIKUM: Endpoint API Mahasiswa
    app.Get("/api/mahasiswa", func(c fiber.Ctx) error {
        return c.JSON(fiber.Map{
            "nim":           "H1DO23040", // GANTI DENGAN NIM KAMU
            "nama":          "Prima Dzaky Hibatulloh", // GANTI DENGAN NAMA KAMU
            "program_studi": "Informatika",
        })
    })

    // Menyalakan server di port 3000
    log.Fatal(app.Listen(":3000"))
}