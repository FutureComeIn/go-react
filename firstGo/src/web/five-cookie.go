package main

import (
	"fmt"
	"net/http"
	"time"
)

func main(){
	expiration := time.Now()
	expiration = expiration.AddDate(1, 0, 0)
	cookie := http.Cookie{Name: "Tom", Value: "Tom123", Expires: expiration}
	// http.SetCookie(w ResponseWriter, cookie *Cookie)，w http.ResponseWriter
	http.SetCookie(w, &cookie)

	// 获取cookie，r *http.Request
	cookie, _ = r.Cookie("Tom")
	fmt.Fprint(w, cookie)

	// 另外一种方式获取
	for _, cookie := range r.Cookies() {
		fmt.Fprint(w, cookie.Name)
	}
}