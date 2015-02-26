package java

import (
	"bytes"
	"fmt"
	"log"

	"v.io/v23/vdl"
	"v.io/v23/vdl/compile"
)

const complexTmpl = `
// This file was auto-generated by the veyron vdl tool.
// Source: {{.Source}}
package {{.PackagePath}};

/**
 * type {{.Name}} {{.VdlTypeString}} {{.Doc}}
 **/
@io.v.v23.vdl.GeneratedFromVdl(name = "{{.VdlTypeName}}")
{{ .AccessModifier }} final class {{.Name}} extends {{.VdlComplex}} {
    private static final long serialVersionUID = 1L;

    public static final io.v.v23.vdl.VdlType VDL_TYPE =
            io.v.v23.vdl.Types.getVdlTypeFromReflect({{.Name}}.class);

    public {{.Name}}({{.ValueType}} real, {{.ValueType}} imag) {
        super(VDL_TYPE, real, imag);
    }

    public {{.Name}}({{.ValueType}} real) {
        this(real, 0);
    }

    public {{.Name}}() {
        this(0, 0);
    }
}
`

// genJavaComplexFile generates the Java class file for the provided user-defined VDL complex type.
func genJavaComplexFile(tdef *compile.TypeDef, env *compile.Env) JavaFileInfo {
	var ValueType string
	switch kind := tdef.Type.Kind(); kind {
	case vdl.Complex64:
		ValueType = "float"
	case vdl.Complex128:
		ValueType = "double"
	default:
		panic(fmt.Errorf("val: unhandled kind: %v", kind))
	}
	javaTypeName := toUpperCamelCase(tdef.Name)
	data := struct {
		AccessModifier string
		Doc            string
		Name           string
		PackagePath    string
		Source         string
		ValueType      string
		VdlComplex     string
		VdlTypeName    string
		VdlTypeString  string
	}{
		AccessModifier: accessModifierForName(tdef.Name),
		Doc:            javaDocInComment(tdef.Doc),
		Name:           javaTypeName,
		PackagePath:    javaPath(javaGenPkgPath(tdef.File.Package.GenPath)),
		Source:         tdef.File.BaseName,
		ValueType:      ValueType,
		VdlComplex:     javaVdlPrimitiveType(tdef.Type.Kind()),
		VdlTypeName:    tdef.Type.Name(),
		VdlTypeString:  tdef.Type.String(),
	}
	var buf bytes.Buffer
	err := parseTmpl("complex", complexTmpl).Execute(&buf, data)
	if err != nil {
		log.Fatalf("vdl: couldn't execute VDL complex template: %v", err)
	}
	return JavaFileInfo{
		Name: javaTypeName + ".java",
		Data: buf.Bytes(),
	}
}
