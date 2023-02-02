package database

import (
	"context"
	"fmt"
	"jwt-authentication/models"
	"log"
)

func RunMigrate(ctx context.Context, dbRepo Repository) {
	fmt.Println("1. Migrate running")
	if err := dbRepo.Migrate(ctx); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Migrate done!")
}

func RunDeleteTable(ctx context.Context, dbRepo Repository) {
	fmt.Println("2. DeleteTable running")
	if err := dbRepo.DeleteTable(ctx); err != nil {
		log.Fatal(err)
	}
	fmt.Println("DeteleTable done!")
}

func RunAdd(ctx context.Context, dbRepo Repository, user models.User) models.User {
	fmt.Println("2. Add running")
	//fmt.Println(user)
	res, err := dbRepo.Add(ctx, user)
	if err != nil {
		log.Println("err = ", err)
	}
	return *res
}

func RunGetAll(ctx context.Context, dbRepo Repository) []models.User {
	fmt.Println("3. GetAll running")
	res, err := dbRepo.GetAll(ctx)
	if err != nil {
		log.Println("err = ", err)
	}
	return res
}
