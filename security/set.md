https://docs.gofiber.io/guide/routing#middleware


## Set in fiber
```go
app.Use(func(c *fiber.Ctx) error {
  // Set some security headers:
  c.Set("X-XSS-Protection", "1; mode=block")
  c.Set("X-Content-Type-Options", "nosniff")
  c.Set("X-Download-Options", "noopen")
  c.Set("Strict-Transport-Security", "max-age=5184000")
  c.Set("X-Frame-Options", "SAMEORIGIN")
  c.Set("X-DNS-Prefetch-Control", "off")

  // Go to next middleware:
  return c.Next()
})

app.Get("/", func(c *fiber.Ctx) error {
  return c.SendString("Hello, World!")
})
```
