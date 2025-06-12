package repositories

import (
	"context"
	"fmt"

	"github.com/davigomesdev/reconfile/internal/domain/contracts"
	"github.com/davigomesdev/reconfile/internal/domain/entities"
	"github.com/davigomesdev/reconfile/internal/domain/errors"
	domain_repositories "github.com/davigomesdev/reconfile/internal/domain/repositories"
	"github.com/jackc/pgx/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lib/pq"
)

type userRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) domain_repositories.UserRepositoryInterface {
	return &userRepository{pool: pool}
}

func (r *userRepository) toEntity(scanner interface {
	Scan(dest ...interface{}) error
}) (*entities.UserEntity, error) {
	entity := &entities.UserEntity{}
	var deletedAt pgtype.Timestamptz

	if err := scanner.Scan(
		&entity.ID,
		&entity.Name,
		&entity.Email,
		&entity.Password,
		&entity.RefreshToken,
		&entity.CreatedAt,
		&entity.UpdatedAt,
		&deletedAt,
	); err != nil {
		return nil, err

	}

	if deletedAt.Status == pgtype.Present {
		entity.DeletedAt = &deletedAt.Time
	}

	return entity, nil
}

func (r *userRepository) EmailExists(ctx context.Context, email string) error {
	var exists bool
	err := r.pool.QueryRow(ctx, `
        SELECT EXISTS (
            SELECT 1 FROM users 
            WHERE email = $1 AND deleted_at IS NULL
		)
	`, email).Scan(&exists)
	if err != nil {
		return err
	}

	if exists {
		return errors.NewConflictError("E-mail já cadastrado.")
	}
	return nil
}

func (r *userRepository) Get(ctx context.Context, id string) (*entities.UserEntity, error) {
	row := r.pool.QueryRow(ctx, `SELECT * FROM users WHERE id = $1 AND deleted_at IS NULL`, id)
	entity, err := r.toEntity(row)
	if err != nil {
		return nil, errors.NewNotFoundError("Usuário não encontrado.")
	}

	return entity, nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*entities.UserEntity, error) {
	row := r.pool.QueryRow(ctx, `SELECT * FROM users WHERE email = $1 AND deleted_at IS NULL`, email)
	entity, err := r.toEntity(row)
	if err != nil {
		return nil, errors.NewNotFoundError("Usuário não encontrado.")
	}

	return entity, nil
}

func (r *userRepository) GetAll(ctx context.Context) ([]*entities.UserEntity, error) {
	rows, err := r.pool.Query(ctx, `SELECT * FROM users WHERE deleted_at IS NULL`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entities []*entities.UserEntity
	for rows.Next() {
		e, err := r.toEntity(rows)
		if err != nil {
			return nil, err
		}
		entities = append(entities, e)
	}
	return entities, nil
}

func (r *userRepository) Search(ctx context.Context, input *contracts.SearchInput) (*contracts.SearchResult[*entities.UserEntity], error) {
	page, perPage := 1, 15
	if input.Page != nil && *input.Page > 0 {
		page = *input.Page
	}
	if input.PerPage != nil && *input.PerPage > 0 {
		perPage = *input.PerPage
	}
	offset := (page - 1) * perPage

	baseQuery := "FROM users WHERE deleted_at IS NULL"
	whereClause := ""
	args := []interface{}{}
	argPos := 1

	if input.Filter != nil && *input.Filter != "" {
		whereClause = fmt.Sprintf(" AND (name ILIKE $%d OR email ILIKE $%d)", argPos, argPos+1)
		args = append(args, "%"+*input.Filter+"%", "%"+*input.Filter+"%")
		argPos += 2
	}

	query := "SELECT * " + baseQuery + whereClause

	if input.Sort != nil && *input.Sort != "" {
		dir := "ASC"
		if input.SortDir != nil && *input.SortDir == contracts.SortDesc {
			dir = "DESC"
		}
		query += fmt.Sprintf(" ORDER BY %s %s", pq.QuoteIdentifier(*input.Sort), dir)
	}

	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argPos, argPos+1)
	args = append(args, perPage, offset)

	rows, err := r.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("query error: %w", err)
	}
	defer rows.Close()

	var items []*entities.UserEntity
	for rows.Next() {
		e, err := r.toEntity(rows)
		if err != nil {
			return nil, fmt.Errorf("mapping error: %w", err)
		}
		items = append(items, e)
	}

	countQuery := "SELECT COUNT(*) " + baseQuery + whereClause
	var total int
	if err := r.pool.QueryRow(ctx, countQuery, args[:len(args)-2]...).Scan(&total); err != nil {
		return nil, fmt.Errorf("count error: %w", err)
	}

	return &contracts.SearchResult[*entities.UserEntity]{
		Items:       items,
		Total:       total,
		CurrentPage: page,
		PerPage:     perPage,
		Sort:        input.Sort,
		SortDir:     input.SortDir,
		Filter:      input.Filter,
	}, nil
}

func (r *userRepository) Create(ctx context.Context, e *entities.UserEntity) error {
	_, err := r.pool.Exec(ctx, `
		INSERT INTO users (id, name, email, password, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`, e.ID, e.Name, e.Email, e.Password, e.CreatedAt, e.UpdatedAt)
	return err
}

func (r *userRepository) Update(ctx context.Context, e *entities.UserEntity) error {
	_, err := r.pool.Exec(ctx, `
		UPDATE users
		SET    name = $2,
		       email = $3,
		       password = $4,
		       updated_at = $5
		WHERE  id = $1 AND deleted_at IS NULL
	`, e.ID, e.Name, e.Email, e.Password, e.UpdatedAt)
	return err
}

func (r *userRepository) Delete(ctx context.Context, id string) error {
	_, err := r.pool.Exec(ctx, `
		UPDATE users
		SET    deleted_at = NOW()
		WHERE  id = $1 AND deleted_at IS NULL
	`, id)
	return err
}
