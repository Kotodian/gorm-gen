package main

// buildList 生成gorm list
func (g *Generator) buildList(pkg string, str *Struct) {
	g.Printf("// List%s list objects\n", str.name)
	g.Printf("func %sList(ctx context.Context, db *gorm.DB) ([]*%s.%s, error) {\n", str.name, pkg, str.name)
	g.Printf("%ss := make([]*%s.%s, 0)\n", firstLower(str.name), pkg, str.name)
	g.Printf("err := db.WithContext(ctx).Where(\"deleted_at=0\").Find(&%ss).Error\n", firstLower(str.name))
	g.Printf("return %ss, err\n", firstLower(str.name))
	g.Printf("}\n")
}

// buildQuery 生成gorm query
func (g *Generator) buildQuery(pkg string, str *Struct) {
	g.Printf("// Query%s query objects by condition\n", str.name)
	g.Printf("func %sQuery(ctx context.Context, db *gorm.DB, args map[string]interface{}) ([]*%s.%s,error) {\n", str.name, pkg, str.name)
	g.Printf("%ss := make([]*%s.%s, 0)\n", firstLower(str.name), pkg, str.name)
	g.Printf("err := db.WithContext(ctx).Where(args).Where(\"deleted_at=0\").Find(&%ss).Error\n", firstLower(str.name))
	g.Printf("return %ss, err\n", firstLower(str.name))
	g.Printf("}\n")
}

// buildCreate 生成gorm create
func (g *Generator) buildCreate(pkg string, str *Struct) {
	g.Printf("// Create%s create object\n", str.name)
	g.Printf("func Create%s(ctx context.Context, db *gorm.DB, %s *%s.%s) error {\n", str.name, firstLower(str.name), pkg, str.name)
	g.Printf("return db.WithContext(ctx).Create(%s).Error\n", firstLower(str.name))
	g.Printf("}\n")
}

// buildUpdate 生成gorm update
func (g *Generator) buildUpdate(pkg string, str *Struct) {
	g.Printf("// Update%s update object\n", str.name)
	g.Printf("func Update%s(ctx context.Context, db *gorm.DB, args map[string]interface{}) error {\n", str.name)
	g.Printf("return db.WithContext(ctx).Model(&%s.%s{}).Updates(args).Error\n", pkg, str.name)
	g.Printf("}\n")
}

// buildDelete 生成gorm delete
func (g *Generator) buildDelete(pkg string, str *Struct) {
	g.Printf("// Delete%s delete object\n", str.name)
	g.Printf("func Delete%s(ctx context.Context, db *gorm.DB, %s %s.%s) error {\n", str.name, firstLower(str.name), pkg, str.name)
	g.Printf("return db.WithContext(ctx).Delete(&%s).Error", firstLower(str.name))
	g.Printf("}\n")
}
