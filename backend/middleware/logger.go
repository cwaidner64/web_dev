
package middleware

import (
	"fmt"
	"net/http"
	"time"
	"os"
)

type ReqLogResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rlrw *ReqLogResponseWriter) WriteHeader(statusCode int) {
	rlrw.statusCode = statusCode
	rlrw.ResponseWriter.WriteHeader(statusCode)
}


func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rlrw := &ReqLogResponseWriter{
			w,
			http.StatusOK,
		}
		next.ServeHTTP(rlrw, r)
		//format details, irrelevant time inputted
		fmt.Println(os.Getenv("APP_NAME"),start.Format("2006-01-02 15:04:05"), r.Method,rlrw.statusCode ,r.URL.Path,time.Since(start))
	})
}