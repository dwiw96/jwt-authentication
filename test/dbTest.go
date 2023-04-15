package test

import (
	"context"
	"fmt"

	"jwt-authentication/database"
	"jwt-authentication/models"
)

func PostgresDemo(ctx context.Context, dbRepo database.Repository) {
	database.RunDeleteTable(ctx, dbRepo)
	database.RunMigrate(ctx, dbRepo)
	res1 := database.RunAdd(ctx, dbRepo, models.User1)
	fmt.Println("Add result = ", res1)
	res2 := database.RunGetAll(ctx, dbRepo)
	fmt.Println("Get All = ", res2)
}
