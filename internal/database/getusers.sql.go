// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: getusers.sql

package database

import (
	"context"
)

const getusers = `-- name: Getusers :many
SELECT id, created_at, updated_at, name, api_key FROM app_user
`

func (q *Queries) Getusers(ctx context.Context) ([]AppUser, error) {
	rows, err := q.db.QueryContext(ctx, getusers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []AppUser
	for rows.Next() {
		var i AppUser
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
			&i.ApiKey,
		); err != nil {
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
