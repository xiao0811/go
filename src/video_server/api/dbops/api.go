package dbops

import (
	_ "github.com/Go-SQL-Driver/MySQL"
	"log"
)

func AddUserCredential(loginName, pwd string) error {
	stmtIns, err := dbConn.Prepare("INSERT INTO users (login_name, pwd) VALUE (?, ?)")
	if err != nil {
		return err
	}

	_, _ = stmtIns.Exec(loginName, pwd)
	_ = stmtIns.Close()
	return nil
}

func GetUserCredential(loginName string) (string, error) {
	stmtOut, err := dbConn.Prepare("SELECT pwd FROM user WHERE login_name = ?")
	if err != nil {
		log.Printf("%s", err)
		return "", err
	}

	var pwd string
	_ = stmtOut.QueryRow(loginName).Scan(&pwd)
	_ = stmtOut.Close()

	return pwd, nil
}

func DeleteUser(loginName, pwd string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM user WHERE login_name = ? AND pwd = ?")
	if err != nil {
		log.Println(err)
		return err
	}

	stmtDel.Exec(loginName, pwd)
	stmtDel.Close()
	return nil
}
