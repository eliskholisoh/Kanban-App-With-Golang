package api

import (
	"a21hc3NpZ25tZW50/entity"
	"a21hc3NpZ25tZW50/service"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type UserAPI interface {
	Login(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)

	Delete(w http.ResponseWriter, r *http.Request)
}

type userAPI struct {
	userService service.UserService
}

func NewUserAPI(userService service.UserService) *userAPI {
	return &userAPI{userService}
}

func (u *userAPI) Login(w http.ResponseWriter, r *http.Request) {
	var user entity.UserLogin

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid decode json"))
		return
	}
	// TODO: answer here
	if user.Email == "" || user.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("email or password is empty"))
		return
	}
	etyUsr := entity.User{Email: user.Email, Password: user.Password}
	login, err := u.userService.Login(r.Context(), &etyUsr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "user_id",
		Value: strconv.Itoa(login),
		Path:  "/",
	})

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"user_id": login, "message": "login success"})
}

func (u *userAPI) Register(w http.ResponseWriter, r *http.Request) {
	var user entity.UserRegister

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid decode json"))
		return
	}
	// TODO: answer here
	if user.Fullname == "" || user.Email == "" || user.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("register data is empty"))
		return
	}
	users := &entity.User{Fullname: user.Fullname, Email: user.Email, Password: user.Password}
	regis, err := u.userService.Register(r.Context(), users)
	if err != nil {
		if err.Error() == "email already exists" {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(entity.NewErrorResponse(err.Error()))
			return
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
			return
		}

	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{"user_id": regis.ID, "message": "register success"})
}

func (u *userAPI) Logout(w http.ResponseWriter, r *http.Request) {
	// TODO: answer here
	http.SetCookie(w, &http.Cookie{})
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "logout success"})
}

func (u *userAPI) Delete(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("user_id")

	if userId == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("user_id is empty"))
		return
	}

	deleteUserId, _ := strconv.Atoi(userId)

	err := u.userService.Delete(r.Context(), int(deleteUserId))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "delete success"})
}
