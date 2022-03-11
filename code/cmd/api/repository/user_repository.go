package repository

import (
	"context"
)

func GetMemberById(id string) (model.user, error) {
	db := db.DB()
	defer db.Close()

	var u model.user
	var strQuery = "SELECT email FROM open_market.public.users WHERE id=$1"

	db.QueryRow(
		context.Background(),
		strQuery,
		id,
	).Scan(
		&u.Id,
		&u.Email,
		&u.Name,
	)

	return u, nil
}
