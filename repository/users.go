package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/ASeegull/edriver-space/model"
	"github.com/lib/pq"
)

type UsersRepos struct {
	db *sql.DB
}

func NewUsersRepos(db *sql.DB) *UsersRepos {
	return &UsersRepos{
		db: db,
	}
}

func (r *UsersRepos) GetUserByCredentials(ctx context.Context, email, password string) (*model.User, error) {

	return r.getUser(ctx, "SELECT * FROM users WHERE email = $1 AND password = $2", email, password)
}

func (r *UsersRepos) GetUserById(ctx context.Context, userId string) (*model.User, error) {

	return r.getUser(ctx, "SELECT * FROM users WHERE id = $1", userId)
}

func (r *UsersRepos) CreateUser(ctx context.Context, newUser model.User) (string, error) {
	_, err := r.db.ExecContext(ctx, "INSERT INTO users(firstname, lastname, email, password) VALUES ($1, $2, $3, $4)",
		newUser.Firstname, newUser.Lastname, newUser.Email, newUser.Password)
	if err != nil {
		// convert to postgres error
		if err, ok := err.(*pq.Error); ok {
			// unique_violation
			if err.Code == "23505" {
				return "", model.ErrUserWithEmailExist
			}
		}
		return "", err
	}

	var userId string

	if err := r.db.QueryRowContext(ctx, "SELECT id FROM users WHERE email = $1", newUser.Email).Scan(&userId); err != nil {
		return "", err
	}
	return userId, nil
}

// CreateUserPolice adds new user with police role to the database
func (r *UsersRepos) CreateUserPolice(ctx context.Context, newUser model.User) (string, error) {
	_, err := r.db.ExecContext(ctx, "INSERT INTO users(firstname, lastname, email, password, role) VALUES ($1, $2, $3, $4, $5)",
		newUser.Firstname, newUser.Lastname, newUser.Email, newUser.Password, "police")
	if err != nil {
		// convert to postgres error
		if err, ok := err.(*pq.Error); ok {
			// unique_violation
			if err.Code == "23505" {
				return "", model.ErrUserWithEmailExist
			}
		}
		return "", err
	}

	var userId string

	if err = r.db.QueryRowContext(ctx, "SELECT id FROM users WHERE email = $1", newUser.Email).Scan(&userId); err != nil {
		return "", err
	}
	return userId, nil
}

func (r *UsersRepos) getUser(ctx context.Context, query string, args ...interface{}) (*model.User, error) {
	user := &model.User{}

	if err := r.db.QueryRowContext(ctx, query, args...).Scan(
		&user.Id, &user.Firstname, &user.Lastname, &user.Email, &user.Password, &user.Role, &user.DriverLicenseNumber,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, model.ErrUserNotFound
		}
		return user, err
	}
	return user, nil
}

func (r *UsersRepos) GetDriverLicence(ctx context.Context, itn string) (string, error) {
	var licenceNumber string
	if err := r.db.QueryRowContext(ctx, "SELECT licence_number FROM drivers WHERE individual_tax_number = $1", itn).Scan(&licenceNumber); err != nil {
		if err == sql.ErrNoRows {
			return "", model.ErrDriverLicenceNotFound
		}
		return "", err
	}
	return licenceNumber, nil
}

func (r UsersRepos) UpdateUserDriverLicence(ctx context.Context, userId, driverLicenceNumber string) error {
	result, err := r.db.ExecContext(ctx, "UPDATE users SET driver_licence_number = $1 WHERE id = $2", driverLicenceNumber, userId)
	if err != nil {
		return err
	}
	if num, _ := result.RowsAffected(); num != 1 {
		return errors.New("noting to update")
	}
	return nil
}

func (r *UsersRepos) GetCarsFines(ctx context.Context, userId string) ([]model.CarsFine, error) {
	var carsFines []model.CarsFine

	rows, err := r.db.QueryContext(ctx,
		"SELECT * FROM cars_fines WHERE id IN (SELECT car_id FROM cars_owners WHERE user_id = $1)", userId)
	if err != nil {
		return []model.CarsFine{}, err
	}

	for rows.Next() {
		var carsFine model.CarsFine

		if err := rows.Scan(
			&carsFine.Id,
			&carsFine.VehicleRegistrationNumber,
			&carsFine.FineNum,
			&carsFine.DataAndTime,
			&carsFine.Place,
			&carsFine.FileLawArticle,
			&carsFine.Price,
			&carsFine.Info,
			&carsFine.ImdUrl,
		); err != nil {
			return []model.CarsFine{}, err
		}

		carsFines = append(carsFines, carsFine)
	}

	if err := rows.Close(); err != nil {
		return []model.CarsFine{}, err
	}

	return carsFines, nil
}

func (r *UsersRepos) GetDriversFines(ctx context.Context, userId string) ([]model.DriversFine, error) {
	var driversFines []model.DriversFine

	rows, err := r.db.QueryContext(ctx, "SELECT * FROM drivers_fines WHERE licence_number IN (SELECT driver_licence_number FROM users WHERE id = $1)", userId)
	if err != nil {
		return []model.DriversFine{}, err
	}

	for rows.Next() {
		var driversFine model.DriversFine

		if err := rows.Scan(
			&driversFine.Id,
			&driversFine.LicenceNumber,
			&driversFine.FineNum,
			&driversFine.DataAndTime,
			&driversFine.Place,
			&driversFine.FileLawArticle,
			&driversFine.Price,
			&driversFine.VehicleRegistrationNumber,
		); err != nil {
			return []model.DriversFine{}, err
		}

		driversFines = append(driversFines, driversFine)
	}

	if err := rows.Close(); err != nil {
		return []model.DriversFine{}, err
	}

	return driversFines, nil
}

func (r *UsersRepos) ConnectCarAndUser(ctx context.Context, car model.Car, userId string) error {
	var checkId string

	row := r.db.QueryRowContext(
		ctx,
		"SELECT user_id FROM cars_owners WHERE car_id IN (SELECT id FROM cars WHERE vin_code = $1 AND registration_number = $2)",
		car.VIN, car.RegistrationNum,
	)

	if err := row.Scan(&checkId); err != nil {
		if err != sql.ErrNoRows {
			return err
		}
	}
	// check if the car is connected to the user
	if userId == checkId {
		return model.ErrCarIsAdded
	}

	result, err := r.db.ExecContext(ctx,
		"INSERT INTO cars_owners(user_id, car_id) (SELECT $1, id FROM cars WHERE vin_code = $2 AND registration_number = $3)",
		userId, car.VIN, car.RegistrationNum)
	if err != nil {
		return err
	}
	// check if the found car is added to the user
	num, _ := result.RowsAffected()
	if num != 1 {
		return model.ErrCarDoesNotExist
	}
	return nil
}
