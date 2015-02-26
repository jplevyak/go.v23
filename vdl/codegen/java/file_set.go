package java

import (
	"bytes"
	"log"

	"v.io/v23/vdl/compile"
)

const setTmpl = `// This file was auto-generated by the veyron vdl tool.
// Source: {{.SourceFile}}

package {{.Package}};

/**
 * {{.Name}} {{.VdlTypeString}} {{.Doc}}
 **/
@io.v.v23.vdl.GeneratedFromVdl(name = "{{.VdlTypeName}}")
{{ .AccessModifier }} final class {{.Name}} extends io.v.v23.vdl.VdlSet<{{.KeyType}}> {
	private static final long serialVersionUID = 1L;

    public static final io.v.v23.vdl.VdlType VDL_TYPE =
            io.v.v23.vdl.Types.getVdlTypeFromReflect({{.Name}}.class);

    public {{.Name}}(java.util.Set<{{.KeyType}}> impl) {
        super(VDL_TYPE, impl);
    }

    public {{.Name}}() {
        this(new java.util.HashSet<{{.KeyType}}>());
    }
}
`

// genJavaSetFile generates the Java class file for the provided named set type.
func genJavaSetFile(tdef *compile.TypeDef, env *compile.Env) JavaFileInfo {
	javaTypeName := toUpperCamelCase(tdef.Name)
	data := struct {
		AccessModifier string
		Doc            string
		KeyType        string
		Name           string
		Package        string
		SourceFile     string
		VdlTypeName    string
		VdlTypeString  string
	}{
		AccessModifier: accessModifierForName(tdef.Name),
		Doc:            javaDocInComment(tdef.Doc),
		KeyType:        javaType(tdef.Type.Key(), true, env),
		Name:           javaTypeName,
		Package:        javaPath(javaGenPkgPath(tdef.File.Package.GenPath)),
		SourceFile:     tdef.File.BaseName,
		VdlTypeName:    tdef.Type.Name(),
		VdlTypeString:  tdef.Type.String(),
	}
	var buf bytes.Buffer
	err := parseTmpl("set", setTmpl).Execute(&buf, data)
	if err != nil {
		log.Fatalf("vdl: couldn't execute set template: %v", err)
	}
	return JavaFileInfo{
		Name: javaTypeName + ".java",
		Data: buf.Bytes(),
	}
}
