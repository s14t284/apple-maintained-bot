package domain

// TODO: 旧フォーマットであるため、blockを使うように修正する

// Message slackへ送る投稿フォーマット
type Message struct {
	Payload Payload `json:"payload"`
}

// Payload slackへ送る投稿の中身
type Payload struct {
	Channel     string      `json:"channel"`
	UserName    string      `json:"username"`
	Text        string      `json:"text"`
	Attachments Attachments `json:"attachments"`
	IconEmoji   string      `json:"icon_emoji"`
}

// Attachments slackへ送るメッセージの本体
type Attachments []Attachment

// Attachment slackへ送るメッセージの1つ
type Attachment struct {
	FallBack   string `json:"fallback"`
	Color      string `json:"color"`
	PreText    string `json:"pretext"`
	AuthorName string `json:"author_name"`
	AuthorLink string `json:"author_link"`
	Title      string `json:"title"`
	TitleLink  string `json:"title_link"`
	Text       string `json:"text"`
}
