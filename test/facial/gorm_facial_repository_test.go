package adapters

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	adapters "github.com/Narutchai01/Project_S-BE/adapters/facial"
	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestGormCreateFacial(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	repo := adapters.NewGormFacialRepository(gormDB)

	expectData := entities.Facial{
		Name: "facial_type1",
		Image: "facial/type1/path",
		CreateBY: 1,
	}

	t.Run("success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT INTO "facials"`).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		mock.ExpectCommit()

		_, err := repo.CreateFacial(expectData)

		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("failure", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT INTO "facials"`).WillReturnError(errors.New("database error"))
		mock.ExpectRollback()

		_, err := repo.CreateFacial(expectData)

		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

}

func TestGormGetfacials(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	repo := adapters.NewGormFacialRepository(gormDB)

	expectData := entities.Facial{
		Model: gorm.Model{
			ID: 1,
		},
		Name: "facial_type1",
		Image: "facial/type1/path",
		CreateBY: 1,
	}

	columns := sqlmock.NewRows([]string{"id", "name", "image", "create_by"}).
		AddRow(expectData.ID, expectData.Name, expectData.Image, expectData.CreateBY)
	expectedSQL := `SELECT (.+) FROM "facials"`

	t.Run("success", func(t *testing.T) {
		mock.ExpectQuery(expectedSQL).WillReturnRows(columns)

		_, err := repo.GetFacials()

		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("failure", func(t *testing.T) {
		mock.ExpectQuery(expectedSQL).WillReturnError(errors.New("database error"))

		_, err := repo.GetFacials()

		assert.Error(t, err)
	})

}

func TestGormGetfacial(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	repo := adapters.NewGormFacialRepository(gormDB)

	expectData := entities.Facial{
		Model: gorm.Model{
			ID: 1,
		},
		Name: "facial_type1",
		Image: "facial/type1/path",
		CreateBY: 1,
	}

	expectedSQL := `SELECT \* FROM "facials" WHERE "facials"\."id" = \$1 AND "facials"\."deleted_at" IS NULL ORDER BY "facials"\."id" LIMIT \$2`
	rows := sqlmock.NewRows([]string{"id", "name", "image", "create_by"}).
		AddRow(expectData.ID, expectData.Name, expectData.Image, expectData.CreateBY)
	t.Run("success", func(t *testing.T) {
		mock.ExpectQuery(expectedSQL).
			WithArgs(1, 1).
			WillReturnRows(rows)

		result, err := repo.GetFacial(int(expectData.ID))

		assert.NoError(t, err)
		assert.Equal(t, expectData, result)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("failure", func(t *testing.T) {
		mock.ExpectQuery(expectedSQL).
			WithArgs(1, 1).
			WillReturnError(errors.New("database error"))

		_, err := repo.GetFacial(int(expectData.ID))

		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestGormUpdateUpdatefacial(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	repo := adapters.NewGormFacialRepository(gormDB)

	expectData := entities.Facial{
		Model: gorm.Model{
			ID: 1,
		},
		Name: "facial_type1",
		Image: "facial/type1/path",
		CreateBY: 1,
	}

	expectedSQL := `UPDATE "facials" SET "id"=\$1,"updated_at"=\$2,"name"=\$3,"image"=\$4,"create_by"=\$5 WHERE id = \$6 AND "facials"."deleted_at" IS NULL`

	t.Run("success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(expectedSQL).
			WithArgs(expectData.ID, sqlmock.AnyArg(), expectData.Name, expectData.Image, expectData.CreateBY, expectData.ID).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectCommit()

		result, err := repo.UpdateFacial(int(expectData.ID), expectData)

		assert.NoError(t, err)
		assert.Equal(t, expectData, result)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("failure", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(expectedSQL).
		WithArgs(expectData.ID, sqlmock.AnyArg(), expectData.Name, expectData.Image, expectData.CreateBY, expectData.ID).
			WillReturnError(errors.New("database error"))
		mock.ExpectRollback()

		_, err := repo.UpdateFacial(int(expectData.ID), expectData)

		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

}

func TestGormDeletefacial(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	repo := adapters.NewGormFacialRepository(gormDB)

	expectData := entities.Facial{
		Model: gorm.Model{
			ID: 1,
		},
		Name: "facial_type1",
		Image: "facial/type1/path",
		CreateBY: 1,
	}

	expectedSQL := `UPDATE "facials" SET "deleted_at"=\$1 WHERE "facials"."id" = \$2 AND "facials"."deleted_at" IS NULL`

	t.Run("success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(expectedSQL).
			WithArgs(sqlmock.AnyArg(), expectData.ID).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectCommit()

		err := repo.DeleteFacial(int(expectData.ID))

		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("failure", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(expectedSQL).
			WithArgs(sqlmock.AnyArg(), expectData.ID).
			WillReturnError(errors.New("database error"))
		mock.ExpectRollback()

		err := repo.DeleteFacial(int(expectData.ID))

		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
