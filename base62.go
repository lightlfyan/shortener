package main

import (
    "crypto/md5"
    "fmt"
    "strconv"

    "bytes"
)

const (
    charset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)


func base62(url string) []string {
    // 32 hex length, 128 bit
    hex := fmt.Sprintf("%x", md5.Sum([]byte(url)))
    result := make([]string, 0, 6)

    for i:=0; i<=24; i+=8 {
        v, _:= strconv.ParseInt(hex[i:i+8], 16, 32)
        v = v & 0x3FFFFFFF   // 30bit

        sb := make([]byte, 6)
        for j := 0; j<6; j++ {
            sb[j] = charset[v % 62]
            v >>= 5
        }
        result = append(result, string(sb))
    }
    return result
}


func base62v2(url string) string {
    hex := fmt.Sprintf("%x", md5.Sum([]byte(url)))
    // v, _:= strconv.ParseInt(hex[16:], 16, 64)

    sb := bytes.NewBuffer(nil)
    for i:=0; i<=16; i+=16 {
        v, _:= strconv.ParseInt(hex[i:i+16], 16, 64)
        for ;v>0; {
            sb.WriteByte(charset[v % 62])
            v /= 62
        }
    }

    return sb.String()
}

func main() {
    urls := []string{
        "http:/www.example.com",
        "http:/www.example.com?a=1",
        "http:/www.example.com?a=2",
        "http:/www.example.com?a=3",
        "http:/www.example.com?a=4",
    }

    for _, url := range urls {
        fmt.Println()
        fmt.Println("v1:", url, base62(url))
        fmt.Println("v2:", url, base62v2(url))
    }

    /*
    v1: http:/www.example.com [4WXb6v vvvvvv 8hSI57 vvvvvv]
    v2: http:/www.example.com D2oiV43rKQaUXnGK31Tfm1

    v1: http:/www.example.com?a=1 [vvvvvv vvvvvv BozzFm vvvvvv]
    v2: http:/www.example.com?a=1 7M85y0N8lZaAZJXgk1jtn9

    v1: http:/www.example.com?a=2 [WrmS7p h0MsBe vvvvvv sVseBa]
    v2: http:/www.example.com?a=2 xwgc5Y9iEm47M85y0N8lZa

    v1: http:/www.example.com?a=3 [vvvvvv vvvvvv 6E17Mb kHXQqc]
    v2: http:/www.example.com?a=3 7M85y0N8lZae5vdDnmlkV1

    v1: http:/www.example.com?a=4 [vvvvvv vvvvvv vvvvvv vvvvvv]
    v2: http:/www.example.com?a=4 7M85y0N8lZa7M85y0N8lZa
    */
}