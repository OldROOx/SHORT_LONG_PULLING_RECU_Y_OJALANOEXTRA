package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func PollTumbas() {
	for {
		resp, err := http.Get("http://localhost:8080/tumbas")
		if err != nil {
			fmt.Println("❌ Error al obtener tumbas:", err)
		} else {
			body, _ := ioutil.ReadAll(resp.Body)
			fmt.Println("✅ Servidor B recibió tumbas:", string(body))
			resp.Body.Close()
		}
		time.Sleep(5 * time.Second)
	}
}
