package rpc

import (
	"apisrv/pkg/newsportal"
	"time"
)

type (
	queryParams struct {
		CategoryId int `query:"categoryId"`
		TagId      int `query:"tagId"`
		PageSize   int `query:"pageSize"`
		Page       int `query:"page"`
	}

	Tag struct {
		ID    int    `json:"tagID"`
		Title string `json:"title"`
	}

	Category struct {
		ID    int    `json:"categoryId"`
		Title string `json:"title"`
	}

	News struct {
		ID          int       `json:"newsID"`
		Title       string    `json:"title"`
		Content     *string   `json:"content"`
		Author      string    `json:"author"`
		Category    *Category `json:"category"`
		Tags        []Tag     `json:"tags"`
		PublishedAt time.Time `json:"publishedAt"`
	}

	NewsSummary struct {
		ID          int       `json:"newsSummaryID"`
		Title       string    `json:"title"`
		PublishedAt time.Time `json:"publishedAt"`
		Category    *Category `json:"category"`
		Tags        []Tag     `json:"tags"`
	}

	User struct {
		Login          string     `json:"login"`
		LastActivityAt *time.Time `json:"lastActivityAt"`
	}

	NewsInput struct {
		Id          *int    `json:"newsID"`
		Title       string  `json:"title"`
		Content     *string `json:"content"`
		Author      string  `json:"author"`
		CategoryID  int     `json:"categoryID"`
		Tags        []int   `json:"tagIds"`
		PublishedAt string  `json:"publishedAt"`
	}

	TagInput struct {
		ID    *int   `json:"tagID"`
		Title string `json:"title"`
	}

	CategoryInput struct {
		ID          *int   `json:"categoryID"`
		Title       string `json:"title"`
		OrderNumber *int   `json:"orderNumber"`
	}
)

func (q *queryParams) NewFilter() newsportal.Filters {
	return newsportal.NewFilters(
		q.CategoryId, q.TagId,
		q.PageSize, q.Page,
	)
}

func newsToManager(in *NewsInput) *newsportal.NewsInput {
	layout := "2006-01-02 15:04:05.000 -0700"
	timeP, _ := time.Parse(layout, in.PublishedAt)

	return &newsportal.NewsInput{
		Id:          in.Id,
		Title:       in.Title,
		Content:     in.Content,
		Author:      in.Author,
		CategoryID:  in.CategoryID,
		TagIDs:      in.Tags,
		PublishedAt: &timeP,
	}
}

func categoryToManager(in *CategoryInput) *newsportal.CategoryInput {
	return &newsportal.CategoryInput{
		ID:          in.ID,
		Title:       in.Title,
		OrderNumber: in.OrderNumber,
	}
}

func tagToManager(in *TagInput) *newsportal.TagInput {
	return &newsportal.TagInput{
		ID:    in.ID,
		Title: in.Title,
	}
}
