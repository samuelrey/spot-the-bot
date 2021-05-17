package cmd

import "fmt"

func Leave() {
	fmt.Println("leave")
	// TODO these should come from an app Context
	userIDs := []string{}
	userID := "ex"

	found := -1
	for i, id := range userIDs {
		if id == userID {
			found = i
			break
		}
	}
	if found != -1 {
		userIDs = append(userIDs[:found], userIDs[found+1:]...)
		fmt.Println("remove user: ", userID)
	}
}
