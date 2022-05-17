package ref

import (
	"github.com/4ra1n/accelerator/classfile"
)

type Field struct {
	ClassMember
}

func newFields(class *Class, cfFields []*classfile.MemberInfo) []*Field {
	fields := make([]*Field, len(cfFields))
	for i, cfField := range cfFields {
		fields[i] = &Field{}
		fields[i].class = class
		fields[i].copyMemberInfo(cfField)
	}
	return fields
}

func (self *Field) IsVolatile() bool {
	return 0 != self.accessFlags&ACC_VOLATILE
}

func (self *Field) IsTransient() bool {
	return 0 != self.accessFlags&ACC_TRANSIENT
}

func (self *Field) IsEnum() bool {
	return 0 != self.accessFlags&ACC_ENUM
}
