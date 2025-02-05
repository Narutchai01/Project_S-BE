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
		Image:  "image_url_test",
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
		Skincare: []entities.Skincare{
			{Model: gorm.Model{ID: 1}},
			{Model: gorm.Model{ID: 2}},
		},
	}				

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
		Skincare: []entities.Skincare{
			{Model: gorm.Model{ID: 1},},
			{Model: gorm.Model{ID: 2},},
		},
	}

	expectedSQL := `SELECT \* FROM "results" WHERE "results"."deleted_at" IS NULL`
	rows := sqlmock.NewRows([]string{"id", "image", "user_id", "skin_type"}).
    		AddRow(int(expectData.ID), expectData.Image, int(expectData.UserId), int(expectData.SkinType))
	resultSkincareRows := sqlmock.NewRows([]string{"result_id", "skincare_id"}).
		AddRow(int(expectData.ID), int(expectData.Skincare[0].ID)).
		AddRow(int(expectData.ID), int(expectData.Skincare[1].ID))
	skincareRows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(int(expectData.Skincare[0].ID), "skincare1").
		AddRow(int(expectData.Skincare[1].ID), "skincare2")

	t.Run("success", func(t *testing.T) {
		mock.ExpectQuery(expectedSQL).
			WillReturnRows(rows)
		
		mock.ExpectQuery(`SELECT \* FROM "result_skincare" WHERE "result_skincare"."result_id" = \$1`).
			WillReturnRows(resultSkincareRows)
		
		// Fix for skincares query
		mock.ExpectQuery(`SELECT "skincares"\."id",.*FROM "skincares" WHERE "skincares"\."id" IN \(\$1,\$2\) AND "skincares"\."deleted_at" IS NULL`).
			WillReturnRows(skincareRows)
		
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

func TestGormGetResultById(t *testing.T) {
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
		Skincare: []entities.Skincare{
			{Model: gorm.Model{ID: 1},},
			{Model: gorm.Model{ID: 2},},
		},
	}

	expectedSQL := `SELECT \* FROM "results" WHERE "results"\."id" = \$1 AND "results"\."deleted_at" IS NULL ORDER BY "results"\."id" LIMIT \$2`
	rows := sqlmock.NewRows([]string{"id", "image", "user_id", "skin_type"}).
    		AddRow(1, "image_url_test", 1, 1)
	resultSkincareRows := sqlmock.NewRows([]string{"result_id", "skincare_id"}).
		AddRow(int(expectData.ID), int(expectData.Skincare[0].ID)).
		AddRow(int(expectData.ID), int(expectData.Skincare[1].ID))
	skincareRows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(int(expectData.Skincare[0].ID), "skincare1").
		AddRow(int(expectData.Skincare[1].ID), "skincare2")
	t.Run("success", func(t *testing.T) {
		mock.ExpectQuery(expectedSQL).
			WithArgs(int(expectData.ID), 1).
			WillReturnRows(rows)
	
		mock.ExpectQuery(`SELECT \* FROM "result_skincare" WHERE "result_skincare"\."result_id" = \$1`).
			WillReturnRows(resultSkincareRows)
	
		mock.ExpectQuery(`SELECT "skincares"\."id",.*FROM "skincares" WHERE "skincares"\."id" IN \(\$1,\$2\) AND "skincares"\."deleted_at" IS NULL`).
			WillReturnRows(skincareRows)
	
		result, err := repo.GetResultById(int(expectData.ID))
	
		assert.NoError(t, err)
		assert.Equal(t, expectData.ID, result.ID)
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
		Model:  gorm.Model{ID: 1},
		Image:  "image_url_test",
		UserId: 1,
		AcneType: []entities.Acne_Facial_Result{
			{ID: 1, Count: 10},
		},
		FacialType: []entities.Acne_Facial_Result{
			{ID: 1, Count: 10},
		},
		SkinType: 1,
		Skincare: []entities.Skincare{
			{Model: gorm.Model{ID: 1}},
			{Model: gorm.Model{ID: 2}},
		},
	}

	acneTypeJSON, _ := json.Marshal(expectData.AcneType)
	facialTypeJSON, _ := json.Marshal(expectData.FacialType)

	expectedSQL := `UPDATE "results" SET "id"=\$1,"updated_at"=\$2,"image"=\$3,"user_id"=\$4,"acne_type"=\$5,"facial_type"=\$6,"skin_type"=\$7 WHERE id = \$8 AND "results"."deleted_at" IS NULL`

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

func TestGormGetResultsByUserId(t *testing.T) {
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
		Skincare: []entities.Skincare{
			{Model: gorm.Model{ID: 1},},
			{Model: gorm.Model{ID: 2},},
		},
	}

	// expectedSQL := `SELECT \* FROM "results" WHERE "results"\."id" = \$1 AND "results"\."deleted_at" IS NULL ORDER BY "results"\."id" LIMIT \$2`
	rows := sqlmock.NewRows([]string{"id", "image", "user_id", "skin_type"}).
    		AddRow(1, "image_url_test", 1, 1)
	resultSkincareRows := sqlmock.NewRows([]string{"result_id", "skincare_id"}).
		AddRow(int(expectData.ID), int(expectData.Skincare[0].ID)).
		AddRow(int(expectData.ID), int(expectData.Skincare[1].ID))
	skincareRows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(int(expectData.Skincare[0].ID), "skincare1").
		AddRow(int(expectData.Skincare[1].ID), "skincare2")

	expectedSQL := `SELECT \* FROM "results" WHERE user_id = \$1 AND "results"\."deleted_at" IS NULL`

	t.Run("success", func(t *testing.T) {
		mock.ExpectQuery(expectedSQL).WithArgs(1).WillReturnRows(rows)

		mock.ExpectQuery(`SELECT \* FROM "result_skincare" WHERE "result_skincare"."result_id" = \$1`).
			WillReturnRows(resultSkincareRows)

		mock.ExpectQuery(`SELECT "skincares"\."id",.*FROM "skincares" WHERE "skincares"\."id" IN \(\$1,\$2\) AND "skincares"\."deleted_at" IS NULL`).
			WillReturnRows(skincareRows)

		_, err := repo.GetResultsByUserId(1)

		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("failure", func(t *testing.T) {
		mock.ExpectQuery(expectedSQL).WithArgs(1).WillReturnError(errors.New("database error"))

		_, err := repo.GetResultsByUserId(1)

		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestGormGetLatestResultsByUserIdFromToken(t *testing.T) {
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
		Skincare: []entities.Skincare{
			{Model: gorm.Model{ID: 1},},
			{Model: gorm.Model{ID: 2},},
		},
	}

	expectedSQL := `SELECT \* FROM "results" WHERE user_id = \$1 AND "results"\."deleted_at" IS NULL ORDER BY "results"\."id" DESC LIMIT \$2`
	rows := sqlmock.NewRows([]string{"id", "image", "user_id", "skin_type"}).
    		AddRow(1, "image_url_test", 1, 1)
	resultSkincareRows := sqlmock.NewRows([]string{"result_id", "skincare_id"}).
		AddRow(int(expectData.ID), int(expectData.Skincare[0].ID)).
		AddRow(int(expectData.ID), int(expectData.Skincare[1].ID))
	skincareRows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(int(expectData.Skincare[0].ID), "skincare1").
		AddRow(int(expectData.Skincare[1].ID), "skincare2")

	t.Run("success", func(t *testing.T) {
		mock.ExpectQuery(expectedSQL).WithArgs(int(expectData.UserId), 1).WillReturnRows(rows)
		mock.ExpectQuery(`SELECT \* FROM "result_skincare" WHERE "result_skincare"\."result_id" = \$1`).
			WillReturnRows(resultSkincareRows)
		mock.ExpectQuery(`SELECT "skincares"\."id",.*FROM "skincares" WHERE "skincares"\."id" IN \(\$1,\$2\) AND "skincares"\."deleted_at" IS NULL`).
			WillReturnRows(skincareRows)

		result, err := repo.GetLatestResultByUserIdFromToken(int(expectData.UserId))

		assert.NoError(t, err)
		assert.Equal(t, expectData.ID, result.ID)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("failure", func(t *testing.T) {
		mock.ExpectQuery(expectedSQL).WithArgs(1, 1).WillReturnError(errors.New("database error"))

		_, err := repo.GetLatestResultByUserIdFromToken(1)

		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
