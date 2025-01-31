package belajargolanggorm

import (
	"belajar-golang-gorm/helpers"
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var db = helpers.OpenConnection()

func TestOpenConnection(t *testing.T) {
	assert.NotNil(t, db)
}

func TestExecuteSQL(t *testing.T) {
	err := db.Exec("insert into sample(id,name) values(?, ?)", "1", "Eko").Error
	assert.Nil(t, err)

	err = db.Exec("insert into sample(id,name) values(?, ?)", "2", "Budi").Error
	assert.Nil(t, err)

	err = db.Exec("insert into sample(id,name) values(?, ?)", "3", "Joko").Error
	assert.Nil(t, err)

	err = db.Exec("insert into sample(id,name) values(?, ?)", "4", "Rully").Error
	assert.Nil(t, err)
}

type Sample struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func TestRawSQL(t *testing.T) {
	var sample Sample
	err := db.Raw("select id, name from sample where id = ?", "1").Scan(&sample).Error
	assert.Nil(t, err)
	assert.Equal(t, "1", sample.Id)
	assert.Equal(t, "Eko", sample.Name)

	var samples []Sample
	err = db.Raw("select id, name from sample").Scan(&samples).Error
	assert.Nil(t, err)
	assert.Equal(t, 4, len(samples))
}

func TestSqlRow(t *testing.T) {
	var samples []Sample

	rows, err := db.Raw("select id, name from sample").Rows()
	assert.Nil(t, err)
	defer rows.Close()

	for rows.Next() {
		var id string
		var name string
		err := rows.Scan(&id, &name)
		assert.Nil(t, err)

		samples = append(samples, Sample{
			Id:   id,
			Name: name,
		})
	}

	assert.Equal(t, 4, len(samples))
}

func TestScanRows(t *testing.T) {
	var samples []Sample

	rows, err := db.Raw("select id, name from sample").Rows()
	assert.Nil(t, err)
	defer rows.Close()

	for rows.Next() {
		err := db.ScanRows(rows, &samples)
		assert.Nil(t, err)
	}

	assert.Equal(t, 4, len(samples))
}

func TestCreateUser(t *testing.T) {
	user := User{
		Id:       "1",
		Password: "rahasia",
		Name: Name{
			FirstName:  "Manuel",
			MiddleName: "Theodore",
			LastName:   "Leleuly",
		},
		Information: "Ini akan di ignore",
	}

	response := db.Create(&user)
	assert.Nil(t, response.Error)
	assert.Equal(t, 1, int(response.RowsAffected))
}

func TestBatchInsert(t *testing.T) {
	var users []User
	for i := 2; i < 10; i++ {
		users = append(users, User{
			Id: strconv.Itoa(i),
			Name: Name{
				FirstName: "User " + strconv.Itoa(i),
			},
			Password: "rahasia_" + strconv.Itoa(i),
		})
	}

	result := db.Create(&users)
	assert.Nil(t, result.Error)
	assert.Equal(t, 8, int(result.RowsAffected))
}

func TestTransactionSuccess(t *testing.T) {
	err := db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&User{
			Id:       "10",
			Password: "rahasia_10",
			Name: Name{
				FirstName: "User 10",
			},
		}).Error
		if err != nil {
			return err
		}

		err = tx.Create(&User{
			Id:       "11",
			Password: "rahasia_11",
			Name: Name{
				FirstName: "User 11",
			},
		}).Error
		if err != nil {
			return err
		}

		err = tx.Create(&User{
			Id:       "12",
			Password: "rahasia_12",
			Name: Name{
				FirstName: "User 12",
			},
		}).Error
		if err != nil {
			return err
		}

		return nil
	})

	assert.Nil(t, err)
}

func TestTransactionError(t *testing.T) {
	err := db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&User{
			Id:       "13",
			Password: "rahasia_13",
			Name: Name{
				FirstName: "User 13",
			},
		}).Error
		if err != nil {
			return err
		}

		err = tx.Create(&User{
			Id:       "11",
			Password: "rahasia_11",
			Name: Name{
				FirstName: "User 11",
			},
		}).Error
		if err != nil {
			return err
		}

		return nil
	})

	assert.NotNil(t, err)
}

