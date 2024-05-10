package main

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"
)

type RssFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Language    string    `xml:"language"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func urlToFeed(url string) (RssFeed, error) {

	httpClient := http.Client{
		Timeout: time.Second * 10, // Maximum of 2 secs
	}

	resp, err := httpClient.Get(url)
	if err != nil {
		return RssFeed{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RssFeed{}, err
	}

	rssFeed := RssFeed{}
	err = xml.Unmarshal(data, &rssFeed)
	if err != nil {
		return RssFeed{}, err
	}
	return rssFeed, nil
}
