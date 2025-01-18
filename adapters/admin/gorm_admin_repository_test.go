package adapters

import (
      "errors"
      "testing"

      "github.com/DATA-DOG/go-sqlmock"
      "github.com/Narutchai01/Project_S-BE/entities"
      "github.com/stretchr/testify/assert"
      "gorm.io/driver/postgres"
      "gorm.io/gorm"
)

func TestGormCreateAdmin(t *testing.T) {
      db, mock, err := sqlmock.New()
      if err != nil {
            t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
      }
      defer db.Close()

      gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
      if err != nil {
            panic("Failed to connect to database")
      }

      repo := NewGormAdminRepository(gormDB)

      expectData := entities.Admin{
            FullName: "aut",
            Email:    "aut@gmail.com",
            Password: "aut1234hashed",
            Image:    "autimage",
      }

      t.Run("success", func(t *testing.T) {
            mock.ExpectBegin()
            mock.ExpectQuery(`INSERT INTO "admins"`).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
            mock.ExpectCommit()

            _, err := repo.CreateAdmin(expectData)

            assert.NoError(t, err)
            assert.NoError(t, mock.ExpectationsWereMet())
      })

      t.Run("failure", func(t *testing.T) {
            mock.ExpectBegin()
            mock.ExpectQuery(`INSERT INTO "admins"`).WillReturnError(errors.New("database error"))
            mock.ExpectRollback()

            _, err := repo.CreateAdmin(expectData)

            assert.Error(t, err)
            assert.NoError(t, mock.ExpectationsWereMet())
      })

}

func TestGormGetAdmins(t *testing.T) {
      db, mock, err := sqlmock.New()
      if err != nil {
            t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
      }
      defer db.Close()

      gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
      if err != nil {
            panic("Failed to connect to database")
      }

      repo := NewGormAdminRepository(gormDB)

      expectData := entities.Admin{
            FullName: "aut",
            Email:    "aut@gmail.com",
            Password: "aut1234hashed",
            Image:    "autimage",
      }

      columns := sqlmock.NewRows([]string{"full_name", "email", "password", "image"}).
            AddRow(expectData.FullName, expectData.Email, expectData.Password, expectData.Image)
      expectedSQL := `SELECT (.+) FROM "admins"`

      t.Run("success", func(t *testing.T) {
            mock.ExpectQuery(expectedSQL).WillReturnRows(columns)

            _, err := repo.GetAdmins()

            assert.NoError(t, err)
            assert.NoError(t, mock.ExpectationsWereMet())
      })

      t.Run("failure", func(t *testing.T) {
            mock.ExpectQuery(expectedSQL).WillReturnError(errors.New("database error"))

            _, err := repo.GetAdmins()

            assert.Error(t, err)
      })

}

func TestGormGetAdmin(t *testing.T) {
      db, mock, err := sqlmock.New()
      if err != nil {
            t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
      }
      defer db.Close()

      gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
      if err != nil {
            panic("Failed to connect to database")
      }

      repo := NewGormAdminRepository(gormDB)

      expectData := entities.Admin{
            Model: gorm.Model{
                  ID: 1,
            },
            FullName: "aut",
            Email:    "aut@gmail.com",
            Password: "aut1234hashed",
            Image:    "autimage",
      }

      expectedSQL := `SELECT \* FROM "admins" WHERE "admins"\."id" = \$1 AND "admins"\."deleted_at" IS NULL ORDER BY "admins"\."id" LIMIT \$2`
      rows := sqlmock.NewRows([]string{"id", "full_name", "email", "password", "image"}).
            AddRow(1, expectData.FullName, expectData.Email, expectData.Password, expectData.Image)
      t.Run("success", func(t *testing.T) {
            mock.ExpectQuery(expectedSQL).
                  WithArgs(1, 1).
                  WillReturnRows(rows)

            result, err := repo.GetAdmin(1)

            assert.NoError(t, err)
            assert.Equal(t, expectData, result)
            assert.NoError(t, mock.ExpectationsWereMet())
      })

      t.Run("failure", func(t *testing.T) {
            mock.ExpectQuery(expectedSQL).
                  WithArgs(1, 1).
                  WillReturnError(errors.New("database error"))

            _, err := repo.GetAdmin(1)

            assert.Error(t, err)
            assert.NoError(t, mock.ExpectationsWereMet())
      })
}

