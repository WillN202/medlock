package web

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Manchester-Dev/medlock/internal/account"
	"github.com/Manchester-Dev/medlock/internal/achievements"
	"github.com/Manchester-Dev/medlock/internal/store"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	r := NewRouter(nil, nil)
	require.NotNil(t, r)
	req, err := http.NewRequest("GET", "/health", nil)
	require.NoError(t, err)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestLogin(t *testing.T) {
	// Logging in via the http router requires the router
	// to have access to an accounts store. We can achieve
	// this by injecting the in-memory accounts store as
	// a dependency of the router.
	//
	// Tip: In GoLand, it is possible to add/remove the
	// parameters of a function across all uses via the
	// Refactor > Change Signature... option.

	accountStore := account.NewInMemoryStore()
	student := account.NewStudent("Test Login")
	err := accountStore.SaveAccount(student)
	require.NoError(t, err)
	r := NewRouter(accountStore, nil)
	require.NotNil(t, r)

	t.Run("should return id and name of stored account", func(t *testing.T) {
		loginReq := []byte(fmt.Sprintf(`{"code": "%s"}`, student.Code()))
		req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(loginReq))
		require.NoError(t, err)
		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)
		require.Equal(t, http.StatusOK, rr.Code)
		var resp loginResponse
		err = json.NewDecoder(rr.Body).Decode(&resp)
		require.NoError(t, err)
		assert.Equal(t, student.ID(), resp.ID)
		assert.Equal(t, student.Name(), resp.Name)
	})

	t.Run("should return unauthorised if account with code is not stored", func(t *testing.T) {
		notStoredAccount := account.NewStudent("Don't Store Me")
		loginReq := loginRequest{Code: notStoredAccount.Code()}
		body, err := json.Marshal(loginReq)
		require.NoError(t, err)
		req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
		require.NoError(t, err)
		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)
		require.Equal(t, http.StatusUnauthorized, rr.Code)
	})

	t.Run("should return bad request when given ill-formatted json", func(t *testing.T) {
		loginReq := []byte(fmt.Sprintf(`{"code": "%s"`, student.Code()))
		req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(loginReq))
		require.NoError(t, err)
		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)
		require.Equal(t, http.StatusBadRequest, rr.Code)
	})

}

func TestGetAchievementsForStudent(t *testing.T) {
	accountStore := account.NewInMemoryStore()
	achievementStore := store.NewInMemory()
	student := account.NewStudent("Test Login")
	err := accountStore.SaveAccount(student)
	require.NoError(t, err)
	aID := achievementStore.CreateAchievement("achievement")
	achievementStore.AddProgression(achievements.StudentAchievement{
		AchievementID: aID,
		StudentID:     student.ID(),
		Progress:      achievements.Started,
	})
	r := NewRouter(accountStore, achievementStore)
	require.NotNil(t, r)

	t.Run("should return not found for student not present in accounts store", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/students/student-not-in-store/achievements", nil)
		require.NoError(t, err)
		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)
		require.Equal(t, http.StatusNotFound, rr.Code)
	})

	t.Run("should return a student's achievement", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/students/"+student.ID()+"/achievements", nil)
		require.NoError(t, err)
		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)
		require.Equal(t, http.StatusOK, rr.Code)
		var resp achievementResponse
		err = json.NewDecoder(rr.Body).Decode(&resp)
		require.NoError(t, err)
		require.Len(t, resp.Achievements, 1)
		achievement := resp.Achievements[0]
		assert.Equal(t, "achievement", achievement.Achievement)
		assert.Equal(t, achievements.Started, achievement.Progress)

		rr.Flush()
		a2ID := achievementStore.CreateAchievement("achievement2")
		achievementStore.AddProgression(achievements.StudentAchievement{
			AchievementID: a2ID,
			StudentID:     student.ID(),
			Progress:      achievements.Finished,
		})

		r.ServeHTTP(rr, req)
		require.Equal(t, http.StatusOK, rr.Code)
		err = json.NewDecoder(rr.Body).Decode(&resp)
		require.NoError(t, err)
		require.Len(t, resp.Achievements, 2)
		achievement = resp.Achievements[1]
		assert.Equal(t, "achievement2", achievement.Achievement)
		assert.Equal(t, achievements.Finished, achievement.Progress)
	})
}

