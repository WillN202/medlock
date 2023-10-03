package web

import (
	"encoding/json"
	"github.com/Manchester-Dev/medlock/internal/account"
	"github.com/Manchester-Dev/medlock/internal/achievements"
	"github.com/go-chi/cors"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

type loginRequest struct {
	Code string `json:"code"`
}

type loginResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

func toType(role string) string {
	switch role {
	case "teacher":
		return "Teacher"
	case "student":
		return "Student"
	default:
		return ""
	}
}

type simpleAchievement struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

type progressResponse struct {
	Achievement simpleAchievement     `json:"achievement"`
	Progress    achievements.Progress `json:"progress"`
}

type achievementResponse struct {
	Achievements []progressResponse `json:"achievements"`
}

func NewRouter(accountStore account.Store, achievementStore achievements.Store) http.Handler {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://localhost*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	router.Get("/health", func(writer http.ResponseWriter, req *http.Request) {
		writer.WriteHeader(http.StatusOK)
	})
	router.Get("/students/{id}/achievements", getStudentAchievements(accountStore, achievementStore))
	router.Put("/students/{id}/achievements/{achievement}/progress", updateAchievementProgress(accountStore, achievementStore))
	router.Post("/achievements", createAchievement(achievementStore))
	router.Get("/achievements", getAllAchievements(achievementStore))
	router.Post("/login", login(accountStore))

	return router
}

type allAchievementsResponse struct {
	Achievements []achievements.Achievement `json:"achievements"`
}

func getAllAchievements(store achievements.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		achvs := store.GetAllAchievements()
		resp := allAchievementsResponse{Achievements: achvs}
		err := json.NewEncoder(w).Encode(resp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

type createAchievementRequest struct {
	Name string `json:"name"`
}

type createAchievementResponse struct {
	ID string `json:"id"`
}

func createAchievement(store achievements.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var achReq createAchievementRequest
		err := json.NewDecoder(req.Body).Decode(&achReq)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if strings.TrimSpace(achReq.Name) == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		id := store.CreateAchievement(achReq.Name)
		err = json.NewEncoder(w).Encode(createAchievementResponse{ID: id})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

type progressUpdateRequest struct {
	Progress string `json:"progress"`
}

func updateAchievementProgress(accountsStore account.Store, achievementsStore achievements.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		id := chi.URLParam(req, "id")
		if !accountsStore.AccountExists(id) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		achievementID := chi.URLParam(req, "achievement")
		if !achievementsStore.AchievementExists(achievementID) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if req.Body == nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		var progReq progressUpdateRequest
		err := json.NewDecoder(req.Body).Decode(&progReq)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if progReq.Progress != string(achievements.NotStarted) &&
			progReq.Progress != string(achievements.Started) &&
			progReq.Progress != string(achievements.Finished) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		achievementsStore.AddProgression(achievements.StudentAchievement{
			AchievementID: achievementID,
			StudentID:     id,
			Progress:      achievements.Progress(progReq.Progress),
		})
	}
}

func getStudentAchievements(accountsStore account.Store, achievementsStore achievements.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		id := chi.URLParam(req, "id")
		if !accountsStore.AccountExists(id) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		achvs, err := achievementsStore.GetStudentAchievements(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		names := make([]string, len(achvs))
		for i, a := range achvs {
			aa, err := achievementsStore.GetAchievement(a.AchievementID)
			if err != nil {
				continue
			}
			names[i] = aa.Name
		}
		err = json.NewEncoder(w).Encode(toAchievementResponse(achvs, names))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func toAchievementResponse(achvs []achievements.StudentAchievement, names []string) achievementResponse {
	a := make([]progressResponse, len(achvs))
	for i, aa := range achvs {
		a[i] = progressResponse{
			Achievement: simpleAchievement{
				Name: names[i],
				ID:   aa.AchievementID,
			},
			Progress: aa.Progress,
		}
	}
	return achievementResponse{Achievements: a}
}

func login(accountStore account.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var loginReq loginRequest
		err := json.NewDecoder(req.Body).Decode(&loginReq)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		loggedIn, err := accountStore.Login(loginReq.Code)
		if _, ok := err.(*account.CodeDoesNotExistError); ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		err = json.NewEncoder(w).Encode(loginResponse{
			ID:   loggedIn.ID(),
			Name: loggedIn.Name(),
			Type: toType(loggedIn.Role()),
		})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
