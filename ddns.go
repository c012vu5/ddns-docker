package main

import (
    "fmt"
    "math/rand"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"
)

var (
    ACC  = ""
    PASS = ""
)

func main() {
    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, syscall.SIGTERM)

    go func() {
        sig := <-sigChan
        fmt.Printf("Received signal: %v\n", sig)

        os.Exit(0)
    }()

    fmt.Println("Global IP notifier started")

    pulseLoop()
}

func pulseLoop() {
    pulse()
    for {
        time.Sleep(time.Duration(getRandomWaitSec()) * time.Second)
        pulse()
    }
}

func pulse() {
    client := &http.Client{}
    req, err := http.NewRequest("GET", fmt.Sprintf("https://%s:%s@ipv4.mydns.jp/login.html", ACC, PASS), nil)
    if err != nil {
        fmt.Println("Error creating request:", err)
        return
    }

    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Error sending request:", err)
        return
    }
    defer resp.Body.Close()

    if resp.StatusCode == http.StatusOK {
        fmt.Println("ADDRESS update success.")
    } else {
        fmt.Println("ADDRESS update failure.")
    }
}

func getRandomWaitSec() int {
    waitSec := rand.Intn(601) + 600
    return waitSec
}
