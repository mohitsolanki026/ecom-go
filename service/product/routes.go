package product

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/mohitsolanki026/econ-go/types"
	"github.com/mohitsolanki026/econ-go/utils"
)

type Handler struct {
	Store types.ProductStore
}

func NewHandler(store types.ProductStore) *Handler{
	return &Handler{Store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router){
	router.HandleFunc("/products",h.handleGetProduct).Methods(http.MethodGet)
	router.HandleFunc("/product",h.handleCreateProduct).Methods(http.MethodPost)
}

func (h *Handler) handleGetProduct(w http.ResponseWriter, r *http.Request){
	products, err := h.Store.GetProducts()

	if err != nil{
		utils.WriteError(w,http.StatusInternalServerError,err)
	}
	
	utils.WriteJSON(w, http.StatusAccepted, products)
}

func (h *Handler) handleCreateProduct(w http.ResponseWriter, r *http.Request){
	var payload types.Product

	if err := utils.ParseJSON(r, &payload); err != nil {
		fmt.Println("error here")
		utils.WriteError(w,http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		error := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest,fmt.Errorf("validation error: %v",error))
		return
	}

	err := h.Store.CreateProduct(&payload)
	
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	
	utils.WriteJSON(w, http.StatusCreated, payload)

}

