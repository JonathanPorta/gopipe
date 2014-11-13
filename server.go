package main

import (
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"log"
	"net/http"
  "strings"
)

type Message struct {
	Callback_url string

	Push_data struct {
		Pushed_at int
		Pusher    string
	}

	Repository struct {
		Status           string
		Description      string
		Is_trusted       bool
		Full_description string
		Repo_url         string
		Owner            string
		Is_official      bool
		Is_private       bool
		Name             string
		Namespace        string
		Star_count       int
		Comment_count    int
		Date_created     int
		Dockerfile       string
		Repo_name        string
	}
}

type Service func(Message)

var services = map[string] Service {
  "dockerhub": Dockerhub,
}

func Dockerhub(m Message) {
  spew.Dump(m)
  log.Println("docker!")
}

func main() {
  http.HandleFunc("/", handler)
  http.ListenAndServe(":3000", nil)
}

func handler(response http.ResponseWriter, request *http.Request) {
  log.Println(request.Method, request.URL.Path)
  if(request.Method == "POST") {
    serviceName := strings.Split(request.URL.Path, "/")[1]
    service,ok := services[serviceName]
    if(ok) {
      decoder := json.NewDecoder(request.Body)
      var m Message
      err := decoder.Decode(&m)
      if err != nil { panic(err) }
      service(m)
    } else {
      log.Println("No service found: ", serviceName)
    }
  }
  fmt.Fprintf(response, "Hi there, I love %s!", request.URL.Path[1:])
}
