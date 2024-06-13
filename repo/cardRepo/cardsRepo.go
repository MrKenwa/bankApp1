package cardRepo

import (
	"bankApp1/models"
	"bankApp1/sqlQueries"
	"bankApp1/txManager"
	"context"
	"database/sql"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"time"
)

type CardRepo struct {
	manager txManager.TxManager
}

func NewCardRepo(manager txManager.TxManager) *CardRepo {
	return &CardRepo{manager: manager}
}

func (r *CardRepo) Create(ctx context.Context, c models.Card) (models.CardID, error) {
	query, args, err := sq.Insert(sqlQueries.CardTable).
		Columns(sqlQueries.InsertCardColumns...).
		Values(
			c.CardNumber,
			c.UserID,
			c.Type,
			c.Pin,
			time.Now(),
		).
		Suffix(fmt.Sprintf("RETURNING %s", sqlQueries.CardIDColumnName)).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return -1, err
	}

	t, err := r.manager.ExtractTXOrDB(ctx)
	if err != nil {
		return -1, err
	}

	var id models.CardID
	if err := t.QueryRow(query, args...).Scan(&id); err != nil {
		return -1, err
	}
	return id, nil
}

func (r *CardRepo) Get(ctx context.Context, filter models.CardFilter) (models.Card, error) {
	cards, err := r.GetMany(ctx, filter)
	if err != nil {
		return models.Card{}, err
	}

	if len(cards) == 0 {
		return models.Card{}, sql.ErrNoRows
	}
	return cards[0], nil
}

func (r *CardRepo) GetMany(ctx context.Context, filter models.CardFilter) (models.ManyCards, error) {
	conds := r.getConds(filter)

	query, args, err := sq.Select(sqlQueries.GetCardColumns...).
		From(sqlQueries.CardTable).
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

	var manyCards models.ManyCards
	if err := t.Select(&manyCards, query, args...); err != nil {
		return nil, err
	}
	return manyCards, nil
}

func (r *CardRepo) Delete(ctx context.Context, id models.CardID) error {
	query, args, err := sq.Update(sqlQueries.CardTable).
		Set(sqlQueries.DeletedAtColumnName, time.Now()).
		Where(sq.Eq{
			sqlQueries.CardIDColumnName: id,
		}).
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

func (r *CardRepo) getConds(filter models.CardFilter) sq.And {
	var sb sq.And

	if len(filter.IDs) != 0 {
		sb = append(sb, sq.Eq{
			sqlQueries.CardIDColumnName: filter.IDs,
		})
	}
	if len(filter.Numbers) != 0 {
		sb = append(sb, sq.Eq{
			sqlQueries.CardNumberColumnName: filter.Numbers,
		})
	}
	if len(filter.UserIDs) != 0 {
		sb = append(sb, sq.Eq{
			sqlQueries.UserIDColumnName: filter.UserIDs,
		})
	}
	if len(filter.Types) != 0 {
		sb = append(sb, sq.Eq{
			sqlQueries.CardTypeColumnName: filter.Types,
		})
	}
	sb = append(sb, sq.Eq{
		sqlQueries.DeletedAtColumnName: nil,
	})
	return sb
}
