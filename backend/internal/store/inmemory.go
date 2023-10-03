package store

import (
	"errors"
	"github.com/Manchester-Dev/medlock/internal/achievements"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func NewInMemory() achievements.Store {
	return &inmemory{
		achievements:    make(map[string]achievements.StudentAchievement),
		achievementList: make(map[string]achievements.Achievement),
	}
}

type inmemory struct {
	achievements    map[string]achievements.StudentAchievement
	achievementList map[string]achievements.Achievement
}

func (i *inmemory) GetAchievement(id string) (*achievements.Achievement, error) {
	a, ok := i.achievementList[id]
	if !ok {
		return nil, errors.New("achievement not found")
	}
	return &a, nil
}

func (i *inmemory) GetAllAchievements() []achievements.Achievement {
	var aa []achievements.Achievement
	for _, a := range i.achievementList {
		aa = append(aa, a)
	}
	return aa
}

func (i *inmemory) CreateAchievement(name string) string {
	ach := achievements.Achievement{
		ID:   gonanoid.Must(),
		Name: name,
	}
	i.achievementList[ach.ID] = ach
	return ach.ID
}

func (i *inmemory) AchievementExists(id string) bool {
	_, ok := i.achievementList[id]
	return ok
}

func (i *inmemory) GetStudentAchievements(studentID string) ([]achievements.StudentAchievement, error) {
	var aa []achievements.StudentAchievement
	for _, sa := range i.achievements {
		if sa.StudentID != studentID {
			continue
		}
		aa = append(aa, sa)
	}
	return aa, nil
}

func (i *inmemory) GetStudentsByAchievement(id string) []string {
	var students []string
	for _, sa := range i.achievements {
		if sa.AchievementID != id {
			continue
		}
		students = append(students, sa.StudentID)
	}
	return students
}

func (i *inmemory) AddProgression(progression achievements.StudentAchievement) {
	i.achievements[progression.StudentID+"#"+progression.AchievementID] = progression
}
