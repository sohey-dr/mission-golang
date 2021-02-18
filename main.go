package main

// go modulesで管理すると絶対パスでのimportが必須です
import (
    "techtrain-go/controllers"
)

func main() {
    controllers.StartWebServer()
}