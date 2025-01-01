-- name: CreateOrganization :one
-- 組織の新規作成
INSERT INTO organizations (
    organization_name
) VALUES (
    @organization_name::text
)
RETURNING *;

-- name: SearchOrganizations :many
-- 組織の検索クエリ
SELECT 
   organization_id,
   organization_name,
   created_at,
   updated_at,
   deleted_at
FROM organizations
WHERE
   CASE 
      WHEN @search_status::text = 'all' THEN true
      WHEN @search_status::text = 'active' THEN deleted_at IS NULL
      WHEN @search_status::text = 'in_active' THEN deleted_at IS NOT NULL
   END
   AND (@organization_id::bigint = 0 OR organization_id = @organization_id::bigint)
   AND (@organization_name::text = '' OR organization_name LIKE '%' || @organization_name::text || '%')
   AND (@created_at_start::timestamp IS NULL OR created_at >= @created_at_start::timestamp)
   AND (@created_at_end::timestamp IS NULL OR created_at <= @created_at_end::timestamp)
ORDER BY organization_id DESC
LIMIT CAST(sqlc.arg('limit') AS INTEGER)
OFFSET CAST(sqlc.arg('offset') AS INTEGER);

-- name: GetOrganizationByID :one
-- IDで組織を取得する
SELECT 
      organization_id,
      organization_name,
      created_at,
      updated_at,
      deleted_at
FROM 
      organizations
WHERE 
      organization_id = @organization_id::bigint 
      AND deleted_at IS NULL;

-- name: CountSearchOrganizations :one
-- 検索結果の総件数を取得
SELECT 
   COUNT(*) as total_count
FROM organizations
WHERE
   CASE 
      WHEN @search_status::text = 'all' THEN true
      WHEN @search_status::text = 'active' THEN deleted_at IS NULL
      WHEN @search_status::text = 'in_active' THEN deleted_at IS NOT NULL
   END
   AND (@organization_id::bigint = 0 OR organization_id = @organization_id::bigint)
   AND (@organization_name::text = '' OR organization_name LIKE '%' || @organization_name::text || '%')
   AND (@created_at_start::timestamp IS NULL OR created_at >= @created_at_start::timestamp)
   AND (@created_at_end::timestamp IS NULL OR created_at <= @created_at_end::timestamp);

-- name: UpdateOrganization :exec
-- 組織情報を更新する
UPDATE organizations
SET 
      organization_name = COALESCE(@organization_name::text, organization_name),
      updated_at = CURRENT_TIMESTAMP
WHERE 
      organization_id = @organization_id::bigint 
      AND deleted_at IS NULL;

-- name: SoftDeleteOrganization :exec
-- 組織を論理削除する
UPDATE organizations
SET 
      deleted_at = CURRENT_TIMESTAMP
WHERE 
      organization_id = @organization_id::bigint 
      AND deleted_at IS NULL;

-- name: RestoreOrganization :exec
-- 論理削除された組織を復元する
UPDATE organizations
SET 
      deleted_at = NULL,
      updated_at = CURRENT_TIMESTAMP
WHERE 
      organization_id = @organization_id::bigint 
      AND deleted_at IS NOT NULL;
