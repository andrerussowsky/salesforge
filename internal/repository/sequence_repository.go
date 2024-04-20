package repository

import (
    "database/sql"
    "fmt"
    "salesforge/internal/models"
)

// CreateSequenceWithSteps creates a new sequence with steps in the database
func CreateSequence(sequence *models.Sequence) (int, error) {
    var sequenceID int

    // Start a transaction
    tx, err := db.Begin()
    if err != nil {
        return 0, fmt.Errorf("could not begin transaction: %v", err)
    }

    // Insert sequence into sequences table
    err = tx.QueryRow("INSERT INTO sequences (name, open_tracking_enabled, click_tracking_enabled) VALUES ($1, $2, $3) RETURNING id",
        sequence.Name, sequence.OpenTrackingEnabled, sequence.ClickTrackingEnabled).Scan(&sequenceID)
    if err != nil {
        tx.Rollback()
        return 0, fmt.Errorf("could not insert sequence: %v", err)
    }

    // Insert steps into steps table
    for _, step := range sequence.Steps {
        _, err := tx.Exec("INSERT INTO steps (sequence_id, email_subject, email_content) VALUES ($1, $2, $3)",
            sequenceID, step.EmailSubject, step.EmailContent)
        if err != nil {
            tx.Rollback()
            return 0, fmt.Errorf("could not insert step: %v", err)
        }
    }

    // Commit the transaction
    if err := tx.Commit(); err != nil {
        return 0, fmt.Errorf("could not commit transaction: %v", err)
    }

    return sequenceID, nil
}

// UpdateSequenceStep updates a sequence step in the database
func UpdateSequenceStep(sequenceID, stepID int, step *models.Step) error {
    _, err := db.Exec("UPDATE steps SET email_subject = $1, email_content = $2 WHERE id = $3 AND sequence_id = $4",
        step.EmailSubject, step.EmailContent, stepID, sequenceID)
    if err != nil {
        return fmt.Errorf("could not update sequence step: %v", err)
    }

    return nil
}

// DeleteSequenceStep deletes a sequence step from the database
func DeleteSequenceStep(sequenceID, stepID int) error {
    _, err := db.Exec("DELETE FROM steps WHERE id = $1 AND sequence_id = $2", stepID, sequenceID)
    if err != nil {
        return fmt.Errorf("could not delete sequence step: %v", err)
    }

    return nil
}

// UpdateSequenceTracking updates sequence open or click tracking in the database
func UpdateSequenceTracking(sequenceID int, tracking *models.Tracking) error {
    _, err := db.Exec("UPDATE sequences SET open_tracking_enabled = $1, click_tracking_enabled = $2 WHERE id = $3",
        tracking.OpenTracking, tracking.ClickTracking, sequenceID)
    if err != nil {
        return fmt.Errorf("could not update sequence tracking: %v", err)
    }

    return nil
}

var db *sql.DB

// InitDB initializes the database connection
func InitDB(database *sql.DB) {
    db = database
}
