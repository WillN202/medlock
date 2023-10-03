package account

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSaveAccount(t *testing.T) {
	t.Parallel()
	store := NewInMemoryStore()
	s := NewStudent("Student A")
	err := store.SaveAccount(s)
	require.NoError(t, err)
	imStore, ok := store.(*inmemory)
	require.True(t, ok)
	assert.Len(t, imStore.accounts, 1)

	err = store.SaveAccount(s)
	require.Error(t, err)
	_, ok = err.(CodeConflictError)
	assert.True(t, ok)
}

func TestLogin(t *testing.T) {
	t.Parallel()
	store := NewInMemoryStore()
	s := NewStudent("Student A")
	err := store.SaveAccount(s)
	require.NoError(t, err)
	loggedIn, err := store.Login(s.Code())
	require.NoError(t, err)
	assert.Equal(t, s.ID(), loggedIn.ID())
	assert.Equal(t, s.Name(), loggedIn.Name())
	assert.Equal(t, s.Role(), loggedIn.Role())
}

func TestAccountExists(t *testing.T) {
	t.Parallel()
	store := NewInMemoryStore()
	s := NewStudent("Student A")
	err := store.SaveAccount(s)
	require.NoError(t, err)
	exists := store.AccountExists(s.ID())
	assert.True(t, exists)
	assert.False(t, store.AccountExists("this-doesn't-exist"))
}
