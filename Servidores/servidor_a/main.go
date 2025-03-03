package main

import (
	"Servidores/Servidores/servidor_a/utils"
	"fmt"
)

func main() {
	fmt.Println("🚀 Servidor A iniciado (Consulta fantasmas)...")
	go utils.PollFantasmas()

	select {}
}
