package tests

//
//import (
//	"login-service/internal/domain/model"
//	"login-service/internal/infrastructure"
//	"os"
//	"testing"
//)
//
//func TestConnectDB(t *testing.T) {
//
//	// Set the required environment variables for the test
//	os.Setenv("HOST_DB", "contadoraccess.cmpk7hpbdidk.sa-east-1.rds.amazonaws.com")
//	os.Setenv("USER_DB", "postgres")
//	os.Setenv("PASSWORD_DB", "SLsSTTrL7YWE5Cn")
//	os.Setenv("NAME_DB", "freebrain")
//	os.Setenv("PORT_DB", "5432")
//
//	// Call the ConnectDB function
//	db, err := infrastructure.ConnectDB()
//	if err != nil {
//		t.Fatalf("failed to connect to database: %v", err)
//	}
//
//	// Check that the DB instance is not nil
//	if db == nil {
//		t.Fatal("DB instance is nil")
//	}
//
//	//Check that we can perform a database query on the DB instance
//	var count int64
//	err = db.Model(&model.User{}).Count(&count).Error
//	if err != nil {
//		t.Fatalf("failed to perform database query: %v", err)
//	}
//}
