package database

import (
	"context"
	"fmt"
	"jwt-authentication/models"
	"log"
)

func RunMigrate(ctx context.Context, dbRepo Repository) error {
	fmt.Println("1. Migrate running")
	if err := dbRepo.Migrate(ctx); err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println("Migrate done!")
	return nil
}

func RunDeleteTable(ctx context.Context, dbRepo Repository) error {
	fmt.Println("2. DeleteTable running")
	if err := dbRepo.DeleteTable(ctx); err != nil {
		log.Println(err)
		return ErrTableNotExist
	}
	fmt.Println("DeteleTable done!")
	return nil
}

func RunAdd(ctx context.Context, dbRepo Repository, user models.User) (*models.User, error) {
	fmt.Println("2. Add running")
	//fmt.Println(user)
	res, err := dbRepo.Add(ctx, user)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return res, nil
}

func RunDelete(ctx context.Context, dbRepo Repository, username string) error {
	fmt.Println("3. Delete running")
	err := dbRepo.Delete(ctx, username)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func RunGetAll(ctx context.Context, dbRepo Repository) []models.User {
	fmt.Println("4. GetAll running")
	res, err := dbRepo.GetAll(ctx)
	if err != nil {
		log.Println("err = ", err)
	}
	return res
}
