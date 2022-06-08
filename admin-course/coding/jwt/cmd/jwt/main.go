package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		if token, err := create("test_user"); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(err.Error()))
		} else {
			_, _ = fmt.Fprint(w, token)
		}
	})

	http.HandleFunc("/secured", func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = fmt.Fprint(w, "Authorization header not set")
			return
		}

		splits := strings.Split(authHeader, "Bearer")
		if len(splits) != 2 {
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = fmt.Fprint(w, "Invalid Authorization header")
			return
		}

		tokenString := strings.TrimSpace(splits[1])

		if token, err := validate(tokenString); err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte(err.Error()))
		} else if !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte("Invalid token"))
		} else {
			_, _ = fmt.Fprint(w, "This content is secured")
		}
	})

	http.ListenAndServe(":8090", nil)
}
