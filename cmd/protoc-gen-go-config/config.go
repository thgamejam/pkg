package main

import (
	"fmt"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
	"strings"
)

type importPackage struct {
	name string
	path string
}

func (i *importPackage) GetName() string {
	return i.name
}

func (i *importPackage) GetPath() string {
	return i.path
}

var (
	goPackage = ""
	importSet = make(map[string]*importPackage)
)

// generateFile generates a _conf.pb.go file.
func generateFile(gen *protogen.Plugin, file *protogen.File) *protogen.GeneratedFile {
	if len(file.Messages) == 0 {
		return nil
	}
	filename := file.GeneratedFilenamePrefix + "_config.pb.go"
	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	g.P("// Code generated by protoc-gen-go-config. DO NOT EDIT.")
	g.P("// versions:")
	g.P(fmt.Sprintf("// protoc-gen-go-config %s", release))
	g.P()
	g.P("package ", file.GoPackageName)
	g.P()

	options, _ := file.Desc.Options().(*descriptorpb.FileOptions)
	goPackage = options.GetGoPackage()

	return g
}

// generateFileContent generates config file content.
func generateFileContent(gen *protogen.Plugin, file *protogen.File, g *protogen.GeneratedFile) {
	for _, message := range file.Messages {
		genMessage(gen, file, g, message)
	}
}

func genMessage(gen *protogen.Plugin, file *protogen.File, g *protogen.GeneratedFile, message *protogen.Message) {
	for _, s := range message.Comments.LeadingDetached {
		g.P(s)
	}

	g.P(message.Comments.Leading, "type ", message.GoIdent.GoName, "Interface interface {")

	for _, field := range message.Fields {
		genField(gen, file, g, field)
	}

	g.P("}")
}

func genField(gen *protogen.Plugin, file *protogen.File, g *protogen.GeneratedFile, field *protogen.Field) {
	if !field.Desc.Kind().IsValid() {
		return
	}

	var funcStr string
	if field.Desc.IsList() {
		funcStr = fmt.Sprint("Get", field.GoName, "() ", "[]", singleType(field.Desc))
	} else if field.Desc.IsMap() {
		funcStr = fmt.Sprint("Get", field.GoName, "() ",
			"map[", field.Desc.MapKey().Kind().String(), "]", singleType(field.Desc.MapValue()),
		)
	} else {
		// ???????????????
		funcStr = fmt.Sprint("Get", field.GoName, "() ", singleType(field.Desc))
	}

	g.P(field.Comments.Leading, funcStr, field.Comments.Trailing)
}

func singleType(desc protoreflect.FieldDescriptor) string {
	if desc.Message() == nil {
		return fmt.Sprint("*", desc.Kind().String())
	}

	if desc.Message().ParentFile() == nil {
		return fmt.Sprint("*", desc.Kind().String())
	}

	options, ok := desc.Message().ParentFile().Options().(*descriptorpb.FileOptions)
	if !ok {
		return fmt.Sprint("*", desc.Kind().String())
	}

	if goPackage == options.GetGoPackage() || options.GetGoPackage() == "" {
		return fmt.Sprint("*", desc.Kind().String())
	}

	return fmt.Sprint("*", getGoPackage(options.GetGoPackage()).GetName(), ".", desc.Message().Name())
}

func getGoPackage(contest string) *importPackage {
	ss := strings.Split(contest, "/")
	path := strings.Join(ss[:len(ss)-1], "/")

	if pack, ok := importSet[path]; ok {
		return pack
	}

	ns := strings.Split(ss[len(ss)-1], ";")
	name := ns[0]
	if len(ns) > 1 {
		name = ns[1]
	}
	pack := &importPackage{
		name: name,
		path: path,
	}
	importSet[pack.path] = pack

	return pack
}
