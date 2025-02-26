package adapter_test

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	adapters "github.com/Narutchai01/Project_S-BE/adapters/favorite"
	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestFavoriteCommnet(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	repo := adapters.NewGormFavoriteRepository(gormDB)

	t.Run("success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT INTO "favorite_comments"`).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		mock.ExpectCommit()

		_, err := repo.FavoriteComment(1, 1)

		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("error", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT INTO "favorite_comments"`).WillReturnError(gorm.ErrInvalidData)
		mock.ExpectRollback()

		_, err := repo.FavoriteComment(1, 1)

		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

}

func TestFindFavoriteComment(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	repo := adapters.NewGormFavoriteRepository(gormDB)

	t.Run("success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "comment_id", "user_id"}).AddRow(1, 1, 1)
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "favorite_comments" WHERE (comment_id = $1 AND user_id = $2) AND "favorite_comments"."deleted_at" IS NULL ORDER BY "favorite_comments"."id" LIMIT $3`)).
			WithArgs(1, 1, 1).
			WillReturnRows(rows)

		favorite, err := repo.FindFavoriteComment(1, 1)

		assert.NoError(t, err)
		assert.Equal(t, uint(1), favorite.ID)
		assert.Equal(t, uint(1), favorite.CommentID)
		assert.Equal(t, uint(1), favorite.UserID)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("error", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "favorite_comments" WHERE (comment_id = $1 AND user_id = $2) AND "favorite_comments"."deleted_at" IS NULL ORDER BY "favorite_comments"."id" LIMIT $3`)).
			WithArgs(1, 1, 1).
			WillReturnError(gorm.ErrRecordNotFound)

		_, err := repo.FindFavoriteComment(1, 1)

		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

}
func TestUpdateFavoriteComments(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	repo := adapters.NewGormFavoriteRepository(gormDB)

	t.Run("success", func(t *testing.T) {
		favoriteComment := entities.FavoriteComment{
			Model:     gorm.Model{ID: 1},
			CommentID: 1,
			UserID:    1,
			Status:    true,
		}

		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta(`UPDATE "favorite_comments" SET "created_at"=$1,"updated_at"=$2,"deleted_at"=$3,"comment_id"=$4,"user_id"=$5,"status"=$6 WHERE "favorite_comments"."deleted_at" IS NULL AND "id" = $7`)).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), favoriteComment.CommentID, favoriteComment.UserID, favoriteComment.Status, favoriteComment.ID).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		_, err := repo.UpdateFavoriteComment(favoriteComment)

		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("error", func(t *testing.T) {
		favoriteComment := entities.FavoriteComment{
			Model:     gorm.Model{ID: 1},
			CommentID: 1,
			UserID:    1,
			Status:    false,
		}

		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta(`UPDATE "favorite_comments" SET "created_at"=$1,"updated_at"=$2,"deleted_at"=$3,"comment_id"=$4,"user_id"=$5,"status"=$6 WHERE "favorite_comments"."deleted_at" IS NULL AND "id" = $7`)).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), favoriteComment.CommentID, favoriteComment.UserID, favoriteComment.Status, favoriteComment.ID).
			WillReturnError(gorm.ErrInvalidData)
		mock.ExpectRollback()

		_, err := repo.UpdateFavoriteComment(favoriteComment)

		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestFavoriteThread(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	repo := adapters.NewGormFavoriteRepository(gormDB)

	t.Run("success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT INTO "favorite_threads"`).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		mock.ExpectCommit()

		_, err := repo.FavoriteThread(1, 1)

		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("error", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT INTO "favorite_threads"`).WillReturnError(gorm.ErrInvalidData)
		mock.ExpectRollback()

		_, err := repo.FavoriteThread(1, 1)

		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

}

func TestFindFavoriteThread(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	repo := adapters.NewGormFavoriteRepository(gormDB)

	t.Run("success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "thread_id", "user_id"}).AddRow(1, 1, 1)
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "favorite_threads" WHERE (thread_id = $1 AND user_id = $2) AND "favorite_threads"."deleted_at" IS NULL ORDER BY "favorite_threads"."id" LIMIT $3`)).
			WithArgs(1, 1, 1).
			WillReturnRows(rows)

		favorite, err := repo.FindFavoriteThread(1, 1)

		assert.NoError(t, err)
		assert.Equal(t, uint(1), favorite.ID)
		assert.Equal(t, uint(1), favorite.ThreadID)
		assert.Equal(t, uint(1), favorite.UserID)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("error", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "favorite_threads" WHERE (thread_id = $1 AND user_id = $2) AND "favorite_threads"."deleted_at" IS NULL ORDER BY "favorite_threads"."id" LIMIT $3`)).
			WithArgs(1, 1, 1).
			WillReturnError(gorm.ErrRecordNotFound)

		_, err := repo.FindFavoriteThread(1, 1)

		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestUpdateFavoriteThreads(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	repo := adapters.NewGormFavoriteRepository(gormDB)

	t.Run("success", func(t *testing.T) {
		favoriteThread := entities.FavoriteThread{
			Model:    gorm.Model{ID: 1},
			ThreadID: 1,
			UserID:   1,
			Status:   true,
		}

		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta(`UPDATE "favorite_threads" SET "created_at"=$1,"updated_at"=$2,"deleted_at"=$3,"thread_id"=$4,"user_id"=$5,"status"=$6 WHERE "favorite_threads"."deleted_at" IS NULL AND "id" = $7`)).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), favoriteThread.ThreadID, favoriteThread.UserID, favoriteThread.Status, favoriteThread.ID).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		_, err := repo.UpdateFavoriteThread(favoriteThread)

		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("error", func(t *testing.T) {
		favoriteThread := entities.FavoriteThread{
			Model:    gorm.Model{ID: 1},
			ThreadID: 1,
			UserID:   1,
			Status:   false,
		}

		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta(`UPDATE "favorite_threads" SET "created_at"=$1,"updated_at"=$2,"deleted_at"=$3,"thread_id"=$4,"user_id"=$5,"status"=$6 WHERE "favorite_threads"."deleted_at" IS NULL AND "id" = $7`)).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), favoriteThread.ThreadID, favoriteThread.UserID, favoriteThread.Status, favoriteThread.ID).
			WillReturnError(gorm.ErrInvalidData)
		mock.ExpectRollback()

		_, err := repo.UpdateFavoriteThread(favoriteThread)

		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
func TestCountFavoriteThread(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	repo := adapters.NewGormFavoriteRepository(gormDB)
	t.Run("success", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT count(*) FROM "favorite_threads" WHERE (thread_id = $1 AND status != false) AND "favorite_threads"."deleted_at" IS NULL`)).
			WithArgs(1).
			WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(5))

		count, err := repo.CountFavoriteThread(1)

		assert.NoError(t, err)
		assert.Equal(t, int64(5), count)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("error", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT count(*) FROM "favorite_threads" WHERE (thread_id = $1 AND status != false) AND "favorite_threads"."deleted_at" IS NULL`)).
			WithArgs(1).
			WillReturnError(gorm.ErrInvalidData)

		_, err := repo.CountFavoriteThread(1)

		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