func TestGormUpdateAdmin(t *testing.T) {
      db, mock, err := sqlmock.New()
      if err != nil {
            t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
      }
      defer db.Close()

      gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
      if err != nil {
            panic("Failed to connect to database")
      }

      repo := NewGormAdminRepository(gormDB)

      expectData := entities.Admin{
            Model: gorm.Model{
                  ID: 1,
            },
            FullName: "aut",
            Email:    "aut@gmail.com",
            Password: "aut1234hashed",
            Image:    "autimage",
      }

      expectedSQL := `UPDATE "admins" SET "id"=\$1,"updated_at"=\$2,"full_name"=\$3,"email"=\$4,"password"=\$5,"image"=\$6 WHERE id = \$7 AND "admins"."deleted_at" IS NULL`

      t.Run("success", func(t *testing.T) {
            mock.ExpectBegin()
            mock.ExpectExec(expectedSQL).
                  WithArgs(expectData.ID, sqlmock.AnyArg(), expectData.FullName, expectData.Email, expectData.Password, expectData.Image, expectData.ID).
                  WillReturnResult(sqlmock.NewResult(0, 1))
            mock.ExpectCommit()

            result, err := repo.UpdateAdmin(int(expectData.ID), expectData)

            assert.NoError(t, err)
            assert.Equal(t, expectData, result)
            assert.NoError(t, mock.ExpectationsWereMet())
      })

      t.Run("failure", func(t *testing.T) {
            mock.ExpectBegin()
            mock.ExpectExec(expectedSQL).
                  WithArgs(expectData.ID, sqlmock.AnyArg(), expectData.FullName, expectData.Email, expectData.Password, expectData.Image, expectData.ID).
                  WillReturnError(errors.New("database error"))
            mock.ExpectRollback()

            _, err := repo.UpdateAdmin(1, expectData)

            assert.Error(t, err)
            assert.NoError(t, mock.ExpectationsWereMet())
      })

}

func TestGormDeleteAdmin(t *testing.T) {
      db, mock, err := sqlmock.New()
      if err != nil {
            t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
      }
      defer db.Close()

      gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
      if err != nil {
            panic("Failed to connect to database")
      }

      repo := NewGormAdminRepository(gormDB)

      expectData := entities.Admin{
            Model: gorm.Model{
                  ID: 1,
            },
            FullName: "aut",
            Email:    "aut@gmail.com",
            Password: "aut1234hashed",
            Image:    "autimage",
      }

      expectedSQL := `UPDATE "admins" SET "deleted_at"=\$1 WHERE id = \$2 AND "admins"."deleted_at" IS NULL`

      t.Run("success", func(t *testing.T) {
            mock.ExpectBegin()
            mock.ExpectExec(expectedSQL).
                  WithArgs(sqlmock.AnyArg(), expectData.ID).
                  WillReturnResult(sqlmock.NewResult(0, 1))
            mock.ExpectCommit()

            result, err := repo.DeleteAdmin(int(expectData.ID))

            assert.NoError(t, err)
            assert.Equal(t, entities.Admin{}, result)
            assert.NoError(t, mock.ExpectationsWereMet())
      })

      t.Run("failure", func(t *testing.T) {
            mock.ExpectBegin()
            mock.ExpectExec(expectedSQL).
                  WithArgs(sqlmock.AnyArg(), expectData.ID).
                  WillReturnError(errors.New("database error"))
            mock.ExpectRollback()

            result, err := repo.DeleteAdmin(int(expectData.ID))

            assert.Error(t, err)
            assert.Equal(t, entities.Admin{}, result)
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

      repo := NewGormAdminRepository(gormDB)

      expectData := entities.Admin{
            Model: gorm.Model{
                  ID: 1,
            },
            FullName: "aut",
            Email:    "aut@gmail.com",
            Password: "aut1234hashed",
            Image:    "autimage",
      }

      // Adjust the expected SQL to use ? placeholders like GORM
      expectedSQL := `SELECT \* FROM "admins" WHERE email = \$1 AND "admins"\."deleted_at" IS NULL ORDER BY "admins"\."id" LIMIT \$2`
      rows := sqlmock.NewRows([]string{"id", "full_name", "email", "password", "image"}).
            AddRow(expectData.ID, expectData.FullName, expectData.Email, expectData.Password, expectData.Image)

      t.Run("success", func(t *testing.T) {
            mock.ExpectQuery(expectedSQL).
                  WithArgs(expectData.Email, 1).
                  WillReturnRows(rows)

            result, err := repo.GetAdminByEmail(expectData.Email)

            assert.NoError(t, err)
            assert.Equal(t, expectData, result)
            assert.NoError(t, mock.ExpectationsWereMet())
      })

      t.Run("failure", func(t *testing.T) {
            mock.ExpectQuery(expectedSQL).
                  WithArgs(expectData.Email, 1). // Adjust based on parameterized LIMIT
                  WillReturnError(errors.New("database error"))

            result, err := repo.GetAdminByEmail(expectData.Email)

            assert.Error(t, err)
            assert.Empty(t, result)
            assert.EqualError(t, err, "database error")
            assert.NoError(t, mock.ExpectationsWereMet())
      })
}