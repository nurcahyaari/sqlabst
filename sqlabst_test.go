package sqlabst_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/nurcahyaari/sqlabst"
	"github.com/stretchr/testify/assert"
)

func TestTxOrDbActived(t *testing.T) {
	t.Run("Test1 - begin not null", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		sql := sqlx.NewDb(db, "mysql")
		mock.ExpectBegin()
		sqlabst := sqlabst.NewSqlAbst(sql)
		err = sqlabst.Begin()

		assert.NoError(t, err)
		assert.NotNil(t, sqlabst.Tx)
	})

	t.Run("Test1 - begin null", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		sql := sqlx.NewDb(db, "mysql")
		mock.ExpectBegin()
		sqlabst := sqlabst.NewSqlAbst(sql)

		assert.NoError(t, err)
		assert.Nil(t, sqlabst.Tx)
	})

	t.Run("Test1 - begin close", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		sql := sqlx.NewDb(db, "mysql")
		mock.ExpectBegin()
		sqlabst := sqlabst.NewSqlAbst(sql)
		err = sqlabst.Begin()
		assert.NoError(t, err)
		assert.NotNil(t, sqlabst.Tx)

		mock.ExpectRollback()
		err = sqlabst.Rollback()
		assert.NoError(t, err)
		assert.Nil(t, sqlabst.Tx)
	})
}
