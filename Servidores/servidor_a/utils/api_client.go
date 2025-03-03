package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func PollSabritas() {
	for {
		resp, err := http.Get("http://localhost:8080/sabritas")
		if err != nil {
			fmt.Println("❌ Error al obtener sabritas:", err)
		} else {
			body, _ := ioutil.ReadAll(resp.Body)
			fmt.Println("✅ Servidor A recibió sabritas:", string(body))
			resp.Body.Close()
		}
		time.Sleep(5 * time.Second) 
	}
}
