package taskmanager_test

import (
	"testing"
)

func TestTaskManagerRepository(t *testing.T) {
	// ctx := context.Background()

	// pgContainer, err := postgres.RunContainer(ctx,
	// 	testcontainers.WithImage("postgres:15.3-alpine"),
	// 	postgres.WithInitScripts(filepath.Join("..", "testdata", "init-db.sql")),
	// 	postgres.WithDatabase("test-db"),
	// 	postgres.WithUsername("postgres"),
	// 	postgres.WithPassword("postgres"),
	// 	testcontainers.WithWaitStrategy(
	// 		wait.ForLog("database system is ready to accept connections").
	// 			WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	// )
	// if err != nil {
	// 	t.Fatal(err)
	// }

	// t.Cleanup(func() {
	// 	if err := pgContainer.Terminate(ctx); err != nil {
	// 		t.Fatalf("failed to terminate pgContainer: %s", err)
	// 	}
	// })

	// connStr, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	// assert.NoError(t, err)
	// db := database.NewPostgresClient()
	// taskRepo, err := taskmanager.NewPostgresRepository(db)
	// assert.NoError(t, err)

	// c, err := taskRepo.CreateCustomer(ctx, taskmanager_gen.Task{
	// 	Name: "Henry",
	// })
	// assert.NoError(t, err)
	// assert.NotNil(t, c)

	// customer, err := taskRepo.GetCustomerByEmail(ctx, "henry@gmail.com")
	// assert.NoError(t, err)
	// assert.NotNil(t, customer)
	// assert.Equal(t, "Henry", customer.Name)
	// assert.Equal(t, "henry@gmail.com", customer.Email)
}
