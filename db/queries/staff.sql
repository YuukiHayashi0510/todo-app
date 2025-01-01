-- name: CreateStaff :one
-- スタッフの新規作成
INSERT INTO staff (
    organization_id,
    email,
    staff_name
) VALUES (
    @organization_id::bigint,
    @email::text,
    @staff_name::text
)
RETURNING *;

-- name: SearchStaff :many
-- スタッフの検索クエリ
SELECT 
    staff_id,
    organization_id,
    email,
    staff_name,
    created_at,
    updated_at,
    deleted_at
FROM staff
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
    AND (@created_at_end::timestamp IS NULL OR created_at <= @created_at_end::timestamp)
ORDER BY staff_id DESC
LIMIT CAST(sqlc.arg('limit') AS INTEGER)
OFFSET CAST(sqlc.arg('offset') AS INTEGER);

-- name: CountSearchStaff :one
-- 検索結果の総件数を取得
SELECT 
    COUNT(*) as total_count
FROM staff
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
