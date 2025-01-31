package result_test

import (
	"encoding/json"
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

func TestGormGetResults(t *testing.T) {
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
		Model: gorm.Model{ID: 1},
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

	// Convert slices to JSON
	acneTypeJSON, _ := json.Marshal(expectData.AcneType)
	facialTypeJSON, _ := json.Marshal(expectData.FacialType)
	skincareJSON, _ := json.Marshal(expectData.Skincare)

	columns := sqlmock.NewRows([]string{"id", "image", "user_id", "acne_type", "facial_type", "skin_type", "skincare"}).
		AddRow(expectData.ID, expectData.Image, expectData.UserId, string(acneTypeJSON), string(facialTypeJSON), expectData.SkinType, string(skincareJSON))

	expectedSQL := `SELECT (.+) FROM "results"`

	t.Run("success", func(t *testing.T) {
		mock.ExpectQuery(expectedSQL).WillReturnRows(columns)

		_, err := repo.GetResults()

		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("failure", func(t *testing.T) {
		mock.ExpectQuery(expectedSQL).WillReturnError(errors.New("database error"))

		_, err := repo.GetResults()

		assert.Error(t, err)
	})
}

func TestGormGetResult(t *testing.T) {
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
		Model: gorm.Model{ID: 1},
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

	// Convert slices to JSON
	acneTypeJSON, _ := json.Marshal(expectData.AcneType)
	facialTypeJSON, _ := json.Marshal(expectData.FacialType)
	skincareJSON, _ := json.Marshal(expectData.Skincare)

	expectedSQL := `SELECT \* FROM "results" WHERE "results"\."id" = \$1 AND "results"\."deleted_at" IS NULL ORDER BY "results"\."id" LIMIT \$2`
	rows := sqlmock.NewRows([]string{"id", "image", "user_id", "acne_type", "facial_type", "skin_type", "skincare"}).
	AddRow(expectData.ID, expectData.Image, expectData.UserId,  string(acneTypeJSON),  string(facialTypeJSON), expectData.SkinType,  string(skincareJSON))
	t.Run("success", func(t *testing.T) {
		mock.ExpectQuery(expectedSQL).
			WithArgs(1, 1).
			WillReturnRows(rows)

		result, err := repo.GetResultById(int(expectData.ID))

		assert.NoError(t, err)
		assert.Equal(t, expectData, result)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("failure", func(t *testing.T) {
		mock.ExpectQuery(expectedSQL).
			WithArgs(1, 1).
			WillReturnError(errors.New("database error"))

		_, err := repo.GetResultById(int(expectData.ID))

		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestGormUpdateResultById(t *testing.T) {
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
		Model: gorm.Model{ID: 1},
		Image: "image_url_test",
		UserId: 1,
		AcneType: []entities.Acne_Facial_Result{
			{ID: 1, Count: 10},
		},
		FacialType: []entities.Acne_Facial_Result{
			{ID: 1, Count: 10},
		},
		SkinType: 1,
		Skincare: []uint{1, 2, 3},
	}

	acneTypeJSON, _ := json.Marshal(expectData.AcneType)
	facialTypeJSON, _ := json.Marshal(expectData.FacialType)
	skincareJSON, _ := json.Marshal(expectData.Skincare)

	expectedSQL := `UPDATE "results" SET "id"=\$1,"updated_at"=\$2,"image"=\$3,"user_id"=\$4,"acne_type"=\$5,"facial_type"=\$6,"skin_type"=\$7,"skincare"=\$8 WHERE id = \$9 AND "results"."deleted_at" IS NULL`

	t.Run("success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(expectedSQL).
			WithArgs(
				expectData.ID, 
				sqlmock.AnyArg(), 
				expectData.Image, 
				expectData.UserId, 
				string(acneTypeJSON), 
				string(facialTypeJSON), 
				expectData.SkinType,
				string(skincareJSON),
				expectData.ID,
		).WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectCommit()

		result, err := repo.UpdateResultById(int(expectData.ID), expectData)

		assert.NoError(t, err)
		assert.Equal(t, expectData, result)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("failure", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(expectedSQL).
			WithArgs(
				expectData.ID, 
				sqlmock.AnyArg(), 
				expectData.Image, 
				expectData.UserId, 
				string(acneTypeJSON), 
				string(facialTypeJSON), 
				expectData.SkinType,
				string(skincareJSON),
				expectData.ID,
		).WillReturnError(errors.New("database error"))
		mock.ExpectRollback()

		_, err := repo.UpdateResultById(int(expectData.ID), expectData)

		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

}

func TestGormDeleteResultById(t *testing.T) {
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
	}

	expectedSQL := `UPDATE "results" SET "deleted_at"=\$1 WHERE "results"."id" = \$2 AND "results"."deleted_at" IS NULL`

	t.Run("success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(expectedSQL).
			WithArgs(sqlmock.AnyArg(), expectData.ID).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectCommit()

		err := repo.DeleteResultById(int(expectData.ID))

		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("failure", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(expectedSQL).
			WithArgs(sqlmock.AnyArg(), expectData.ID).
			WillReturnError(errors.New("database error"))
		mock.ExpectRollback()

		err := repo.DeleteResultById(int(expectData.ID))

		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}