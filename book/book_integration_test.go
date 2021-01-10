// +build integration

package book_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/elliotforbes/go-fiber-tutorial/book"
	"github.com/elliotforbes/go-fiber-tutorial/database"
	"github.com/elliotforbes/go-fiber-tutorial/transport"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type BookTestSuite struct {
	suite.Suite
	dbConn *gorm.DB
	app    *fiber.App
}

func (suite *BookTestSuite) SetupSuite() {
	var err error
	suite.dbConn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Database connection successfully opened")

	suite.app = transport.Setup()
	database.InitDatabase()
	database.DBConn.AutoMigrate(&book.Book{})
}

func (suite *BookTestSuite) TestCreateBook() {
	req := httptest.NewRequest(
		"POST",
		"/api/v1/book",
		strings.NewReader(`{"ISIN": 12345, "title":"Test Book", "author": "Elliot", "rating": 5}`),
	)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.FormatInt(req.ContentLength, 10))

	res, err := suite.app.Test(req, -1)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, res.StatusCode)

	var bookTest book.Book
	database.DBConn.Where("title = ?", "Test Book").First(&bookTest)
	fmt.Println(bookTest)
	assert.Equal(suite.T(), bookTest.Title, "Test Book")
}

func (suite *BookTestSuite) TestReadBook() {
	req := httptest.NewRequest(
		"GET",
		"/api/v1/book/1",
		nil,
	)
	res, err := suite.app.Test(req, -1)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, res.StatusCode)

	var testbook book.Book
	body, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &testbook)

	assert.Equal(suite.T(), "Test Book", testbook.Title)
}

func (suite *BookTestSuite) TestDeleteBook() {
	req := httptest.NewRequest(
		"POST",
		"/api/v1/book",
		strings.NewReader(`{"ISIN": 45678, "title":"Test Book", "author": "Elliot", "rating": 5}`),
	)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.FormatInt(req.ContentLength, 10))

	res, err := suite.app.Test(req, -1)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, res.StatusCode)

	req = httptest.NewRequest(
		"DELETE",
		"/api/v1/book/45678",
		nil,
	)
	res, err = suite.app.Test(req, -1)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, res.StatusCode)

	// var testbook book.Book
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
	// json.Unmarshal(body, &testbook)

	// assert.Equal(suite.T(), "Test Book", testbook.Title)
}

func TestBookTestSuite(t *testing.T) {
	suite.Run(t, new(BookTestSuite))
}
