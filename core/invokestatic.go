package core

import (
	"fmt"
	"github.com/4ra1n/accelerator/classfile"
	"github.com/4ra1n/accelerator/global"
)

type INVOKESTATIC struct{ Index16Instruction }

func (self *INVOKESTATIC) GetOperands() []string {
	name := global.CP.GetConstantInfo(uint16(self.Index))
	className := name.(*classfile.ConstantMethodRefInfo).ClassName()
	methodName, desc := name.(*classfile.ConstantMethodRefInfo).NameAndDescriptor()
	ret := make([]string, 1)
	out := fmt.Sprintf("%s %s %s", className, methodName, desc)
	ret[0] = out
	return ret
}
