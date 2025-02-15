package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type SkillSet struct {
	LanguagesKnown []string
	Passion        []string
}

type UserData struct {
	UserId         int
	UserName       string
	IsActive       bool
	Roles          []string
	TrainingCredit float32 `json:"training-credit"` // json file contains different key(training-credit), hence we map to TrainingCredit
	Skills         SkillSet
}

func testMe(num1, num2 int) int {
	return num1 + num2
}

func main() {
	var userInfo []UserData
	fileByte, err := os.ReadFile("resources/test.json")
	if err != nil {
		fmt.Printf("Error occured while reading hr json. Err:%s", err)
	}
	err = json.Unmarshal(fileByte, &userInfo)
	if err != nil {
		fmt.Printf("Error occured while unmarshalling hr json. Err:%s", err)
	}
	// for loop to iterate over the records
	for _, value := range userInfo {
		fmt.Printf("------Employee Details: %s------\n", value.UserName)
		fmt.Println("User ID:", value.UserId)
		fmt.Println("Is Active:", value.IsActive)
		fmt.Println("Available training credit", value.TrainingCredit)
		fmt.Println("Languages known:", value.Skills.LanguagesKnown)
		fmt.Println("Passionate about:", value.Skills.Passion)
	}
}
