package newsportal

import (
	"apisrv/pkg/db"
	"context"
	"fmt"
)

type Manager struct {
	repoNews db.NewsRepo
	repoUser db.CommonRepo
	repoVFS  db.VfsRepo
}

func NewManager(dbc *db.DB) *Manager {
	return &Manager{
		repoNews: db.NewNewsRepo(dbc),
		repoUser: db.NewCommonRepo(dbc),
		repoVFS:  db.NewVfsRepo(dbc),
	}
}

/*** News ***/

func (m *Manager) GetNewsByFilters(ctx context.Context, fil Filters) ([]News, error) {
	dbNews, err := m.repoNews.NewsByFilters(
		ctx, fil.NewsToDB(), db.Pager{Page: fil.Page, PageSize: fil.PageSize},
		db.WithColumns(db.Columns.News.Category),
	)
	if err != nil {
		return nil, fmt.Errorf("news fetch failed: %w", err)
	}
	result := NewNewsList(dbNews)

	// collect everything in 1 news
	tags, err := m.GetTagsByID(ctx, result.UniqueTagIDs())
	if err != nil {
		return nil, fmt.Errorf("tags fetch failed: %w", err)
	}

	// collect everything in a news array
	result.SetTags(tags)

	return result, nil
}

func (m *Manager) GetNewsByID(ctx context.Context, id int) (*News, error) {
	// receiving news by ID
	news, err := m.repoNews.NewsByID(ctx, id, db.WithColumns(db.Columns.News.Category))
	if err != nil {
		return nil, fmt.Errorf("news fetch failed: %w", err)
	}
	result := NewNews(news)

	tags, err := m.GetTagsByID(ctx, result.TagIDs)
	if err != nil {
		return nil, fmt.Errorf("tags fetch failed: %w", err)
	}

	result.Tags = tags

	return result, nil
}

func (m *Manager) GetNewsCount(ctx context.Context, fil Filters) (int, error) {
	return m.repoNews.CountNews(ctx, fil.NewsToDB())
}

/*** Category ***/

func (m *Manager) GetAllCategory(ctx context.Context) ([]Category, error) {
	categories, err := m.repoNews.CategoriesByFilters(ctx, &db.CategorySearch{}, db.PagerNoLimit)

	return NewCategories(categories), err
}

/*** Tag ***/

func (m *Manager) GetAllTag(ctx context.Context) ([]Tag, error) {
	tags, err := m.repoNews.TagsByFilters(ctx, &db.TagSearch{}, db.PagerNoLimit)

	return NewTags(tags), err
}

func (m *Manager) GetTagsByID(ctx context.Context, ids []int) (Tags, error) {
	fil := db.TagSearch{}
	if len(ids) > 0 {
		fil.IDs = ids
	}
	tags, err := m.repoNews.TagsByFilters(
		ctx, &fil, db.PagerNoLimit,
	)

	return NewTags(tags), err
}

/*** User ***/

func (m *Manager) GetUserByID(ctx context.Context, id int) (*User, error) {
	item, err := m.repoUser.UserByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("user fetch failed: %w", err)
	}
	return NewUser(item), err
}
