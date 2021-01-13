package main

import (
	"encoding/json"
	_type "github.com/wang1137095129/my-k8s-controller/type"
	"io"
	"log"
	"net/http"
)

func main() {
	log.Println("test controller started")
	for true {
		resp, err := http.Get("http://localhost:8001/apis/v1/pod?watch=true")
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		decoder := json.NewDecoder(resp.Body)
		for {
			var event WebsiteWatchEvent
			if err := decoder.Decode(&event); err == io.EOF {
				break
			} else if err != nil {
				log.Fatal(err)
			}
			log.Printf("Recevied watch event: %s.%s.%s", event.Type, event.Object.Metadata.Namespace, event.Object.Metadata.Name)
		}
	}
}

type WebsiteWatchEvent struct {
	Type   string
	Object _type.Deployment
}
