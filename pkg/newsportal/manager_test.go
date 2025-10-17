package newsportal

import (
	"apisrv/pkg/db"
	"apisrv/pkg/db/test"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestManager_GetNewsById(t *testing.T) {
	dbc, _ := test.Setup(t)

	n1, cleaner1 := test.News(t, dbc, &db.News{Title: "test"}, test.WithFakeNews, test.WithNewsRelations)
	defer cleaner1()
	n2, cleaner2 := test.News(t, dbc, &db.News{Title: "test"}, test.WithFakeNews, test.WithNewsRelations)
	defer cleaner2()
	n3, cleaner3 := test.News(t, dbc, &db.News{Title: "test"}, test.WithFakeNews, test.WithNewsRelations)
	defer cleaner3()

	m := &Manager{repoNews: db.NewNewsRepo(dbc)}

	res1, err := m.GetNewsById(t.Context(), n1.ID)
	if err != nil {
		t.Error(err)
	}
	if res1.Title != "test" {
		t.Errorf("res1.ID = %v, want %v", res1.ID, n1.ID)
	}

	res2, err := m.GetNewsById(t.Context(), n2.ID)
	if err != nil {
		t.Error(err)
	}
	if res2.Title != "test" {
		t.Errorf("res2.ID = %v, want %v", res2.ID, n2.ID)
	}

	res3, err := m.GetNewsById(t.Context(), n3.ID)
	if err != nil {
		t.Error(err)
	}
	if res3.Title != "test" {
		t.Errorf("res3.ID = %v, want %v", res3.ID, n3.ID)
	}
}

func TestManager_GetNewsByFilters(t *testing.T) {
	dbc, _ := test.Setup(t)

	timeStr := "2025-09-17 12:10:00.000"
	layout := "2006-01-02 15:04:05"
	parseT, _ := time.Parse(layout, timeStr)

	n1, cleaner1 := test.News(t, dbc, &db.News{Title: "test", CategoryID: 12, TagIDs: []int{14, 15}, PublishedAt: parseT}, test.WithFakeNews, test.WithNewsRelations)
	defer cleaner1()
	n2, cleaner2 := test.News(t, dbc, &db.News{Title: "test", CategoryID: 12, TagIDs: []int{15, 16}, PublishedAt: parseT}, test.WithFakeNews, test.WithNewsRelations)
	defer cleaner2()
	n3, cleaner3 := test.News(t, dbc, &db.News{Title: "test", CategoryID: 12, TagIDs: []int{16, 14}}, test.WithFakeNews, test.WithNewsRelations)
	defer cleaner3()

	tests := []struct {
		name    string
		args    Filters
		want    []News
		wantErr bool
	}{
		{
			name: "test 1",
			args: Filters{
				CategoryId: 12,
				TagId:      15,
			},
			want: []News{
				{News: *n1},
				{News: *n2},
			},
			wantErr: false,
		},
		{
			name: "test 1",
			args: Filters{
				CategoryId: 12,
				TagId:      16,
			},
			want: []News{
				{News: *n2},
				{News: *n3},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Manager{repoNews: db.NewNewsRepo(dbc)}

			list, err := m.GetNewsByFilters(t.Context(), tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetNewsByFilters() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if len(list) == 0 {
				t.Error("not foud:", len(list))
				return
			}

			for i, v := range list {
				if v.Title != tt.want[i].Title {
					t.Errorf("GetNewsByFilters() list[i].Title = %v, want %v", v.Title, tt.want[i].Title)
				}
				if v.Category.ID != tt.want[i].CategoryID {
					t.Errorf("GetNewsByFilters() list[i].Category = %v, want %v", v.Category, tt.want[i].Category)
				}
				timeNow := time.Now()
				if v.PublishedAt.After(timeNow) {
					t.Errorf("GetNewsByFilters() list[i].PublishedAt = %v, want %v", v.PublishedAt, timeNow)
				}
				for j, tagID := range list[i].TagIDs {
					if len(tt.want[i].TagIDs) < j {
						t.Errorf("GetNewsByFilters() list[i].TagIDs = %v, want %v", v.TagIDs, tagID)
						break
					}
					if tagID != tt.want[i].TagIDs[j] {
						t.Errorf("GetNewsByFilters() list[i].TagID = %v, want %v", list[i].TagIDs[i], tagID)
					}
				}
			}
		})
	}
}
