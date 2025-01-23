package adapters

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	adapters "github.com/Narutchai01/Project_S-BE/adapters/acne"
	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestGormCreateAcne(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	repo := adapters.NewGormAcneRepository(gormDB)

	expectData := entities.Acne{
		Name: "innisfree",
		Image: "innisfree/image/path",
		CreateBY: 1,
	}

	t.Run("success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT INTO "acnes"`).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		mock.ExpectCommit()

		_, err := repo.CreateAcne(expectData)

		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("failure", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT INTO "acnes"`).WillReturnError(errors.New("database error"))
		mock.ExpectRollback()

		_, err := repo.CreateAcne(expectData)

		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

}

func TestGormGetAcnes(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	repo := adapters.NewGormAcneRepository(gormDB)

	expectData := entities.Acne{
		Model: gorm.Model{
			ID: 1,
		},
		Name: "innisfree",
		Image: "innisfree/image/path",
		CreateBY: 1,
	}

	columns := sqlmock.NewRows([]string{"id", "name", "image", "create_by"}).
		AddRow(expectData.ID, expectData.Name, expectData.Image, expectData.CreateBY)
	expectedSQL := `SELECT (.+) FROM "acnes"`

	t.Run("success", func(t *testing.T) {
		mock.ExpectQuery(expectedSQL).WillReturnRows(columns)

		_, err := repo.GetAcnes()

		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("failure", func(t *testing.T) {
		mock.ExpectQuery(expectedSQL).WillReturnError(errors.New("database error"))

		_, err := repo.GetAcnes()

		assert.Error(t, err)
	})

}

func TestGormGetAcne(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	repo := adapters.NewGormAcneRepository(gormDB)

	expectData := entities.Acne{
		Model: gorm.Model{
			ID: 1,
		},
		Name: "innisfree",
		Image: "innisfree/image/path",
		CreateBY: 1,
	}

	expectedSQL := `SELECT \* FROM "acnes" WHERE "acnes"\."id" = \$1 AND "acnes"\."deleted_at" IS NULL ORDER BY "acnes"\."id" LIMIT \$2`
	rows := sqlmock.NewRows([]string{"id", "name", "image", "create_by"}).
		AddRow(expectData.ID, expectData.Name, expectData.Image, expectData.CreateBY)
	t.Run("success", func(t *testing.T) {
		mock.ExpectQuery(expectedSQL).
			WithArgs(1, 1).
			WillReturnRows(rows)

		result, err := repo.GetAcne(int(expectData.ID))

		assert.NoError(t, err)
		assert.Equal(t, expectData, result)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("failure", func(t *testing.T) {
		mock.ExpectQuery(expectedSQL).
			WithArgs(1, 1).
			WillReturnError(errors.New("database error"))

		_, err := repo.GetAcne(int(expectData.ID))

		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestGormUpdateUpdateAcne(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	repo := adapters.NewGormAcneRepository(gormDB)

	expectData := entities.Acne{
		Model: gorm.Model{
			ID: 1,
		},
		Name: "innisfree",
		Image: "innisfree/image/path",
		CreateBY: 1,
	}

	expectedSQL := `UPDATE "acnes" SET "id"=\$1,"updated_at"=\$2,"name"=\$3,"image"=\$4,"create_by"=\$5 WHERE id = \$6 AND "acnes"."deleted_at" IS NULL`

	t.Run("success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(expectedSQL).
			WithArgs(expectData.ID, sqlmock.AnyArg(), expectData.Name, expectData.Image, expectData.CreateBY, expectData.ID).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectCommit()

		result, err := repo.UpdateAcne(int(expectData.ID), expectData)

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

		_, err := repo.UpdateAcne(int(expectData.ID), expectData)

		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

}

func TestGormDeleteAcne(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	repo := adapters.NewGormAcneRepository(gormDB)

	expectData := entities.Acne{
		Model: gorm.Model{
			ID: 1,
		},
		Name: "innisfree",
		Image: "innisfree/image/path",
		CreateBY: 1,
	}

	expectedSQL := `UPDATE "acnes" SET "deleted_at"=\$1 WHERE "acnes"."id" = \$2 AND "acnes"."deleted_at" IS NULL`

	t.Run("success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(expectedSQL).
			WithArgs(sqlmock.AnyArg(), expectData.ID).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectCommit()

		err := repo.DeleteAcne(int(expectData.ID))

		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("failure", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(expectedSQL).
			WithArgs(sqlmock.AnyArg(), expectData.ID).
			WillReturnError(errors.New("database error"))
		mock.ExpectRollback()

		err := repo.DeleteAcne(int(expectData.ID))

		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
