package httpapi

import (
	"errors"
	"net/http"
	"strconv"

	"example.com/office-template/internal/modules/tasks"
	"example.com/office-template/internal/platform/httpx"
)

type taskHandler struct{ service *tasks.Service }

func (h taskHandler) list(w http.ResponseWriter, r *http.Request) {
	result, err := h.service.List(r.Context())
	if err != nil {
		httpx.WriteError(w, 500, "INTERNAL", "internal server error")
		return
	}
	httpx.WriteJSON(w, 200, map[string]any{"data": result})
}
func (h taskHandler) create(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title string `json:"title"`
	}
	if err := httpx.DecodeJSON(w, r, &input); err != nil {
		httpx.WriteError(w, 400, "INVALID_JSON", "invalid JSON body or content type")
		return
	}
	result, err := h.service.Create(r.Context(), input.Title)
	if errors.Is(err, tasks.ErrInvalidTitle) {
		httpx.WriteError(w, 422, "INVALID_TITLE", err.Error())
		return
	}
	if err != nil {
		httpx.WriteError(w, 500, "INTERNAL", "internal server error")
		return
	}
	httpx.WriteJSON(w, http.StatusCreated, map[string]any{"data": result})
}
func (h taskHandler) update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil || id <= 0 {
		httpx.WriteError(w, 400, "INVALID_ID", "invalid task id")
		return
	}
	var input struct {
		Completed *bool `json:"completed"`
	}
	if err := httpx.DecodeJSON(w, r, &input); err != nil || input.Completed == nil {
		httpx.WriteError(w, 400, "INVALID_JSON", "completed must be a boolean")
		return
	}
	result, err := h.service.SetCompleted(r.Context(), id, *input.Completed)
	if errors.Is(err, tasks.ErrNotFound) {
		httpx.WriteError(w, 404, "NOT_FOUND", "task not found")
		return
	}
	if err != nil {
		httpx.WriteError(w, 500, "INTERNAL", "internal server error")
		return
	}
	httpx.WriteJSON(w, 200, map[string]any{"data": result})
}
func (h taskHandler) delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil || id <= 0 {
		httpx.WriteError(w, 400, "INVALID_ID", "invalid task id")
		return
	}
	err = h.service.Delete(r.Context(), id)
	if errors.Is(err, tasks.ErrNotFound) {
		httpx.WriteError(w, 404, "NOT_FOUND", "task not found")
		return
	}
	if err != nil {
		httpx.WriteError(w, 500, "INTERNAL", "internal server error")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
