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

-- name: GetMyProfile :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;


-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: UpdateAvatar :one
UPDATE users SET 
avt = $2
WHERE id = $1
RETURNING id, fullname, gender, avt, lat, lng;

-- name: UpdatePosition :one
UPDATE users SET 
lat = $2, 
lng = $3
WHERE id = $1
RETURNING id, fullname, gender, avt, lat, lng;

-- name: SearchUser :many
SELECT id, fullname, gender, avt, lat, lng
FROM users
WHERE fullname ILIKE '%' || $1 || '%'
LIMIT $2
OFFSET $3;

-- name: CountUser :one
SELECT count(id)
FROM users
WHERE fullname ILIKE '%' || $1 || '%';

-- name: GetUsers :many
SELECT 
  id, 
  username, 
  fullname, 
  avt, 
  lat, 
  lng, 
  created_at,
  CAST(
    6371 * acos(
      cos(radians($1)) * cos(radians(lat)) * 
      cos(radians(lng) - radians($2)) + 
      sin(radians($1)) * sin(radians(lat))
    ) AS float8
  ) AS distance
FROM users
WHERE 
  (6371 * acos(
    cos(radians($1)) * cos(radians(lat)) * 
    cos(radians(lng) - radians($2)) + 
    sin(radians($1)) * sin(radians(lat))
  )) <= $3
ORDER BY distance;