func TestManualTransactionSuccess(t *testing.T) {
	tx := db.Begin()
	defer tx.Rollback()

	err := tx.Create(&User{
		Id:       "13",
		Password: "rahasia_13",
		Name: Name{
			FirstName: "User 13",
		},
	}).Error
	assert.Nil(t, err)

	err = tx.Create(&User{
		Id:       "14",
		Password: "rahasia_14",
		Name: Name{
			FirstName: "User 14",
		},
	}).Error
	assert.Nil(t, err)

	if err == nil {
		tx.Commit()
	}
}

func TestManualTransactionError(t *testing.T) {
	tx := db.Begin()
	defer tx.Rollback()

	err := tx.Create(&User{
		Id:       "15",
		Password: "rahasia_15",
		Name: Name{
			FirstName: "User 15",
		},
	}).Error
	assert.Nil(t, err)

	err = tx.Create(&User{
		Id:       "14",
		Password: "rahasia_14",
		Name: Name{
			FirstName: "User 14",
		},
	}).Error
	assert.NotNil(t, err)

	if err == nil {
		tx.Commit()
	}
}

func TestQuerySingleObject(t *testing.T) {
	user := User{}
	result := db.First(&user)
	assert.Nil(t, result.Error)
	assert.Equal(t, "1", user.Id)

	user = User{}
	result = db.Last(&user)
	assert.Nil(t, result.Error)
	assert.Equal(t, "9", user.Id)
}

func TestQuerySingleObjectInlineCondition(t *testing.T) {
	user := User{}
	result := db.Take(&user, "id = ?", "5")
	assert.Nil(t, result.Error)
	assert.Equal(t, "5", user.Id)
}

func TestQueryAllObjects(t *testing.T) {
	var users []User
	result := db.Find(&users, "id in ?", []string{"1", "2", "3", "4"})
	for _, user := range users {
		fmt.Println(user)
	}
	assert.Nil(t, result.Error)
	assert.Equal(t, 4, len(users))
}

func TestAndOperator(t *testing.T) {
	var users []User
	result := db.Where("first_name like ?", "%User%").Where("password like ?", "%rahasia%").Find(&users)
	assert.Nil(t, result.Error)

	assert.Equal(t, 13, len(users))
}

func TestOrOperator(t *testing.T) {
	var users []User
	result := db.Where("first_name like ?", "%User%").Or("password = ?", "rahasia").Find(&users)
	assert.Nil(t, result.Error)

	assert.Equal(t, 14, len(users))
}

func TestNotOperator(t *testing.T) {
	var users []User
	result := db.Not("first_name like ?", "%User%").Where("password like ?", "%rahasia%").Find(&users)
	assert.Nil(t, result.Error)

	assert.Equal(t, 1, len(users))
}

func TestSelectFields(t *testing.T) {
	var users []User
	result := db.Select("id", "first_name").Find(&users)
	assert.Nil(t, result.Error)

	for _, user := range users {
		assert.NotNil(t, user.Id)
		assert.NotEqual(t, "", user.Name.FirstName)
	}

	assert.Equal(t, 14, len(users))
}

func TestStructCondition(t *testing.T) {
	userCondition := User{
		Name: Name{
			FirstName: "User 5",
			LastName:  "", // tidak bisa, karena dianggap default value
		},
	}

	var users []User
	result := db.Where(userCondition).Find(&users)
	assert.Nil(t, result.Error)
	assert.Equal(t, 1, len(users))
}

func TestMapCondition(t *testing.T) {
	mapCondition := map[string]any{
		"middle_name": "",
	}

	var users []User
	result := db.Where(mapCondition).Find(&users)
	assert.Nil(t, result.Error)
	assert.Equal(t, 13, len(users))
}

func TestOrderLimitOffset(t *testing.T) {
	var users []User
	result := db.Order("id asc, first_name asc").Limit(5).Offset(5).Find(&users)
	assert.Nil(t, result.Error)
	assert.Equal(t, 5, len(users))
	assert.Equal(t, "14", users[0].Id)
}

