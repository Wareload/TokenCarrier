package main

import (
	"github.com/charmbracelet/log"
	_ "github.com/joho/godotenv/autoload"
	"net/http"
	"os"
	"tokencarrier/internal/proxy"
)

const address = ":3000"

func init() {
	log.SetReportTimestamp(true)
	log.SetReportCaller(false)
	logLevel, _ := log.ParseLevel(os.Getenv("LOG_LEVEL"))
	log.SetLevel(logLevel)
}

func main() {
	handler, err := proxy.GetProxyHandler()
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("Server is listening on port %s.", address)
	if err := http.ListenAndServe(address, handler); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
