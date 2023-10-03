package main

import (
	"fmt"
	"github.com/Manchester-Dev/medlock/internal/account"
	"github.com/Manchester-Dev/medlock/internal/achievements"
	"github.com/Manchester-Dev/medlock/internal/store"
	"github.com/Manchester-Dev/medlock/internal/web"
	"math/rand"
	"net/http"
)

func main() {
	aa := []achievements.Progress{
		achievements.Started,
		achievements.Finished,
	}
	achvs := []string{
		"Set Up Recycling Boxes",
		"Bring in Eco Friendly Water Bottle",
		"Fix a Broken Toy",
		"Stitch up a Hole in some Clothing",
		"Start a Compost Heap",
		"Plant some Seeds",
		"Recycle 10 Batteries",
		"Make a Sculpture out of Bottle Caps",
		"Donate Old Clothing",
	}
	accountStore := account.NewInMemoryStore()
	student := account.NewStudent("Test Login")
	teacher := account.NewTeacher("Test Teacher")
	err := accountStore.SaveAccount(student)
	check(err)
	err = accountStore.SaveAccount(teacher)
	check(err)
	achvStore := store.NewInMemory()
	for i := 0; i < 9; i++ {
		id := achvStore.CreateAchievement(achvs[i])
		r := rand.Int() % 2
		achvStore.AddProgression(achievements.StudentAchievement{
			AchievementID: id,
			StudentID:     student.ID(),
			Progress:      aa[r],
		})
	}
	fmt.Printf("stored student with code: %s\n", student.Code())
	fmt.Printf("stored teacher with code: %s\n", teacher.Code())
	r := web.NewRouter(accountStore, achvStore)
	http.ListenAndServe(":4000", r)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
