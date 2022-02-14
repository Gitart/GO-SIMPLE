package main

import (
    "bytes"
    "go/ast"
    "go/parser"
    "go/printer"
    "go/token"
    "golang.org/x/tools/go/ast/inspector"
    "log"
    "os"
    "text/template"
)

//Шаблон, на основе которого будем генерировать
//.EntityName, .PrimaryType — параметры,
//в которые будут установлены данные, добытые из AST-модели
var repositoryTemplate = template.Must(template.New("").Parse(`
package main

import (
    "github.com/jinzhu/gorm"
)

type {{ .EntityName }}Repository struct {
    db *gorm.DB
}

func New{{ .EntityName }}Repository(db *gorm.DB) {{ .EntityName }}Repository {
    return {{ .EntityName }}Repository{ db: db}
}

func (r {{ .EntityName }}Repository) Get({{ .PrimaryName }} {{ .PrimaryType}}) (*{{ .EntityName }}, error) {
    entity := new({{ .EntityName }})
    err := r.db.Limit(1).Where("{{ .PrimarySQLName }} = ?", {{ .PrimaryName }}).Find(entity).Error()
    return entity, err
}


func (r {{ .EntityName }}Repository) Create(entity *{{ .EntityName }}) error {
    return r.db.Create(entity).Error
}

func (r {{ .EntityName }}Repository) Update(entity *{{ .EntityName }}) error {
    return r.db.Model(entity).Update.Error
}

func (r {{ .EntityName }}Repository) Update(entity *{{ .EntityName }}) error {
    return r.db.Model(entity).Update.Error
}

func (r {{ .EntityName }}Repository) Delete(entity *{{ .EntityName }}) error {
    return r.db.Delete.Error
}
`))

//Агрегатор данных для установки параметров в шаблоне
type repositoryGenerator struct{
    typeSpec    *ast.TypeSpec
    structType  *ast.StructType
}

//Просто helper-функция для печати замысловатого ast.Expr в обычный string
func expr2string(expr ast.Expr) string {
    var buf bytes.Buffer
    err := printer.Fprint(&buf, token.NewFileSet(), expr)
    if err !- nil {
        log.Fatalf("error print expression to string: #{err}")
    return buf.String()
}

//Helper для извлечения поля структуры,
//которое станет первичным ключом в таблице DB
//Поиск поля ведётся по тегам
//Ищем то, что мы пометили gorm:"primary_key"
func (r repositoryGenerator) primaryField() (*ast.Field, error) {
    for _, field := range r.structType.Fields.List {
        if !strings.Contains(field.Tag.Value, "primary")
            continue
        }
        return field, nil
    }
    return nil, fmt.Errorf("has no primary field")
}

//Собственно, генератор
//оформлен методом структуры repositoryGenerator,
//так что параметры передавать не нужно:
//они уже аккумулированы в ресивере метода r repositoryGenerator
//Передаём ссылку на ast.File,
//в котором и окажутся плоды трудов
func (r repositoryGenerator) Generate(outFile *ast.File) error {
    //Находим первичный ключ
    primary, err := r.primaryField()
    if err != nil {
        return err
    }
    //Аллокация и установка параметров для template
    params := struct {
        EntityName      string
        PrimaryName     string
        PrimarySQLName  string
        PrimaryType     string
    }{
        //Параметры извлекаем из ресивера метода
        EntityName      r.typeSpec.Name.Name,
        PrimaryName     primary.Names[0].Name,
        PrimarySQLName  primary.Names[0].Name,
        PrimaryType     expr2string(primary.Type),
    }
    //Аллокация буфера,
    //куда будем заливать выполненный шаблон
    var buf bytes.Buffer
    //Процессинг шаблона с подготовленными параметрами
    //в подготовленный буфер
    err = repositoryTemplate.Execute(&buf, params)
    if err != nil {
        return fmt.Errorf("execute template: %v", err)
    }
    //Теперь сделаем парсинг обработанного шаблона,
    //который уже стал валидным кодом Go,
    //в дерево разбора,
    //получаем AST этого кода
    templateAst, err := parser.ParseFile(
        token.NewFileSet(),
        //Источник для парсинга лежит не в файле,
        "",
        //а в буфере
        buf.Bytes(),
        //mode парсинга, нас интересуют в основном комментарии
        parser.ParseComments,
    )
    if err != nil {
        return fmt.Errorf("parse template: %v", err)
    }
    //Добавляем декларации из полученного дерева
    //в результирующий outFile *ast.File,
    //переданный нам аргументом
    for _, decl := range templateAst.Decls {
        outFile.Decls = append(outFile.Decls, decl)
    }
    return nil
}

