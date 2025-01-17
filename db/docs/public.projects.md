# public.projects

## Description

## Columns

| Name            | Type                     | Default                                      | Nullable | Children | Parents | Comment |
| --------------- | ------------------------ | -------------------------------------------- | -------- | -------- | ------- | ------- |
| project_id      | bigint                   | nextval('projects_project_id_seq'::regclass) | false    |          |         |         |
| organization_id | bigint                   |                                              | false    |          |         |         |
| project_name    | varchar(255)             |                                              | false    |          |         |         |
| description     | text                     |                                              | true     |          |         |         |
| created_at      | timestamp with time zone | CURRENT_TIMESTAMP                            | false    |          |         |         |
| updated_at      | timestamp with time zone | CURRENT_TIMESTAMP                            | false    |          |         |         |
| deleted_at      | timestamp with time zone |                                              | true     |          |         |         |

## Constraints

| Name          | Type        | Definition               |
| ------------- | ----------- | ------------------------ |
| projects_pkey | PRIMARY KEY | PRIMARY KEY (project_id) |

## Indexes

| Name          | Definition                                                                    |
| ------------- | ----------------------------------------------------------------------------- |
| projects_pkey | CREATE UNIQUE INDEX projects_pkey ON public.projects USING btree (project_id) |

## Relations

![er](public.projects.svg)

---

> Generated by [tbls](https://github.com/k1LoW/tbls)