type UserResponse struct {
	Id        string
	FirstName string
	LastName  string
}

func TestQueryNonModel(t *testing.T) {
	var users []UserResponse
	result := db.Model(&User{}).Select("id", "first_name", "last_name").Find(&users)
	assert.Nil(t, result.Error)
	assert.Equal(t, 14, len(users))
}

func TestUpdate(t *testing.T) {
	user := User{}
	result := db.First(&user, "id = ?", "1")
	assert.Nil(t, result.Error)

	user.Name.FirstName = "Manuel"
	user.Name.MiddleName = ""
	user.Name.LastName = "Leleuly"
	user.Password = "password123"
	result = db.Save(&user)
	assert.Nil(t, result.Error)

	updatedUser := User{}
	result = db.Where("id = ?", "1").Find(&updatedUser)
	assert.Nil(t, result.Error)
	assert.Equal(t, user.Id, updatedUser.Id)
	assert.Equal(t, user.Name.FirstName, updatedUser.Name.FirstName)
	assert.Equal(t, user.Name.MiddleName, updatedUser.Name.MiddleName)
	assert.Equal(t, user.Name.LastName, updatedUser.Name.LastName)
	assert.Equal(t, user.Password, updatedUser.Password)
}

func TestSelectedColumns(t *testing.T) {
	result := db.Model(&User{}).Where("id = ?", "1").Updates(map[string]any{
		"middle_name": "",
		"last_name":   "Morro",
	})
	assert.Nil(t, result.Error)

	result = db.Model(&User{}).Where("id = ?", "1").Update("password", "diubahlagi")
	assert.Nil(t, result.Error)

	result = db.Where("id = ?", "1").Updates(User{
		Name: Name{
			FirstName: "Eko",
			LastName:  "Khannedy",
		},
	})
	assert.Nil(t, result.Error)
}

func TestAutoIncrement(t *testing.T) {
	for i := 0; i < 10; i++ {
		userLog := UserLog{
			UserId: "1",
			Action: "Test Action" + strconv.Itoa(i),
		}
		result := db.Create(&userLog)
		assert.Nil(t, result.Error)
		assert.NotEqual(t, 0, userLog.Id)
		fmt.Println(userLog.Id)
	}
}

func TestSaveOrUpdate(t *testing.T) {
	userLog := UserLog{
		UserId: "1",
		Action: "Test Action",
	}
	result := db.Save(&userLog) // create
	assert.Nil(t, result.Error)

	userLog.UserId = "2"
	result = db.Save(&userLog) // update
	assert.Nil(t, result.Error)
}

func TestSaveOrUpdateNonAutoIncrement(t *testing.T) {
	user := User{
		Id: "99",
		Name: Name{
			FirstName: "User 99",
		},
	}
	result := db.Save(&user) // create
	assert.Nil(t, result.Error)

	user.Name.FirstName = "User 99 Updated"
	result = db.Save(&user) // update
	assert.Nil(t, result.Error)
}

func TestConflict(t *testing.T) {
	user := User{
		Id: "88",
		Name: Name{
			FirstName: "User 88",
		},
	}
	result := db.Clauses(clause.OnConflict{UpdateAll: true}).Create(&user) // create
	assert.Nil(t, result.Error)
}

func TestDelete(t *testing.T) {
	var user User
	result := db.Take(&user, "id = ?", "88")
	assert.Nil(t, result.Error)
	result = db.Delete(&user)
	assert.Nil(t, result.Error)

	result = db.Delete(&User{}, "id = ?", "99")
	assert.Nil(t, result.Error)

	result = db.Where("id = ?", "77").Delete(&User{})
	assert.Nil(t, result.Error)
}

