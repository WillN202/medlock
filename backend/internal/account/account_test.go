package account

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewStudent(t *testing.T) {
	t.Parallel()

	t.Run("should create an account with a given name", func(t *testing.T) {
		student := NewStudent("Student A")
		require.NotNil(t, student)
		assert.Equal(t, "Student A", student.Name())
	})

	t.Run("should create an account with a unique id", func(t *testing.T) {
		student := NewStudent("Student B")
		require.NotNil(t, student)
		studentWithSameName := NewStudent("Student B")
		require.NotNil(t, studentWithSameName)
		assert.NotEqual(t, student.ID(), studentWithSameName.ID())
	})

	t.Run("should create an account with the student role", func(t *testing.T) {
		student := NewStudent("Student C")
		require.NotNil(t, student)
		assert.Equal(t, "student", student.Role())
	})
}

func TestNewTeacher(t *testing.T) {
	t.Parallel()

	t.Run("should create an account with a given name", func(t *testing.T) {
		teacher := NewTeacher("Teacher A")
		require.NotNil(t, teacher)
		assert.Equal(t, "Teacher A", teacher.Name())
	})

	t.Run("should create an account with a unique id", func(t *testing.T) {
		teacher := NewTeacher("Teacher B")
		require.NotNil(t, teacher)
		teacherWithSameName := NewTeacher("Teacher B")
		require.NotNil(t, teacherWithSameName)
		assert.NotEqual(t, teacher.ID(), teacherWithSameName.ID())
	})

	t.Run("should create an account with the teacher role", func(t *testing.T) {
		teacher := NewTeacher("Teacher C")
		require.NotNil(t, teacher)
		assert.Equal(t, "teacher", teacher.Role())
	})
}
