package adapters_test

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

func TestGormCreateThread(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	repo := adapters.NewGormThreadRepository(gormDB)

	expectData := entities.Thread{
		Model:   gorm.Model{ID: 1},
		Title:   "title",
		Caption: "caption",
		UserID:  1,
	}

	t.Run("success", func(t *testing.T) {
		mock.ExpectQuery(`INSERT INTO "threads"`).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		mock.ExpectCommit()

		thread, err := repo.CreateThread(expectData)

		assert.NoError(t, err)
		assert.Equal(t, expectData.ID, thread.ID)
		assert.Equal(t, expectData.Title, thread.Title)
		assert.Equal(t, expectData.Caption, thread.Caption)
		assert.Equal(t, expectData.UserID, thread.UserID)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("error", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT INTO "threads"`).WillReturnError(gorm.ErrInvalidData)
		mock.ExpectRollback()

		_, err := repo.CreateThread(expectData)

		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

}

func TestGormGetThread(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	repo := adapters.NewGormThreadRepository(gormDB)

	expectData := entities.Thread{
		Model:   gorm.Model{ID: 1},
		Title:   "title",
		Caption: "caption",
		UserID:  1,
	}

	t.Run("success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id"}).AddRow(expectData.ID)
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "threads" WHERE id = $1 AND "threads"."deleted_at" IS NULL ORDER BY "threads"."id" LIMIT $2`)).
			WithArgs(expectData.ID, 1).
			WillReturnRows(rows)

		thread, err := repo.GetThread(expectData.ID)

		assert.NoError(t, err)
		assert.Equal(t, expectData.ID, thread.ID)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("thread not found", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "threads" WHERE id = $1 AND "threads"."deleted_at" IS NULL ORDER BY "threads"."id" LIMIT $2`)).
			WithArgs(expectData.ID, 1).
			WillReturnError(gorm.ErrRecordNotFound)

		_, err := repo.GetThread(expectData.ID)

		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestGormCreateThreadImage(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	repo := adapters.NewGormThreadRepository(gormDB)

	expectData := entities.ThreadImage{
		Model:    gorm.Model{ID: 1},
		ThreadID: 1,
		Image:    "image",
	}

	t.Run("success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "thread_images"`)).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		mock.ExpectCommit()

		threadImage, err := repo.CreateThreadImage(expectData)

		assert.NoError(t, err)
		assert.Equal(t, expectData.ID, threadImage.ID)
		assert.Equal(t, expectData.ThreadID, threadImage.ThreadID)
		assert.Equal(t, expectData.Image, threadImage.Image)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("error", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "thread_images"`)).WillReturnError(gorm.ErrInvalidData)
		mock.ExpectRollback()

		_, err := repo.CreateThreadImage(expectData)

		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestGormGetThreadImages(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	repo := adapters.NewGormThreadRepository(gormDB)

	expectData := []entities.ThreadImage{
		{Model: gorm.Model{ID: 1}, ThreadID: 1, Image: "image1"},
		{Model: gorm.Model{ID: 2}, ThreadID: 1, Image: "image2"},
	}

	t.Run("success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "thread_id", "image"}).
			AddRow(expectData[0].ID, expectData[0].ThreadID, expectData[0].Image).
			AddRow(expectData[1].ID, expectData[1].ThreadID, expectData[1].Image)
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "thread_images" WHERE thread_id = $1 AND "thread_images"."deleted_at" IS NULL`)).
			WithArgs(expectData[0].ThreadID).
			WillReturnRows(rows)

		threadImages, err := repo.GetThreadImages(expectData[0].ThreadID)

		assert.NoError(t, err)
		assert.Equal(t, expectData[0].ID, threadImages[0].ID)
		assert.Equal(t, expectData[0].ThreadID, threadImages[0].ThreadID)
		assert.Equal(t, expectData[0].Image, threadImages[0].Image)
		assert.Equal(t, expectData[1].ID, threadImages[1].ID)
		assert.Equal(t, expectData[1].ThreadID, threadImages[1].ThreadID)
		assert.Equal(t, expectData[1].Image, threadImages[1].Image)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("thread images not found", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "thread_images" WHERE thread_id = $1 AND "thread_images"."deleted_at" IS NULL`)).
			WithArgs(expectData[0].ThreadID).
			WillReturnError(gorm.ErrRecordNotFound)

		_, err := repo.GetThreadImages(expectData[0].ThreadID)

		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

}
