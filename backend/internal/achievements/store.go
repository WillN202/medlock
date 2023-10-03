package achievements

type Store interface {
	GetStudentAchievements(id string) ([]StudentAchievement, error)
	GetStudentsByAchievement(achievement string) []string
	AddProgression(progression StudentAchievement)
	AchievementExists(id string) bool
	CreateAchievement(name string) string
	GetAllAchievements() []Achievement
	GetAchievement(id string) (*Achievement, error)
}
