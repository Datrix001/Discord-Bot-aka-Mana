package sources

import (
	"errors"
	"fmt"
	"log"
	"strings"

	m "discordBot/bot/manga"

	"github.com/gocolly/colly"
)

func GetManga(name string) (m.Comic, error) {
	comicName := normalize(name)
	comicName = comicName + "-" + "6f7fe6eb"
	var comic m.Comic

	c := colly.NewCollector(colly.AllowedDomains("asurascans.com"))

	comic.Name = name

	// ==================================================
	// GET AUTHOR
	// ==================================================

	c.OnHTML(`a[href^="/browse?author"]`, func(h *colly.HTMLElement) {

		if h.Request.URL.Path == "/comics/" {
			return
		}

		comic.Author = strings.TrimSpace(h.Text)

		fmt.Printf(
			"Author Name: %s\n",
			comic.Author,
		)
	})

	// ==================================================
	// GET CHAPTER NAME
	// ==================================================

	c.OnHTML("h2.text-lg.font-bold", func(h *colly.HTMLElement) {

		if h.Request.URL.Path == "/comics/" {
			return
		}

		comic.LatestChapter = strings.TrimSpace(h.Text)

		fmt.Printf(
			"Chapter: %s\n",
			comic.LatestChapter,
		)
	})

	// ==================================================
	// GET COVER IMAGE
	// ==================================================

	c.OnHTML("#cover-viewer-img", func(h *colly.HTMLElement) {

		if h.Request.URL.Path == "/comics/" {
			return
		}

		link := h.Attr("data-full-src")

		comic.CoverURL = link

		fmt.Printf(
			"Image link: %s\n",
			link,
		)

	})

	url := fmt.Sprintf("https://asurascans.com/comics/%s", comicName)
	fmt.Println(url)
	err := c.Visit(url)
	if err != nil {
		// log.Panic(err)

		return comic, errors.New("No Manhwa like this exist!!")
	}
	c.Wait()
	fmt.Println("Done With Something")

	return comic, nil
}

func GetChapter(name string, chapNum string) (m.Chapter, error) {

	var chap m.Chapter

	chap.MangaId = name
	chap.ChapterId = chapNum

	c := colly.NewCollector(colly.AllowedDomains("asurascans.com"))

	// ==================================================
	// GET TOP CHAPTER URL
	// ==================================================

	c.OnHTML(`div.divide-y a[href*="/chapter/"]`, func(h *colly.HTMLElement) {

		link := h.Attr("href")

		fmt.Printf(
			"Top Chapter: %s -> %s\n\n",
			strings.TrimSpace(h.Text),
			link,
		)

		chap.Url = link

		err := h.Request.Visit(
			h.Request.AbsoluteURL(link),
		)

		if err != nil {
			log.Println(
				"Failed to visit chapter:",
				err,
			)
		}
	})
	name = normalize(name) + "-" + "6f7fe6eb"
	url := fmt.Sprintf("https://asurascans.com/comics/%s/chapter/%s", name, chapNum)
	fmt.Println(url)

	err := c.Visit(url)
	if err != nil {
		// log.Panic(err)

		return chap, errors.New("Failed to visit chapter")
	}

	c.Wait()
	fmt.Println("Done With Something")

	return chap, nil
}

func normalize(name string) string {
	name = strings.ReplaceAll(name, " ", "-")
	name = strings.ReplaceAll(name, "_", "-")
	name = strings.ToLower(name)
	return name
}
