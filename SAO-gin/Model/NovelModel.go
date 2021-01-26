package model

// Novel 小说实体
type Novel struct {
	BookName   string `db:"BookName" json:"book"`
	Author     string `db:"Author" json:"author"`
	Translator string `db:"Translator" json:"translator"`
	Novel      string `db:"Novel" json:"content"`
}

type NovelList struct {
	BookNum  string `db:"Key" json:"book_num"`
	BookName string `db:"BookName" json:"book"`
	Url      string `db:"Url" json:"url"`
}
