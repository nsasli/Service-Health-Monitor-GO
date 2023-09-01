package main

import (
    "fmt"
    "net/http"
)

func main() {
    urls := []string{
        "https://www.google.com",
        "https://www.github.com",
        "https://www.example.com", // Ganti dengan URL layanan yang ingin Anda periksa
    }

    fmt.Println("Memeriksa status kesehatan layanan:")

    for _, url := range urls {
        status := checkHealth(url)
        fmt.Printf("URL: %s, Status: %s\n", url, status)
    }
}

func checkHealth(url string) string {
    response, err := http.Get(url)
    if err != nil {
        return "Tidak dapat terhubung"
    }
    defer response.Body.Close()

    if response.StatusCode == http.StatusOK {
        return "Sehat"
    }

    return "Tidak sehat"
}
