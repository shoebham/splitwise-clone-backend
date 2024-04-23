package router

import (
	"github.com/go-faker/faker/v4"
	"splitwise-backend/models"
)

func CreateFakeUsers() []models.User {
	usersArr := []models.User{}
	for i := 0; i < 10; i++ {
		name := faker.Name()
		// owes := make(map[*models.User]float64) //rand.Intn(500)
		// owed := make(map[*models.User]float64) //rand.Intn(500)
		// balance := 0                           //owed - owes
		number := faker.Phonenumber()
		usersArr = append(usersArr, models.User{
			Name:   name,
			Number: number,
		})
	}
	return usersArr
}

func createFakeGroups() []models.Group {
	groupsArr := []models.Group{}
	fakeUsersId := make([]int, len(fakeUsers))
	for _, user := range fakeUsers {
		fakeUsersId = append(fakeUsersId, user.Uid)
	}
	for i := 0; i < 10; i++ {
		group_name := faker.Word()

		groupsArr = append(groupsArr, models.Group{
			GroupName: group_name,
			Members:   fakeUsersId,
		})
	}
	return groupsArr
}
