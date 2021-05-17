package cmd

import "fmt"

func Join() {
	fmt.Println("join")
	// TODO these should come from an app Context
	userIDs := []string{}
	userID := "ex"

	enrolled := false
	for _, id := range userIDs {
		if id == userID {
			enrolled = true
			break
		}
	}
	if !enrolled {
		userIDs = append(userIDs, userID)
		fmt.Println("new user: ", userID)
	}
}
