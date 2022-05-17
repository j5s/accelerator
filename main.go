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
	//start("Test.class")
}

func start(class string) {
	data, err := ioutil.ReadFile(class)
	cf, err := classfile.Parse(data)
	if err != nil {
		panic(err)
	}

	cl := ref.NewClass(cf)
	global.CP = cf.ConstantPool()
	fmt.Println(cl.Name())

	for _, method := range cf.Methods() {
		codeAttr := method.CodeAttribute()
		if codeAttr == nil {
			continue
		}
		bytecode := codeAttr.Code()
		thread := core.Thread{}
		reader := &core.BytecodeReader{}
		fmt.Println(method.Name())
		for {
			if thread.PC() >= len(bytecode) {
				fmt.Println()
				break
			}
			reader.Reset(bytecode, thread.PC())
			opcode := reader.ReadUint8()
			inst := core.NewInstruction(opcode)
			inst.FetchOperands(reader)
			ops := inst.GetOperands()
			instName := getInstructionName(inst)
			var out string
			if len(ops) == 0 {
				out = fmt.Sprintf("%s", instName)
			} else {
				out = fmt.Sprintf("%s %s", instName, ops[0])
			}
			fmt.Println(out)
			thread.SetPC(reader.PC())
		}
	}
}

func getInstructionName(instruction core.Instruction) string {
	i := fmt.Sprintf("%T", instruction)
	return strings.Split(i, ".")[1]
}
