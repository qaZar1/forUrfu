package database

import (
	"fmt"

	_ "github.com/Impisigmatus/service_core/postgres"
	"github.com/jmoiron/sqlx"
	"github.com/qaZar1/forUrfu/ratings/internal/models"
)

type Database struct {
	db *sqlx.DB
}

func NewDatabase(db *sqlx.DB) *Database {
	return &Database{
		db: db,
	}
}

func (pg *Database) AddRating(rating models.Rating) (bool, error) {
	const query = `
INSERT INTO main.ratings (
	from_user,
	to_user,
	rating
) VALUES (
	:from_user,
	:to_user,
	:rating
) ON CONFLICT (from_user, to_user) DO UPDATE 
SET rating = EXCLUDED.rating;`

	result, err := pg.db.NamedExec(query, rating)
	if err != nil {
		return false, fmt.Errorf("Invalid INSERT INTO main.ratings: %s", err)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("Invalid INSERT INTO main.ratings: %s", err)
	}

	return affected == 1, nil
}

func (pg *Database) CheckRatingByUsername(username string) (models.Rating, error) {
	const query = `
SELECT to_user,
	AVG(rating) AS rating
FROM main.ratings
WHERE to_user = $1
GROUP BY to_user;`

	var rating models.Rating
	if err := pg.db.Get(&rating, query, username); err != nil {
		return models.Rating{}, fmt.Errorf("User does not exist in main.ratings: %w", err)
	}

	return rating, nil
}

// func (pg *Database) GetAllVacancies() ([]models.Vacancy, error) {
// 	const query = `
// 	SELECT vacancy_id, company, title, description, employer_id, status
// 	FROM main.vacancies
// 	ORDER BY
//     CASE
//         WHEN status = 'Открыта' THEN 1
//         ELSE 2
//     END,
//     vacancy_id DESC;`

// 	var vacancies []models.Vacancy
// 	if err := pg.db.Select(&vacancies, query); err != nil {
// 		return nil, fmt.Errorf("Invalid SELECT main.vacancies: %s", err)
// 	}

// 	return vacancies, nil
// }

// func (pg *Database) AddVacancy(vacancy autogen.Vacancy) (int64, error) {
// 	const query = `
// INSERT INTO main.vacancies (
// 	company,
// 	title,
// 	description,
// 	chat_id_employer
// ) VALUES (
// 	$1,
// 	$2,
// 	$3,
// 	$4
// ) ON CONFLICT (vacancy_id) DO NOTHING
// RETURNING vacancy_id;`

// 	var vacancy_id int64
// 	err := pg.db.QueryRowx(query, vacancy.Company, vacancy.Title, vacancy.Description, vacancy.ChatIdEmployer).Scan(&vacancy_id)
// 	if err != nil {
// 		return 0, fmt.Errorf("Invalid INSERT INTO main.vacancies: %s", err)
// 	}

// 	return vacancy_id, nil
// }

// func (pg *Database) RemoveVacancy(vacancyId int64) (bool, error) {
// 	const query = "DELETE FROM main.vacancies WHERE vacancy_id = $1"

// 	exec, err := pg.db.Exec(query, vacancyId)
// 	if err != nil {
// 		return false, fmt.Errorf("Invalid DELETE main.vacancies: %s", err)
// 	}

// 	affected, err := exec.RowsAffected()
// 	if err != nil {
// 		return false, fmt.Errorf("Invalid affected vacancies: %s", err)
// 	}

// 	return affected == 0, nil
// }
