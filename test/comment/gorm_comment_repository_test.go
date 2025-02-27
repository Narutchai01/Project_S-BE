package adapters_test

import (
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	adapters "github.com/Narutchai01/Project_S-BE/adapters/comment"
	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestCreateCommentThreadGorm(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	repo := adapters.NewGormCommentRepository(gormDB)

	t.Run("Create Comment Success", func(t *testing.T) {
		comment := entities.CommentThread{
			ThreadID: 1,
			UserID:   2,
			Text:     "Hello, World!",
		}

		mock.ExpectBegin()
		mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "comment_threads" ("created_at","updated_at","deleted_at","thread_id","user_id","favorite","text") VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING "id"`)).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		mock.ExpectCommit()

		_, err := repo.CreateCommentThread(comment)

		assert.NoError(t, err)
	})

	t.Run("Create Comment Failure", func(t *testing.T) {
		comment := entities.CommentThread{
			ThreadID: 1,
			UserID:   1,
			Text:     "Hello, World!",
		}

		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO \"comment_threads\"").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), comment.ThreadID, comment.UserID, comment.Text).WillReturnError(errors.New("some error"))
		mock.ExpectRollback()

		_, err := repo.CreateCommentThread(comment)
		assert.Error(t, err)
	})
}
func TestCreateCommentReviewSkicnare(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	repo := adapters.NewGormCommentRepository(gormDB)

	expectData := entities.CommentReviewSkicare{
		Model:            gorm.Model{ID: 1},
		ReviewSkincareID: 1,
		UserID:           1,
		Favorite:         true,
		Content:          "Hello, World!",
	}

	t.Run("Create Comment Review Skicnare Success", func(t *testing.T) {

		mock.ExpectBegin()
		mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "comment_review_skicares"`)).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		mock.ExpectCommit()

		review, err := repo.CreateCommentReviewSkicnare(expectData)

		assert.NoError(t, err)
		assert.Equal(t, expectData.ID, review.ID)
		assert.Equal(t, expectData.ReviewSkincareID, review.ReviewSkincareID)
		assert.Equal(t, expectData.UserID, review.UserID)
		assert.Equal(t, expectData.Favorite, review.Favorite)
		assert.Equal(t, expectData.Content, review.Content)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Create Comment Review Skicnare Failure", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT INTO "comment_review_skicares"`).WillReturnError(gorm.ErrInvalidData)
		mock.ExpectRollback()

		_, err := repo.CreateCommentReviewSkicnare(expectData)

		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
func TestGetCommentsThread(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	repo := adapters.NewGormCommentRepository(gormDB)

	t.Run("Get Comments Thread Success", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "comment_threads" WHERE thread_id = $1 AND "comment_threads"."deleted_at" IS NULL`)).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id"}))

		_, err := repo.GetCommentsThread(1)

		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Get Comments Thread Failure", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "comment_threads" WHERE thread_id = $1 AND "comment_threads"."deleted_at" IS NULL`)).WithArgs(1).WillReturnError(errors.New("some error"))

		_, err := repo.GetCommentsThread(1)

		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
func TestGetCommentsReviewSkincare(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	repo := adapters.NewGormCommentRepository(gormDB)

	t.Run("Get Comments Review Skincare Success", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "comment_review_skicares" WHERE review_skincare_id = $1 AND "comment_review_skicares"."deleted_at" IS NULL`)).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id"}))

		_, err := repo.GetCommentsReviewSkincare(1)

		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Get Comments Review Skincare Failure", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "comment_review_skicares" WHERE review_skincare_id = $1 AND "comment_review_skicares"."deleted_at" IS NULL`)).WithArgs(1).WillReturnError(errors.New("some error"))

		_, err := repo.GetCommentsReviewSkincare(1)

		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
