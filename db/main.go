package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"io/ioutil"
)

type Config struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbname"`
	} `yaml:"database"`
}

// main
// @Description 使用gorm-gen生成model和query
// @Author Oberl-Fitzgerald 2023-12-07 10:50:27
func main() {
	var cfg Config

	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		panic(fmt.Sprintf("Error reading YAML file: %v", err))
	}

	err = yaml.Unmarshal(yamlFile, &cfg)
	if err != nil {
		panic(fmt.Sprintf("Error unmarshalling YAML: %v", err))
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(fmt.Sprintf("GORM connect MYSQL ERROR: %+v", err))
	}

	g := gen.NewGenerator(gen.Config{
		OutPath:           "./query", /*生成的查询类代码的输出路径，默认./query*/
		ModelPkgPath:      "model",   /*生成DAO代码的包名，默认是：model*/
		WithUnitTest:      true,      /*是否为DAO包生成单元测试代码，默认：false*/
		FieldNullable:     true,      /*数据库中的字段可为空，则生成struct字段为指针类型*/
		FieldCoverable:    false,     /*如果数据库中字段有默认值，则生成指针类型的字段，以避免零值（zero-value）问题*/
		FieldSignable:     false,     /*Use signable type as field’s type based on column’s data type in database*/
		FieldWithIndexTag: true,      /*为结构体生成gorm index tag，如gorm:"index:idx_name"，默认：false*/
		FieldWithTypeTag:  true,      /*为结构体生成gorm type tag，如：gorm:"type:varchar(12)"，默认：false*/
		Mode:              gen.WithDefaultQuery | gen.WithQueryInterface,
		/*gen.WithDefaultQuery	是否生成全局变量Q作为DAO接口，如果开启，你可以通过这样的方式查询数据dal.Q.User.First()*/
		/*gen.WithQueryInterface	生成查询API代码，而不是struct结构体。通常用来MOCK测试*/
		/*gen.WithoutContext	生成无需传入context参数的代码。如果无需传入context，则代码调用方式如：dal.User.First()，否则，调用方式要像这样：dal.User.WithContext(ctx).First()*/
	})
	g.UseDB(db)

	// 自定义字段的数据类型
	// 统一数字类型为int64,兼容protobuf和thrift
	dataMap := map[string]func(columnType gorm.ColumnType) (dataType string){
		"tinyint":   func(columnType gorm.ColumnType) (dataType string) { return "int64" },
		"smallint":  func(columnType gorm.ColumnType) (dataType string) { return "int64" },
		"mediumint": func(columnType gorm.ColumnType) (dataType string) { return "int64" },
		"bigint":    func(columnType gorm.ColumnType) (dataType string) { return "int64" },
		"int":       func(columnType gorm.ColumnType) (dataType string) { return "int64" },
	}
	// 要先于`ApplyBasic`执行
	g.WithDataTypeMap(dataMap)

	allModel := g.GenerateAllTable()
	g.ApplyBasic(allModel...)
	g.Execute()
}
