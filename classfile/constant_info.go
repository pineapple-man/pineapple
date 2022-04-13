package classfile

const (
	// CONSTANT_Utf8_info UTF-8 编码的字符串
	CONSTANT_Utf8_info = 1

	// CONSTANT_Integer_info 整形字面量
	CONSTANT_Integer_info = 3

	// CONSTANT_Float_info 浮点型字面量
	CONSTANT_Float_info = 4

	// CONSTANT_Long_info 长整型字面量
	CONSTANT_Long_info = 5

	// CONSTANT_Double_info 双精度浮点型字面量
	CONSTANT_Double_info = 6

	// CONSTANT_Class_info 类或接口的符号引用
	CONSTANT_Class_info = 7

	// CONSTANT_String_info 字符串类型的字面量
	CONSTANT_String_info = 8

	// CONSTANT_Fieldref_info 字段的符号引用
	CONSTANT_Fieldref_info = 9

	// CONSTANT_Methodref_info 类中方法的符号引用
	CONSTANT_Methodref_info = 10

	// CONSTANT_Interfacemethodref_info 接口中方法的符号引用
	CONSTANT_Interfacemethodref_info = 11

	// CONSTANT_NameAndType_info 字段或方法的符号引用
	CONSTANT_NameAndType_info = 12

	// CONSTANT_MethodHandle_info 表示方法句柄
	CONSTANT_MethodHandle_info = 15

	// CONSTANT_MethodType_info 标志方法类型
	CONSTANT_MethodType_info = 16
	// CONSTANT_Dynamic_info 表示一个动态计算常量
	CONSTANT_Dynamic_info = 17
	// CONSTANT_InvokeDynamic_info 表示一个动态方法调用点
	CONSTANT_InvokeDynamic_info = 18
	// CONSTANT_Module_info 表示一个模块
	CONSTANT_Module_info = 19
	// CONSTANT_Package_info 表示一个模块中开放后者导出的包
	CONSTANT_Package_info = 20
)
