package middleware

import (
	"context"
	"net/http"
)

// CounterKey はctxに格納するカウンターのキーです。
type CounterKey string

// Counter は閲覧者のカウントを行うミドルウェアです。
func (m *middleware) Counter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// カウントアップ処理をして結果を保存しctxに格納
		counter := 0
		ctx := r.Context()

		ctx = context.WithValue(ctx, CounterKey("counter"), counter)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
