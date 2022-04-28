package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/klyngen/starbook/person"
	"github.com/klyngen/starbook/star"
	"log"
	"net/http"
	"strconv"
)

type Presentation struct {
	personRepo person.PersonRepository
	starRepo   star.StarRepository
	router     *mux.Router
}

func NewPresentationLayer(personRepo person.PersonRepository, starRepo star.StarRepository) Presentation {
	presentation := Presentation{
		personRepo: personRepo,
		starRepo:   starRepo,
	}

	router := mux.NewRouter()
	presentation.router = router
	presentation.CreateRoutes()
	router.Use(mux.CORSMethodMiddleware(router))
	return presentation
}

func (p *Presentation) CreateRoutes() {
	p.router.HandleFunc("/*", func(writer http.ResponseWriter, r *http.Request) {
		writer.Header().Set("Access-Control-Allow-Headers", "*")
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Content-Type", "application/json")
	}).Methods(http.MethodOptions)

	p.router.HandleFunc("/person", p.HandleCreatePerson).Methods("POST", http.MethodOptions)
	p.router.HandleFunc("/person", p.HandleGetPersons).Methods("GET")
	p.router.HandleFunc("/star", p.HandleAddStar).Methods("POST", http.MethodOptions)
	p.router.HandleFunc("/star/{id}", p.HandleGetPersonStars).Methods("GET")
}

func (p *Presentation) StartServer(port string) {
	server := &http.Server{
		Handler: p.router,
		Addr:    "0.0.0.0:" + port,
	}

	log.Fatal(server.ListenAndServe())
}
func (p *Presentation) HandleCreatePerson(writer http.ResponseWriter, r *http.Request) {
	var person person.Person
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Headers", "*")
	writer.Header().Set("Access-Control-Allow-Origin", "*")

	json.NewDecoder(r.Body).Decode(&person)

	log.Println(person)

	encoder := json.NewEncoder(writer)

	if !person.IsValidPerson() {
		encoder.Encode(struct{ message string }{
			message: "Missing fileds",
		})
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err := p.personRepo.Create(&person)

	if err != nil {
		encoder.Encode(struct{ message string }{
			message: "Unknown server error",
		})
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (p *Presentation) HandleGetPersonStars(writer http.ResponseWriter, r *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Headers", "*")
	writer.Header().Set("Access-Control-Allow-Origin", "*")

	variables := mux.Vars(r)
	id, _ := strconv.Atoi(variables["id"])

	encoder := json.NewEncoder(writer)

	stars, _ := p.starRepo.GetStarsByPerson(uint(id))

	encoder.Encode(stars)
}

func (p *Presentation) HandleAddStar(writer http.ResponseWriter, r *http.Request) {
	writer.Header().Set("Access-Control-Allow-Headers", "*")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Content-Type", "application/json")

	var star star.Star
	json.NewDecoder(r.Body).Decode(&star)
	encoder := json.NewEncoder(writer)

	if !star.IsValidStar() {
		encoder.Encode(struct{ message string }{
			message: "Missing fileds",
		})
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err := p.starRepo.Create(&star)

	if err != nil {
		encoder.Encode(struct{ message string }{
			message: "Unknown server error",
		})
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
}

func (p *Presentation) HandleGetPersons(writer http.ResponseWriter, r *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	encoder := json.NewEncoder(writer)

	persons, err := p.personRepo.AllPersons()

	if err != nil {
		encoder.Encode(struct{ message string }{
			message: "Cannot fetch persons",
		})
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	encoder.Encode(&persons)
}
