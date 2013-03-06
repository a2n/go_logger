package logger

import (
    "testing"
    "fmt"
    "crypto/md5"
    "time"
    "io"
)

func TestInfo(t  *testing.T) {
    h := md5.New()
    io.WriteString(h, time.Now().String())
    msg := fmt.Sprintf("%x", h.Sum(nil))
    New("a2n ").Info(msg)
}