func TestSoftDelete(t *testing.T) {
	todo := Todo{
		UserId:      "1",
		Title:       "Todo 1",
		Description: "Isi todo 1",
	}
	result := db.Create(&todo)
	assert.Nil(t, result.Error)

	result = db.Delete(&todo)
	assert.Nil(t, result.Error)
	assert.NotNil(t, todo.DeletedAt)

	var todos []Todo
	result = db.Find(&todos)
	assert.Nil(t, result.Error)
	assert.Equal(t, 0, len(todos))
}

func TestUnscoped(t *testing.T) {
	var todo Todo
	result := db.Unscoped().First(&todo, "id = ?", "1")
	assert.Nil(t, result.Error)

	result = db.Unscoped().Delete(&todo)
	assert.Nil(t, result.Error)

	var todos []Todo
	result = db.Unscoped().Find(&todos)
	assert.Nil(t, result.Error)
	assert.Equal(t, 0, len(todos))
}

func TestLock(t *testing.T) {
	err := db.Transaction(func(tx *gorm.DB) error {
		var user User
		err := tx.Clauses(clause.Locking{Strength: clause.LockingStrengthUpdate}).Take(&user, "id = ?", "2").Error
		if err != nil {
			return err
		}

		user.Name.FirstName = "Joko"
		user.Name.LastName = "Morro"
		return tx.Save(&user).Error
	})

	assert.Nil(t, err)
}

func TestCreateWallet(t *testing.T) {
	wallet := Wallet{
		Id:      "1",
		UserId:  "1",
		Balance: 1000000,
	}
	result := db.Create(&wallet)
	assert.Nil(t, result.Error)
}

func TestRetrieveRelation(t *testing.T) {
	var user User
	err := db.Model(&User{}).Preload("Wallet").Take(&user, "id = ?", "1").Error
	assert.Nil(t, err)

	fmt.Println(user)
	assert.Equal(t, "1", user.Id)
	assert.Equal(t, "1", user.Wallet.Id)
}

func TestRetrieveRelationJoin(t *testing.T) {
	var user User
	err := db.Model(&User{}).Joins("Wallet").Take(&user, "users.id = ?", "1").Error
	assert.Nil(t, err)

	fmt.Println(user)
	assert.Equal(t, "1", user.Id)
	assert.Equal(t, "1", user.Wallet.Id)
}

func TestAutoCreateUpdate(t *testing.T) {
	user := User{
		Id:       "20",
		Password: "rahasia",
		Name: Name{
			FirstName: "User 20",
		},
		Wallet: Wallet{
			Id:      "20",
			UserId:  "20",
			Balance: 1000000,
		},
	}
	err := db.Create(&user).Error
	assert.Nil(t, err)
}

func TestSkipCreateUpdate(t *testing.T) {
	user := User{
		Id:       "21",
		Password: "rahasia",
		Name: Name{
			FirstName: "User 21",
		},
		Wallet: Wallet{
			Id:      "21",
			UserId:  "21",
			Balance: 1000000,
		},
	}
	err := db.Omit(clause.Associations).Create(&user).Error
	assert.Nil(t, err)
}

func TestUserAndAddresses(t *testing.T) {
	user := User{
		Id:       "2",
		Password: "rahasia",
		Name:     Name{FirstName: "User 50"},
		Wallet: Wallet{
			Id:      "2",
			UserId:  "2",
			Balance: 1000000,
		},
		Addresses: []Address{
			{
				UserId:  "2",
				Address: "Jalan A",
			},
			{
				UserId:  "2",
				Address: "Jalan B",
			},
		},
	}

	result := db.Save(&user)
	assert.Nil(t, result.Error)
}

func TestPreloadJoinOneToMany(t *testing.T) {
	// var usersPreload []User
	var user User
	result := db.Model(&User{}).Preload("Addresses").Joins("Wallet").Take(&user, "users.id = ?", "50")
	fmt.Println(user)
	// for _, user := range usersPreload {
	// 	fmt.Println(user)
	// }
	assert.Nil(t, result.Error)
}

