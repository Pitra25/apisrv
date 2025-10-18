package newsportal

import (
	"apisrv/pkg/db"
	"time"
)

type (
	Filters struct {
		CategoryId int
		TagId      int
		PageSize   int
		Page       int
		UserId     int
	}

	Tag      struct{ db.Tag }
	TagInput struct {
		ID    *int
		Title string
	}

	Category      struct{ db.Category }
	CategoryInput struct {
		ID          *int
		Title       string
		OrderNumber *int
	}

	News struct {
		db.News
		Category *Category
		Tags     []Tag
	}
	NewsInput struct {
		Id          *int
		Title       string
		Content     *string
		Author      string
		CategoryID  int
		TagIDs      []int
		PublishedAt *time.Time
	}

	User struct{ db.User }
)

func NewFilters(categoryId, tagId, pageSize, page int) Filters {
	return Filters{
		CategoryId: categoryId,
		TagId:      tagId,
		PageSize:   pageSize,
		Page:       page,
	}
}

func (f *Filters) NewsToDB() *db.NewsSearch {
	statusID := db.StatusEnabled
	timeNow := time.Now()
	filter := db.NewsSearch{
		StatusIDEQuals: &statusID,
		PublishedAtLE:  &timeNow,
	}
	if f.TagId != 0 {
		filter.TagID = &f.TagId
	}
	if f.CategoryId != 0 {
		filter.CategoryID = &f.CategoryId
	}
	return &filter
}

func (f *Filters) UserToDB() *db.UserSearch {
	statusID := db.StatusEnabled
	filter := db.UserSearch{
		StatusIDEQuals: &statusID,
	}
	if f.UserId != 0 {
		filter.ID = &f.UserId
	}
	return &filter
}
