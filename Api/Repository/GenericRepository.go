package repository

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"regexp"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"github.com/uptrace/bun/extra/bundebug"
)

var (
	DB  *bun.DB
	Ctx = context.Background()
)

type GenericRepositoryInterface interface {
	FindOneById(entity interface{}, id int) error
	FindOneBy(entity interface{}, params map[string]interface{}) error
	FindAll(entities interface{}) error
	FindAllBy(entities interface{}, params map[string]interface{}) error
	CountAll(entity interface{}) (int, error)
	Create(entity interface{}) (sql.Result, error)
	Update(entity interface{}) (sql.Result, error)
}

type GenericRepository struct {
	GenericRepositoryInterface
}

func (r *GenericRepository) Init(entities []interface{}) {
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		fmt.Println("DB_URL environment variable is required")
		os.Exit(1)
	}
	sqldb, err := sql.Open("mysql", dbURL)
	if err != nil {
		panic(err)
	}

	DB = bun.NewDB(sqldb, mysqldialect.New())

	file, err := os.Create("bundebug.log")
	if err != nil {
		panic(err)
	}

	DB.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.WithWriter(file),
	))

	if err := DB.Ping(); err != nil {
		panic(err)
	}

	for i := range entities {
		DB.RegisterModel(&entities[i])
		_, err := DB.NewCreateTable().Model(&entities[i]).IfNotExists().Exec(Ctx)
		if err != nil {
			panic(err)
		}
	}
}

func (r *GenericRepository) FindOneById(entity interface{}, id int) error {
	return DB.NewSelect().Model(entity).Where("id = ?", id).Scan(Ctx)
}

func (r *GenericRepository) applyParams(model *bun.SelectQuery, params map[string]interface{}) *bun.SelectQuery {
	for key, value := range params {
		if key == "limit" || key == "offset" {
			limit, boolErr := value.(string)
			if !boolErr {
				continue
			}

			limitInt, err := strconv.Atoi(limit)
			if err != nil {
				continue
			}

			if key == "offset" {
				model.Offset(limitInt)
			} else if key == "limit" {
				model.Limit(limitInt)
			}
			continue
		}

		regex := regexp.MustCompile(`^[a-z_]+$`)
		if !regex.MatchString(key) {
			continue
		}
		model.Where(key+" = ?", value)
	}
	return model
}

func (r *GenericRepository) FindOneBy(entity interface{}, params map[string]interface{}) error {
	model := DB.NewSelect().Model(entity)
	params["limit"] = 1
	params["offset"] = 0
	model = r.applyParams(model, params)
	return model.Scan(Ctx)
}

func (r *GenericRepository) FindAll(entities interface{}) error {
	return DB.NewSelect().Model(entities).Scan(Ctx)
}

func (r *GenericRepository) FindAllBy(entities interface{}, params map[string]interface{}) error {
	model := DB.NewSelect().Model(entities)
	model = r.applyParams(model, params)
	return model.Scan(Ctx)
}

func (r *GenericRepository) CountAll(entity interface{}) (int, error) {
	return DB.NewSelect().Model(entity).Count(Ctx)
}

func (r *GenericRepository) Create(entity interface{}) (sql.Result, error) {
	return DB.NewInsert().Model(entity).Exec(Ctx)
}

func (r *GenericRepository) Update(entity interface{}) (sql.Result, error) {
	return DB.NewUpdate().Model(entity).Exec(Ctx)
}
