package main

import "github.com/kitanoyoru/kitaShop-backend/internal/app"

const (
  configDir = "configs"
)

func main() {
  app.Run(configDir)
}
