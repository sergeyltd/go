package main

import (
	"bufio"
	"os"
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
}

func main() {
	// u := User{
	// 	FirstName: "Tom",
	// 	LastName:  "Riddle",
	// }

	// AddNewUser(&u)

	println("Read users from files...")
	users,_ := FileToListOfUsers("data.csv")

	println("Send users to server")
	for _,u := range users{
		AddNewUser(&u)
	}

	println("Client successfully sent all users")
}

func FileToListOfUsers(path string) ([]User, error) {
	lines, err := readLines(path)
	if err != nil {
		return nil, err
	}

	var users []User
	for _,v := range lines{
		spl := strings.Split(v, ",")
		u := User{
			FirstName: spl[0],
			LastName: spl[1],
		}
		users = append(users, u)
	}

	return users, nil
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func AddNewUser(user *User) {
	jsonByteArr, _ := json.Marshal(user)
	jsonStr := string(jsonByteArr)
	jsonStrBytes := []byte(jsonStr)
	buf := bytes.NewBuffer(jsonStrBytes)
	resp, err := http.Post("http://localhost:3000/users", "application/json", buf)
	if err != nil {
		println(err, resp.StatusCode)
		panic("Error on stage to add new user")
	}
}
