package handler

import (
	"errors"
	"net/http"

	"github.com/ruhollahh/go-progressive-rendering/api/httperrors"
	"github.com/ruhollahh/go-progressive-rendering/internal/service"
	"github.com/ruhollahh/go-progressive-rendering/web/view/components"
	"github.com/ruhollahh/go-progressive-rendering/web/view/pages"
)

func (h *Handler) ShowProduct(w http.ResponseWriter, r *http.Request) {
	id, err := h.getIDParam(r)
	if err != nil {
		httperrors.NotFoundError(w)
		return
	}

	product, err := h.Services.Products.Get(id)
	if err != nil {
		if errors.Is(err, service.ErrRecordNotFound) {
			httperrors.NotFoundError(w)
			return
		}

		httperrors.ServerError(h.Logger, w, r, err)
		return
	}

	rc := http.NewResponseController(w)

	model := pages.ProductViewModel{Product: product}

	err = pages.Product(model).Render(r.Context(), w)
	if err != nil {
		httperrors.ServerError(h.Logger, w, r, err)
		return
	}

	if err := rc.Flush(); err != nil {
		httperrors.ServerError(h.Logger, w, r, err)
		return
	}

	reviews, err := h.Services.Reviews.GetAllForProduct(id)
	if err != nil {
		httperrors.ServerError(h.Logger, w, r, err)
		return
	}

	err = components.Swap("#reviews", components.Reviews(reviews)).Render(r.Context(), w)
	if err != nil {
		httperrors.ServerError(h.Logger, w, r, err)
		return
	}
}
