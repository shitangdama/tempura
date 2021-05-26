package sql

var TabelSQL = `SELECT
c.oid :: int8 AS id,
nc.nspname AS schema,
c.relname AS name,
c.relrowsecurity AS rls_enabled,
c.relforcerowsecurity AS rls_forced,
pg_total_relation_size(format('%I.%I', nc.nspname, c.relname)) :: int8 AS bytes,
pg_size_pretty(
  pg_total_relation_size(format('%I.%I', nc.nspname, c.relname))
) AS size,
pg_stat_get_live_tuples(c.oid) AS live_rows_estimate,
pg_stat_get_dead_tuples(c.oid) AS dead_rows_estimate,
obj_description(c.oid) AS comment
FROM
pg_namespace nc
JOIN pg_class c ON nc.oid = c.relnamespace
WHERE
c.relkind IN ('r', 'p')
AND NOT pg_is_other_temp_schema(nc.oid)
AND (
  pg_has_role(c.relowner, 'USAGE')
  OR has_table_privilege(
	c.oid,
	'SELECT, INSERT, UPDATE, DELETE, TRUNCATE, REFERENCES, TRIGGER'
  )
  OR has_any_column_privilege(c.oid, 'SELECT, INSERT, UPDATE, REFERENCES')
)`

var SchemaSql = `
SELECT
  n.oid :: int8 AS id,
  n.nspname AS name,
  u.rolname AS owner
FROM
  pg_namespace n,
  pg_authid u
WHERE
  n.nspowner = u.oid
  AND (
    pg_has_role(n.nspowner, 'USAGE')
    OR has_schema_privilege(n.oid, 'CREATE, USAGE')
  )
`

var VersionSql = `
SELECT
  version(),
  current_setting('server_version_num') :: int8 AS version_number,
  (
    SELECT
      COUNT(*) AS active_connections
    FROM
      pg_stat_activity
  ) AS active_connections,
  current_setting('max_connections') :: int8 AS max_connections
`

var TypeSql = `
SELECT
  t.oid :: int8 AS id,
  t.typname AS name,
  n.nspname AS schema,
  format_type (t.oid, NULL) AS format,
  array_to_json(
    array(
      SELECT
        e.enumlabel
      FROM
        pg_enum e
      WHERE
        e.enumtypid = t.oid
      ORDER BY
        e.oid
    )
  ) AS enums,
  obj_description (t.oid, 'pg_type') AS comment
FROM
  pg_type t
  LEFT JOIN pg_namespace n ON n.oid = t.typnamespace
WHERE
  (
    t.typrelid = 0
    OR (
      SELECT
        c.relkind = 'c'
      FROM
        pg_class c
      WHERE
        c.oid = t.typrelid
    )
  )
  AND NOT EXISTS (
    SELECT
      1
    FROM
      pg_type el
    WHERE
      el.oid = t.typelem
      AND el.typarray = t.oid
  )
`

var GrantsSql = `
SELECT
  c.oid :: int8 AS table_id,
  u_grantor.rolname AS grantor,
  grantee.rolname AS grantee,
  nc.nspname AS schema,
  c.relname AS table_name,
  c.prtype AS privilege_type,
  CASE
    WHEN pg_has_role(grantee.oid, c.relowner, 'USAGE')
    OR c.grantable THEN TRUE
    ELSE FALSE
  END AS is_grantable,
  CASE
    WHEN c.prtype = 'SELECT' THEN TRUE
    ELSE FALSE
  END AS with_hierarchy
FROM
  (
    SELECT
      pg_class.oid,
      pg_class.relname,
      pg_class.relnamespace,
      pg_class.relkind,
      pg_class.relowner,
      (
        aclexplode(
          COALESCE(
            pg_class.relacl,
            acldefault('r', pg_class.relowner)
          )
        )
      ).grantor AS grantor,
      (
        aclexplode(
          COALESCE(
            pg_class.relacl,
            acldefault('r', pg_class.relowner)
          )
        )
      ).grantee AS grantee,
      (
        aclexplode(
          COALESCE(
            pg_class.relacl,
            acldefault('r', pg_class.relowner)
          )
        )
      ).privilege_type AS privilege_type,
      (
        aclexplode(
          COALESCE(
            pg_class.relacl,
            acldefault('r', pg_class.relowner)
          )
        )
      ).is_grantable AS is_grantable
    FROM
      pg_class
  ) c(
    oid,
    relname,
    relnamespace,
    relkind,
    relowner,
    grantor,
    grantee,
    prtype,
    grantable
  ),
  pg_namespace nc,
  pg_authid u_grantor,
  (
    SELECT
      pg_authid.oid,
      pg_authid.rolname
    FROM
      pg_authid
    UNION ALL
    SELECT
      0 :: oid AS oid,
      'PUBLIC'
  ) grantee(oid, rolname)
WHERE
  c.relnamespace = nc.oid
  AND (c.relkind IN ('r', 'v', 'f', 'p'))
  AND c.grantee = grantee.oid
  AND c.grantor = u_grantor.oid
  AND (
    c.prtype IN (
      'INSERT',
      'SELECT',
      'UPDATE',
      'DELETE',
      'TRUNCATE',
      'REFERENCES',
      'TRIGGER'
    )
  )
  AND (
    pg_has_role(u_grantor.oid, 'USAGE')
    OR pg_has_role(grantee.oid, 'USAGE')
  )
`
