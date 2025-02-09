package adapters

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	adapters "github.com/Narutchai01/Project_S-BE/adapters/skincare"
	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestGormCreateSkincare(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	repo := adapters.NewGormSkincareRepository(gormDB)

	expectData := entities.Skincare{
		Image:       "innisfree/image/path",
		Name:        "innisfree",
		Description: "green tea seed serum",
		CreateBY:    1,
	}

	t.Run("success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT INTO "skincares"`).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		mock.ExpectCommit()

		_, err := repo.CreateSkincare(expectData)

		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("failure", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT INTO "skincares"`).WillReturnError(errors.New("database error"))
		mock.ExpectRollback()

		_, err := repo.CreateSkincare(expectData)

		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

}

func TestGormGetSkincares(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	repo := adapters.NewGormSkincareRepository(gormDB)

	expectData := entities.Skincare{
		Model: gorm.Model{
			ID: 1,
		},
		Image:       "innisfree/image/path",
		Name:        "innisfree",
		Description: "green tea seed serum",
		CreateBY:    1,
	}

	columns := sqlmock.NewRows([]string{"id", "image", "name", "description", "create_by"}).
		AddRow(expectData.ID, expectData.Image, expectData.Name, expectData.Description, expectData.CreateBY)
	expectedSQL := `SELECT (.+) FROM "skincares"`

	t.Run("success", func(t *testing.T) {
		mock.ExpectQuery(expectedSQL).WillReturnRows(columns)

		_, err := repo.GetSkincares()

		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("failure", func(t *testing.T) {
		mock.ExpectQuery(expectedSQL).WillReturnError(errors.New("database error"))

		_, err := repo.GetSkincares()

		assert.Error(t, err)
	})

}

func TestGormGetskin(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	repo := adapters.NewGormSkincareRepository(gormDB)

	expectData := entities.Skincare{
		Model: gorm.Model{
			ID: 1,
		},
		Image:       "innisfree/image/path",
		Name:        "innisfree",
		Description: "green tea seed serum",
		CreateBY:    1,
	}

	expectedSQL := `SELECT \* FROM "skincares" WHERE "skincares"\."id" = \$1 AND "skincares"\."deleted_at" IS NULL ORDER BY "skincares"\."id" LIMIT \$2`
	rows := sqlmock.NewRows([]string{"id", "image", "name", "description", "create_by"}).
		AddRow(expectData.ID, expectData.Image, expectData.Name, expectData.Description, expectData.CreateBY)
	t.Run("success", func(t *testing.T) {
		mock.ExpectQuery(expectedSQL).
			WithArgs(1, 1).
			WillReturnRows(rows)

		result, err := repo.GetSkincareById(int(expectData.ID))

		assert.NoError(t, err)
		assert.Equal(t, expectData, result)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("failure", func(t *testing.T) {
		mock.ExpectQuery(expectedSQL).
			WithArgs(1, 1).
			WillReturnError(errors.New("database error"))

		_, err := repo.GetSkincareById(int(expectData.ID))

		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestGormUpdateUpdateskin(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	repo := adapters.NewGormSkincareRepository(gormDB)

	expectData := entities.Skincare{
		Model: gorm.Model{
			ID: 1,
		},
		Image:       "innisfree/image/path",
		Name:        "innisfree",
		Description: "green tea seed serum",
		CreateBY:    1,
	}

	expectedSQL := `UPDATE "skincares" SET "id"=\$1,"updated_at"=\$2,"image"=\$3,"name"=\$4,"description"=\$5,"create_by"=\$6 WHERE id = \$7 AND "skincares"."deleted_at" IS NULL`

	t.Run("success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(expectedSQL).
			WithArgs(
				expectData.ID,
				sqlmock.AnyArg(),
				expectData.Image,
				expectData.Name,
				expectData.Description,
				expectData.CreateBY,
				expectData.ID,
			).WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectCommit()

		result, err := repo.UpdateSkincareById(int(expectData.ID), expectData)

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
				expectData.Name,
				expectData.Description,
				expectData.CreateBY,
				expectData.ID,
			).WillReturnError(errors.New("database error"))
		mock.ExpectRollback()

		_, err := repo.UpdateSkincareById(int(expectData.ID), expectData)

		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

}

func TestGormDeleteskin(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	repo := adapters.NewGormSkincareRepository(gormDB)

	expectData := entities.Skincare{
		Model: gorm.Model{
			ID: 1,
		},
		Image:       "innisfree/image/path",
		Name:        "innisfree",
		Description: "green tea seed serum",
		CreateBY:    1,
	}

	expectedSQL := `UPDATE "skincares" SET "deleted_at"=\$1 WHERE id = \$2 AND "skincares"."deleted_at" IS NULL`

	t.Run("success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(expectedSQL).
			WithArgs(sqlmock.AnyArg(), expectData.ID).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectCommit()

		_, err := repo.DeleteSkincareById(int(expectData.ID))

		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("failure", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(expectedSQL).
			WithArgs(sqlmock.AnyArg(), expectData.ID).
			WillReturnError(errors.New("database error"))
		mock.ExpectRollback()

		_, err := repo.DeleteSkincareById(int(expectData.ID))

		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
