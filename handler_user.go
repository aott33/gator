package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/aott33/gator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("username is required")
	}
	
	name := cmd.args[0]

	_, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		return err
	}

	err = s.cfg.SetUser(name)
	if err != nil {
		return err
	}

	fmt.Printf("Set user to %s\n",cmd.args[0])
	
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("name is required")
	}

	name := cmd.args[0]

	_, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:			uuid.New(),
		CreatedAt: 	time.Now(),
		UpdatedAt: 	time.Now(),
		Name:		name,
	})
	if err != nil {
		return err
	}

	err = s.cfg.SetUser(name)
	if err != nil {
		return err
	}

	fmt.Printf("User: %s created in database\n", name)

	return nil
}

func handlerGetUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		fmt.Println("Unable to get users")
		return err
	}

	for i := range users {
		if s.cfg.CurrentUserName == users[i] {
			fmt.Printf("* %s (current)\n",users[i])
		} else {
			fmt.Printf("* %s\n",users[i])
		}
	}

	return nil
}
