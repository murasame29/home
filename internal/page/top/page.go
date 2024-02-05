package top

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

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
		Styles: []string{
			"/web/assets/reset.css",
			"/web/assets/style.css",
		},
	}
}

// Index はトップページのコンポーネントです。
type Index struct {
	app.Compo

	// ナビゲーションのインデックス
	navIndex int

	viewcount int
}

// NewComponent はトップページのコンポーネントを生成します。
func NewComponent() *Index {
	return &Index{}
}

// Render はコンポーネントを描画します。
func (i *Index) Render() app.UI {
	return app.Div().Body(
		i.body(),
	).Class("contents")
}

func (i *Index) body() app.UI {
	return app.
		Table().
		Attr("width", "964").
		Attr("cellpadding", "0").
		Attr("cellspacing", "0").
		Attr("height", "100%").
		Body(
			app.TBody().Body(
				i.boundary(),
				app.Td().Body(),
				i.main(),
				i.boundary(),
			),
		)
}

func (i *Index) boundary() app.UI {
	return app.Td().Attr("width", "2").Attr("bgcolor", "#ffffff")
}

func (i *Index) main() app.UI {
	return app.Td().Attr("width", "960").Attr("bgcolor", "#999999").Body(
		i.header(),
		i.navigate(),
	)
}

func (i *Index) header() app.UI {
	return app.Table().Attr("width", "960").Attr("cellpadding", "0").Attr("cellspacing", "0").Attr("height", "692").Body(
		app.TBody().Body(
			app.Tr().Body(
				app.Td().Attr("width", "370").Attr("height", "92").Body(),
				app.Td().Attr("width", "220").Body(
					app.A().
						Attr("href", "/").
						Body(
							app.Img().Src("/web/assets/onmoused-logo.png").Attr("border", "0").Attr("name", "_HPB_ROLLOVER1").Alt("Murasame29のホームページのロゴ"),
						),
				),
				app.Td().Attr("width", "370").Body(),
			),
			app.Tr().Attr("colspan", "3").Attr("width", "960").Attr("valign", "middle").Body(
				app.Td().Attr("width", "960").Attr("colspan", "3").Body(
					app.Raw(fmt.Sprintf("<marquee>あなたは%d人目の訪問者です。キリ番踏んだ方はコメントかTweetしてください</marquee>", i.viewcount)),
				),
			),
			app.Tr().Attr("height", "600").Attr("width", "960").Attr("colspan", "3").Attr("align", "right").Attr("valign", "bottom").Body(
				app.Td().Attr("width", "960").Attr("height", "600").Attr("colspan", "3").Body(
					app.Img().Src("/web/assets/admin.jpg").Alt("サイトの管理者の画像"),
				),
			),
		),
	)
}

func (i *Index) navigate() app.UI {
	return app.Table().Attr("width", "960").Attr("cellpadding", "0").Attr("cellspacing", "0").Body(
		app.TBody().Body(
			app.Tr().Body(
				app.Td().Attr("width", "100").Body(),
				app.Td().Attr("width", "100").Body(
					app.A().Attr("href", "/about").Body().Text("自己紹介"),
				),
				app.Td().Attr("width", "100").Body(
					app.A().Attr("href", "/portfolio").Body().Text("ポートフォリオ"),
				),
				app.Td().Attr("width", "100").Body(),
			),
		),
	)
}
