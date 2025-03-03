package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func PollRefrescos() {
	for {
		resp, err := http.Get("http://localhost:8080/refrescos") 
		if err != nil {
			fmt.Println("❌ Error al obtener refrescos:", err)
		} else {
			body, _ := ioutil.ReadAll(resp.Body)
			fmt.Println("✅ Servidor B recibió refrescos:", string(body))
			resp.Body.Close()
		}
		time.Sleep(5 * time.Second) 
	}
}
