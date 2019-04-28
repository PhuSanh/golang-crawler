package crawler

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"go-crawler/config"
	"go-crawler/models"
	"log"
	"net/http"
	"sync"
)

var wg sync.WaitGroup
var count int
var countLink int

func GetListNews() (err error) {

	fmt.Println("running crawl list ...")
	crawlerUrl := config.Config.Crawler.Url
	for i:= 1; i <= 500; i++ {
		wg.Add(1)
		urlPage := fmt.Sprintf("%s/p%d", crawlerUrl, i)
		go getListNewsPerPage(urlPage)
	}

	wg.Wait()
	fmt.Println("DONE: ", count, countLink)

	return
}

func GetNewsContent() (err error) {
	fmt.Println("running crawl content ...")
	listNews, err := models.GetList()
	if err != nil {
		return
	}

	for _, v := range listNews {
		wg.Add(1)
		go getContentAndUpdateDB(v)
	}

	wg.Wait()
	fmt.Println("DONE")

	return
}

func getContentAndUpdateDB(news models.News) {
	doc, err := getDom(news.Link)
	if err != nil {
		wg.Done()
		return
	}

	html, err := doc.Find("#left_calculator .content_detail").Html()
	if err != nil {
		log.Println("[Error]: Get content: ", err.Error())
		wg.Done()
		return
	}

	news.Content = html
	fmt.Println(news.ID)
	_ = models.UpdateById(news)
	wg.Done()
}

func getDom(url string) (doc *goquery.Document, err error) {
	res, err := http.Get(url)
	if err != nil {
		log.Println("[Error]: ", err.Error())
		return doc, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Println("[Error] error code: ", res.StatusCode, url)
		return doc, err
	}

	doc, err = goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Println("[Error]: ", err.Error())
		return
	}

	return doc, nil
}

func getListNewsPerPage(url string) {

	doc, err := getDom(url)
	if err != nil {
		wg.Done()
		return
	}

	countLink += 1

	doc.Find(".container .sidebar_1 .list_news").Each(func(i int, s *goquery.Selection) {
		title := s.Find(".title_news a:first-child").Text()
		link, _ := s.Find(".title_news a:first-child").Attr("href")
		thumb, _ := s.Find(".thumb_art img").Attr("src")
		description := s.Find(".description").Text()
		news := &models.News{
			Title: title,
			Link: link,
			Thumb: thumb,
			Description: description,
		}
		models.Insert(news)
		count += 1
		wg.Done()
	})
}

