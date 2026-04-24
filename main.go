package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/adminevolsystemcl/hola-jungla-de-bits/internal/hostinfo"
)

func main() {
	addr := flag.String("addr", ":8080", "direccion del servidor HTTP")
	cliMode := flag.Bool("cli", false, "imprime la informacion del host en vez de iniciar el servidor")
	format := flag.String("format", "json", "formato de salida para -cli: json o text")
	flag.Parse()

	if *cliMode {
		runCLI(*format)
		return
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/api/host-info", hostInfoHandler)
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		if err := json.NewEncoder(w).Encode(map[string]string{"status": "ok"}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	log.Printf("host info server listening on %s", *addr)
	if err := http.ListenAndServe(*addr, mux); err != nil {
		log.Fatal(err)
	}
}

func runCLI(format string) {
	info, err := hostinfo.Collect()
	if err != nil {
		log.Fatal(err)
	}

	switch format {
	case "json":
		encoder := json.NewEncoder(os.Stdout)
		encoder.SetIndent("", "  ")
		if err := encoder.Encode(info); err != nil {
			log.Fatal(err)
		}
	case "text":
		if _, err := os.Stdout.WriteString(hostinfo.FormatText(info)); err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatalf("unsupported format %q, use json or text", format)
	}
}

func hostInfoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	info, err := hostinfo.Collect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := hostinfo.WriteHTTP(w, r, info); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
