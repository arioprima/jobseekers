package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/arioprima/jobseeker/tree/main/backend/models"
)

type AdminRepository interface {
	Save(ctx context.Context, tx *sql.Tx, admin *models.AdminUser) (*models.AdminUser, error)
	Update(ctx context.Context, tx *sql.Tx, admin *models.AdminUser) (*models.AdminUser, error)
	Delete(ctx context.Context, tx *sql.Tx, userID string) error
	FindByID(ctx context.Context, tx *sql.Tx, userID string) (*models.AdminUser, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]*models.AdminUser, error)
}

type adminRepositoryImpl struct {
	DB *sql.DB
}

func NewAdminRepositoryImpl(db *sql.DB) AdminRepository {
	return &adminRepositoryImpl{DB: db}
}

func (adminUser *adminRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, admin *models.AdminUser) (*models.AdminUser, error) {
	if admin.UserID == "" {
		return nil, errors.New("user id is empty")
	}

	// Periksa apakah user_id sudah ada di tabel users
	var userExists bool
	SQL := "SELECT EXISTS (SELECT 1 FROM users WHERE id = ?)"
	err := tx.QueryRowContext(ctx, SQL, admin.UserID).Scan(&userExists)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if !userExists {
		tx.Rollback()
		return nil, errors.New("user not found")
	}

	// Sisipkan data ke admin_user hanya jika user_id tidak kosong
	SQL = "INSERT INTO admin (id, birth_place, date_of_birth, phone, address, user_id, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	_, err = tx.ExecContext(ctx, SQL, admin.AdminID, admin.BirthPlace, admin.BirthDate, admin.Phone, admin.Address, admin.UserID, admin.CreatedAt, admin.UpdatedAt)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// Perbarui status first_user di user menjadi false
	SQL = "UPDATE users SET first_user = false WHERE id = ? and first_user = true"
	_, err = tx.ExecContext(ctx, SQL, admin.UserID)
	if err != nil {
		// Rollback transaksi jika terjadi kesalahan
		tx.Rollback()
		return nil, err
	}

	// Commit transaksi jika semuanya berhasil
	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return admin, nil
}

func (adminUser *adminRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, admin *models.AdminUser) (*models.AdminUser, error) {
	//TODO implement me
	if admin.UserID == "" {
		return nil, errors.New("user id is empty")
	}

	var userExists bool
	SQL := "SELECT EXISTS (SELECT 1 FROM users WHERE id = ?)"
	err := tx.QueryRowContext(ctx, SQL, admin.UserID).Scan(&userExists)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if !userExists {
		tx.Rollback()
		return nil, errors.New("user not found")
	}

	SQL = "UPDATE admin SET birth_place = ?, date_of_birth = ?, phone = ?, address = ?, updated_at = ? WHERE user_id = ?"
	_, err = tx.ExecContext(ctx, SQL, admin.BirthPlace, admin.BirthDate, admin.Phone, admin.Address, admin.UpdatedAt, admin.UserID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	SQL = "UPDATE users SET first_name = ?, last_name = ?, updated_at = ? WHERE id = ?"
	_, err = tx.ExecContext(ctx, SQL, admin.FirstName, admin.LastName, admin.UpdatedAt, admin.UserID)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return admin, nil

}

func (adminUser *adminRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, userID string) error {
	//TODO implement me
	if userID == "" {
		return errors.New("user id is empty")
	}

	var userExists bool
	SQL := "SELECT EXISTS (SELECT 1 FROM users WHERE id = ?)"
	err := tx.QueryRowContext(ctx, SQL, userID).Scan(&userExists)
	if err != nil {
		tx.Rollback()
		return err
	}

	if !userExists {
		tx.Rollback()
		return errors.New("user not found")
	}

	SQL = "UPDATE users SET is_active = false WHERE id = ?"
	_, err = tx.ExecContext(ctx, SQL, userID)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (adminUser *adminRepositoryImpl) FindByID(ctx context.Context, tx *sql.Tx, userID string) (*models.AdminUser, error) {
	// TODO implement me
	if userID == "" {
		return nil, errors.New("user id is empty")
	}

	var userExists bool
	SQL := "SELECT EXISTS (SELECT 1 FROM users WHERE id = ?)"
	err := tx.QueryRowContext(ctx, SQL, userID).Scan(&userExists)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if !userExists {
		tx.Rollback()
		return nil, errors.New("user not found")
	}

	// Adjust the SQL query to select only the necessary columns
	SQL = "SELECT admin.*, users.first_name, users.last_name, users.email, users.role_id FROM admin INNER JOIN users ON admin.user_id = users.id WHERE admin.user_id = ?"
	row := tx.QueryRowContext(ctx, SQL, userID)

	var user models.AdminUser
	err = row.Scan(
		&user.AdminID,
		&user.BirthPlace,
		&user.BirthDate,
		&user.Phone,
		&user.Address,
		&user.UserID,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.RoleID,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (adminUser *adminRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]*models.AdminUser, error) {
	//TODO implement me
	panic("implement me")
}
