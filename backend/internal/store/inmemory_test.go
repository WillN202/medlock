package store

import (
	"github.com/Manchester-Dev/medlock/internal/account"
	"github.com/Manchester-Dev/medlock/internal/achievements"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetStudentAchievements(t *testing.T) {
	s := NewInMemory()
	targetStudent := account.NewStudent("Target Student")
	notTargetStudent := account.NewStudent("Not Target Student")

	t.Run("should be empty when student is not progressing on any achievements", func(t *testing.T) {
		studentAchievements, err := s.GetStudentAchievements(targetStudent.ID())
		require.NoError(t, err)
		assert.Empty(t, studentAchievements)
	})

	t.Run("should return the achievements assigned to a student", func(t *testing.T) {
		aID := gonanoid.Must()
		s.AddProgression(achievements.StudentAchievement{
			StudentID:     targetStudent.ID(),
			AchievementID: aID,
			Progress:      achievements.Started,
		})
		s.AddProgression(achievements.StudentAchievement{
			StudentID:     notTargetStudent.ID(),
			AchievementID: aID,
			Progress:      achievements.Started,
		})
		studentAchievements, err := s.GetStudentAchievements(targetStudent.ID())
		require.NoError(t, err)
		require.Len(t, studentAchievements, 1)
		ach := studentAchievements[0]
		assert.Equal(t, targetStudent.ID(), ach.StudentID)
	})
}

func TestGetStudentsByAchievement(t *testing.T) {

	s := NewInMemory()

	t.Run("should return empty when there are no students progressing with the achievement", func(t *testing.T) {
		aID := gonanoid.Must()
		all := s.GetStudentsByAchievement(aID)
		assert.Empty(t, all)
	})

	t.Run("should return all students which are progressing on a particular achievement", func(t *testing.T) {
		student := account.NewStudent("Test Student")
		aID := gonanoid.Must()
		progressing := achievements.StudentAchievement{
			StudentID:     student.ID(),
			AchievementID: aID,
			Progress:      achievements.Started,
		}

		s.AddProgression(progressing)
		all := s.GetStudentsByAchievement(aID)

		require.Len(t, all, 1)
		assert.Equal(t, student.ID(), all[0])

		student2 := account.NewStudent("Test Student")
		s.AddProgression(achievements.StudentAchievement{
			AchievementID: aID,
			StudentID:     student2.ID(),
			Progress:      achievements.Started,
		})
		studentNotProgressingAchievement := account.NewStudent("Test Student")
		s.AddProgression(achievements.StudentAchievement{
			AchievementID: "a123",
			StudentID:     studentNotProgressingAchievement.ID(),
			Progress:      achievements.Started,
		})
		all = s.GetStudentsByAchievement(aID)

		require.Len(t, all, 2)
		assert.Contains(t, all, student.ID())
		assert.Contains(t, all, student2.ID())
		assert.NotContains(t, all, studentNotProgressingAchievement.ID())
	})

}

func TestAddProgression(t *testing.T) {

	s := NewInMemory()

	t.Run("should add a student's progression on a particular achievement", func(t *testing.T) {
		student := account.NewStudent("Test Student")
		aID := gonanoid.Must()
		progressing := achievements.StudentAchievement{
			StudentID:     student.ID(),
			AchievementID: aID,
		}
		s.AddProgression(progressing)

		ss, ok := s.(*inmemory)
		require.True(t, ok)
		p, ok := ss.achievements[student.ID()+"#"+progressing.AchievementID]
		require.True(t, ok)
		assert.Equal(t, progressing, p)
	})

	t.Run("should update a student's progression on a particular achievement", func(t *testing.T) {
		student := account.NewStudent("Test Student")
		aID := gonanoid.Must()
		progressing := achievements.StudentAchievement{
			StudentID:     student.ID(),
			AchievementID: aID,
		}
		s.AddProgression(progressing)

		updated := achievements.StudentAchievement{
			StudentID:     student.ID(),
			AchievementID: aID,
		}
		updated.Progress = achievements.Started
		s.AddProgression(updated)
		ss, ok := s.(*inmemory)
		require.True(t, ok)
		p, ok := ss.achievements[student.ID()+"#"+progressing.AchievementID]
		require.True(t, ok)
		assert.Equal(t, updated, p)
	})

}

func TestAchievementExists(t *testing.T) {
	s := NewInMemory()

	t.Run("should return true for existing achievement", func(t *testing.T) {
		ach := "THIS_IS_AN_ACHIEVEMENT"
		id := s.CreateAchievement(ach)
		assert.True(t, s.AchievementExists(id))
	})
}

func TestGetAchievement(t *testing.T) {
	t.Parallel()
	store := NewInMemory()

	t.Run("should throw an error if achievement does not exist", func(t *testing.T) {
		_, err := store.GetAchievement("non-existent-id")
		require.Error(t, err)
		assert.Contains(t, err.Error(), "not found")
	})
}
