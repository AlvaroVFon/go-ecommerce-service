package roles

import (
	"context"
	"database/sql"
)

type RoleRepository struct {
	db *sql.DB
}

func NewRoleRepository(db *sql.DB) *RoleRepository {
	return &RoleRepository{db: db}
}

func (r *RoleRepository) FindByID(ctx context.Context, id int) (*Role, error) {
	query := "SELECT id, name FROM roles WHERE id = $1"
	row := r.db.QueryRowContext(ctx, query, id)
	var role Role

	if err := row.Scan(&role.ID, &role.Name); err != nil {
		return nil, err
	}

	return &role, nil
}

func (r *RoleRepository) FindByName(ctx context.Context, name string) (*Role, error) {
	query := "SELECT id, name FROM roles WHERE name = $1"
	row := r.db.QueryRowContext(ctx, query, name)
	var role Role
	if err := row.Scan(&role.ID, &role.Name); err != nil {
		return nil, err
	}

	return &role, nil
}

func (r *RoleRepository) FindAll(ctx context.Context) ([]Role, error) {
	query := "SELECT id, name FROM roles"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var roles []Role
	for rows.Next() {
		var role Role
		if err := rows.Scan(&role.ID, &role.Name); err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}

	return roles, nil
}
