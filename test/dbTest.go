package test

import (
	"context"
	"fmt"

	pg "jwt-authentication/database"
	"jwt-authentication/models"
)

func PostgresDemo(ctx context.Context, dbRepo pg.Repository) {
	pg.RunDeleteTable(ctx, dbRepo)
	pg.RunMigrate(ctx, dbRepo)
	res1 := pg.RunAdd(ctx, dbRepo, models.User1)
	fmt.Println("Add result = ", res1)
	res2 := pg.RunGetAll(ctx, dbRepo)
	fmt.Println("Get All = ", res2)
}
