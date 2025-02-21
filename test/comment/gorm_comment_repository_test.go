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

func TestCreateComment(t *testing.T) {
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
		comment := entities.Comment{
			ThreadID: 1,
			UserID:   2,
			Text:     "Hello, World!",
		}

		mock.ExpectBegin()
		mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "comments" ("created_at","updated_at","deleted_at","thread_id","user_id","text") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id"`)).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		mock.ExpectCommit()

		_, err := repo.CreateComment(comment)

		assert.NoError(t, err)
	})

	t.Run("Create Comment Failure", func(t *testing.T) {
		comment := entities.Comment{
			ThreadID: 1,
			UserID:   1,
			Text:     "Hello, World!",
		}

		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO \"comments\"").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), comment.ThreadID, comment.UserID, comment.Text).WillReturnError(errors.New("some error"))
		mock.ExpectRollback()

		_, err := repo.CreateComment(comment)
		assert.Error(t, err)
	})
}
