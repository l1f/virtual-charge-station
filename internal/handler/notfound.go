package handler

import (
	"context"
	"net/http"

	"github.com/l1f/virtual-charge-station/web/templates"
)

func (h *Handler) NotFound(w http.ResponseWriter, r *http.Request) {
	component := templates.NotFound()

	component.Render(context.Background(), w)
}
