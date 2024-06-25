package cardRepo

import (
	models2 "bankApp1/internal/models"
	sqlQueries2 "bankApp1/pkg/sqlQueries"
	"bankApp1/txManager"
	"context"
	"database/sql"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"math/rand"
	"time"
)

type CardRepo struct {
	manager *txManager.TxManager
}

func NewCardRepo(manager *txManager.TxManager) *CardRepo {
	return &CardRepo{manager: manager}
}

func (r *CardRepo) Create(ctx context.Context, c models2.Card) (models2.CardID, error) {
	cNumber := rand.Intn(99999999-10000000+1) + 10000000
	query, args, err := sq.Insert(sqlQueries2.CardTable).
		Columns(sqlQueries2.InsertCardColumns...).
		Values(
			cNumber,
			c.UserID,
			c.Type,
			c.Pin,
			time.Now(),
		).
		Suffix(fmt.Sprintf("RETURNING %s", sqlQueries2.CardIDColumnName)).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return -1, err
	}

	t, err := r.manager.ExtractTXOrDB(ctx)
	if err != nil {
		return -1, err
	}

	var id models2.CardID
	if err := t.QueryRow(query, args...).Scan(&id); err != nil {
		return -1, err
	}
	return id, nil
}

func (r *CardRepo) Get(ctx context.Context, filter models2.CardFilter) (models2.Card, error) {
	cards, err := r.GetMany(ctx, filter)
	if err != nil {
		return models2.Card{}, err
	}

	if len(cards) == 0 {
		return models2.Card{}, sql.ErrNoRows
	}
	return cards[0], nil
}

func (r *CardRepo) GetMany(ctx context.Context, filter models2.CardFilter) (models2.ManyCards, error) {
	conds := r.getConds(filter)

	query, args, err := sq.Select(sqlQueries2.GetCardColumns...).
		From(sqlQueries2.CardTable).
		Where(conds).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}

	t, err := r.manager.ExtractTXOrDB(ctx)
	if err != nil {
		return nil, err
	}

	var manyCards models2.ManyCards
	if err := t.Select(&manyCards, query, args...); err != nil {
		return nil, err
	}
	return manyCards, nil
}

func (r *CardRepo) Delete(ctx context.Context, id models2.CardID) error {
	query, args, err := sq.Update(sqlQueries2.CardTable).
		Set(sqlQueries2.DeletedAtColumnName, time.Now()).
		Where(sq.Eq{
			sqlQueries2.CardIDColumnName: id,
		}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	t, err := r.manager.ExtractTXOrDB(ctx)
	if err != nil {
		return err
	}

	if _, err := t.Exec(query, args...); err != nil {
		return err
	}
	return nil
}

func (r *CardRepo) getConds(filter models2.CardFilter) sq.And {
	var sb sq.And

	if len(filter.IDs) != 0 {
		sb = append(sb, sq.Eq{
			sqlQueries2.CardIDColumnName: filter.IDs,
		})
	}
	if len(filter.Numbers) != 0 {
		sb = append(sb, sq.Eq{
			sqlQueries2.CardNumberColumnName: filter.Numbers,
		})
	}
	if len(filter.UserIDs) != 0 {
		sb = append(sb, sq.Eq{
			sqlQueries2.UserIDColumnName: filter.UserIDs,
		})
	}
	if len(filter.Types) != 0 {
		sb = append(sb, sq.Eq{
			sqlQueries2.CardTypeColumnName: filter.Types,
		})
	}
	sb = append(sb, sq.Eq{
		sqlQueries2.DeletedAtColumnName: nil,
	})
	return sb
}
