package result_test

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	adapters "github.com/Narutchai01/Project_S-BE/adapters/result"
	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestGormCreateResult(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	repo := adapters.NewGormResultRepository(gormDB)

	expectData := entities.Result{
		Model: gorm.Model{
			ID: 1,
		},
		Image: "image_url_test",
		UserId: 1,
		AcneType: []entities.Acne_Facial_Result{
			{ID: 1, Count: 10},
			{ID: 2, Count: 5},
		},
		FacialType: []entities.Acne_Facial_Result{
			{ID: 1, Count: 10},
			{ID: 2, Count: 5},
		},
		SkinType: 1,
		Skincare: []uint{1, 2, 3},
	}

	t.Run("success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT INTO "results"`).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		mock.ExpectCommit()

		_, err := repo.CreateResult(expectData)

		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("failure", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT INTO "results"`).WillReturnError(errors.New("database error"))
		mock.ExpectRollback()

		_, err := repo.CreateResult(expectData)

		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

}