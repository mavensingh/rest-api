package models

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

var CurrentID int

type Article struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	SubTitle  string    `json:"subtitle"`
	CreatedAt time.Time `json:"created_at"`
}

// make Article accessable for all the packages for additional use(like get article by id,get all articles from memory)
var Articles []Article

// Give us some seed data
func init() {
	CreateArticle(Article{Title: "Learning Golang", Content: "Go", SubTitle: "Golang"})
	CreateArticle(Article{Title: "Golang Data Types", Content: "Go", SubTitle: "Golang"})
}

func GetArticleByID(id int) Article {
	for _, article := range Articles {
		if article.Id == id {
			return article
		}
	}
	// return empty Todo if not found
	return Article{}
}

// CreateArticle with logic to generate unique id based content
func CreateArticle(article Article) (Article, error) {
	if len(Articles) == 0 {
		CurrentID += 1
		article.Id = CurrentID
		article.CreatedAt = time.Now()
		Articles = append(Articles, article)
		return article, nil
	}
	NotExists := []int{}
	for _, keys := range Articles {
		if keys.Id == article.Id {
			return Article{}, fmt.Errorf("article %d id already exists", keys.Id)
		} else {
			NotExists = append(NotExists, article.Id)
		}
		if article.Id == 0 {
			CurrentID += 1
			article.Id = CurrentID
			article.CreatedAt = time.Now()
			Articles = append(Articles, article)
			return article, nil
		}
	}
	for _, id := range NotExists {
		CurrentID = id
		Articles = append(Articles, article)
		return article, nil
	}
	return Article{}, nil
}

func SearchArticles(query string) []Article {
	var articles []Article
	for _, keys := range Articles {
		if strings.Contains(strings.ToLower(keys.Title), query) || strings.Contains(strings.ToLower(keys.SubTitle), query) || strings.Contains(strings.ToLower(keys.Content), query) {
			articles = append(articles, keys)
			continue
		}
	}
	return articles
}

func PaginationLogic(Limit, LastCheckedID string) ([]Article, error) {
	limit := 0
	LastId := 0
	if Limit == "" || LastCheckedID == "" {
		limit = 2
		LastId = 1
	} else {
		newLimit, err := strconv.Atoi(Limit)
		if err != nil {
			return nil, err
		}
		limit = newLimit
		LastId, err = strconv.Atoi(LastCheckedID)
		if err != nil {
			return nil, err
		}
	}

	if limit < 0 || LastId < 0 {
		return Articles[0:2], nil
	}

	// handling two way if id not found then index will gonna be zero.
	index := 0
	for idx, article := range Articles {
		if article.Id == LastId {
			index = idx
			break
		}
	}
	// if id and index are greater then
	if index+limit > len(Articles) {
		return Articles[index : len(Articles)-1], nil
	}

	articles := Articles[index : index+limit]
	return articles, nil
}
