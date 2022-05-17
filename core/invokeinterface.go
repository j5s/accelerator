package core

import (
	"fmt"
	"github.com/4ra1n/accelerator/classfile"
	"github.com/4ra1n/accelerator/global"
)

type INVOKEINTERFACE struct {
	index uint
}

func (self *INVOKEINTERFACE) FetchOperands(reader *BytecodeReader) {
	self.index = uint(reader.ReadUint16())
	reader.ReadUint8()
	reader.ReadUint8()
}

func (self *INVOKEINTERFACE) GetOperands() []string {
	name := global.CP.GetConstantInfo(uint16(self.index))
	className := name.(*classfile.ConstantMethodRefInfo).ClassName()
	methodName, desc := name.(*classfile.ConstantMethodRefInfo).NameAndDescriptor()
	ret := make([]string, 1)
	out := fmt.Sprintf("%s %s %s", className, methodName, desc)
	ret[0] = out
	return ret
}