package achievements

type Progress string

var NotStarted Progress = ""
var Started Progress = "STARTED"
var Finished Progress = "FINISHED"

// StudentAchievement represents a particular student's progress of a
// set Achievement.
type StudentAchievement struct {
	AchievementID string   `json:"achievement"`
	StudentID     string   `json:"studentId"`
	Progress      Progress `json:"progress"`
}

// Achievement represents a real life achievement of
// which student can progress.
type Achievement struct {
	ID   string
	Name string
}
