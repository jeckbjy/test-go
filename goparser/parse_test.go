package goparser

import (
	"fmt"
	"go/ast"
	"go/doc"
	"go/parser"
	"go/token"
	"log"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

const fileSource = `
package annotation

import "context"

//aa
type ZDao interface {
	Get()
}

// ChatDao aaa
// @orm(driver=mysql,table=foo,model=Foo)
type ChatDao interface {
	// 通过ChatID和权限查询ChatterID
	// @sql("select chatter_id from chat where chat_id=? and post_permission=?")
	GetByChatID(ctx context.Context, chatID *int64, permission int) ([]int64, error)

	// 通过ChatID和UserID查询ChatterID
	// @sql("select chatter_id from chat
	// where chat_id=? and chatter_id in (?) and post_permission=?", mode=named)
	// @sql(select)
	GetByChatIDUserID(ctx context.Context, chatID int64, 
		uids []int64, permission1 int) ([]int64, error)

	// @sql(update xx where id = ? set type = ?)
	Update(ctx context.Context, id int64, type_ string) error

	// @sql(insert)
	Insert(ctx context.Context)

	// @sql(delete)
	Delete(ctx context.Context, id int64)
}
`

const fileSource1 = `
package main
// This documents FirstType and SecondType together
type (
    // FirstType docs
    FirstType struct {
        // FirstMember docs
        FirstMember string
    }

    // SecondType docs
    SecondType struct {
        // SecondMember docs
        SecondMember string
	}
)
`

func TestParser(t *testing.T) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "src.go", fileSource, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	p, err := doc.NewFromFiles(fset, []*ast.File{f}, "", doc.AllDecls|doc.AllMethods|doc.PreserveAST)
	if err != nil {
		log.Fatal(err)
	}

	offset := f.Pos()

	fmt.Printf("aaa%+v\n", len(p.Funcs))
	for _, t1 := range p.Types {
		fmt.Printf("name:%+v\n", t1.Name)
		fmt.Printf("docs:%+v", t1.Doc)
		for _, s := range t1.Decl.Specs {
			switch x := s.(type) {
			case *ast.TypeSpec:
				fmt.Printf("x:%+v\n", x.Name)
				fmt.Printf("d:%+v\n", x.Doc.Text())
				i := x.Type.(*ast.InterfaceType)
				for _, m := range i.Methods.List {
					fmt.Printf("func:%+v,%+v\n", m.Names[0].Name, m.Doc.Text())
					f := m.Type.(*ast.FuncType)
					for _, p := range f.Params.List {
						fmt.Printf("name: %+v\n", p.Names[0].Name)
						// fmt.Printf("type:%T\n", p.Type)
						fmt.Printf("type: %s\n",
							fileSource[p.Type.Pos()-offset+1:p.Type.End()-offset+1])
					}
				}
			}
			// spew.Dump(s)
		}
		for _, m := range t1.Methods {
			fmt.Printf("func name:%s\n", m.Name)
			fmt.Printf("func docs:%s\n", m.Doc)
		}
	}

	fmt.Printf("-------------------\n")

	// ast.Inspect(f, func(n ast.Node) bool {
	// 	switch x := n.(type) {
	// 	case *ast.TypeSpec:
	// 		if x.Name.IsExported() {
	// 			switch z := x.Type.(type) {
	// 			case *ast.InterfaceType:
	// 				fmt.Printf("name:%s\n", x.Name.Name)
	// 				fmt.Printf("docs:%s\n", x.Doc.Text())
	// 				for _, m := range z.Methods.List {
	// 					fmt.Printf("func:%+v,%+v\n", m.Names[0].Name, m.Doc.Text())
	// 					f := m.Type.(*ast.FuncType)
	// 					for _, p := range f.Params.List {
	// 						fmt.Printf("name: %+v\n", p.Names[0].Name)
	// 						// fmt.Printf("type:%T\n", p.Type)
	// 						fmt.Printf("type: %s\n",
	// 							fileSource[p.Type.Pos()-offset+1:p.Type.End()-offset+1])
	// 					}

	// 					for _, p := range f.Results.List {
	// 						fmt.Printf("type: %s\n",
	// 							fileSource[p.Type.Pos()-offset+1:p.Type.End()-offset+1])
	// 						// fmt.Printf("type:%T\n", p.Type)
	// 					}

	// 					// if f, ok := m.Type.(*ast.FuncType); ok {
	// 					// 	fmt.Printf("%s\n", m.Names)
	// 					// }
	// 				}
	// 			}
	// 		}
	// 	}
	// 	return true
	// })

	// ast.Inspect(f, func(n ast.Node) bool {
	// 	switch x := n.(type) {
	// 	case *ast.FuncDecl:
	// 		fmt.Printf("%s:\tFuncDecl %s\t%s\n", fset.Position(n.Pos()), x.Name, x.Doc.Text())
	// 	case *ast.TypeSpec:
	// 		fmt.Printf("%s:\tTypeSpec %s\t%s\n", fset.Position(n.Pos()), x.Name, x.Doc.Text())
	// 	case *ast.Field:
	// 		fmt.Printf("%s:\tField %s\t%s\n", fset.Position(n.Pos()), x.Names, x.Doc.Text())
	// 	case *ast.GenDecl:
	// 		fmt.Printf("%s:\tGenDecl %s\n", fset.Position(n.Pos()), x.Doc.Text())
	// 	}
	// 	return true
	// })

	// printer.Fprint(os.Stdout, fs, f)
	// spew.Dump(f)
	// v := visitor{locals: make(map[string]int)}
	// ast.Walk(v, f)
	// fmt.Printf("%+v\n", v.locals)

	// v := &visitor{}
	// ast.Walk(v, f)
}

