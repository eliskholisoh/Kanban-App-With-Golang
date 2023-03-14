package api

import (
	"a21hc3NpZ25tZW50/entity"
	"a21hc3NpZ25tZW50/service"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type CategoryAPI interface {
	GetCategory(w http.ResponseWriter, r *http.Request)
	CreateNewCategory(w http.ResponseWriter, r *http.Request)
	DeleteCategory(w http.ResponseWriter, r *http.Request)
	GetCategoryWithTasks(w http.ResponseWriter, r *http.Request)
}

type categoryAPI struct {
	categoryService service.CategoryService
}

func NewCategoryAPI(categoryService service.CategoryService) *categoryAPI {
	return &categoryAPI{categoryService}
}

func (c *categoryAPI) GetCategory(w http.ResponseWriter, r *http.Request) {
	// TODO: answer here
	usrId := r.Context().Value("id")
	if usrId == nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid user id"))
		return
	}
	idUser, err := strconv.Atoi(usrId.(string))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
		return
	}

	categoris, err := c.categoryService.GetCategories(r.Context(), idUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(categoris)
}

func (c *categoryAPI) CreateNewCategory(w http.ResponseWriter, r *http.Request) {
	var category entity.CategoryRequest

	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid category request"))
		return
	}

	// TODO: answer here
	if category.Type == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid category request"))
		return
	}

	usrId := r.Context().Value("id")
	if usrId == nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid user id"))
		return
	}

	idUser, err := strconv.Atoi(usrId.(string))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
		return
	}

	ctg := entity.Category{Type: category.Type, UserID: idUser}
	categoris, err := c.categoryService.StoreCategory(r.Context(), &ctg)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{"user_Id": categoris.UserID, "category_id": categoris.ID, "message": "success create new category"})
}

func (c *categoryAPI) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	// TODO: answer here
	idURL := r.URL.Query().Get("category_id")
	ctgId, err := strconv.Atoi(idURL)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	err = c.categoryService.DeleteCategory(r.Context(), ctgId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	idUsr := r.Context().Value("id").(string)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"user_id": idUsr, "category_id": ctgId, "message": "success delete category"})
}

func (c *categoryAPI) GetCategoryWithTasks(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("id")

	idLogin, err := strconv.Atoi(userId.(string))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("get category task", err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid user id"))
		return
	}

	categories, err := c.categoryService.GetCategoriesWithTasks(r.Context(), int(idLogin))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("internal server error"))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(categories)

}
