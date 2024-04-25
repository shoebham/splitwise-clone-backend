package router

import (
	"github.com/gofiber/fiber/v3"
	"splitwise-backend/handlers"
	"splitwise-backend/models"
)

func SetupGroupRoutes(app *fiber.App) {
	groups := app.Group("/group")
	getGroupDetails(groups)
	createNewGroup(app, groups)
	updateGroup(app, groups)
	deleteGroup(groups)
	addGroupMember(groups)
	deleteGroupMember(groups)
	updateTransactions(groups)

}

func updateTransactions(groups fiber.Router) fiber.Router {
	return groups.Post("/:id/updateTransactions", func(c fiber.Ctx) error {
		return nil
	})
}

func deleteGroupMember(groups fiber.Router) {
	type DeleteMemberRequest struct {
		Members []int `json:"members"`
	}
	// delete group member with member id
	groups.Delete("/:id/member", func(c fiber.Ctx) error {
		idInt, idErr := CheckId(c)
		if idErr != nil {
			return idErr
		}

		var deleteMemberRequest DeleteMemberRequest
		var groupMemberList []int
		if err := c.Bind().Body(&deleteMemberRequest); err != nil {
			return InternalError(c, err)
		}
		groupMemberList = deleteMemberRequest.Members
		if err := handlers.DeleteMembersFromGroup(idInt, groupMemberList); err != nil {
			return InternalError(c, err)
		}
		return SuccessfulRequest(c, "Group Member Deleted")

	})
}

func addGroupMember(groups fiber.Router) {
	// expected body { members: ["id1","id2","id3","id4"]}
	// add group member and get member id in return

	type AddMemberRequest struct {
		Members []int `json:"members"`
	}
	groups.Put("/:id/member", func(c fiber.Ctx) error {

		var addMemberRequest AddMemberRequest
		idInt, idErr := CheckId(c)
		if idErr != nil {
			return idErr
		}
		var groupMemberList []int
		if err := c.Bind().Body(&addMemberRequest); err != nil {
			return InternalError(c, err)
		}
		groupMemberList = addMemberRequest.Members
		if err := handlers.AddMembersToGroup(idInt, groupMemberList); err != nil {
			return InternalError(c, err)
		}
		return SuccessfulRequest(c, "Group Member Added")
	})
}

func deleteGroup(groups fiber.Router) {
	// delete group with group id
	groups.Delete("/:id", func(c fiber.Ctx) error {

		idInt, idErr := CheckId(c)
		if idErr != nil {
			return idErr
		}

		if err := handlers.DeleteGroup(idInt); err != nil {

			return InternalError(c, err)
		}

		return SuccessfulRequest(c, "Group Deleted")
	})
}

func updateGroup(app *fiber.App, groups fiber.Router) {
	// update group with group id
	groups.Put("/:id", func(c fiber.Ctx) error {

		_, idErr := CheckId(c)
		if idErr != nil {
			return idErr
		}
		var updatedGroup models.Group
		if err := c.Bind().Body(&updatedGroup); err != nil {
			return err
		}
		if err := handlers.UpdateGroup(updatedGroup); err != nil {
			return InternalError(c, err)
		}

		return SuccessfulRequest(c, "Group Updated")

	})
}

func createNewGroup(app *fiber.App, groups fiber.Router) {
	// create new group
	groups.Post("/", func(c fiber.Ctx) error {
		groupsArr := createFakeGroups()
		for _, group := range groupsArr {
			handlers.CreateGroup(app, group)
		}

		return SuccessfulRequest(c, "Group Created")
	})
}

func getGroupDetails(groups fiber.Router) {
	// get group details
	groups.Get("/", func(c fiber.Ctx) error {
		handlers.GetAllGroups(c.App())
		return nil
	})
}
