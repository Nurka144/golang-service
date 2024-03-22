package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/example/test/internal/models"
	"github.com/sirupsen/logrus"
)

type Test struct {
	db *sqlx.DB
}

func InitTestRepository(db *sqlx.DB) Test {
	return Test{
		db: db,
	}
}

func (r Test) Create(body models.TestCreateBody, AuthData models.AuthMiddlewareData) (int, error) {
	tx, err := r.db.Beginx()

	if err != nil {
		return 0, errors.New("Получена ошибка при открытие транзакции : " + err.Error())
	}

	defer func() {
		if err != nil {
			logrus.Info("[ROLLBACK] Причина = ", err.Error())
			tx.Rollback()
			return
		}

		err = tx.Commit()

		if err != nil {
			logrus.Info("Ошибка при [COMMIT] : ", err.Error())
		}
	}()

	return 1, nil
}

func (r Test) Find(id int) (*models.TestFind, error) {

	conn, err := r.db.Connx(context.Background())

	if err != nil {
		return nil, errors.Wrap(err, "Получена ошибка при получение пула из соединения")
	}

	defer func() {
		conn.Close()
	}()

	return &models.TestFind{}, nil
}
