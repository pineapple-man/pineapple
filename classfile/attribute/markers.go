package attribute

import (
	"pineapple/classfile"
)

/*Deprecated_attribute {
	u2 attribute_name_index;
	u4 attribute_length;
}*/

type DeprecatedAttribute struct {
	MarketAttribute
}

/*Synthetic_attribute {
	u2 attribute_name_index;
	u4 attribute_length;
}*/

type SyntheticAttribute struct {
	MarketAttribute
}

type MarketAttribute struct {
}

func (m *MarketAttribute) readInfo(reader *classfile.ClassReader) {
	// do nothing
}
