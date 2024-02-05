package top

import "github.com/maxence-charriere/go-app/v9/pkg/app"

// NewHandler はトップページのハンドラーを生成します。
func NewHandler() *app.Handler {
	return &app.Handler{
		Author:          "murasame29",
		BackgroundColor: "#f5f5f5",
		Icon: app.Icon{
			Default: "/web/assets/icon.png",
		},
		Image: "/web/assets/icon.png",
		Title: "Murasame29のホームページ",
	}
}

// Index はトップページのコンポーネントです。
type Index struct {
	app.Compo
}

// NewComponent はトップページのコンポーネントを生成します。
func NewComponent() *Index {
	return &Index{}
}

// Render はコンポーネントを描画します。
func (i *Index) Render() app.UI {
	return app.Div().Body(
		app.H1().Text("Murasame29のホームページ"),
		app.P().Text("ようこそ！"),
	)
}
