package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func main() {
	// Iniciar el servidor HTTP en una goroutine
	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "¡Hola! Servidor funcionando correctamente en el puerto 8080")
		})

		fmt.Println("Servidor iniciado en http://localhost:8080")
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	// Ejecutar cloudflared tunnel y mostrar salida en tiempo real
	cmd := exec.Command("cloudflared", "tunnel", "--url", "http://localhost:8080")

	// Capturar stdout
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal("Error al crear stdout pipe:", err)
	}

	// Capturar stderr
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal("Error al crear stderr pipe:", err)
	}

	// Iniciar el comando
	if err := cmd.Start(); err != nil {
		log.Fatal("Error al iniciar cloudflared:", err)
	}

	// Leer stdout en tiempo real
	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			fmt.Println("[STDOUT]", scanner.Text())
		}
	}()

	// Leer stderr en tiempo real
	go func() {
		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
			fmt.Println("[STDERR]", scanner.Text())
		}
	}()

	// Esperar a que el comando termine
	if err := cmd.Wait(); err != nil {
		log.Fatal("Error al ejecutar cloudflared:", err)
	}
}
