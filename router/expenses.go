package router

import (
	"github.com/gofiber/fiber/v3"
	"github.com/mitchellh/mapstructure"
	"splitwise-backend/handlers"
	"splitwise-backend/models"
)

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
	expenses.Post("/:id/settleUp", func(c fiber.Ctx) error {
		return nil
	})
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
			handlers.CreateExpense(expense)
		}

		return SuccessfulRequest(c, "Expense Created")

	})
}
func updateExpense(expenses fiber.Router) {
	expenses.Put("/:id", func(c fiber.Ctx) error {

		idInt, idErr := CheckId(c)
		if idErr != nil {
			return idErr
		}

		var updatedExpense map[string]interface{}

		if err := c.Bind().Body(&updatedExpense); err != nil {
			return err
		}
		var expense models.Expense
		expense.Eid = idInt
		if err := mapstructure.Decode(updatedExpense, &expense); err != nil {
			return err
		}

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

}
