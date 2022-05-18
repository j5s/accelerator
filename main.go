package main

import (
	"flag"
	"fmt"
	"github.com/4ra1n/accelerator/classfile"
	"github.com/4ra1n/accelerator/core"
	"github.com/4ra1n/accelerator/files"
	"github.com/4ra1n/accelerator/global"
	"github.com/4ra1n/accelerator/ref"
	"github.com/4ra1n/accelerator/rule"
	"io/ioutil"
	"strings"
)

var (
	ruleFile string
	jarsDir  string
)

func main() {
	// resolve input
	flag.StringVar(&ruleFile, "rule", "rule.txt", "rule file")
	flag.StringVar(&jarsDir, "jars", ".", "your jars dir")
	flag.Parse()
	// delete temp dir
	files.RemoveTempFiles()
	// extract all jars to temp dir
	files.UnzipJars(jarsDir)
	// traverse to get all class files
	classes := files.ReadAllClasses()
	for _, c := range classes {
		// analyze each class file separately
		start(c)
	}
	files.RemoveTempFiles()
}

func start(class string) {
	data, err := ioutil.ReadFile(class)
	cf, err := classfile.Parse(data)
	if err != nil {
		panic(err)
	}
	cl := ref.NewClass(cf)
	// constant pool cache
	// refresh cp in each new class file
	global.CP = cf.ConstantPool()
	for _, method := range cf.Methods() {
		codeAttr := method.CodeAttribute()
		if codeAttr == nil {
			// interface or abstract
			continue
		}
		bytecode := codeAttr.Code()
		// virtual thread
		thread := core.Thread{}
		reader := &core.BytecodeReader{}
		// save all instructions to struct
		instSet := core.InstructionSet{}
		instSet.ClassName = cl.Name()
		instSet.MethodName = method.Name()
		instSet.Desc = method.Descriptor()
		for {
			// read finish
			if thread.PC() >= len(bytecode) {
				break
			}
			// offset
			reader.Reset(bytecode, thread.PC())
			// read instruction
			opcode := reader.ReadUint8()
			inst := core.NewInstruction(opcode)
			// read operands of the instruction
			inst.FetchOperands(reader)
			ops := inst.GetOperands()
			instEntry := core.InstructionEntry{
				Instrument: getInstructionName(inst),
				Operands:   ops,
			}
			instSet.InstArray = append(instSet.InstArray, instEntry)
			// offset++
			thread.SetPC(reader.PC())
		}
		doAnalysis(instSet)
	}
}

func getInstructionName(instruction core.Instruction) string {
	// type name -> instruction name
	i := fmt.Sprintf("%T", instruction)
	return strings.Split(i, ".")[1]
}

func doAnalysis(instSet core.InstructionSet) {
	analysisRule := rule.GetRule(ruleFile)
	if len(analysisRule) == 0 {
		panic("error rule file")
	}
	ruleLen := len(analysisRule)
	var ruleLenState []bool
	i := 0
	for _, inst := range instSet.InstArray {
		if inst.Instrument != analysisRule[i][0] {
			continue
		}
		splits := strings.Split(inst.Operands[0], " ")
		stdName := splits[0]
		stdDesc := splits[1]
		if analysisRule[i][1] != stdName {
			continue
		}
		if analysisRule[i][2] == stdDesc ||
			analysisRule[i][2] == "*" {
			ruleLenState = append(ruleLenState, true)
			if i != ruleLen-1 {
				i++
			} else {
				break
			}
		}
	}
	if len(ruleLenState) == ruleLen {
		fmt.Println(instSet.ClassName, instSet.MethodName)
	}
}
