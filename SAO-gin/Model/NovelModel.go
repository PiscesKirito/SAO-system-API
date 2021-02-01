package model

type NovelFilter struct {
	Key        string `json:"bookNum", form:"Key"`
	ChapterNum int    `json:"chapterNum", form:"chapterNum"`
}
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

type NovelChapterNumFilter struct {
	Key string `json:"Key", form:"Key"`
}
type NovelChapterNum struct {
	ChapterNum int `db:"chapterNum", json:"chapterNum"`
}
