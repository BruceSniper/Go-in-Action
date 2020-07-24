package sql2struct

import (
	"fmt"
	"github.com/go-programming-tour-book/tour/internal/word"
	"html/template"
	"os"
)

const structTpl = `type {{.TableName | ToCamelCase}} struct {
{{range .Columns}}	{{ $length := len .Comment}} {{ if gt $length 0 }}// {{.Comment}} {{else}}// {{.Name}} {{ end }}
	{{ $typeLen := len .Type }} {{ if gt $typeLen 0 }}{{.Name | ToCamelCase}}	{{.Type}}	{{.Tag}}{{ else }}{{.Name}}{{ end }}
{{end}}}
func (model {{.TableName | ToCamelCase}}) TableName() string {
	return "{{.TableName}}"
}`

type StructTemplate struct {
	structTpl string
}

type StructColumn struct { //用来存储转换后的Go结构体中的所有字段信息
	Name    string
	Type    string
	Tag     string
	Comment string
}

type StructTemplateDB struct { //用来存储最终用于渲染的模板对象信息
	TableName string
	Columns   []*StructColumn
}

func NewStructTemplate() *StructTemplate {
	return &StructTemplate{structTpl: structTpl}
}

//对通过查询COLUMNS表所组装得到的 tbColumns 进行进一步的分解和转换
func (t *StructTemplate) AssemblyColumns(tbColumns []*TableColumn) []*StructColumn {
	tplColumns := make([]*StructColumn, 0, len(tbColumns))
	for _, column := range tbColumns {
		tplColumns = append(tplColumns, &StructColumn{
			Name:    column.ColumnName,
			Type:    DBTypeToStructType[column.DataType],
			Tag:     fmt.Sprintf("`json:"+"%s"+"`", column.ColumnName),
			Comment: column.ColumnComment,
		})
	}
	return tplColumns
}

func (t *StructTemplate) Generate(tableName string, tplColumns []*StructColumn) error {
	tpl := template.Must(template.New("sql2struct").Funcs(template.FuncMap{ //首先声明了一个名为sql2struct的新模板对象
		"ToCamelCase": word.UnderscoreToUpperCamelCase, //定义了自定义函数ToCamelCase，并与word.UnderscoreToUpperCamelCase方法进行绑定
	}).Parse(t.structTpl))

	tplDB := StructTemplateDB{ //组装符合预定义模板的模板对象
		TableName: tableName,
		Columns:   tplColumns,
	}
	err := tpl.Execute(os.Stdout, tplDB) //调用Execute方法进行渲染
	if err != nil {
		return err
	}
	return nil
}
