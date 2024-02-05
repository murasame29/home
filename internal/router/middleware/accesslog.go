package middleware

import (
	"encoding/json"
	"net/http"
)

// AccessLog はアクセスログを定義する構造体です。
type AccessLog struct {
	// RemoteAddr はリクエストを送信したクライアントのアドレスです。
	RemoteAddr string `json:"remote_addr"`
	// Method はリクエストのHTTPメソッドです。
	Method string `json:"method"`
	// StatusCode はリクエストに対するレスポンスのステータスコードです。
	StatusCode int `json:"status_code"`
	// URI はリクエストのURIです。
	URI string `json:"uri"`
	// UserAgent はリクエストを送信したクライアントのUser-Agentです。
	UserAgent string `json:"user_agent"`
	// Error はリクエスト処理中に発生したエラーです。
	Error error `json:"error"`
}

// Marshal はアクセスログをJSON形式に変換します。
func (al AccessLog) Marshal() ([]byte, error) {
	return json.Marshal(al)
}

// StatusCodeKey はステータスコードを格納するためのキーです。
type StatusCodeKey struct{}

// ErrorKey はエラーを格納するためのキーです。
type ErrorKey struct{}

// AccessLog はアクセスログを出力するミドルウェアです。
func (m *middleware) AccessLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var al AccessLog
		al = AccessLog{
			RemoteAddr: r.RemoteAddr,
			Method:     r.Method,
			URI:        r.RequestURI,
			UserAgent:  r.UserAgent(),
		}
		next.ServeHTTP(w, r)

		if StatusCode, ok := r.Context().Value(StatusCodeKey{}).(int); ok {
			al.StatusCode = StatusCode
		} else {
			al.StatusCode = 200
		}

		if err, ok := r.Context().Value(ErrorKey{}).(error); ok {
			al.Error = err
		}

		// アクセスログの出力
		m.l.Info(al)
	})
}
