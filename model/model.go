package model

// Portfolio はポートフォリオの情報を表します。
type Portfolio struct {
	// ポートフォリオのタイトル ex. Hack-Portal
	Title string
	// ポートフォリオの説明 ex. Hack-Portalは、ハッカソンの情報を共有するためのポータルサイトです。
	About string
	// ポートフォリオのURL ex. https://hack-portal.com
	DeploymentURL string
	// ポートフォリオのリポジトリURL ex.
	RepositoryURL string
	// ポートフォリオの技術スタック
	TechStack []TechStack
	// ポートフォリオのスクリーンショット
	Screenshots []string
	// ポートフォリオの作成日
	CreatedAt string
	// ポートフォリオの更新日
	UpdatedAt string
}

// MySite は自分のポートフォリオサイトの情報を表します。
type MySite struct {
	// 閲覧カウント
	Count string
	// 自分の技術スタック
	TechStack []TechStack
	// 自分のリンク
	Links []Link
	// サイトについたコメント
	Comments []Comment
}

// TechStack は技術スタックの情報を表します。
type TechStack struct {
	// 技術スタックの名前
	Name string
	// 技術スタックのアイコン
	Icon string
}

// Link はリンクの情報を表します。
type Link struct {
	// リンクの名前
	Name string
	// リンクのURL
	URL string
}

// Comment はコメントの情報を表します。
type Comment struct {
	// コメントのID
	ID string
	// コメントの名前
	Name string
	// コメントの内容
	Content string
	// コメントの作成日
	CreatedAt string
}
