package rpc

import (
	"apisrv/pkg/newsportal"
)

//go:generate colgen -imports=apisrv/pkg/newsportal -funcpkg=newsportal
//colgen:News:MapP(newsportal)
//colgen:NewsSummary:MapP(newsportal.News)
//colgen:Tag:MapP(newsportal)
//colgen:Category:MapP(newsportal)
//colgen:User:MapP(newsportal)

func NewCategory(in *newsportal.Category) *Category {
	if in == nil {
		return nil
	}
	return &Category{
		ID:    in.ID,
		Title: in.Title,
	}
}

func NewNews(in *newsportal.News) *News {
	if in == nil {
		return nil
	}
	return &News{
		ID:          in.ID,
		Title:       in.Title,
		Author:      in.Author,
		Content:     in.Content,
		PublishedAt: in.PublishedAt,
		Category:    NewCategory(in.Category),
	}
}

func NewTag(in *newsportal.Tag) *Tag {
	if in == nil {
		return nil
	}
	return &Tag{
		ID:    in.ID,
		Title: in.Title,
	}
}

func NewNewsSummary(in *newsportal.News) *NewsSummary {
	if in == nil {
		return nil
	}
	return &NewsSummary{
		ID:          in.ID,
		Title:       in.Title,
		PublishedAt: in.PublishedAt,
		Category:    NewCategory(in.Category),
		Tags:        NewTags(in.Tags),
	}
}

func NewUser(in *newsportal.User) *User {
	if in == nil {
		return nil
	}
	return &User{
		Login:          in.Login,
		LastActivityAt: in.LastActivityAt,
	}
}
