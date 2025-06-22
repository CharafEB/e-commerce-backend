package model

import (
	"context"
)

func (artC *Prod) UpdateEvent(ctx context.Context, queryVal []string) error {
	if len(queryVal) == 0 {
		return nil
	}
	QueryVal := `UPDATE events SET is_published = false`
	_, err := artC.DB.ExecContext(ctx, QueryVal)
	if err != nil {
		return err
	}

	tx, err := artC.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.PrepareContext(ctx, `UPDATE events SET is_published = true WHERE event_id = $1`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, val := range queryVal {
		_, err := stmt.ExecContext(ctx, val)
		if err != nil {
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

// -->

func (artC *Prod) Modifierevent(ctx context.Context, queryVal []string) error {
	if len(queryVal) == 0 {
		return nil
	}
	QueryVal := `UPDATE events SET is_published = false`
	_, err := artC.DB.ExecContext(ctx, QueryVal)
	if err != nil {
		return err
	}

	tx, err := artC.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.PrepareContext(ctx, `UPDATE events SET is_published = true WHERE event_id = $1`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, val := range queryVal {
		_, err := stmt.ExecContext(ctx, val)
		if err != nil {
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

// -->
