package router

import (
	"github.com/gofiber/fiber/v3"
	"splitwise-backend/handlers"
	"splitwise-backend/models"
)

// after create expense, we should update members and their balance.
func SetupExpenseRoutes(app *fiber.App) {
	expenses := app.Group("/expense")

	// get expense
	getExpenseDetails(expenses)
	// create new expense
	createNewExpense(expenses)
	// update expense with expense id
	updateExpense(expenses)
	// delete expense with expense id
	deleteExpense(expenses)
	// settleUp transaction
	settleUp(expenses)
}

func getExpenseDetails(expenses fiber.Router) {
	expenses.Get("/", func(c fiber.Ctx) error {
		handlers.GetAllExpense()
		return nil
	})
}

func createNewExpense(expenses fiber.Router) {
	expenses.Post("/", func(c fiber.Ctx) error {
		expenseArr := createFakeExpenses()
		for _, expense := range expenseArr {
			err := handlers.CreateExpense(expense)
			if err != nil {
				return err
			}
		}

		return SuccessfulRequest(c, "Expense Created")

	})
}
func updateExpense(expenses fiber.Router) {
	expenses.Patch("/:id", func(c fiber.Ctx) error {

		idInt, idErr := CheckId(c)
		if idErr != nil {
			return idErr
		}

		//var updatedExpense map[string]interface{}
		var expense models.Expense

		if err := c.Bind().Body(&expense); err != nil {
			return err
		}
		expense.Eid = idInt

		//if err := mapstructure.Decode(updatedExpense, &expense); err != nil {
		//	return err
		//}

		if err := handlers.UpdateExpense(expense); err != nil {
			return InternalError(c, err)
		}
		return SuccessfulRequest(c, "Expense Updated")

	})

}

func deleteExpense(expenses fiber.Router) {
	expenses.Delete("/:id", func(c fiber.Ctx) error {
		idInt, idErr := CheckId(c)
		if idErr != nil {
			return idErr
		}
		if err := handlers.DeleteExpense(idInt); err != nil {
			return InternalError(c, err)
		}
		return SuccessfulRequest(c, "Expense Deleted")

	})

}

func settleUp(expenses fiber.Router) {
	expenses.Post("/:id/settleUp", func(c fiber.Ctx) error {
		// settle expense between n members
		// expense.issettled = true
		// update user balance, owes,owed
		idInt, idErr := CheckId(c)
		if idErr != nil {
			return idErr
		}

		var expense models.Expense
		expense.Eid = idInt
		expense.IsSettled = true
		if err := handlers.SettleExpense(expense); err != nil {
			return InternalError(c, err)
		}
		return SuccessfulRequest(c, "Expense Settled")

		return nil
	})

}
