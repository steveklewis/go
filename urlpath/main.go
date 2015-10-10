package main

import (
       "fmt"
       "net/http"

       "github.com/labstack/echo"
       mw "github.com/labstack/echo/middleware"
)

func handle(c *echo.Context) error {
     fmt.Println("RawPath", c.Request().URL.RawPath)
     return c.String(http.StatusOK, c.Request().URL.Path)
}


func main() {

     e := echo.New()

     e.Use(mw.Logger())
     e.Use(mw.Recover())

     e.Get("/", handle)
     e.Get("/index/type", handle)

     e.Run(":1323")
     
}