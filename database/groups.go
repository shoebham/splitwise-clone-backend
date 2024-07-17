package database

import (
	"errors"
	"github.com/gofiber/fiber/v3/log"
	"github.com/lib/pq"
	"math/rand"
	"splitwise-backend/models"
	"strings"
)

func GetGroups() {
}

func InsertInGroupTable(group models.Group) {
	groupId := make([]int, 0, rand.Intn(len(group.Members)))
	for _, member := range group.Members {
		if rand.Float64() < 0.5 { // Randomly decide to add a member or not
			groupId = append(groupId, member)
		}
	}
	query := "INSERT INTO GROUPS (group_name, members) VALUES ($1,$2)"
	_, err := db.Exec(query, group.GroupName, pq.Array(groupId))
	if err != nil {
		panic(err)
	}
}

func UpdateGroup(group models.Group) error {

	exists, _ := CheckIdExists("groups", "gid", group.Gid)
	if exists {
		query := "UPDATE groups SET group_name= $1, members =$2 where gid = $3"

		_, err := db.Query(query, group.GroupName, pq.Array(group.Members), group.Gid)
		if err != nil {
			panic(err)
		}
		log.Info("Group updated", group.Gid)
	} else {
		return errors.New("group not found")
	}

	return nil
}

func DeleteGroup(gid int) error {

	exists, _ := CheckIdExists("groups", "gid", gid)
	if exists {
		query := "DELETE FROM groups WHERE gid = $1"

		_, err := db.Query(query, gid)
		if err != nil {
			panic(err)
		}
		log.Warn("Deleted group ", gid)
	} else {
		return errors.New("group not found")
	}
	return nil
}

func AddMembersToGroup(gid int, members []int) error {

	exists, _ := CheckIdExists("groups", "gid", gid)

	if exists {
		query := "UPDATE groups SET members = $1 where gid = $2"
		_, err := db.Exec(query, pq.Array(members), gid)
		if err != nil {
			panic(err)
		}
		log.Info("Added members %v", members)
	} else {
		return errors.New("group not found")
	}
	return nil

}

func DeleteMembersFromGroup(gid int, members []int) error {
	exists, _ := CheckIdExists("groups", "gid", gid)
	if exists {
		for _, member := range members {
			query := "UPDATE groups SET members = array_remove(members,$1) where gid = $2"
			_, err := db.Exec(query, member, gid)
			if err != nil {
				panic(err)
			}
			log.Warn("Deleted Member:", member, " From group: ", gid)
		}
	} else {
		return errors.New("group not found")
	}
	return nil
}

func SelectFromGroups(gid []string) []models.Group {
	var query string
	if len(gid) == 0 {
		query = "SELECT * FROM groups"
	} else {
		query = "SELECT * from groups where gid in (" + strings.Join(gid, ",") + ")"
	}
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	groups := []models.Group{}
	return groups
}
