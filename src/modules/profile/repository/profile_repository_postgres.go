package repository

import (
	"database/sql"

	"github.com/CRUD-postgre/src/modules/profile/model"
)

type profileRepositoryPostgres struct {
	db *sql.DB
}

// can ngerti maksudna ieu, jeung naha kudu aya pointer trus
func NewProfileRepositoryPostgres(db *sql.DB) *profileRepositoryPostgres {
	return &profileRepositoryPostgres{db}
}

// Khusus funcation save
func (r *profileRepositoryPostgres) Save(profile *model.Profile) error {
	query := `INSERT INTO "profile"("id","first_name","last_name","email","password","created_at","updated_at")
		VALUES($1,$2,$3,$4,$5,$6,$7)`

	statment, err := r.db.Prepare(query)

	if err != nil {
		return err
	}
	// defer adalah ketika ada kyword difer paling terakhir di eksekusi.dan urutan nya dari bawah ketika eksekusinya
	defer statment.Close()

	_, err = statment.Exec(profile.ID, profile.FristName, profile.LastName, profile.Email, profile.Password, profile.CreatedAt, profile.UpdatedAt)

	if err != nil {
		return err
	}

	return nil
}

// Khusus funcation Update
func (r *profileRepositoryPostgres) Update(id string, profile *model.Profile) error {
	query := `UPDATE "profile" SET "firs_name"=$1, "last_name"=$2, "email"=$3, "password"=$4, "updated_at"=$5 WHERE "id"=$6`

	statment, err := r.db.Prepare(query)

	if err != nil {
		return err
	}

	defer statment.Close()

	_, err = statment.Exec(profile.FristName, profile.LastName, profile.Email, profile.Password, profile.UpdatedAt, id)

	if err != nil {
		return err
	}

	return nil

}

// Delete
func (r *profileRepositoryPostgres) Delete(id string) error {

	query := `DELETE FROM "profile" WHERE "id" =$1`

	statment, err := r.db.Prepare(query)

	if err != nil {
		return nil
	}

	defer statment.Close()

	_, err = statment.Exec(id)

	if err != nil {
		return err
	}

	return nil
}

// FindByID

func (r *profileRepositoryPostgres) FindByID(id string) (*model.Profile, error) {
	query := `SELECT * FROM "profile" WHERE id=$1`

	var profile model.Profile

	statment, err := r.db.Prepare(query)

	if err != nil {
		return nil, err
	}

	defer statment.Close()

	err = statment.QueryRow(id).Scan(&profile.ID, &profile.FristName, &profile.LastName, &profile.Email, &profile.Password, &profile.UpdatedAt, &profile.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &profile, nil

}

// FindAll

// func (r *profileRepositoryPostgres) FindAll() (*model.Profiles, error) {
// 	query := `SELECT * FROM "profile"`

// 	var profile model.Profiles

// 	rows, err := r.db.Query(query)

// 	if err != nil {
// 		return nil, err
// 	}

// 	defer rows.Close()

// 	for rows.Next() {
// 		var profile model.Profiles

// 		err = rows.Scan(&profile.ID, &profile.FristName, &profile.LastName, &profile.Email, &profile.Password, &profile.UpdatedAt, &profile.CreatedAt)

// 		if err != nil {
// 			return nil, err
// 		}

// 		profiles = append(profiles, profile)
// 	}

// 	return profiles, nil

// }
