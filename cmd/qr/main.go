package main

import (
	"log"
	"net/http"

	"github.com/skip2/go-qrcode"
)

func qrHandler(w http.ResponseWriter, r *http.Request) {
	data := r.URL.Query().Get("data")
	if data == "" {
		http.Error(w, "Missing 'data' parameter", http.StatusBadRequest)
		return
	}

	// Generate a 256x256 QR code image
	png, err := qrcode.Encode(data, qrcode.Medium, 256)
	if err != nil {
		http.Error(w, "Failed to generate QR code", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.Write(png)
}

func main() {
	http.HandleFunc("/generate", qrHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
