package main

import (
	"github.com/codegangsta/martini"
	"github.com/tanner/isgtwifidown.com/gtwifi"
	"html/template"
	"log"
	"net/http"
	"time"
)

type PageData struct {
	Green  bool
	Yellow bool
	Red    bool
	Reason string
}

type LastData struct {
	Status        int
	Reason        string
	TimeRetrieved time.Time
}

var lastData LastData

func main() {
	m := martini.Classic()

	template := template.Must(template.ParseFiles("index.tmpl"))

	schedule(checkData, 5*time.Minute)
	checkData()

	m.Get("/", func(res http.ResponseWriter, req *http.Request) {
		data := PageData{}
		data.Green = lastData.Status == gtwifi.GREEN
		data.Yellow = lastData.Status == gtwifi.YELLOW
		data.Red = lastData.Status == gtwifi.RED
		data.Reason = lastData.Reason

		template.Execute(res, data)
	})

	m.Run()
}

func checkData() {
	log.Println("Retrieving new data...")

	status, err := gtwifi.GetStatus()

	if err != nil {
		log.Println(err)
		return
	}

	log.Println("New data retrieved!")

	lastData = LastData{status.Status, status.Reason, time.Now()}
}

func schedule(caller func(), delay time.Duration) chan bool {
	ticker := time.NewTicker(delay)
	stop := make(chan bool)

	go func() {
		for {
			select {
			case <-ticker.C:
				caller()
			case <-stop:
				ticker.Stop()
				return
			}
		}
	}()

	return stop
}
