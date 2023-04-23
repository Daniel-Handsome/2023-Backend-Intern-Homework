package repository

import (
	"github.com/Daniel-Handsome/2023-Backend-intern-Homework/model"
	"gorm.io/gorm"
)

//baseRepo interface 可以寫crud之類的

type BaseRepo struct {
	orm *gorm.DB
}

func NewRepo(orm *gorm.DB, model model.BaseModel) *BaseRepo {
	return &BaseRepo{orm: orm}
}

func (r *BaseRepo) Transaction(execute func(tx *gorm.DB) error) error {
	return r.orm.Transaction(execute)
}

func (r *BaseRepo) Begin() *gorm.DB {
	return r.orm.Begin()
}

type FilterDecorator interface {
	ApplyFilter(query *gorm.DB, value interface{}) *gorm.DB
}

type Filters []FilterDecorator

func (filters Filters) Apply(query *gorm.DB, value interface{}) *gorm.DB {
	for _, filter := range filters {
		query = filter.ApplyFilter(query, value)
	}
	return query
}

//func (r *BaseRepo) FilterDecoratorsFromRequest(request map[string]string, namespace string) *gorm.DB {
//	filters := []FilterDecorator{}
//
//	// 取得所有 filter 名稱
//	filterNames := []string{"name", "age"}
//
//	// 產生指定 namespace
//
//	// 產生所有 filters
//	for _, name := range filterNames {
//		pkgPath := "github.com/Daniel-Handsome/2023-Backend-intern-Homework/repository/filter"
//		typeName := "NameFilter"
//
//		// 動態載入 package
//		pkg, err := reflect.Import(pkgPath, nil)
//		if err != nil {
//			panic(err)
//		}
//
//		// 取得 type
//		typeFullName := fmt.Sprintf("%s.%s", pkgPath, typeName)
//		t := pkg.Type(typeFullName)
//		// 判斷 filter 是否存在於指定 namespace
//		if t.Implements(filterType) {
//			// 創建 filter 實例
//			f := reflect.New(t.Elem()).Interface().(FilterDecorator)
//
//			// 加入 filters 切片
//			filters = append(filters, f)
//		}
//	}
//
//	// 使用 filters 執行過濾
//	var query *gorm.DB
//	query = filters[0].ApplyFilter(query, "Daniel")
//	query = filters[1].ApplyFilter(query, 20)
//
//	return query
//}

//func getTypeByPath(path string) (reflect.Type, error) {
//	p, err := buildPath(path)
//	if err != nil {
//		return nil, err
//	}
//	return reflect.TypeOf(p), nil
//}
//
//func buildPath(path string) (interface{}, error) {
//	// 分割路径
//	parts := strings.Split(path, "/")
//	// 获取路径的根部分
//	root := parts[0]
//	// 构造要执行的命令
//	cmd := fmt.Sprintf("package %s; import %s \"%s\"; var v %s.%s", root, root, parts[1], parts[1], parts[2])
//	// 解析命令
//	fset := token.NewFileSet()
//	file, err := parser.ParseFile(fset, "", cmd, parser.Mode(0))
//	if err != nil {
//		return nil, err
//	}
//	// 获取变量 v 的值
//	v := reflect.ValueOf(file.Scope.Lookup("v").Decl.(*ast.ValueSpec).Values[0]).Interface()
//	return v, nil
//}