func TestUpdateAchievementProgress(t *testing.T) {
	accountStore := account.NewInMemoryStore()
	achievementStore := store.NewInMemory()
	student := account.NewStudent("Test Login")
	err := accountStore.SaveAccount(student)
	require.NoError(t, err)
	achievementID := givenAchievement(achievementStore)
	achievementStore.AddProgression(achievements.StudentAchievement{
		AchievementID: achievementID,
		StudentID:     student.ID(),
		Progress:      achievements.Started,
	})
	r := NewRouter(accountStore, achievementStore)
	require.NotNil(t, r)

	t.Run("should return not found for student not present in accounts store", func(t *testing.T) {
		req, err := http.NewRequest("PUT", "/students/student-not-in-store/achievements/doesnt-matter/progress", nil)
		require.NoError(t, err)
		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)
		require.Equal(t, http.StatusNotFound, rr.Code)
	})

	t.Run("should return not found for achievement not present in accounts store", func(t *testing.T) {
		req, err := http.NewRequest("PUT", fmt.Sprintf("/students/%s/achievements/achievement-not-exist/progress", student.ID()), nil)
		require.NoError(t, err)
		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)
		require.Equal(t, http.StatusNotFound, rr.Code)
	})

	t.Run("should return bad request if missing progress in body", func(t *testing.T) {
		req, err := http.NewRequest("PUT", fmt.Sprintf("/students/%s/achievements/%s/progress", student.ID(), achievementID), nil)
		require.NoError(t, err)
		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)
		require.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("should return bad request if given ill-formatted json", func(t *testing.T) {
		body := bytes.NewBufferString(`{"progress": "STARTED`)
		req, err := http.NewRequest("PUT", fmt.Sprintf("/students/%s/achievements/%s/progress", student.ID(), achievementID), body)
		require.NoError(t, err)
		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)
		require.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("should return bad request if given non-compatible progress", func(t *testing.T) {
		body := bytes.NewBufferString(`{"progress": "STARTEDO"}`)
		req, err := http.NewRequest("PUT", fmt.Sprintf("/students/%s/achievements/%s/progress", student.ID(), achievementID), body)
		require.NoError(t, err)
		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)
		require.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("should update progress for valid request", func(t *testing.T) {
		body := bytes.NewBufferString(`{"progress": "FINISHED"}`)
		req, err := http.NewRequest("PUT", fmt.Sprintf("/students/%s/achievements/%s/progress", student.ID(), achievementID), body)
		require.NoError(t, err)
		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)
		require.Equal(t, http.StatusOK, rr.Code)

		achvs := getStudentAchievementsFromAPI(t, r, student)
		require.Len(t, achvs.Achievements, 1)
		assert.Equal(t, achievements.Finished, achvs.Achievements[0].Progress)
	})
}

func givenAchievement(achievementStore achievements.Store) string {
	aName := "Achievement_" + gonanoid.Must(4)
	return achievementStore.CreateAchievement(aName)
}

func TestGetAllAchievements(t *testing.T) {
	accountStore := account.NewInMemoryStore()
	achievementStore := store.NewInMemory()
	achievement1 := givenAchievement(achievementStore)
	achievement2 := givenAchievement(achievementStore)
	achievement3 := givenAchievement(achievementStore)
	r := NewRouter(accountStore, achievementStore)
	require.NotNil(t, r)

	t.Run("should return all achievements present in store", func(t *testing.T) {
		resp := getAllAchievementsFromAPI(t, r)

		assert.Len(t, resp.Achievements, 3)
		ids := make([]string, len(resp.Achievements))
		for i, a := range resp.Achievements {
			ids[i] = a.ID
		}
		assert.Contains(t, ids, achievement1)
		assert.Contains(t, ids, achievement2)
		assert.Contains(t, ids, achievement3)
	})
}

func getAllAchievementsFromAPI(t *testing.T, r http.Handler) allAchievementsResponse {
	req, err := http.NewRequest(http.MethodGet, "/achievements", nil)
	require.NoError(t, err)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	require.Equal(t, http.StatusOK, rr.Code)
	var resp allAchievementsResponse
	err = json.NewDecoder(rr.Body).Decode(&resp)
	require.NoError(t, err)
	return resp
}

func TestCreateAchievement(t *testing.T) {
	accountStore := account.NewInMemoryStore()
	achievementStore := store.NewInMemory()
	student := account.NewStudent("Test Login")
	err := accountStore.SaveAccount(student)
	require.NoError(t, err)
	r := NewRouter(accountStore, achievementStore)
	require.NotNil(t, r)

	t.Run("should return achievement id on success", func(t *testing.T) {
		body := []byte(`{"name": "Write some code"}`)
		req, err := http.NewRequest(http.MethodPost, "/achievements", bytes.NewBuffer(body))
		require.NoError(t, err)
		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)
		require.Equal(t, http.StatusOK, rr.Code)
		var resp createAchievementResponse
		err = json.NewDecoder(rr.Body).Decode(&resp)
		require.NoError(t, err)
		require.NotEmpty(t, resp.ID)

		assert.True(t, achievementStore.AchievementExists(resp.ID))
		allResp := getAllAchievementsFromAPI(t, r)
		assert.Len(t, allResp.Achievements, 1)
		assert.Contains(t, allResp.Achievements, achievements.Achievement{
			ID:   resp.ID,
			Name: "Write some code",
		})
	})
}

func getStudentAchievementsFromAPI(t *testing.T, r http.Handler, student account.Account) achievementResponse {
	req, err := http.NewRequest(http.MethodGet, "/students/"+student.ID()+"/achievements", nil)
	require.NoError(t, err)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	require.Equal(t, http.StatusOK, rr.Code)
	var resp achievementResponse
	err = json.NewDecoder(rr.Body).Decode(&resp)
	require.NoError(t, err)
	return resp
}
