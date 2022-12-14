package svc

import (
	"api/internal/config"
	"api/model/favorite"
	"api/model/ingredient"
	"api/model/recipe"
	"api/model/user"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/postgres"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/golang-migrate/migrate/v4"
	postgresmigrate "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

type ServiceContext struct {
	Config          config.Config
	IngredientModel ingredient.IngredientsModel
	// RecipeModel give access to stages and quantities models
	RecipeModel   recipe.RecipesModel
	FavoriteModel favorite.FavoritesModel
	UserModel     user.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	conn := postgres.New(c.Postgre.Datasource)

	// DB automatic migration
	if err := initDB(conn, c.Postgre.MigrationFolder); err != nil {
		fmt.Println(err.Error())
	}

	return &ServiceContext{
		Config:          c,
		IngredientModel: ingredient.NewIngredientsModel(conn),
		RecipeModel:     recipe.NewRecipesModel(conn),
		FavoriteModel:   favorite.NewFavoritesModel(conn),
		UserModel:       user.NewUsersModel(conn),
	}
}

// Return an error if the migration fails.
// Do not alter the db if the schema is up-to-date.
// Based on the schema_migration table versus content of the sql folder.
func initDB(conn sqlx.SqlConn, migrationFolder string) error {
	db, err := conn.RawDB()
	if err != nil {
		return err
	}
	driver, err := postgresmigrate.WithInstance(db, &postgresmigrate.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://" + migrationFolder,
		"postgres", driver)
	if err != nil {
		return err
	}

	if err = m.Up(); err != nil {
		version, dirty, _ := m.Version()
		fmt.Println("DB Migration Version:", version, " isDirty:", dirty)
		return err
	}
	version, dirty, err := m.Version()
	if err != nil {
		return err
	}
	fmt.Println("DB Migration Version:", version, " isDirty:", dirty)

	return nil
}
