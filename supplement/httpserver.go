package supplement

import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, r *http.Request) {
	for name, headers := range r.Header {
		for _, header := range headers {
			fmt.Fprintf(w, "header name: %s header value: %s\n", name, header)
		}
	}
}

func cancelableHello(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Printf("cancelableHello starts\n")
	defer fmt.Printf("cancelableHello ends\n")

	select {
	case <-time.After(10 * time.Second): //模拟一个耗时操作
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done(): //canceled
		err := ctx.Err()
		fmt.Printf("error: %s\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func HttpServer() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/cancelable-hello", cancelableHello)

	http.ListenAndServe(":9000", nil)
}
