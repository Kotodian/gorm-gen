package main

import "strings"

func (g *Generator) buildScope(pkg string, s *Struct) {
	for _, field := range s.fields {
		if t := field.tags.Exists("query"); t != nil {
			value := t.Value("query")
			if value == "" {
				continue
			}
			for _, v := range strings.Split(value, ",") {
				if v == "equal" {
					g.Printf("// Query%sBy%s query object by one condition\n", s.name, field.name)
					g.Printf("func Query%sBy%s(ctx context.Context, db *gorm.DB, %s) ([]*%s.%s, error) {\n", s.name, field.name, field.String(), pkg, s.name)
					g.Printf("res := make([]*%s.%s, 0)\n", pkg, s.name)
					g.Printf("err := db.WithContext(ctx).Where(\"%s = ?\", %s).Find(&res).Error\n", field.DB(), firstLower(field.name))
					g.Printf("return res, err\n")
					g.Printf("}\n")
				} else if v == "scope" {
					switch field.goType {
					case "int64", "int32", "int",
						"uint64", "uint32", "uint",
						"float64", "float32":
						g.Printf("// Query%sBetween%s query object by one condition\n", s.name, field.name)
						g.Printf("func Query%sBetween%s(ctx context.Context, db *gorm.DB, min %s, max %s) ([]*%s.%s, error) {\n", s.name, field.name, field.goType, field.goType, pkg, s.name)
						g.Printf("res := make([]*%s.%s, 0)\n", pkg, s.name)
						g.Printf("err := db.WithContext(ctx).Where(\"%s >= ?\", min).Where(\"%s <= ?\", max).Find(&res).Error\n", field.DB(), field.DB())
						g.Printf("return res, err\n")
						g.Printf("}\n")
					default:
						g.Printf("// Query%sIn%s query object by one condition\n", s.name, field.name)
						g.Printf("func Query%sIn%s(ctx context.Context, db *gorm.DB, %s) ([]*%s.%s, error) {\n", s.name, field.name, field.Strings(), pkg, s.name)
						g.Printf("res := make([]*%s.%s, 0)\n", pkg, s.name)
						g.Printf("err := db.WithContext(ctx).Where(\"%s in ?\", %s).Find(&res).Error\n", field.DB(), firstLower(field.name))
						g.Printf("return res, err\n")
						g.Printf("}\n")
					}
				}
			}
		}
	}
}
