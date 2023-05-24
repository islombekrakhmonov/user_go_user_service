package postgres

import (
	"context"
	"user_service/genproto/user_service"
	"github.com/jackc/pgx/v4/pgxpool"
)


type userRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *userRepo {
	return &userRepo{
		db: db,
	}
}

func (u *userRepo) Create(ctx context.Context, req *user_service.CreateUserRequest) (pKey *user_service.UserPKey, err error) {

	query := `
	INSERT INTO users (
		first_name,
		last_name,
		phone_number
	) values (
		$1,
		$2,
		$3
	)`

	_, err = u.db.Exec(ctx, query,
		req.FirstName,
		req.LastName,
		req.PhoneNumber,
	)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (u *userRepo) Get(ctx context.Context, pKey *user_service.UserPKey) (resp *user_service.User, err error) {
	var (
		query  string
		user user_service.User
	)

	query = `
	SELECT
		id,
		first_name,
		last_name,
		phone_number
	FROM users
	WHERE id = $1`

	err = u.db.QueryRow(ctx, query, pKey.Id).Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.PhoneNumber,
	)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (u *userRepo) GetAll(ctx context.Context, req *user_service.GetAllUsersRequest) (resp *user_service.GetAllUsersResponse, err error) {
	resp = &user_service.GetAllUsersResponse{}
	var (
		query  string
		filter = " WHERE TRUE "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
	SELECT
		id,
		first_name,
		last_name,
		phone_number
	FROM users `

	query += filter + offset + limit
	
	rows, err := u.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user user_service.User
		err := rows.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.PhoneNumber,
		)
		if err != nil {
			return nil, err
		}
		resp.Users = append(resp.Users, &user)
	}

	return nil, nil
}

func (u *userRepo) Delete(ctx context.Context, pKey *user_service.UserPKey) (err error) {
	query := `
	DELETE FROM users
	WHERE id = $1`

	_, err = u.db.Exec(ctx, query, pKey.Id)
	if err != nil {
		return err
	}

	return nil
}	

func (u *userRepo) Update(ctx context.Context, pKey *user_service.UserPKey) (err error) {
	var (
		query  string
		user user_service.User
	)

	query = `
	UPDATE users
	SET
		first_name = $1,
		last_name = $2,
		phone_number = $3
	WHERE id = $4`

	_, err = u.db.Exec(ctx, query, user.FirstName, user.LastName, user.PhoneNumber, pKey.Id)
	if err != nil {
		return err
	}

	return nil
}