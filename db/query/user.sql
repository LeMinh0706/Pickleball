-- name: CreateUser :one
INSERT INTO users(
    id,
    username,
    password,
    fullname,
    gender,
    avt,
    lat,
    lng
) VALUES(
    $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING fullname, gender, avt, lat, lng;

-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;
