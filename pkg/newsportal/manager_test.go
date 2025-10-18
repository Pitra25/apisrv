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

func createTestData(t *testing.T) ([]db.News, []test.Cleaner) {
	dbc, _ := test.Setup(t)

	var (
		res     []db.News
		c       []test.Cleaner
		timeStr = "2025-10-17 12:10:00.000"
		layout  = "2006-01-02 15:04:05"
	)

	parseT, _ := time.Parse(layout, timeStr)

	n1, cleaner1 := test.News(t, dbc, &db.News{Title: "test", CategoryID: 12, TagIDs: []int{14, 15}, PublishedAt: parseT}, test.WithFakeNews, test.WithNewsRelations)
	n2, cleaner2 := test.News(t, dbc, &db.News{Title: "test", CategoryID: 12, TagIDs: []int{15, 16}, PublishedAt: parseT}, test.WithFakeNews, test.WithNewsRelations)
	n3, cleaner3 := test.News(t, dbc, &db.News{Title: "test", CategoryID: 12, TagIDs: []int{16, 14}}, test.WithFakeNews, test.WithNewsRelations)

	res = append(res, *n1, *n2, *n3)
	c = append(c, cleaner1, cleaner2, cleaner3)

	return res, c
}

func TestManager_GetNewsById(t *testing.T) {
	dbc, _ := test.Setup(t)

	list, cleaner := createTestData(t)
	defer func() {
		for _, v := range cleaner {
			v()
		}
	}()

	m := &Manager{repoNews: db.NewNewsRepo(dbc)}

	res1, err := m.GetNewsByID(t.Context(), list[0].ID)
	if err != nil {
		t.Error(err)
	}
	if res1.Title != "test" {
		t.Errorf("res1.ID = %v, want %v", res1.ID, list[0].ID)
	}

	res2, err := m.GetNewsByID(t.Context(), list[1].ID)
	if err != nil {
		t.Error(err)
	}
	if res2.Title != "test" {
		t.Errorf("res2.ID = %v, want %v", res2.ID, list[1].ID)
	}

	res3, err := m.GetNewsByID(t.Context(), list[2].ID)
	if err != nil {
		t.Error(err)
	}
	if res3.Title != "test" {
		t.Errorf("res3.ID = %v, want %v", res3.ID, list[2].ID)
	}
}

func TestManager_GetNewsByFilters(t *testing.T) {
	dbc, _ := test.Setup(t)

	testData, cleaner := createTestData(t)
	defer func() {
		for _, v := range cleaner {
			v()
		}
	}()

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
				{News: testData[0]},
				{News: testData[1]},
			},
			wantErr: false,
		},
		{
			name: "test 2",
			args: Filters{
				CategoryId: 12,
				TagId:      16,
			},
			want: []News{
				{News: testData[1]},
				//{News: testData[2]}, // Раскомментировать если надо проверить на фильтр даты публикации
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Manager{repoNews: db.NewNewsRepo(dbc)}

			list, err := m.GetNewsByFilters(t.Context(), tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if len(list) == 0 {
				t.Error("not foud:", len(list))
				return
			}

			if len(list) != len(tt.want) {
				t.Errorf("list = %v, want %v", len(list), len(tt.want))
			}

			for i, v := range list {
				if v.Title != tt.want[i].Title {
					t.Errorf("list[i].Title = %v, want %v", v.Title, tt.want[i].Title)
				}
				if v.Category.ID != tt.want[i].CategoryID {
					t.Errorf("list[i].Category = %v, want %v", v.Category, tt.want[i].Category)
				}
				for j, tagID := range list[i].TagIDs {
					if len(tt.want[i].TagIDs) <= j {
						t.Errorf("list[i].TagIDs = %v, want %v", v.TagIDs, tagID)
						break
					}
					if tagID != tt.want[i].TagIDs[j] {
						t.Errorf("list[i].TagID = %v, want %v", list[i].TagIDs[i], tagID)
					}
				}
			}
		})
	}
}
