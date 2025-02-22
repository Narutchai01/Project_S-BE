package adapter_test

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	adapters "github.com/Narutchai01/Project_S-BE/adapters/thread"
	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestGetThreadDetails(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	repo := adapters.NewGormThreadRepository(gormDB)
	t.Run("GetThreadDetails", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "thread_id", "skincare_id"}).
			AddRow(1, 1, 1)

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "thread_details" WHERE thread_id = $1 AND "thread_details"."deleted_at" IS NULL`)).
			WithArgs(1).
			WillReturnRows(rows)

		skincareRows := sqlmock.NewRows([]string{"id"}).
			AddRow(1)

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "skincares" WHERE "skincares"."id" = $1 AND "skincares"."deleted_at" IS NULL`)).
			WithArgs(1).
			WillReturnRows(skincareRows)

		threadDetails, err := repo.GetThreadDetails(1)

		assert.NoError(t, err)
		assert.Len(t, threadDetails, 1)
		assert.Equal(t, uint(1), threadDetails[0].ThreadID)
		assert.Equal(t, uint(1), threadDetails[0].SkincareID)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("GetThreadDetails Error", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "thread_details" WHERE thread_id = $1 AND "thread_details"."deleted_at" IS NULL`)).
			WithArgs(1).
			WillReturnError(gorm.ErrRecordNotFound)

		threadDetails, err := repo.GetThreadDetails(1)

		assert.Error(t, err)
		assert.Len(t, threadDetails, 0)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

}

func TestCreateThread(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	repo := adapters.NewGormThreadRepository(gormDB)
	t.Run("CreateThread", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "threads"`)).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		mock.ExpectCommit()

		thread, err := repo.CreateThread(1)

		assert.NoError(t, err)
		assert.Equal(t, uint(1), thread.UserID)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("CreateThread Error", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "threads"`)).WillReturnError(gorm.ErrInvalidData)
		mock.ExpectRollback()

		thread, err := repo.CreateThread(1)

		assert.Error(t, err)
		assert.Equal(t, uint(0), thread.UserID)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

}

func TestCreateThreadDetails(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	repo := adapters.NewGormThreadRepository(gormDB)
	t.Run("CreateThreadDetails", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "thread_details" ("created_at","updated_at","deleted_at","thread_id","skincare_id","caption") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id"`)).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), 1, 1, "Test").
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		mock.ExpectCommit()

		threadDetail, err := repo.CreateThreadDetail(entities.ThreadDetail{ThreadID: 1, SkincareID: 1, Caption: "Test"})

		assert.NoError(t, err)
		assert.Equal(t, uint(1), threadDetail.ThreadID)
		assert.Equal(t, uint(1), threadDetail.SkincareID)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("CreateThreadDetails Error", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "thread_details" ("created_at","updated_at","deleted_at","thread_id","skincare_id","caption") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id"`)).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), 1, 1, "Test").
			WillReturnError(gorm.ErrInvalidData)
		mock.ExpectRollback()

		threadDetail, err := repo.CreateThreadDetail(entities.ThreadDetail{ThreadID: 1, SkincareID: 1, Caption: "Test"})

		assert.Error(t, err)
		assert.Equal(t, uint(0), threadDetail.ThreadID)
		assert.Equal(t, uint(0), threadDetail.SkincareID)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

}
func TestGetThread(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	repo := adapters.NewGormThreadRepository(gormDB)
	t.Run("GetThread", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "user_id"}).
			AddRow(1, 1)

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "threads" WHERE "threads"."id" = $1 AND "threads"."deleted_at" IS NULL ORDER BY "threads"."id" LIMIT $2`)).
			WithArgs(1, 1).
			WillReturnRows(rows)

		threadDetailRows := sqlmock.NewRows([]string{"id", "thread_id", "skincare_id"}).
			AddRow(1, 1, 1)

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "thread_details" WHERE "thread_details"."thread_id" = $1 AND "thread_details"."deleted_at" IS NULL`)).
			WithArgs(1).
			WillReturnRows(threadDetailRows)

		userRows := sqlmock.NewRows([]string{"id"}).
			AddRow(1)

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."id" = $1`)).
			WithArgs(1).
			WillReturnRows(userRows)

		_, err := repo.GetThread(1)

		assert.NoError(t, err)
	})

	t.Run("GetThread Error", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "threads" WHERE "threads"."id" = $1 AND "threads"."deleted_at" IS NULL ORDER BY "threads"."id" LIMIT $2`)).
			WithArgs(1, 1).
			WillReturnError(gorm.ErrRecordNotFound)

		_, err := repo.GetThread(1)

		assert.Error(t, err)
	})

}
func TestGetThreadss(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	repo := adapters.NewGormThreadRepository(gormDB)
	t.Run("GetThreads", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "user_id"}).
			AddRow(1, 1).
			AddRow(2, 2)

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "threads" WHERE "threads"."deleted_at" IS NULL`)).
			WillReturnRows(rows)

		userRows := sqlmock.NewRows([]string{"id"}).
			AddRow(1).
			AddRow(2)

		threadDetailsRows := sqlmock.NewRows([]string{"id", "thread_id", "skincare_id"}).
			AddRow(1, 1, 1).
			AddRow(2, 2, 2)

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "thread_details" WHERE "thread_details"."thread_id" IN ($1,$2)`)).
			WithArgs(1, 2).
			WillReturnRows(threadDetailsRows)

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."id" IN ($1,$2)`)).
			WithArgs(1, 2).
			WillReturnRows(userRows)

		threads, err := repo.GetThreads()

		assert.NoError(t, err)
		assert.Len(t, threads, 2)
		assert.Equal(t, uint(1), threads[0].UserID)
		assert.Equal(t, uint(2), threads[1].UserID)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("GetThreads Error", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "threads" WHERE "threads"."deleted_at" IS NULL`)).
			WillReturnError(gorm.ErrRecordNotFound)

		threads, err := repo.GetThreads()

		assert.Error(t, err)
		assert.Len(t, threads, 0)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

}
func TestDeleteThreadGorm(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	repo := adapters.NewGormThreadRepository(gormDB)
	t.Run("DeleteThread", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta(`UPDATE "threads" SET "deleted_at"=$1 WHERE id = $2 AND "threads"."deleted_at" IS NULL`)).
			WithArgs(sqlmock.AnyArg(), 1).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		err := repo.DeleteThread(1)

		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("DeleteThread Error", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta(`UPDATE "threads" SET "deleted_at"=$1 WHERE id = $2 AND "threads"."deleted_at" IS NULL`)).
			WithArgs(sqlmock.AnyArg(), 1).
			WillReturnError(gorm.ErrInvalidData)
		mock.ExpectRollback()

		err := repo.DeleteThread(1)

		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	// create test case data not found
	t.Run("DeleteThread not found", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta(`UPDATE "threads" SET "deleted_at"=$1 WHERE id = $2 AND "threads"."deleted_at" IS NULL`)).
			WithArgs(sqlmock.AnyArg(), 1).
			WillReturnError(gorm.ErrRecordNotFound)
		mock.ExpectRollback()

		err := repo.DeleteThread(1)

		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

}
