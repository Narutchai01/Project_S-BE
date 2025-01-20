package adapters

import (
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	adapters "github.com/Narutchai01/Project_S-BE/adapters/user"
	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestGormCreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	repo := adapters.NewGormUserRepository(gormDB)

	expectData := entities.User{
		FullName: "aut",
		Email:    "aut@gmail.com",
		Password: "aut1234hashed",
	}

	t.Run("success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT INTO "users"`).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		mock.ExpectCommit()

		_, err := repo.CreateUser(expectData)

		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("failure", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT INTO "users"`).WillReturnError(errors.New("database error"))
		mock.ExpectRollback()

		_, err := repo.CreateUser(expectData)

		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

}

func TestGormUpdateUserPasswordById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	repo := adapters.NewGormUserRepository(gormDB)

	expectData := entities.User{
		Model: gorm.Model{
			ID: 1,
		},
		Password: "aut1234hashed",
	}

	expectedSQL := `UPDATE "users" SET "password"=$1,"updated_at"=$2 WHERE id = $3 AND "users"."deleted_at" IS NULL RETURNING *`
	columns := sqlmock.NewRows([]string{"id", "password"}).AddRow(expectData.ID, expectData.Password)

	t.Run("success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(regexp.QuoteMeta(expectedSQL)).
			WithArgs(expectData.Password, sqlmock.AnyArg(), int(expectData.ID)).
			WillReturnRows(columns,)
		mock.ExpectCommit()

		result, err := repo.UpdateUserPasswordById(int(expectData.ID), expectData.Password)

		assert.NoError(t, err)
		assert.Equal(t, expectData.Password, result.Password)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("failure", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(regexp.QuoteMeta(expectedSQL)).
			WithArgs(expectData.Password, sqlmock.AnyArg(), int(expectData.ID)).
			WillReturnError(errors.New("database error"))
		mock.ExpectRollback()

		_, err := repo.UpdateUserPasswordById(int(expectData.ID), expectData.Password)

		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

}

func TestGormGetAdminByEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	repo := adapters.NewGormUserRepository(gormDB)

	expectData := entities.User{
		Model: gorm.Model{
			ID: 1,
		},
		FullName: "aut",
		Email:    "aut@gmail.com",
		Password: "aut1234hashed",
		Image:    "autimage",
	}

	expectedSQL := `SELECT \* FROM "users" WHERE email = \$1 AND "users"\."deleted_at" IS NULL ORDER BY "users"\."id" LIMIT \$2`
	rows := sqlmock.NewRows([]string{"id", "full_name", "email", "password", "image"}).
		AddRow(expectData.ID, expectData.FullName, expectData.Email, expectData.Password, expectData.Image)

	t.Run("success", func(t *testing.T) {
		mock.ExpectQuery(expectedSQL).
			WithArgs(expectData.Email, 1).
			WillReturnRows(rows)

		result, err := repo.GetUserByEmail(expectData.Email)

		assert.NoError(t, err)
		assert.Equal(t, expectData, result)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("failure", func(t *testing.T) {
		mock.ExpectQuery(expectedSQL).
			WithArgs(expectData.Email, 1).
			WillReturnError(errors.New("database error"))

		result, err := repo.GetUserByEmail(expectData.Email)

		assert.Error(t, err)
		assert.Empty(t, result)
		assert.EqualError(t, err, "database error")
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
