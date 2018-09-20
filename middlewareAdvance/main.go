package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type handleFunc func(w http.ResponseWriter, r *http.Request)

type Middleware func(handleFunc) handleFunc

func Logging() Middleware {
	return func(f handleFunc) handleFunc {

		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			defer func() {
				log.Println(r.URL.Path, time.Since(start))
			}()

			f(w, r)

		}
	}
}

func Method(m string) Middleware {
	return func(f handleFunc) handleFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if r.Method != m {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}

			f(w, r)
		}
	}
}

func Chain(f handleFunc, middlewares ...Middleware) handleFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world")
}

func main() {
	http.HandleFunc("/", Chain(Hello, Method("GET"), Logging()))
	http.ListenAndServe(":8888", nil)
}
