package api

import (
	"cs361new/pkg/db"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-pg/pg/v10"
	"log"
	"net/http"
	"strconv"
)

func NewAPI(pgdb *pg.DB) *chi.Mux {

	//router setup below:
	r := chi.NewRouter()
	r.Use(middleware.Logger, middleware.WithValue("DB", pgdb))

	r.Route("/homes", func(r chi.Router) {
		r.Post("/", createHome)
		r.Get("/{homeID}", getHomeByID)
		r.Get("/", getHomes)
		r.Put("/{homeID}", updateHomeByID)
		r.Delete("/{homeID}", deleteHomeByID)
	})
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		//		w.Write([]byte("hello world"))
	})

	return r
}

type CreateHomeRequest struct {
	ID      int64 `json:"id"`
	Price   int64 `json:"price"`
	AgentID int64 `json:"agent_id"`
}

type HomeResponse struct {
	Success bool     `json:"success"`
	Error   string   `json:"error"`
	Home    *db.Home `json:"home"`
}

type CreateHomeResponse struct {
	Success bool     `json:"success"`
	Error   string   `json:"error"`
	Home    *db.Home `json:"home"`
}

//handlers begin crud create get update delete homes from tut switch to gym goers or some
// variant **
func createHome(w http.ResponseWriter, r *http.Request) {

	//parse step:
	req := &CreateHomeRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		res := &HomeResponse{
			Success: false,
			Error:   err.Error(),
			Home:    nil,
		}
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response: %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//get db context:
	pgd, ok := r.Context().Value("DB").(*pg.DB)
	if !ok {
		res := &HomeResponse{
			Success: false,
			Error:   "could not get db from context",
			Home:    nil,
		}
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response: %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//insert db obj:
	home, err := db.CreateHome(pgd, &db.Home{
		Price:   req.Price,
		AgentID: req.AgentID,
		//Agent:   req.Agent,    ////////////////////////////////////////////////
	})
	if err != nil {
		res := &HomeResponse{
			Success: false,
			Error:   err.Error(),
			Home:    nil,
		}
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response: %v\n", err)

		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//return response:
	res := &HomeResponse{
		Success: true,
		Error:   "",
		Home:    home,
	}

	_ = json.NewEncoder(w).Encode(res)
	w.WriteHeader(http.StatusOK)
}

//type GetHomesResponse struct {
//	Homes []db.Home `json:"homes"`
//}

type HomesResponse struct {
	Success bool       `json:"success"`
	Error   string     `json:"error"`
	Homes   []*db.Home `json:"homes"`
}

func getHomes(w http.ResponseWriter, r *http.Request) {
	//obtain db context:
	pgd, ok := r.Context().Value("DB").(*pg.DB)
	if !ok {
		res := &HomeResponse{
			Success: false,
			Error:   "could not get db from context",
			Home:    nil,
		}
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response: %v\n", err)

		}

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	homes, err := db.GetHomes(pgd)
	if err != nil {
		res := &HomesResponse{
			Success: false,
			Error:   err.Error(),
			Homes:   nil,
		}
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response: %v\n", err)

		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//ret response:
	//return response:
	res := &HomesResponse{
		Success: true,
		Error:   "",
		Homes:   homes,
	}

	_ = json.NewEncoder(w).Encode(res)
	w.WriteHeader(http.StatusOK)

}

//type GetHomeByIDResponse struct {
//}

//w.Write([]byte("Create GYm User"))
//}

func getHomeByID(w http.ResponseWriter, r *http.Request) {
	homeID := chi.URLParam(r, "homeID")

	//obtain db context:
	pgd, ok := r.Context().Value("DB").(*pg.DB)
	if !ok {
		res := &HomeResponse{
			Success: false,
			Error:   "could not get db from context",
			Home:    nil,
		}
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response: %v\n", err)

		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	home, err := db.GetHome(pgd, homeID)
	if err != nil {
		res := &HomeResponse{
			Success: false,
			Error:   err.Error(),
			Home:    nil,
		}
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response: %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//return response:
	res := &HomeResponse{
		Success: true,
		Error:   "",
		Home:    home,
	}
	_ = json.NewEncoder(w).Encode(res)
	w.WriteHeader(http.StatusOK)

	//w.Write([]byte(fmt.Sprintf("get home: %s", homeID)))

}

type updateHomeByIDRequest struct {
	ID      int64 `json:"id"`
	Price   int64 `json:"price"`
	AgentID int64 `json:"agent_id"`
}

func updateHomeByID(w http.ResponseWriter, r *http.Request) {
	homeID := chi.URLParam(r, "homeID")
	//parse step:
	req := &updateHomeByIDRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		res := &HomeResponse{
			Success: false,
			Error:   err.Error(),
			Home:    nil,
		}
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response: %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//get db context:
	pgd, ok := r.Context().Value("DB").(*pg.DB)
	if !ok {
		res := &HomeResponse{
			Success: false,
			Error:   "could not get db from context",
			Home:    nil,
		}
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response: %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//return err response:
	intHomeID, err := strconv.ParseInt(homeID, 10, 64)
	if err != nil {
		res := &HomeResponse{
			Success: false,
			Error:   err.Error(),
			Home:    nil,
		}
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response: %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	home, err := db.UpdateHome(pgd, &db.Home{
		ID:      intHomeID,
		Price:   req.Price,
		AgentID: req.AgentID,
	})
	//return err response:
	if err != nil {
		res := &HomeResponse{
			Success: false,
			Error:   err.Error(),
			Home:    nil,
		}
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response: %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//return home response:
	res := &HomeResponse{
		Success: true,
		Error:   "",
		Home:    home,
	}

	_ = json.NewEncoder(w).Encode(res)
	w.WriteHeader(http.StatusOK)
}

func deleteHomeByID(w http.ResponseWriter, r *http.Request) {

	homeID := chi.URLParam(r, "homeID")

	//parse step:
	req := &updateHomeByIDRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		res := &HomeResponse{
			Success: false,
			Error:   err.Error(),
			Home:    nil,
		}
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response: %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//get db context:
	pgd, ok := r.Context().Value("DB").(*pg.DB)
	if !ok {
		res := &HomeResponse{
			Success: false,
			Error:   "could not get db from context",
			Home:    nil,
		}
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response: %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//return err response:
	intHomeID, err := strconv.ParseInt(homeID, 10, 64)
	if err != nil {
		res := &HomeResponse{
			Success: false,
			Error:   err.Error(),
			Home:    nil,
		}
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response: %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//update home dont pass anything back on delete
	err = db.DeleteHome(pgd, intHomeID)
	//return err response:
	if err != nil {
		res := &HomeResponse{
			Success: false,
			Error:   err.Error(),
			Home:    nil,
		}
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response: %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//return home response:
	res := &HomeResponse{
		Success: true,
		Error:   "",
		Home:    nil,
	}

	_ = json.NewEncoder(w).Encode(res)
	w.WriteHeader(http.StatusOK)
}
