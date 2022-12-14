package dbconn

import (
	"log"

	"github.com/Gurpreet-Bacancy/bcl/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Postgres struct {
	DB *gorm.DB
}

// New Instantiates Postgres service.
func NewPostgres(dbUrl string) (*Postgres, error) {
	db := getDBInstance(dbUrl)

	return &Postgres{DB: db}, nil
}

// connection with postgres database
func getDBInstance(dbUrl string) *gorm.DB {
	connection, err := gorm.Open("postgres", dbUrl)
	if err != nil {
		log.Fatalln(err)
	}

	sqldb := connection.DB()

	err = sqldb.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	return connection
}

// close database connetion
func Closedatabase(connection *gorm.DB) {
	sqldb := connection.DB()
	sqldb.Close()
}

// migrate customer model to database
func Initialmigration(dbUrl string) {
	connection := getDBInstance(dbUrl)
	connection.AutoMigrate(&model.UserItem{})
	connection.AutoMigrate(&model.Coordinates{})
	connection.Model(&model.Coordinates{}).AddForeignKey("user_id", "user_items(id)", "CASCADE", "CASCADE")
	defer Closedatabase(connection)
	log.Println("database migration done")
}

func Seeder(dbUrl string) {
	connection := getDBInstance(dbUrl)
	defer Closedatabase(connection)

	users := []model.UserItem{
		{
			Name:  "Steven victo",
			Email: "steven@gmail.com",
		},
		{
			Name:  "Martin Luther",
			Email: "luther@gmail.com",
		},
		{
			Name:  "Mar",
			Email: "mar@gmail.com",
		},
		{
			Name:  "tin",
			Email: "tin@gmail.com",
		},
		{
			Name:  "Lu",
			Email: "lu@gmail.com",
		},
		{
			Name:  "in",
			Email: "in@gmail.com",
		},
		{
			Name:  "ther",
			Email: "ther@gmail.com",
		},
		{
			Name:  "Ste",
			Email: "Ste@gmail.com",
		},
		{
			Name:  "even",
			Email: "even@gmail.com",
		},
		{
			Name:  "en",
			Email: "en@gmail.com",
		},
		{
			Name:  "envi",
			Email: "envi@gmail.com",
		},
		{
			Name:  "vic",
			Email: "vic@gmail.com",
		},
		{
			Name:  "venvi",
			Email: "venvi@gmail.com",
		},
		{
			Name:  "nvicto",
			Email: "nvicto@gmail.com",
		},
		{
			Name:  "cto",
			Email: "cto@gmail.com",
		},
	}

	for _, user := range users {
		connection.Create(&user)
	}

	connection.Find(&users)
	latitude := 10
	longitude := 10
	altitude := 10

	for _, user := range users {

		cordinate := model.Coordinates{
			UserID:    user.ID,
			Latitude:  float32(latitude),
			Longitude: float32(longitude),
			Altitude:  float32(altitude),
		}
		connection.Create(&cordinate)

		latitude += 10
		longitude += 10
		altitude += 10

	}

	log.Println("database seeding done")
}
