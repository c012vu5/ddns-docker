package main

import (
    "fmt"
    "io/ioutil"
    "math/rand"
    "net/http"
    "os"
    "os/signal"
    "regexp"
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
        remoteAddr, err := getRemoteAddr(resp)
        if err != nil {
            fmt.Println("Error getting remote address:", err)
            return
        }
        fmt.Println("ADDRESS update success:", remoteAddr)
    } else {
        fmt.Println("ADDRESS update failure.")
    }
}

func getRandomWaitSec() int {
    waitSec := rand.Intn(601) + 600
    return waitSec
}

func getRemoteAddr(resp *http.Response) (string, error) {
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error reading response:", err)
        return "", err
    }

    bodyStr := string(body)
    re := regexp.MustCompile(`<DT>REMOTE ADDRESS:</DT><DD>([^<]+)</DD>`)
    match := re.FindStringSubmatch(bodyStr)
    if len(match) == 2 {
        return match[1], nil
    }

    return "", fmt.Errorf("Remote address not found")
}
