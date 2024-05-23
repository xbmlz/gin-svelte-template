package main

import (
	"github.com/xbmlz/gin-svelte-template/internal/config"
	"github.com/xbmlz/gin-svelte-template/internal/dal"
	"github.com/xbmlz/gin-svelte-template/internal/logger"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {

	conf := config.NewConfig()
	log := logger.NewLogger(conf)
	db := dal.NewDatabase(conf, log)

	g := gen.NewGenerator(gen.Config{
		// 相对执行`go run`时的路径, 会自动创建目录
		OutPath: "internal/dal/dao",
		// WithDefaultQuery 生成默认查询结构体(作为全局变量使用), 即`Q`结构体和其字段(各表模型)
		// WithoutContext 生成没有context调用限制的代码供查询
		// WithQueryInterface 生成interface形式的查询代码(可导出), 如`Where()`方法返回的就是一个可导出的接口类型
		Mode: gen.WithDefaultQuery | gen.WithQueryInterface,
		// 表字段可为 null 值时, 对应结体字段使用指针类型
		FieldNullable: true,
		// 表字段默认值与模型结构体字段零值不一致的字段, 在插入数据时需要赋值该字段值为零值的, 结构体字段须是指针类型才能成功, 即`FieldCoverable:true`配置下生成的结构体字段.
		// 因为在插入时遇到字段为零值的会被GORM赋予默认值. 如字段`age`表默认值为10, 即使你显式设置为0最后也会被GORM设为10提交.
		// 如果该字段没有上面提到的插入时赋零值的特殊需要, 则字段为非指针类型使用起来会比较方便.
		FieldCoverable: false,
		// 模型结构体字段的数字类型的符号表示是否与表字段的一致, `false`指示都用有符号类型
		FieldSignable: false,
		// 生成 gorm 标签的字段索引属性
		FieldWithIndexTag: false,
		// 生成 gorm 标签的字段类型属性
		FieldWithTypeTag: true,
	})

	g.UseDB(db)

	// 自定义字段的数据类型
	// 统一数字类型为int64,兼容protobuf
	dataMap := map[string]func(columnType gorm.ColumnType) (dataType string){
		"tinyint":   func(columnType gorm.ColumnType) (dataType string) { return "int64" },
		"smallint":  func(columnType gorm.ColumnType) (dataType string) { return "int64" },
		"mediumint": func(columnType gorm.ColumnType) (dataType string) { return "int64" },
		"bigint":    func(columnType gorm.ColumnType) (dataType string) { return "int64" },
		"int":       func(columnType gorm.ColumnType) (dataType string) { return "int64" },
	}

	// 要先于`ApplyBasic`执行
	g.WithDataTypeMap(dataMap)

	g.ApplyBasic(g.GenerateModel("users"))

	g.Execute()
}
