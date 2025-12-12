package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.Reset(context.Background())
	if err != nil {
		fmt.Println("Unsuccessful!")
		return err
	}

	fmt.Println("Successful: users table reset")
	return nil
}