func TestBelongsTo(t *testing.T) {
	fmt.Println("Preload")
	var addresses []Address
	result := db.Preload("User").Find(&addresses)
	assert.Nil(t, result.Error)
	assert.Equal(t, 4, len(addresses))

	fmt.Println("Joins")
	addresses = []Address{}
	result = db.Model(&Address{}).Joins("User").Find(&addresses)
	assert.Nil(t, result.Error)
	assert.Equal(t, 4, len(addresses))
}

func TestBelongsToWallet(t *testing.T) {
	fmt.Println("Preload")
	var wallets []Wallet
	result := db.Preload("User").Find(&wallets)
	assert.Nil(t, result.Error)

	fmt.Println("Joins")
	wallets = []Wallet{}
	result = db.Model(&Wallet{}).Joins("User").Find(&wallets)
	assert.Nil(t, result.Error)
}

func TestCreateManyToMany(t *testing.T) {
	const relationDBName = "user_like_product"
	product := Product{
		Id:    "P001",
		Name:  "Contoh Product",
		Price: 1000000,
	}
	result := db.Save(&product)
	assert.Nil(t, result.Error)

	result = db.Table(relationDBName).Create(map[string]any{
		"user_id":    "1",
		"product_id": "P001",
	})
	assert.Nil(t, result.Error)

	result = db.Table(relationDBName).Create(map[string]any{
		"user_id":    "2",
		"product_id": "P001",
	})
	assert.Nil(t, result.Error)
}

func TestPreloadManyToMany(t *testing.T) {
	var product Product
	result := db.Preload("LikedByUsers").First(&product, "id = ?", "P001")
	assert.Nil(t, result.Error)
	for _, user := range product.LikedByUsers {
		fmt.Println(user)
	}
	assert.Equal(t, 2, len(product.LikedByUsers))
}

func TestPreloadManyToManyUser(t *testing.T) {
	var user User
	result := db.Preload("LikeProducts").First(&user, "id = ?", "1")
	assert.Nil(t, result.Error)
	assert.Equal(t, 1, len(user.LikeProducts))
}

func TestAssociationFind(t *testing.T) {
	var product Product
	result := db.First(&product, "id = ?", "P001")
	assert.Nil(t, result.Error)

	var users []User
	err := db.Model(&product).Where("users.first_name LIKE ?", "User%").Association("LikedByUsers").Find(&users)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(users))
}

func TestAssociationAdd(t *testing.T) {
	var user User
	err := db.Take(&user, "id = ?", "3").Error
	assert.Nil(t, err)

	var product Product
	err = db.Take(&product, "id = ?", "P001").Error
	assert.Nil(t, err)

	err = db.Model(&product).Association("LikedByUsers").Append(&user)
	assert.Nil(t, err)
}

func TestAssociationReplace(t *testing.T) {
	err := db.Transaction(func(tx *gorm.DB) error {
		var user User
		err := tx.Take(&user, "id = ?", "1").Error
		assert.Nil(t, err)

		wallet := Wallet{
			Id:      "01",
			UserId:  "1",
			Balance: 1000000,
		}
		err = tx.Model(&user).Association("Wallet").Replace(&wallet)
		return err
	})
	assert.Nil(t, err)
}

func TestAssociationDelete(t *testing.T) {
	var user User
	err := db.Take(&user, "id = ?", "3").Error
	assert.Nil(t, err)

	var product Product
	err = db.Take(&product, "id = ?", "P001").Error
	assert.Nil(t, err)

	err = db.Model(&product).Association("LikedByUsers").Delete(&user)
	assert.Nil(t, err)
}

func TestAssociationClear(t *testing.T) {
	var product Product
	err := db.Take(&product, "id = ?", "P001").Error
	assert.Nil(t, err)

	err = db.Model(&product).Association("LikedByUsers").Clear()
	assert.Nil(t, err)
}

func TestPreloadingWithCondition(t *testing.T) {
	var user User
	result := db.Preload("Wallet", "balance > ?", 1000000).First(&user, "id = ?", "1")
	assert.Nil(t, result.Error)
}

func TestNestedPreloading(t *testing.T) {
	var wallet Wallet
	result := db.Preload("User.Addresses").Find(&wallet, "id = ?", "2")
	assert.Nil(t, result.Error)
	fmt.Println(wallet)
	fmt.Println(wallet.User)
	fmt.Println(wallet.User.Addresses)
}

