package postgres

// SQL queries for users table.
const (
	queryFindByEmail = `
		SELECT id, email, hashed_password, role 
		FROM users WHERE email = $1`

	queryCreate = `
		INSERT INTO users (id, email, hashed_password, role)
		VALUES ($1, $2, $3, $4)`

	queryExistsByEmail = `
		SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)`
)
