package database

import (
	"github.com/lib/pq"
	"math/rand"
	"splitwise-backend/models"
)

func GetGroups() {
}

func InsertInGroupTable(group models.Group) {
	userId := make([]int, 0, rand.Intn(len(group.Members)))
	for _, member := range group.Members {
		if rand.Float64() < 0.5 { // Randomly decide to add a member or not
			userId = append(userId, member)
		}
	}
	query := "INSERT INTO GROUPS (group_name, members) VALUES ($1,$2)"
	_, err := db.Exec(query, group.GroupName, pq.Array(userId))
	if err != nil {
		panic(err)
	}
}

func UpdateGroup(group models.Group) {
	query := "UPDATE groups SET group_name= $1, members =$2 where gid = $3"

	_, err := db.Query(query, group.GroupName, pq.Array(group.Members), group.Gid)
	if err != nil {
		panic(err)
	}

}

func DeleteGroup(gid int) {

	query := "DELETE FROM groups WHERE gid = $1"

	_, err := db.Query(query, gid)
	if err != nil {
		panic(err)
	}
}
