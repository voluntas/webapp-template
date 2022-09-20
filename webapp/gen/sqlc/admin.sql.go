// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: admin.sql

package sqlc

import (
	"context"
)

const listAccounts = `-- name: ListAccounts :many
SELECT pk, display_name, created_at FROM account
`

func (q *Queries) ListAccounts(ctx context.Context) ([]Account, error) {
	rows, err := q.db.QueryContext(ctx, listAccounts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Account{}
	for rows.Next() {
		var i Account
		if err := rows.Scan(&i.Pk, &i.DisplayName, &i.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}