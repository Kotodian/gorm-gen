package main

type File struct {
	// 文件路径
	dir string
	// 文件名称
	name string
	// 这个文件属于哪个包
	pkg string
	// 这个文件包含的所有结构体
	structs []*Struct
	// 这个文件的输出路径
	output string
}
