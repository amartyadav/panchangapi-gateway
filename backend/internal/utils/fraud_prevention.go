package utils

import (
	"database/sql"
	"panchangapi-gateway/internal/database"
	"time"
)

const (
	maxAttempts   = 5
	blockDuration = 24 * time.Hour
)

func IsSignupAttemptBlocked(email string) (bool, error) {
	var blocked_until sql.NullTime

	err := database.DB.QueryRow("SELECT blocked_until FROM signup_attempts WHERE email =  $1", email).Scan(&blocked_until)
	if err == sql.ErrNoRows {
		return false, nil // no record found, hence not blocked
	}
	if err != nil {
		return false, err // database error
	}

	if blocked_until.Valid && blocked_until.Time.After(time.Now()) {
		return true, nil
	}

	return false, nil
}

func AddSignupAttempt(email string) error {
	tx, err := database.DB.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	var attemptCount int
	var lastAttempt time.Time
	err = tx.QueryRow("SELECT attempt_count, last_attempt_at FROM signup_attempts WHERE email = $1", email).Scan(&attemptCount, &lastAttempt)

	if err == sql.ErrNoRows {
		_, err = tx.Exec("INSERT INTO signup_attempts (email, attempt_count, last_attempt_at) VALUES ($1, 1, $2)", email, time.Now())
	} else if err != nil {
		return err
	} else {
		attemptCount++
		now := time.Now()

		if attemptCount >= maxAttempts {
			blockedUntil := now.Add(blockDuration)
			_, err = tx.Exec("UPDATE signup_attempts SET attempt_count = $1, last_attempt_at=$2, blocked_until=$3 WHERE email=$4", attemptCount, lastAttempt, blockedUntil, email)
		} else {
			_, err = tx.Exec("UPDATE signup_attempts SET attempt_count=$1, last_attempt_at=$2 WHERE email=$3", attemptCount, lastAttempt, email)
		}
	}

	if err != nil {
		return err
	}

	return tx.Commit()

}