func TestPreloadingAll(t *testing.T) {
	var user User
	result := db.Preload(clause.Associations).Take(&user, "id = ?", "1")
	assert.Nil(t, result.Error)
}

func TestJoinQuery(t *testing.T) {
	var users []User
	result := db.Joins("join wallets on wallets.user_id = users.id").Find(&users)
	assert.Nil(t, result.Error)
	assert.Equal(t, 4, len(users))

	users = []User{}
	result = db.Joins("Wallet").Find(&users) // default menggunakan left join
	assert.Nil(t, result.Error)
	assert.Equal(t, 17, len(users))

	users = []User{}
	result = db.InnerJoins("Wallet").Find(&users) // default menggunakan left join
	assert.Nil(t, result.Error)
	assert.Equal(t, 4, len(users))
}

func TestJoinQueryCondition(t *testing.T) {
	var users []User
	result := db.Joins("join wallets on wallets.user_id = users.id AND wallets.balance > ?", 500000).Find(&users)
	assert.Nil(t, result.Error)
	assert.Equal(t, 4, len(users))

	users = []User{}
	result = db.Joins("Wallet").Where("Wallet.balance > ?", 500000).Find(&users)
	assert.Nil(t, result.Error)
	assert.Equal(t, 4, len(users))
}

func TestCount(t *testing.T) {
	var count int64
	result := db.Model(&User{}).Joins("Wallet").Where("Wallet.balance > ?", 500000).Count(&count)
	assert.Nil(t, result.Error)
	assert.Equal(t, int64(4), count)
}

func TestAggregation(t *testing.T) {
	var result AggregationResult
	err := db.Model(&Wallet{}).Select("sum(balance) as total_balance", "min(balance) as min_balance", "max(balance) as max_balance", "avg(balance) as avg_balance").Take(&result).Error
	assert.Nil(t, err)
	assert.Equal(t, int64(4000000), result.TotalBalance)
	assert.Equal(t, int64(1000000), result.MinBalance)
	assert.Equal(t, int64(1000000), result.MaxBalance)
	assert.Equal(t, float64(1000000), result.AvgBalance)
}

func TestAggregationGroupByAndHaving(t *testing.T) {
	var results []AggregationResult
	err := db.Model(&Wallet{}).Select("sum(balance) as total_balance", "min(balance) as min_balance", "max(balance) as max_balance", "avg(balance) as avg_balance").
		Joins("User").Group("User.id").Having("sum(balance) > ?", 500000).
		Find(&results).Error

	assert.Nil(t, err)
	assert.Equal(t, 4, len(results))
}

func TestContext(t *testing.T) {
	ctx := context.Background()

	var users []User
	result := db.WithContext(ctx).Find(&users)
	assert.Nil(t, result.Error)
	assert.Equal(t, 17, len(users))
}

func BrokeWalletBalance(db *gorm.DB) *gorm.DB {
	return db.Where("balance = ?", 0)
}

func SultanWalletBalance(db *gorm.DB) *gorm.DB {
	return db.Where("balance > ?", 1000000)
}

func TestScopes(t *testing.T) {
	var wallets []Wallet
	result := db.Scopes(BrokeWalletBalance).Find(&wallets)
	assert.Nil(t, result.Error)

	wallets = []Wallet{}
	result = db.Scopes(SultanWalletBalance).Find(&wallets)
	assert.Nil(t, result.Error)
}

func TestConnectionPool(t *testing.T) {
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)
}

func TestMigrator(t *testing.T) {
	err := db.Migrator().AutoMigrate(&GuestBook{})
	assert.Nil(t, err)
}

func TestHook(t *testing.T) {
	user := User{
		Password: "rahasia",
		Name: Name{
			FirstName: "User 100",
		},
	}

	err := db.Create(&user).Error
	assert.Nil(t, err)
	assert.NotEqual(t, "", user.Id)
	fmt.Println(user)
}
