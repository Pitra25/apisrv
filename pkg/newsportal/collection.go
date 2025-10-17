package newsportal

import "apisrv/pkg/db"

//go:generate colgen -imports=apisrv/pkg/db
//colgen:News,Category,Tag,User
//colgen:News:UniqueTagIDs,MapP(db)
//colgen:Tag:MapP(db)
//colgen:Category:MapP(db)
//colgen:User:MapP(db)

// MapP converts slice of type T to slice of type M with given converter with pointers.
func MapP[T, M any](a []T, f func(*T) *M) []M {
	n := make([]M, len(a))
	for i := range a {
		n[i] = *f(&a[i])
	}
	return n
}

func (nl NewsList) SetTags(tags Tags) {
	tagIndex := tags.Index()
	for i, v := range nl {
		for _, tag := range v.TagIDs {
			if t, ok := tagIndex[tag]; ok {
				nl[i].Tags = append(nl[i].Tags, t)
			}
		}
	}
}

func NewCategory(in *db.Category) *Category {
	if in == nil {
		return nil
	}
	return &Category{
		Category: *in,
	}
}

func NewNews(in *db.News) *News {
	if in == nil {
		return nil
	}
	return &News{
		News:     *in,
		Category: NewCategory(in.Category),
	}
}

func NewTag(in *db.Tag) *Tag {
	if in == nil {
		return nil
	}
	return &Tag{Tag: *in}
}

func NewUser(in *db.User) *User {
	if in == nil {
		return nil
	}
	return &User{User: *in}
}
