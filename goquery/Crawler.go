package goquery

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Tmp struct {
	title string
	time  string
}

func RequestGoquery(CrawlURL string) (tmp map[int]string, err error) {
	var li []Tmp
	tmp = make(map[int]string)
	doc, err := goquery.NewDocument(CrawlURL)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
		return nil, err
	}

	title := doc.Find("ul.m-list_contest li")
	title.Each(func(i int, s *goquery.Selection) {
		conStatus := s.Find("div.status").Text()
		if conStatus == "予定" {
			conTime := s.Find("time.fixtime-short").Text()
			conTitle := s.Find("div.m-list_contest_ttl").Text()
			t, err := time.Parse("2006-01-02 15:04:05+0900", conTime)
			if err != nil {
				log.Fatalf("failed to serve: %v", err)
				return
			}
			li = append(li, Tmp{title: strings.TrimSpace(conTitle), time: t.Format("2006年01月02日 15時04分")})
			for i, t := range li {
				tmp[i] = fmt.Sprintf("%s\n日時:%s", t.title, t.time)
			}
		}
	},
	)
	return tmp, nil
}
