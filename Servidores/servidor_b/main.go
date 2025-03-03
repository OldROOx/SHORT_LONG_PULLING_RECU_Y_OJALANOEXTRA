package main

import (
	"Servidores/Servidores/servidor_b/utils"
	"fmt"
	
)

func main() {
	fmt.Println("🚀 Servidor B iniciado (Consulta refrescos)...")
	
	go utils.PollRefrescos()

	select {} 
}
