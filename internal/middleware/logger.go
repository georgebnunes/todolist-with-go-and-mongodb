package middleware

import (
	"fmt"
	"net/http"
)

func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("=== Incoming Request ===")
		fmt.Printf("Method  : %s\n", r.Method)
		fmt.Printf("URL     : %s\n", r.URL.Path)
		fmt.Println("Headers:")

		for key, values := range r.Header {
			for _, value := range values {
				fmt.Printf("   %s: %s\n", key, value)
			}
		}

		fmt.Println("======================")

		next.ServeHTTP(w, r)
	})
}
