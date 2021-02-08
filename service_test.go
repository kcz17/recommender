package recommender

import (
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/go-kit/kit/log"
	"github.com/jmoiron/sqlx"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

var (
	s1 = Sock{ID: "1", Name: "name1", Description: "description1", Price: 1.1, Count: 1, ImageURL: []string{"ImageUrl_11", "ImageUrl_21"}, ImageURL_1: "ImageUrl_11", ImageURL_2: "ImageUrl_21", Tags: []string{"odd", "prime"}, TagString: "odd,prime"}
	s2 = Sock{ID: "2", Name: "name2", Description: "description2", Price: 1.2, Count: 2, ImageURL: []string{"ImageUrl_12", "ImageUrl_22"}, ImageURL_1: "ImageUrl_12", ImageURL_2: "ImageUrl_22", Tags: []string{"even", "prime"}, TagString: "even,prime"}
	s3 = Sock{ID: "3", Name: "name3", Description: "description3", Price: 1.3, Count: 3, ImageURL: []string{"ImageUrl_13", "ImageUrl_23"}, ImageURL_1: "ImageUrl_13", ImageURL_2: "ImageUrl_23", Tags: []string{"odd", "prime"}, TagString: "odd,prime"}
	s4 = Sock{ID: "4", Name: "name4", Description: "description4", Price: 1.4, Count: 4, ImageURL: []string{"ImageUrl_14", "ImageUrl_24"}, ImageURL_1: "ImageUrl_14", ImageURL_2: "ImageUrl_24", Tags: []string{"even"}, TagString: "even"}
	s5 = Sock{ID: "5", Name: "name5", Description: "description5", Price: 1.5, Count: 5, ImageURL: []string{"ImageUrl_15", "ImageUrl_25"}, ImageURL_1: "ImageUrl_15", ImageURL_2: "ImageUrl_25", Tags: []string{"odd", "prime"}, TagString: "odd,prime"}
)

var logger log.Logger

func TestRecommenderServiceGet(t *testing.T) {
	logger = log.NewLogfmtLogger(os.Stderr)
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening stub database connection", err)
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	var cols = []string{"id", "name", "description", "price", "count", "image_url_1", "image_url_2", "tag_name"}

	// Test Case 1
	mock.ExpectQuery("SELECT *").WillReturnRows(sqlmock.NewRows(cols))

	// Test Case 2
	mock.ExpectQuery("SELECT *").WillReturnRows(sqlmock.NewRows(cols).
		AddRow(s1.ID, s1.Name, s1.Description, s1.Price, s1.Count, s1.ImageURL[0], s1.ImageURL[1], strings.Join(s1.Tags, ",")).
		AddRow(s2.ID, s2.Name, s2.Description, s2.Price, s2.Count, s2.ImageURL[0], s2.ImageURL[1], strings.Join(s2.Tags, ",")).
		AddRow(s3.ID, s3.Name, s3.Description, s3.Price, s3.Count, s3.ImageURL[0], s3.ImageURL[1], strings.Join(s3.Tags, ",")).
		AddRow(s4.ID, s4.Name, s4.Description, s4.Price, s4.Count, s4.ImageURL[0], s4.ImageURL[1], strings.Join(s4.Tags, ",")).
		AddRow(s5.ID, s5.Name, s5.Description, s5.Price, s5.Count, s5.ImageURL[0], s5.ImageURL[1], strings.Join(s5.Tags, ",")))

	s := NewRecommenderService(sqlxDB, logger)

	// Test Case 1
	have, err := s.Get()
	if err != nil {
		t.Errorf("Get(): returned error %s", err.Error())
	}

	if !reflect.DeepEqual(Sock{}, have) {
		t.Errorf("Get(): want %v, have %v", Sock{}, have)
	}

	// Test Case 2
	have, err = s.Get()
	if err != nil {
		t.Errorf("Get(): returned error %s", err.Error())
	}

	if have.ID == "" {
		t.Errorf("Get(): want non-empty ID, have %s", have.ID)
	}
}
