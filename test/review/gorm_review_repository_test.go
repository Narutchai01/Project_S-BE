package adapters_test

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	adapters "github.com/Narutchai01/Project_S-BE/adapters/review"
	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestGormCreateReview(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	repo := adapters.NewGormReviewRepository(gormDB)

	expectData := entities.ReviewSkincare{
		Model:      gorm.Model{ID: 1},
		Title:      "title",
		Content:    "content",
		SkincareID: []int{1, 2},
		UserID:     1,
		Image:      "image",
	}

	t.Run("success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT INTO "review_skincares"`).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		mock.ExpectCommit()

		review, err := repo.CreateReviewSkincare(expectData)

		assert.NoError(t, err)
		assert.Equal(t, expectData.ID, review.ID)
		assert.Equal(t, expectData.Title, review.Title)
		assert.Equal(t, expectData.Content, review.Content)
		assert.Equal(t, expectData.SkincareID, review.SkincareID)
		assert.Equal(t, expectData.UserID, review.UserID)
		assert.Equal(t, expectData.Image, review.Image)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("failed", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT INTO "review_skincares"`).WillReturnError(gorm.ErrInvalidData)
		mock.ExpectRollback()

		_, err := repo.CreateReviewSkincare(expectData)

		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
func TestGormGetReview(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	repo := adapters.NewGormReviewRepository(gormDB)

	expectData := entities.ReviewSkincare{
		Model:      gorm.Model{ID: 1},
		Title:      "title",
		Content:    "content",
		SkincareID: []int{1, 2},
		UserID:     1,
		Image:      "image",
		User:       entities.User{Model: gorm.Model{ID: 1}, FullName: "John Doe"},
	}

	t.Run("success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id"}).AddRow(expectData.ID)
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "review_skincares" WHERE id = $1 AND "review_skincares"."deleted_at" IS NULL ORDER BY "review_skincares"."id" LIMIT $2`)).
			WithArgs(expectData.ID, 1).
			WillReturnRows(rows)

		thread, err := repo.GetReviewSkincare(expectData.ID)

		assert.NoError(t, err)
		assert.Equal(t, expectData.ID, thread.ID)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("failed", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "review_skincares" WHERE id = $1 AND "review_skincares"."deleted_at" IS NULL ORDER BY "review_skincares"."id" LIMIT $2`)).
			WithArgs(expectData.ID, 1).WillReturnError(gorm.ErrRecordNotFound)

		_, err := repo.GetReviewSkincare(expectData.ID)

		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
