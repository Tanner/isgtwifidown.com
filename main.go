package main

import (
	"net/http"
	"html/template"
	"log"
	"github.com/codegangsta/martini"
	"github.com/tanner/isgtwifidown.com/gtwifi"
)

type PageData struct {
	Green bool
	Yellow bool
	Red bool
	Reason string
}

func main() {
	m := martini.Classic()

	template := template.Must(template.ParseFiles("index.tmpl"))

	m.Get("/", func(res http.ResponseWriter, req *http.Request) {
		status, err := gtwifi.GetStatus()
		if err != nil {
			log.Println(err);
		}

		data := PageData{}
		data.Green = status.Status == gtwifi.GREEN
		data.Yellow = status.Status == gtwifi.YELLOW
		data.Red = status.Status == gtwifi.RED
		data.Reason = status.Reason

		template.Execute(res, data)
	})

	m.Run()
}
