package main

import (
	"io"
	"html/template"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Template struct {
    tmpl *template.Template
}

type Weather struct {
    Temperature int
    Rain int
}

type PageData struct {
    Weather Weather
    City string
}

type City struct {
    City string
}

func NewData() *Weather {
    return &Weather{
        Temperature: 30,
        Rain: 20,
    }
}

func NewCity() *City {
    return &City{
        City: "bla",
    }
}

func NewWeatherData(data Weather, city string) PageData {
    return PageData{
        Weather: data,
        City: city,
    }
}

func newTemplate() *Template {
    return &Template{
        tmpl: template.Must(template.ParseGlob("views/*.html")),
    }
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.tmpl.ExecuteTemplate(w, name, data)
}

func main() {
    e := echo.New()
    e.Use(middleware.Logger())

    e.Renderer = newTemplate()
    e.Static("/css", "css")

    data := NewData()
    city := ""

    e.GET("/", func(c echo.Context) error {
        return c.Render(200, "index.html", nil)
    })

    e.POST("city", func(c echo.Context) error {
        city = c.FormValue("city")

        return c.Render(200, "oob-display", NewWeatherData(*data, city))
    })

    e.Logger.Fatal(e.Start(":42069"))
}
