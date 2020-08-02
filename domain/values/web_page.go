package values

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

type WebPage struct {
	Url   string
	Title string
}

func NewWebPage(url string) (WebPage, error) {
	title, err := getTitle(url)
	if err != nil {
		return WebPage{}, err
	}

	return WebPage{
		Url:   url,
		Title: title,
	}, nil
}

func getTitle(url string) (string, error) {
	r := regexp.MustCompile(`http(s)?://([\w-]+\.)+[\w-]+(/[\w- ./?%&=~]*)?`)

	if !r.MatchString(url) {
		return "", errors.New(fmt.Sprintf("Invalid URL regexp: %v", url))
	}

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	bs, _ := ioutil.ReadAll(resp.Body)
	s := string(bs)

	reg := regexp.MustCompile(`<title>(.+)</title>`)

	t := reg.FindString(s)
	t = strings.Replace(t, "</title>", "", 1)
	t = strings.Replace(t, "<title>", "", 1)

	return t, nil
}
