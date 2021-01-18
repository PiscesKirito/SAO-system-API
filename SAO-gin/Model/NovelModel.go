package model

// Novel 小说实体
type Novel struct {
	BookName    string `db:"BookName" json:"book"`
	ChapterName *string `db:"ChapterName" json:"chapter"`
	Author      string `db:"Author" json:"author"`
	Translator  string `db:"Translator" json:"translator"`
	Novel       string `db:"Novel" json:"content"`
}
