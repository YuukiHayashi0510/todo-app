-- name: CreateStaff :one
-- スタッフの新規作成
INSERT INTO staffs (
    organization_id,
    email,
    staff_name
) VALUES (
    @organization_id,
    @email,
    @staff_name
)
RETURNING *;

-- name: SearchStaffs :many
-- スタッフの検索クエリ
SELECT 
   s.*,
   sqlc.embed(organizations)
FROM staffs s
LEFT JOIN organizations ON s.organization_id = organizations.organization_id
WHERE
   CASE 
       WHEN @search_status::text = 'all' THEN true
       WHEN @search_status::text = 'active' THEN s.deleted_at IS NULL
       WHEN @search_status::text = 'in_active' THEN s.deleted_at IS NOT NULL
   END
   AND (@organization_id::bigint = 0 OR s.organization_id = @organization_id::bigint)
   AND (@email::text = '' OR s.email LIKE '%' || @email::text || '%')
   AND (@staff_name::text = '' OR s.staff_name LIKE '%' || @staff_name::text || '%')
   AND (@created_at_start::timestamp IS NULL OR s.created_at >= @created_at_start::timestamp)
   AND (@created_at_end::timestamp IS NULL OR s.created_at <= @created_at_end::timestamp)
ORDER BY s.staff_id DESC
LIMIT CAST(sqlc.arg('limit') AS INTEGER)
OFFSET CAST(sqlc.arg('offset') AS INTEGER);

-- name: CountSearchStaffs :one
-- 検索結果の総件数を取得
SELECT 
    COUNT(*) as total_count
FROM staffs
WHERE
    CASE 
        WHEN @search_status::text = 'all' THEN true
        WHEN @search_status::text = 'active' THEN deleted_at IS NULL
        WHEN @search_status::text = 'in_active' THEN deleted_at IS NOT NULL
    END
    AND (@organization_id::bigint = 0 OR organization_id = @organization_id::bigint)
    AND (@email::text = '' OR email LIKE '%' || @email::text || '%')
    AND (@staff_name::text = '' OR staff_name LIKE '%' || @staff_name::text || '%')
    AND (@created_at_start::timestamp IS NULL OR created_at >= @created_at_start::timestamp)
    AND (@created_at_end::timestamp IS NULL OR created_at <= @created_at_end::timestamp);

-- name: GetStaffByID :one
-- IDでスタッフを取得する
SELECT 
    s.*,
    sqlc.embed(o) as "organization"
FROM staffs s
JOIN organizations o ON s.organization_id = o.organization_id
WHERE s.staff_id = $1 AND s.deleted_at IS NULL
LIMIT 1;

-- name: UpdateStaff :exec
UPDATE staffs
SET 
    email = @email,
    staff_name = @staff_name,
    organization_id = @organization_id,
    updated_at = CURRENT_TIMESTAMP
WHERE 
    staff_id = @staff_id
    AND deleted_at IS NULL;

-- name: SoftDeleteStaff :exec
UPDATE staffs
SET 
    deleted_at = CURRENT_TIMESTAMP,
    updated_at = CURRENT_TIMESTAMP
WHERE 
    staff_id = @staff_id
    AND deleted_at IS NULL;


-- name: RestoreStaff :exec
UPDATE staffs
SET 
    deleted_at = NULL,
    updated_at = CURRENT_TIMESTAMP
WHERE 
    staff_id = @staff_id
    AND deleted_at IS NOT NULL;
