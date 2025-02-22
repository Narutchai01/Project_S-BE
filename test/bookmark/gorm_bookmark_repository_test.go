package adapters_test

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	adapters "github.com/Narutchai01/Project_S-BE/adapters/bookmark"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestCreateBookmark(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	repo := adapters.NewGormBookmarkRepository(gormDB)

	t.Run("Create Bookmark Success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "bookmarks" ("created_at","updated_at","deleted_at","thread_id","user_id","status") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id"`)).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), 1, 1, true).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		mock.ExpectCommit()
		result, err := repo.CreateBookmark(1, 1)
		assert.NoError(t, err)
		assert.Equal(t, uint(1), result.ID)
		assert.Equal(t, uint(1), result.UserID)
		assert.Equal(t, uint(1), result.ThreadID)
		assert.Equal(t, true, result.Status)
	})

	t.Run("Create Bookmark Failed", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "bookmarks" ("created_at","updated_at","deleted_at","thread_id","user_id","status") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id"`)).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), 1, 1, true).WillReturnError(assert.AnError)
		mock.ExpectRollback()
		_, err := repo.CreateBookmark(1, 1)
		assert.Error(t, err)
	})
}
