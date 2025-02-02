// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: organization.sql

package rdb

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const countSearchOrganizations = `-- name: CountSearchOrganizations :one
SELECT 
   COUNT(*) as total_count
FROM organizations
WHERE
   CASE 
      WHEN $1::text = 'all' THEN true
      WHEN $1::text = 'active' THEN deleted_at IS NULL
      WHEN $1::text = 'in_active' THEN deleted_at IS NOT NULL
   END
   AND ($2::bigint = 0 OR organization_id = $2::bigint)
   AND ($3::text = '' OR organization_name LIKE '%' || $3::text || '%')
   AND ($4::timestamp IS NULL OR created_at >= $4::timestamp)
   AND ($5::timestamp IS NULL OR created_at <= $5::timestamp)
`

type CountSearchOrganizationsParams struct {
	SearchStatus     string
	OrganizationID   int64
	OrganizationName string
	CreatedAtStart   pgtype.Timestamp
	CreatedAtEnd     pgtype.Timestamp
}

// 検索結果の総件数を取得
func (q *Queries) CountSearchOrganizations(ctx context.Context, arg CountSearchOrganizationsParams) (int64, error) {
	row := q.db.QueryRow(ctx, countSearchOrganizations,
		arg.SearchStatus,
		arg.OrganizationID,
		arg.OrganizationName,
		arg.CreatedAtStart,
		arg.CreatedAtEnd,
	)
	var total_count int64
	err := row.Scan(&total_count)
	return total_count, err
}

const createOrganization = `-- name: CreateOrganization :one
INSERT INTO organizations (
    organization_name
) VALUES (
    $1::text
)
RETURNING organization_id, organization_name, created_at, updated_at, deleted_at
`

// 組織の新規作成
func (q *Queries) CreateOrganization(ctx context.Context, organizationName string) (Organization, error) {
	row := q.db.QueryRow(ctx, createOrganization, organizationName)
	var i Organization
	err := row.Scan(
		&i.OrganizationID,
		&i.OrganizationName,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getOrganizationByID = `-- name: GetOrganizationByID :one
SELECT 
      organization_id,
      organization_name,
      created_at,
      updated_at,
      deleted_at
FROM 
      organizations
WHERE 
      organization_id = $1::bigint
`

// IDで組織を取得する
func (q *Queries) GetOrganizationByID(ctx context.Context, organizationID int64) (Organization, error) {
	row := q.db.QueryRow(ctx, getOrganizationByID, organizationID)
	var i Organization
	err := row.Scan(
		&i.OrganizationID,
		&i.OrganizationName,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const restoreOrganization = `-- name: RestoreOrganization :exec
UPDATE organizations
SET 
      deleted_at = NULL,
      updated_at = CURRENT_TIMESTAMP
WHERE 
      organization_id = $1::bigint 
      AND deleted_at IS NOT NULL
`

// 論理削除された組織を復元する
func (q *Queries) RestoreOrganization(ctx context.Context, organizationID int64) error {
	_, err := q.db.Exec(ctx, restoreOrganization, organizationID)
	return err
}

const searchOrganizations = `-- name: SearchOrganizations :many
SELECT 
   organization_id,
   organization_name,
   created_at,
   updated_at,
   deleted_at
FROM organizations
WHERE
   CASE 
      WHEN $1::text = 'all' THEN true
      WHEN $1::text = 'active' THEN deleted_at IS NULL
      WHEN $1::text = 'in_active' THEN deleted_at IS NOT NULL
   END
   AND ($2::bigint = 0 OR organization_id = $2::bigint)
   AND ($3::text = '' OR organization_name LIKE '%' || $3::text || '%')
   AND ($4::timestamp IS NULL OR created_at >= $4::timestamp)
   AND ($5::timestamp IS NULL OR created_at <= $5::timestamp)
ORDER BY organization_id DESC
LIMIT CAST($7 AS INTEGER)
OFFSET CAST($6 AS INTEGER)
`

type SearchOrganizationsParams struct {
	SearchStatus     string
	OrganizationID   int64
	OrganizationName string
	CreatedAtStart   pgtype.Timestamp
	CreatedAtEnd     pgtype.Timestamp
	Offset           int32
	Limit            int32
}

// 組織の検索クエリ
func (q *Queries) SearchOrganizations(ctx context.Context, arg SearchOrganizationsParams) ([]Organization, error) {
	rows, err := q.db.Query(ctx, searchOrganizations,
		arg.SearchStatus,
		arg.OrganizationID,
		arg.OrganizationName,
		arg.CreatedAtStart,
		arg.CreatedAtEnd,
		arg.Offset,
		arg.Limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Organization{}
	for rows.Next() {
		var i Organization
		if err := rows.Scan(
			&i.OrganizationID,
			&i.OrganizationName,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const softDeleteOrganization = `-- name: SoftDeleteOrganization :exec
UPDATE organizations
SET 
      updated_at = CURRENT_TIMESTAMP,
      deleted_at = CURRENT_TIMESTAMP
WHERE 
      organization_id = $1::bigint 
      AND deleted_at IS NULL
`

// 組織を論理削除する
func (q *Queries) SoftDeleteOrganization(ctx context.Context, organizationID int64) error {
	_, err := q.db.Exec(ctx, softDeleteOrganization, organizationID)
	return err
}

const updateOrganization = `-- name: UpdateOrganization :exec
UPDATE organizations
SET 
      organization_name = COALESCE($1::text, organization_name),
      updated_at = CURRENT_TIMESTAMP
WHERE 
      organization_id = $2::bigint 
      AND deleted_at IS NULL
`

type UpdateOrganizationParams struct {
	OrganizationName string
	OrganizationID   int64
}

// 組織情報を更新する
func (q *Queries) UpdateOrganization(ctx context.Context, arg UpdateOrganizationParams) error {
	_, err := q.db.Exec(ctx, updateOrganization, arg.OrganizationName, arg.OrganizationID)
	return err
}
