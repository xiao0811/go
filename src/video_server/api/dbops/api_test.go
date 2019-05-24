package dbops

import (
	"fmt"
	"testing"
)

func celarTables()  {
	dbConn.Exec("truncate user")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

func TestMain(m *testing.M) {
	celarTables()
	m.Run()
	celarTables()
}

func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", testAddUser)
	t.Run("Get", testGetUser)
	t.Run("Add", testDeleteUser)
	t.Run("Add", testRegetUser)
}

func testAddUser(t testing.T)  {
	err := AddUserCredential("xiaosha", "123456")

	if err != nil {
		t.Errorf("Error of AddUser: %v", err)
	}
}

func testGetUser(t testing.T)  {
	pwd, err := GetUserCredential("xiaosha")
	if err != nil {
		t.Errorf("Error of GetUser:%v", err)
	}
	fmt.Println(pwd)
}

func testDeleteUser(t *testing.T) {
	err := DeleteUser("xiaosha", "123456")
	if err != nil {
		t.Errorf("Error of DeleteUser: %v", err)
	}
}

func testRegetUser(t testing.T)  {
	pwd, err := GetUserCredential("xiaosha")
	if err != nil {
		t.Errorf("Error of RegerUser: %v", err)
	}

	if pwd != "" {
		t.Errorf("Deleting user test failed")
	}
}
