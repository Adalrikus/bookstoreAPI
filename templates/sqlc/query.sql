-- name: GetBook :one
SELECT * FROM books
WHERE id = ? LIMIT 1;

-- name: GetBooks :many
SELECT * FROM books
ORDER BY Title;

-- name: CreateBook :one
INSERT INTO books (
  Title, Author, Publisher
) VALUES (
  ?, ?, ?
)
RETURNING *;

-- name: UpdateBook :exec
UPDATE books
set Title = ?,
    Author = ?,
    Publisher = ?
WHERE id = ?;

-- name: DeleteBook :exec
DELETE FROM books
WHERE id = ?;
