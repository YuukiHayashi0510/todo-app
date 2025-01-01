-- 組織テーブル
CREATE TABLE organizations (
    organization_id BIGSERIAL PRIMARY KEY,
    organization_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL
);

-- スタッフテーブル
CREATE TABLE staff (
    staff_id BIGSERIAL PRIMARY KEY,
    organization_id BIGINT NOT NULL,
    email VARCHAR(255) NOT NULL,
    staff_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    CONSTRAINT idx_org_email UNIQUE (organization_id, email)
);

-- セッションテーブル（物理削除）
CREATE TABLE staff_sessions (
    session_id BIGSERIAL PRIMARY KEY,
    staff_id BIGINT NOT NULL,
    session_data TEXT,
    expires_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- プロジェクトテーブル
CREATE TABLE projects (
    project_id BIGSERIAL PRIMARY KEY,
    organization_id BIGINT NOT NULL,
    project_name VARCHAR(255) NOT NULL,
    description TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL
);

-- タスクテーブル
CREATE TABLE tasks (
    task_id BIGSERIAL PRIMARY KEY,
    project_id BIGINT NOT NULL,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    status VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL
);

-- ロールテーブル
CREATE TABLE roles (
    role_id BIGSERIAL PRIMARY KEY,
    role_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL
);

-- 機能テーブル
CREATE TABLE features (
    feature_id BIGSERIAL PRIMARY KEY,
    feature_name VARCHAR(255) NOT NULL,
    operation VARCHAR(255) NOT NULL,
    description TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL
);

-- ロールと機能の中間テーブル
CREATE TABLE role_features (
    role_id BIGINT NOT NULL,
    feature_id BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (role_id, feature_id)
);

-- プロジェクトメンバーテーブル
CREATE TABLE project_members (
    project_member_id BIGSERIAL PRIMARY KEY,
    project_id BIGINT NOT NULL,
    staff_id BIGINT NOT NULL,
    role_id BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    CONSTRAINT idx_project_staff UNIQUE (project_id, staff_id)
);
