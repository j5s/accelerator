package main

import (
	"fmt"
	"github.com/4ra1n/accelerator/classfile"
	"github.com/4ra1n/accelerator/core"
	"github.com/4ra1n/accelerator/files"
	"github.com/4ra1n/accelerator/global"
	"github.com/4ra1n/accelerator/ref"
	"io/ioutil"
	"strings"
)

func main() {
	files.RemoveTempFiles()
	files.UnzipJars(".")
	classes := files.ReadAllClasses()
	for _, c := range classes {
		start(c)
	}
}

func start(class string) {
	data, err := ioutil.ReadFile(class)
	cf, err := classfile.Parse(data)
	if err != nil {
		panic(err)
	}

	cl := ref.NewClass(cf)
	global.CP = cf.ConstantPool()

	for _, method := range cf.Methods() {
		codeAttr := method.CodeAttribute()
		if codeAttr == nil {
			continue
		}
		bytecode := codeAttr.Code()
		thread := core.Thread{}
		reader := &core.BytecodeReader{}
		instSet := core.InstructionSet{}
		instSet.ClassName = cl.Name()
		instSet.MethodName = method.Name()
		instSet.Desc = method.Descriptor()

		for {
			if thread.PC() >= len(bytecode) {
				break
			}
			reader.Reset(bytecode, thread.PC())
			opcode := reader.ReadUint8()
			inst := core.NewInstruction(opcode)
			inst.FetchOperands(reader)
			ops := inst.GetOperands()

			instEntry := core.InstructionEntry{
				Instrument: getInstructionName(inst),
				Operands:   ops,
			}
			instSet.InstArray = append(instSet.InstArray, instEntry)

			thread.SetPC(reader.PC())
		}

		fmt.Println(instSet)
	}
}

func getInstructionName(instruction core.Instruction) string {
	i := fmt.Sprintf("%T", instruction)
	return strings.Split(i, ".")[1]
}
