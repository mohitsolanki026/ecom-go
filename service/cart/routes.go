package cart

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/mohitsolanki026/econ-go/service/auth"
	"github.com/mohitsolanki026/econ-go/types"
	"github.com/mohitsolanki026/econ-go/utils"
)

type Handler struct {
	Store        types.OrderStore
	ProductStore types.ProductStore
	UserStore types.UserStore
}

func NewHandler(store types.OrderStore, productStore types.ProductStore, userStore types.UserStore) *Handler {
	return &Handler{Store: store, ProductStore: productStore, UserStore: userStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/cart/checkout",auth.WithJwtAuth(h.handleCheckout, h.UserStore)).Methods(http.MethodPost)
}

func (h *Handler) handleCheckout(w http.ResponseWriter, r *http.Request) {

	userId := auth.GetUserIDFromContext(r.Context())

	var cart types.CartCheckoutPayload
	if err := utils.ParseJSON(r, &cart); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(cart); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	productIds, err := GetCartItemsIDs(cart.Items)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", err))
		return
	}

	ps, err := h.ProductStore.GetProductByIDs(productIds)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", err))
		return
	}

	orderId, totalPrice, err := h.createOrder(ps, cart.Items, userId)
	utils.WriteJSON(w, http.StatusCreated, map[string]any{
		"total_price": totalPrice,
		"order_id":    orderId,
	})

	return
}
