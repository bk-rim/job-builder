package repository

import (
	"database/sql"
	"fmt"

	"github.com/bk-rim/job-service/domain"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (r *PostgresRepository) Create(job *domain.Job) (domain.Job, error) {
	var id int64

	err := r.db.QueryRow("INSERT INTO jobs (name, type, frequency, created_on, status, executed_on, webhook_slack) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id",
		job.Name, job.Type, job.Frequency, job.CreatedOn, job.Status, job.ExecutedOn, job.WebhookSlack).Scan(&id)
	if err != nil {
		return domain.Job{}, fmt.Errorf("error creating job: %w", err)
	}
	job.ID = int(id)
	return *job, nil
}

func (r *PostgresRepository) GetByID(id int) (domain.Job, error) {
	var job domain.Job
	err := r.db.QueryRow("SELECT * FROM jobs WHERE id = $1", id).
		Scan(&job.ID, &job.Name, &job.Type, &job.Frequency, &job.CreatedOn, &job.Status, &job.ExecutedOn, &job.WebhookSlack)
	if err != nil {
		return domain.Job{}, fmt.Errorf("error getting job by ID: %w", err)
	}
	return job, nil
}

func (r *PostgresRepository) GetAll() ([]domain.Job, error) {
	rows, err := r.db.Query("SELECT * FROM jobs")
	if err != nil {
		return nil, fmt.Errorf("error getting all jobs: %w", err)
	}
	defer rows.Close()

	var jobs []domain.Job
	for rows.Next() {
		var job domain.Job
		err := rows.Scan(&job.ID, &job.Name, &job.Type, &job.Frequency, &job.CreatedOn, &job.Status, &job.ExecutedOn, &job.WebhookSlack)
		if err != nil {
			return nil, fmt.Errorf("error scanning job: %w", err)
		}
		jobs = append(jobs, job)
	}
	return jobs, nil
}

func (r *PostgresRepository) Update(id int, job domain.Job) (domain.Job, error) {
	_, err := r.db.Exec("UPDATE jobs SET name=$1, type=$2, frequency=$3, created_on=$4, status=$5, executed_on=$6, webhook_slack=$7 WHERE id=$8",
		job.Name, job.Type, job.Frequency, job.CreatedOn, job.Status, job.ExecutedOn, job.WebhookSlack, id)
	if err != nil {
		return domain.Job{}, fmt.Errorf("error updating job: %w", err)
	}
	return job, nil
}

func (r *PostgresRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM jobs WHERE id=$1", id)
	if err != nil {
		return fmt.Errorf("error deleting job: %w", err)
	}
	return nil
}
