package main

import (
	"fmt"
	"log"

	bullhorn "github.com/intiqo/bullhorn-go-sdk"
)

func main() {
	// Initialize Bullhorn client
	params := &bullhorn.AuthParams{
		ClientId:          "your-client-id",
		ClientSecret:      "your-client-secret", 
		Username:          "your-username",
		Password:          "your-password",
		AuthenticationUrl: "https://auth.bullhornstaffing.com/oauth/authorize",
		LoginUrl:          "https://rest.bullhornstaffing.com/rest-services/login",
		RestTokenTTL:      "60",
	}

	client, err := bullhorn.NewClient(params)
	if err != nil {
		log.Fatalf("Failed to create Bullhorn client: %v", err)
	}

	candidateId := 3084

	// 1. GET - Retrieve to-many associations (e.g., get skills for a candidate)
	fmt.Println("=== Getting To-Many Associations ===")
	options := bullhorn.QueryOptions{
		Fields: []string{"id", "name", "skillLevel"},
		Count:  10,
		Start:  0,
	}
	
	response, skills, err := client.GetToManyAssociations(
		bullhorn.CandidateEntity,
		candidateId,
		"primarySkills",
		options,
	)
	if err != nil {
		log.Printf("Failed to get skills: %v", err)
	} else {
		fmt.Printf("Found skills for candidate %d. HTTP Status: %d\n", candidateId, response.StatusCode())
		fmt.Printf("Skills data: %v\n", skills)
	}

	// 2. PUT - Add new associations (e.g., add skills to a candidate)
	fmt.Println("\n=== Adding To-Many Associations ===")
	skillIdsToAdd := []string{"964", "684", "253"}
	
	response, result, err := client.AssociateEntities(
		bullhorn.CandidateEntity,
		candidateId,
		"primarySkills",
		skillIdsToAdd,
	)
	if err != nil {
		log.Printf("Failed to add skills: %v", err)
	} else {
		fmt.Printf("Successfully added skills. HTTP Status: %d\n", response.StatusCode())
		fmt.Printf("Change Type: %s, Changed Entity ID: %d\n", result.ChangeType, result.ChangedEntityId)
	}


	// Example with JobOrder submissions
	fmt.Println("\n=== JobOrder Submissions Example ===")
	jobOrderId := 1234
	
	// Get submissions for a job order
	options = bullhorn.QueryOptions{
		Fields: []string{"id", "candidate", "status", "dateAdded"},
		Count:  20,
	}
	
	response, submissions, err := client.GetToManyAssociations(
		bullhorn.JobOrderEntity,
		jobOrderId,
		"submissions",
		options,
	)
	if err != nil {
		log.Printf("Failed to get job submissions: %v", err)
	} else {
		fmt.Printf("Found submissions for job order %d. HTTP Status: %d\n", jobOrderId, response.StatusCode())
		fmt.Printf("Submissions data: %v\n", submissions)
	}
}