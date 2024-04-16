package storage

import (
	"context"
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"iTask/modules/project_member_invited/model"
	"testing"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/sbank?sslmode=disable"
)

var testDB *sql.DB

func TestCreateProjectMemberInvited(t *testing.T) {
	var err error

	// load config and create a db connection
	//cfg, err := config.LoadConfig()
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.Mysql.User, cfg.Mysql.Password,
	//	cfg.Mysql.Host, cfg.Mysql.Port, cfg.Mysql.DBName)

	//Host: khanhtran-rds.c9seqcum6oc9.ap-southeast-1.rds.amazonaws.com
	//Port: 3306
	//User: root
	//Password: Trankhanh47
	//DBName: iTask
	dbHost := "khanhtran-rds.c9seqcum6oc9.ap-southeast-1.rds.amazonaws.com"
	dbPort := "3306"
	dbUser := "root"
	dbPassword := "Trankhanh47"
	dbName := "iTask"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
	t.Log(dsn)
	log := logger.Default.LogMode(logger.Info)

	testDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: log,
	})
	if err != nil {
		t.Errorf("cannot connect to db: %v", err)
	}

	store := NewSQLStore(testDB)
	dataTest := &model.ProjectMemberInvited{
		ProjectId:          999,
		VerificationMailId: 999,
	}

	err = store.CreateProjectMemberInvited(context.Background(), dataTest)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	t.Logf("Success")

	// clean up
	//err = testDB.Exec("DELETE FROM project_member_invited WHERE project_id = ?", 999).Error
	//if err != nil {
	//	t.Errorf("Error: %v", err)
	//}
}
