package gorm

import (
	"majordome/config"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"

	"fmt"
)

type Gorm struct {
	c *config.Config
}

func (g *Gorm) NewDb() []*gorm.DB {
	g.c = &config.Config{}

	c, _ := g.c.NewConfig()

	arrayConnections := make([]*gorm.DB, 0)

	var db *gorm.DB

	var err error

	if len(c.DB.Postgres) > 0 {
		for _, v := range c.DB.Postgres {
			psqlConnStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
				v.Host,
				v.Port,
				v.UserName,
				v.Database,
				v.Password)
			db, err = gorm.Open("postgres", psqlConnStr)

			if err != nil {
				fmt.Printf("err postgres = %v \n", err)
			}

			arrayConnections = append(arrayConnections, db)

			db = &gorm.DB{}
		}
	}
	return arrayConnections
}
