package adapters

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	adapters "github.com/Narutchai01/Project_S-BE/adapters/recovery"
	"github.com/Narutchai01/Project_S-BE/entities"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestGormCreateRecovery(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	repo := adapters.NewGormRecoveryRepository(gormDB)

	expectData := entities.Recovery{
		OTP:    "123456",
		UserId: 1,
	}

	t.Run("success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT INTO "recoveries"`).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		mock.ExpectCommit()

		_, err := repo.CreateRecovery(expectData)

		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("failure", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT INTO "recoveries"`).WillReturnError(errors.New("database error"))
		mock.ExpectRollback()

		_, err := repo.CreateRecovery(expectData)

		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

}

func TestGormDeleteRecoveryById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	repo := adapters.NewGormRecoveryRepository(gormDB)

	expectData := entities.Recovery{
		Model: gorm.Model{
			ID: 1,
		},
		OTP:    "123456",
		UserId: 1,
	}

	expectedSQL := `UPDATE "recoveries" SET "deleted_at"=\$1 WHERE id = \$2 AND "recoveries"."deleted_at" IS NULL`

	t.Run("success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(expectedSQL).
			WithArgs(sqlmock.AnyArg(), expectData.ID).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectCommit()

		result, err := repo.DeleteRecoveryById(int(expectData.ID))

		assert.NoError(t, err)
		assert.Equal(t, entities.Recovery{}, result)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("failure", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(expectedSQL).
			WithArgs(sqlmock.AnyArg(), expectData.ID).
			WillReturnError(errors.New("database error"))
		mock.ExpectRollback()

		result, err := repo.DeleteRecoveryById(int(expectData.ID))

		assert.Error(t, err)
		assert.Equal(t, entities.Recovery{}, result)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestGormGetRecoveries(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	repo := adapters.NewGormRecoveryRepository(gormDB)

	expectData := entities.Recovery{
		Model: gorm.Model{
			ID: 1,
		},
		OTP:    "123456",
		UserId: 1,
	}

	columns := sqlmock.NewRows([]string{"id", "otp", "user_id"}).
		AddRow(expectData.ID, expectData.OTP, expectData.UserId)
	expectedSQL := `SELECT (.+) FROM "recoveries"`

	t.Run("success", func(t *testing.T) {
		mock.ExpectQuery(expectedSQL).WillReturnRows(columns)

		_, err := repo.GetRecoveries()

		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("failure", func(t *testing.T) {
		mock.ExpectQuery(expectedSQL).WillReturnError(errors.New("database error"))

		_, err := repo.GetRecoveries()

		assert.Error(t, err)
	})

}

func TestGormGetRecoveryById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	repo := adapters.NewGormRecoveryRepository(gormDB)

	expectData := entities.Recovery{
		Model: gorm.Model{
			ID: 1,
		},
		OTP:    "123456",
		UserId: 1,
	}

	expectedSQL := `SELECT \* FROM "recoveries" WHERE "recoveries"\."id" = \$1 AND "recoveries"\."deleted_at" IS NULL ORDER BY "recoveries"\."id" LIMIT \$2`
	rows := sqlmock.NewRows([]string{"id", "otp", "user_id"}).
		AddRow(expectData.ID, expectData.OTP, expectData.UserId)
	t.Run("success", func(t *testing.T) {
		mock.ExpectQuery(expectedSQL).
			WithArgs(int(expectData.ID), 1).
			WillReturnRows(rows)

		result, err := repo.GetRecoveryById(int(expectData.ID))

		assert.NoError(t, err)
		assert.Equal(t, expectData, result)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("failure", func(t *testing.T) {
		mock.ExpectQuery(expectedSQL).
			WithArgs(int(expectData.ID), 1).
			WillReturnError(errors.New("database error"))

		_, err := repo.GetRecoveryById(int(expectData.ID))

		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestGormGetRecoveryByUserId(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	repo := adapters.NewGormRecoveryRepository(gormDB)

	expectData := entities.Recovery{
		Model: gorm.Model{
			ID: 1,
		},
		OTP:    "123456",
		UserId: 1,
	}

	expectedSQL := `SELECT \* FROM "recoveries" WHERE user_id = \$1 AND "recoveries"\."deleted_at" IS NULL ORDER BY "recoveries"\."id" LIMIT \$2`
	rows := sqlmock.NewRows([]string{"id", "otp", "user_id"}).
		AddRow(expectData.ID, expectData.OTP, expectData.UserId)
	t.Run("success", func(t *testing.T) {
		mock.ExpectQuery(expectedSQL).
			WithArgs(int(expectData.ID), 1).
			WillReturnRows(rows)

		result, err := repo.GetRecoveryByUserId(int(expectData.ID))

		assert.NoError(t, err)
		assert.Equal(t, expectData, result)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("failure", func(t *testing.T) {
		mock.ExpectQuery(expectedSQL).
			WithArgs(int(expectData.ID), 1).
			WillReturnError(errors.New("database error"))

		_, err := repo.GetRecoveryByUserId(int(expectData.ID))

		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestGormUpdateRecoveryOtpById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	repo := adapters.NewGormRecoveryRepository(gormDB)

	expectData := entities.Recovery{
		Model: gorm.Model{
			ID: 1,
		},
		OTP:    "123456",
		UserId: 1,
	}

	expectedSQL := `UPDATE "recoveries" SET "otp"=\$1,"updated_at"=\$2 WHERE id = \$3 AND "recoveries"."deleted_at" IS NULL RETURNING *`
	rows := sqlmock.NewRows([]string{"id", "otp", "user_id"}).
		AddRow(expectData.ID, expectData.OTP, expectData.UserId)

	t.Run("success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(expectedSQL).
			WithArgs(expectData.OTP, sqlmock.AnyArg(), expectData.ID).
			WillReturnRows(rows)
		mock.ExpectCommit()

		result, err := repo.UpdateRecoveryOtpById(int(expectData.ID), expectData.OTP)

		assert.NoError(t, err)
		assert.Equal(t, expectData.OTP, result.OTP)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("failure", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(expectedSQL).
			WithArgs(expectData.OTP, sqlmock.AnyArg(), expectData.ID).
			WillReturnError(errors.New("database error"))
		mock.ExpectRollback()

		_, err := repo.UpdateRecoveryOtpById(int(expectData.ID), expectData.OTP)

		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
