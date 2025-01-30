package repository

import (
	"context"
	"github.com/bayuuat/tutuplapak/domain"
	"github.com/doug-martin/goqu/v9"
)

type UserRepository interface {
	Save(ctx context.Context, user *domain.User) error
	Update(ctx context.Context, user *domain.User) error
	FindById(ctx context.Context, id string) (domain.User, error)
	FindByColumns(ctx context.Context, ids []string, filter []interface{}) (user []domain.User, err error)
	FindByEmail(ctx context.Context, email string) (domain.User, error)
	FindByPhone(ctx context.Context, phone string) (domain.User, error)
}

type userRepository struct {
	db *goqu.Database
}

func NewUser(db *goqu.Database) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u userRepository) Save(ctx context.Context, user *domain.User) error {
	executor := u.db.Insert("users").Rows(user).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (u userRepository) Update(ctx context.Context, user *domain.User) error {
	executor := u.db.Update("users").Where(goqu.C("id").Eq(user.Id)).Set(user).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (u userRepository) FindById(ctx context.Context, id string) (user domain.User, err error) {
	dataset := u.db.From("users").Where(goqu.Ex{
		"id": id,
	})

	_, err = dataset.ScanStructContext(ctx, &user)
	return
}

func (u userRepository) FindByColumns(ctx context.Context, ids []string, filter []interface{}) (user []domain.User, err error) {
	dataset := u.db.From("users").Select(filter...).Where(goqu.C("id").In(ids))
	err = dataset.ScanStructsContext(ctx, &user)
	return
}

func (u userRepository) FindByEmail(ctx context.Context, email string) (user domain.User, err error) {
	dataset := u.db.From("users").Where(goqu.Ex{
		"email": email,
	})
	_, err = dataset.ScanStructContext(ctx, &user)
	return
}

func (u userRepository) FindByPhone(ctx context.Context, phone string) (user domain.User, err error) {
	dataset := u.db.From("users").Where(goqu.Ex{
		"phone": phone,
	})
	_, err = dataset.ScanStructContext(ctx, &user)
	return
}
