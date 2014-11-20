package javascript

import (
	"fmt"

	"veyron.io/veyron/veyron2/vdl"
	"veyron.io/veyron/veyron2/vdl/vdlutil"
)

// makeTypeDefinitionsString generates a string that defines the specified types.
// It consists of the following sections:
// - Definitions. e.g. "var _typeNamedBool = new Type();"
// - Field assignments. e.g. "_typeNamedBool.name = \"NamedBool\""
// - Constructor definitions. e.g. "types.NamedBool = Registry.lookupOrCreateConstructor(_typeNamedBool)"
func makeTypeDefinitionsString(jsnames typeNames) string {
	str := ""
	sortedDefs := jsnames.SortedList()

	for _, def := range sortedDefs {
		str += makeDefString(def.Name)
	}

	for _, def := range sortedDefs {
		str += makeTypeFieldAssignmentString(def.Name, def.Type, jsnames)
	}

	for _, def := range sortedDefs {
		if def.Type.Name() != "" {
			str += makeConstructorDefinitionString(def.Name, def.Type)
		}
	}

	return str
}

// makeDefString generates a type definition for the specified type name.
// e.g. "var _typeNamedBool = new Type();"
func makeDefString(jsname string) string {
	return fmt.Sprintf("var %s = new Type();\n", jsname)
}

// makeTypeFieldAssignmentString generates assignments for type fields.
// e.g. "_typeNamedBool.name = \"NamedBool\""
func makeTypeFieldAssignmentString(jsname string, t *vdl.Type, jsnames typeNames) string {
	// kind
	str := fmt.Sprintf("%s.kind = %s;\n", jsname, jsKind(t.Kind()))

	// name
	if t.Name() != "" {
		str += fmt.Sprintf("%s.name = %q;\n", jsname, t.Name())
	}

	// labels
	if t.Kind() == vdl.Enum {
		str += fmt.Sprintf("%s.labels = [", jsname)
		for i := 0; i < t.NumEnumLabel(); i++ {
			if i > 0 {
				str += ", "
			}
			str += fmt.Sprintf("%q", t.EnumLabel(i))
		}
		str += "];\n"
	}

	// len
	if t.Kind() == vdl.Array { // Array is the only type where len is valid.
		str += fmt.Sprintf("%s.len = %d;\n", jsname, t.Len())
	}

	// elem
	switch t.Kind() {
	case vdl.Optional, vdl.Array, vdl.List, vdl.Map:
		str += fmt.Sprintf("%s.elem = %s;\n", jsname, jsnames.LookupName(t.Elem()))
	}

	// key
	switch t.Kind() {
	case vdl.Set, vdl.Map:
		str += fmt.Sprintf("%s.key = %s;\n", jsname, jsnames.LookupName(t.Key()))
	}

	// fields
	if t.Kind() == vdl.Struct {
		str += fmt.Sprintf("%s.fields = [", jsname)
		for i := 0; i < t.NumField(); i++ {
			if i > 0 {
				str += ", "
			}
			field := t.Field(i)
			str += fmt.Sprintf("{name: %q, type: %s}", vdlutil.ToCamelCase(field.Name), jsnames.LookupName(field.Type))
		}
		str += "];\n"
	}

	// types
	if t.Kind() == vdl.OneOf {
		str += fmt.Sprintf("%s.types = [", jsname)
		for i := 0; i < t.NumOneOfType(); i++ {
			if i > 0 {
				str += ", "
			}
			str += fmt.Sprintf("%s", jsnames.LookupName(t.OneOfType(i)))
		}
		str += "];\n"
	}

	return str
}

// makeConstructorDefinitionString creates a string that defines the constructor for the type.
// e.g. "types.NamedBool = RegistrylookupOrCreateConstructor(_typeNamedBool)"
func makeConstructorDefinitionString(jsname string, t *vdl.Type) string {
	_, name := vdl.SplitIdent(t.Name())
	return fmt.Sprintf("types.%s = Registry.lookupOrCreateConstructor(%s, %q);\n", name, jsname, name)
}

func jsKind(k vdl.Kind) string {
	switch k {
	case vdl.Any:
		return "Kind.ANY"
	case vdl.OneOf:
		return "Kind.ONEOF"
	case vdl.Optional:
		return "Kind.NILABLE"
	case vdl.Bool:
		return "Kind.BOOL"
	case vdl.Byte:
		return "Kind.BYTE"
	case vdl.Uint16:
		return "Kind.UINT16"
	case vdl.Uint32:
		return "Kind.UINT32"
	case vdl.Uint64:
		return "Kind.UINT64"
	case vdl.Int16:
		return "Kind.INT16"
	case vdl.Int32:
		return "Kind.INT32"
	case vdl.Int64:
		return "Kind.INT64"
	case vdl.Float32:
		return "Kind.FLOAT32"
	case vdl.Float64:
		return "Kind.FLOAT64"
	case vdl.Complex64:
		return "Kind.COMPLEX64"
	case vdl.Complex128:
		return "Kind.COMPLEX128"
	case vdl.String:
		return "Kind.STRING"
	case vdl.Enum:
		return "Kind.ENUM"
	case vdl.TypeObject:
		return "Kind.TYPEOBJECT"
	case vdl.Array:
		return "Kind.ARRAY"
	case vdl.List:
		return "Kind.LIST"
	case vdl.Set:
		return "Kind.SET"
	case vdl.Map:
		return "Kind.MAP"
	case vdl.Struct:
		return "Kind.STRUCT"
	}
	panic(fmt.Errorf("val: unhandled kind: %d", k))
}
