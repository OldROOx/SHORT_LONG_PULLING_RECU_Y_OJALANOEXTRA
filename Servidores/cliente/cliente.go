package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	fmt.Println("🕐 Cliente esperando actualización en API Central...")

	for {
	
		resp, err := http.Get("http://localhost:8080/refrescos/check-deleted")
		if err != nil {
			fmt.Println("❌ Error en long polling:", err)
			time.Sleep(5 * time.Second) 
			continue
		}

		if resp.StatusCode == http.StatusNoContent {
			fmt.Println("⏳ Sin cambios, esperando...")
		} else {
			body, _ := ioutil.ReadAll(resp.Body)
			fmt.Println("🚀 Refresco eliminado:", string(body))
		}

		resp.Body.Close()
	}
}
