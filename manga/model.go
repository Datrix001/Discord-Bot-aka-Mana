package manga

type Comic struct {
	Name          string
	Author        string
	LatestChapter string
	CoverURL      string
}

type Chapter struct {
	MangaId   string
	ChapterId string
	Url       string
}
