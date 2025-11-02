-- name: GetUser :one
SELECT id, full_name, password_hash
  FROM the_user
  WHERE id = $1;
