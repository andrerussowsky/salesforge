package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"

    "salesforge/internal/models"
    "salesforge/internal/repository"

    "github.com/gorilla/mux"
)

// CreateSequence creates a new sequence with steps
func CreateSequence(w http.ResponseWriter, r *http.Request) {
    var sequence models.Sequence
    err := json.NewDecoder(r.Body).Decode(&sequence)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    sequenceID, err := repository.CreateSequence(&sequence)
    if err != nil {
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]int{"sequence_id": sequenceID})
}

// UpdateSequenceStep updates a sequence step
func UpdateSequenceStep(w http.ResponseWriter, r *http.Request) {
    sequenceID, err := strconv.Atoi(mux.Vars(r)["sequenceID"])
    if err != nil {
        http.Error(w, "Invalid sequence ID", http.StatusBadRequest)
        return
    }

    stepID, err := strconv.Atoi(mux.Vars(r)["stepID"])
    if err != nil {
        http.Error(w, "Invalid step ID", http.StatusBadRequest)
        return
    }

    var step models.Step
    err = json.NewDecoder(r.Body).Decode(&step)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    step.ID = stepID
    err = repository.UpdateSequenceStep(sequenceID, stepID, &step)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}

// DeleteSequenceStep deletes a sequence step
func DeleteSequenceStep(w http.ResponseWriter, r *http.Request) {
    sequenceID, err := strconv.Atoi(mux.Vars(r)["sequenceID"])
    if err != nil {
        http.Error(w, "Invalid sequence ID", http.StatusBadRequest)
        return
    }

    stepID, err := strconv.Atoi(mux.Vars(r)["stepID"])
    if err != nil {
        http.Error(w, "Invalid step ID", http.StatusBadRequest)
        return
    }

    err = repository.DeleteSequenceStep(sequenceID, stepID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}

// UpdateSequenceTracking updates sequence tracking
func UpdateSequenceTracking(w http.ResponseWriter, r *http.Request) {
    sequenceID, err := strconv.Atoi(mux.Vars(r)["sequenceID"])
    if err != nil {
        http.Error(w, "Invalid sequence ID", http.StatusBadRequest)
        return
    }

    var tracking models.Tracking
    err = json.NewDecoder(r.Body).Decode(&tracking)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    err = repository.UpdateSequenceTracking(sequenceID, &tracking)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}
