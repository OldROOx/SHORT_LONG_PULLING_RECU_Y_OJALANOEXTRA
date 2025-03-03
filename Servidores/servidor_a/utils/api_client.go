package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func PollFantasmas() {
	for {
		resp, err := http.Get("http://localhost:8080/fantasmas")
		if err != nil {
			fmt.Println("❌ Error al obtener fantasmas:", err)
		} else {
			body, _ := ioutil.ReadAll(resp.Body)
			fmt.Println("✅ Servidor A recibió fantasmas:", string(body))
			resp.Body.Close()
		}
		time.Sleep(5 * time.Second)
	}
}
