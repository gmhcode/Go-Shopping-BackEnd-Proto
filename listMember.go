package main

//ListMember - Connects the lists to the users, so we can know which user is a member of their respective lists
type ListMember struct {
	ListID string `json:"listID"`
	UserID string `json:"userID"`
}