func main() {
    //Цель генерации передаётся переменной окружения
    path := os.Getenv("GOFILE")
    if path == "" {
        log.Fatal("GOFILE must be set")
    }
    //Разбираем целевой файл в AST
    astInFile, err := parser.ParseFile(
        token.NewFileSet(),
        path,
        src: nil,
        //Нас интересуют комментарии
        parser.ParseComments,
    )
    if err != nil {
        log.Fatalf("parse file: %v", err)
    }
    //Для выбора интересных нам деклараций
    //используем Inspector из golang.org/x/tools/go/ast/inspector
    i := inspector.New([]*ast.File{astInFile})
    //Подготовим фильтр для этого инспектора
    iFilter := []ast.Node{
        //Нас интересуют декларации
        &ast.GenDecl{},
    }
    //Выделяем список заданий генерации
    var genTasks []repositoryGenerator
    //Запускаем инспектор с подготовленным фильтром
    //и литералом фильтрующей функции
    i.Nodes(iFilter, func(node ast.Node, push bool) (proceed bool){
        genDecl := node.(*ast.GenDecl)
        //Код без комментариев не нужен,
        if genDecl.Doc == nil {
            return false
        }
        //интересуют спецификации типов,
        typeSpec, ok := genDecl.Specs[0].(*ast.TypeSpec)
        if !ok {
            return false
        }
        //а конкретно структуры
        structType, ok := typeSpec.Type.(*ast.StructType)
        if !ok {
            return false
        }
        //Из оставшегося
        for _, comment := range genDecl.Doc.List {
            switch comment.Text {
            //выделяем структуры, помеченные комментарием repogen:entity,
            case "//repogen:entity":
                //и добавляем в список заданий генерации
                genTasks = append(genTasks, repositoryGenerator{
                    typeSpec: typeSpec,
                    structType: structType,
                })
            }
        }
        return false
    })
    //Аллокация результирующего дерева разбора
    astOutFile := &ast.File{
        Name: astInFile.Name,
    }
    //Запускаем список заданий генерации
    for _, task := range genTask {
        //Для каждого задания вызываем написанный нами генератор
        //как метод этого задания
        //Сгенерированные декларации помещаются в результирующее дерево разбора
        err = task.Generate(astOutFile)
        if err != nil {
            log.Fatalf("generate: %v", err)
        }
    }
    //Подготовим файл конечного результата всей работы,
    //назовем его созвучно файлу модели, добавим только суффикс _gen
    outFile, err := os.Create(strings.TrimSuffix(path, ".go") + "_gen.go")
    if err != nil {
        log.Fatalf("create file: %v", err)
    }
    //Не забываем прибраться
    defer outFile.Close()
    //Печатаем результирующий AST в результирующий файл исходного кода
    //«Печатаем» не следует понимать буквально,
    //дерево разбора нельзя просто переписать в файл исходного кода,
    //это совершенно разные форматы
    //Мы здесь воспользуемся специализированным принтером из пакета ast/printer
    err = printer.Fprint(outFile, token.NewFileSet(), astOutFile)
    if err != nil {
        log.Fatalf("print file: %v", err)
    }
}
