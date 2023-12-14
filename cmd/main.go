package main

import (
	"io"
	"text/template"

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

type WeatherData struct {
    Data Weather
}

func NewData() *Weather {
    return &Weather{
        Temperature: 30,
        Rain: 20,
    }
}

func NewWeatherData(data Weather) WeatherData {
    return WeatherData{
        Data: data,
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

    e.GET("/", func(c echo.Context) error {
        return c.Render(200, "index.html", NewWeatherData(*data))
    })

    e.Logger.Fatal(e.Start(":42069"))
}