type visitor struct {
}

func (v *visitor) Visit(n ast.Node) ast.Visitor {
	if n == nil {
		return nil
	}

	switch t := n.(type) {
	case *ast.Comment:
		fmt.Printf("comment:%+v\n", t.Text)
	case *ast.CommentGroup:
		for _, c := range t.List {
			fmt.Printf("comment group:%+v\n", c)
		}
	case *ast.InterfaceType:
		for _, m := range t.Methods.List {
			spew.Dump(m)
		}
		// fmt.Printf("interface:%+v\n", t.Methods)
	}

	return v
}

// type visitor struct{}

// func (v visitor) Visit(n ast.Node) ast.Visitor {
// 	fmt.Printf("%T\n", n)
// 	return v
// }

// type visitor int

// func (v visitor) Visit(n ast.Node) ast.Visitor {
// 	if n == nil {
// 		return nil
// 	}
// 	fmt.Printf("%s%T\n", strings.Repeat("\t", int(v)), n)
// 	return v + 1
// }

// type visitor struct {
// 	locals map[string]int
// }

// func (v visitor) Visit(n ast.Node) ast.Visitor {
// 	if n == nil {
// 		return nil
// 	}
// 	switch d := n.(type) {
// 	case *ast.AssignStmt:
// 		for _, name := range d.Lhs {
// 			if ident, ok := name.(*ast.Ident); ok {
// 				if ident.Name == "_" {
// 					continue
// 				}
// 				if ident.Obj != nil && ident.Obj.Pos() == ident.Pos() {
// 					v.locals[ident.Name]++
// 				}
// 			}
// 		}
// 	}
// 	return v
// }

// func (v visitor) local(n ast.Node) {
// 	ident, ok := n.(*ast.Ident)
// 	if !ok {
// 		return
// 	}
// 	if ident.Name == "_" || ident.Name == "" {
// 		return
// 	}
// 	if ident.Obj != nil && ident.Obj.Pos() == ident.Pos() {
// 		v.locals[ident.Name]++
// 	}
// }

// func (v visitor) localList(fs []*ast.Field) {
// 	for _, f := range fs {
// 		for _, name := range f.Names {
// 			v.local(name)
// 		}
// 	}
// }

// type visitor struct {
// 	pkgDecls map[*ast.GenDecl]bool
// 	globals  map[string]int
// 	locals   map[string]int
// }

// func newVisitor(f *ast.File) visitor {
// 	decls := make(map[*ast.GenDecl]bool)
// 	for _, decl := range f.Decls {
// 		if d, ok := decl.(*ast.GenDecl); ok {
// 			decls[d] = true
// 		}
// 	}
// 	return visitor{
// 		decls,
// 		make(map[string]int),
// 		make(map[string]int),
// 	}
// }
