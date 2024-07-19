package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"splitwise-backend/constants"
	"splitwise-backend/database"
	"splitwise-backend/router"
	"testing"
)

// group tests
//1. create group with fake users
//2. add members to group
//3. add expenses in group

func setupRouter() (*fiber.App, error) {
	app := fiber.New()
	err := godotenv.Load("../.env")
	if err != nil {
		return nil, err
	}
	constants.InitVars()
	err = database.InitDb()
	if err != nil {
		return nil, err
	}
	router.SetupUserRoutes(app)
	router.SetupGroupRoutes(app)
	router.SetupExpenseRoutes(app)

	return app, nil
}

func runTests(t *testing.T) {
	app, err := setupRouter()
	if err != nil {
		t.Fatalf("Failed to setup router: %v", err)
	}
	testGetAllUsers(app, t)
	//testCreateUsers(app, t)
	testCreateGroup(app, t)

}

func testGetAllUsers(app *fiber.App, t *testing.T) {
	req, _ := http.NewRequest("GET", "/user", nil)
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	fmt.Println(resp.Body)
}
func testCreateUsers(app *fiber.App, t *testing.T) {
	users := router.CreateFakeUsers()
	for _, user := range users {
		body, _ := json.Marshal(user)
		req, _ := http.NewRequest("POST", "/user", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
	}
}
func testCreateGroup(app *fiber.App, t *testing.T) {
	groups := router.CreateFakeGroups()
	for _, group := range groups {
		body, _ := json.Marshal(group)
		req, _ := http.NewRequest("POST", "/group", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		readResponse(resp)

	}
}
func readResponse(resp *http.Response) {
	// Read and print the response body
	respBody, _ := io.ReadAll(resp.Body)

	var response map[string]interface{}
	_ = json.Unmarshal(respBody, &response)

	// Print the message
	fmt.Println("Response message:", response["message"])
}

func TestCreateUserGroups(t *testing.T) {
	runTests(t)
}
