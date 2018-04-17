package main

import (
    "crypto/md5"
    "fmt"
    "strconv"

    "bytes"
)


func base62(url string) []string {
    charset := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    // 32 length
    hex := fmt.Sprintf("%x", md5.Sum([]byte(url)))
    result := make([]string, 0, 6)

    for i:=0; i<=24; i+=8 {
        v, _:= strconv.ParseInt(hex[i:i+8], 16, 32)
        v = v & 0x3FFFFFFF   // 2 << 31 - 1

        sb := make([]byte, 6)
        for j := 0; j<6; j++ {
            sb[j] = charset[v % 62]
            v = v >> 5
        }
        result = append(result, string(sb))
    }
    return result
}


func base62v2(url string) string {
    charset := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    // 32 length
    hex := fmt.Sprintf("%x", md5.Sum([]byte(url)))

    v, _:= strconv.ParseInt(hex[:16], 16, 64)
    // v, _:= strconv.ParseInt(hex[16:], 16, 64)

    sb := bytes.NewBuffer(nil)
    for ;v>0; {
        sb.WriteByte(charset[v % 62])
        v /= 62
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
        fmt.Println(url, base62(url))
        fmt.Println(url, base62v2(url))
        return
    }
}