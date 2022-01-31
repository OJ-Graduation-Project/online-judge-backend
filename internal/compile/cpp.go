package compile

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	"github.com/OJ-Graduation-Project/online-judge-backend/pkg/entities"
)

func getCompilingCommand(language string, submissionId string) *exec.Cmd {
	switch language {
	case "cpp":
		fmt.Println("compiling c++")
		return exec.Command("g++", "-o", submissionId+".out", submissionId+".cpp")
	case "java":
		return exec.Command("javac", submissionId+"/Main.java")
	case "c":
		return exec.Command("gcc", submissionId+".c", "-o", submissionId+".out")
	default: // newly supported languages will be inserted here
		return nil
	}
}
func getExecutionCommand(language string, submissionId string) *exec.Cmd {
	wd, _ := os.Getwd()
	switch language {
	case "cpp":
		return exec.Command("./" + submissionId + ".out")
	case "java":
		return exec.Command("java", wd+"/"+submissionId+"/Main.java")
	case "python":
		return exec.Command("python3", wd+"/"+submissionId+".py")
	case "c":
		return exec.Command("./" + submissionId + ".out")
	default:
		return nil
	}
}
func createCodeFile(code string, submissionId string, language string) error {
	wd, _ := os.Getwd()
	var f *os.File
	var err error = nil
	switch language {
	case "cpp":
		f, err = os.Create(wd + "/" + submissionId + ".cpp")
	case "java":
		err = os.Mkdir(wd+"/"+submissionId, 0775)
		if err != nil {
			fmt.Println("error in creating directory for the submission")
		}
		f, err = os.Create(wd + "/" + submissionId + "/Main.java")
	case "python":
		f, err = os.Create(wd + "/" + submissionId + ".py")
	case "c":
		f, err = os.Create(wd + "/" + submissionId + ".c")
	default: // newly supported languages will be inserted here
		fmt.Println("Language is not supported")
	}
	if err != nil {
		fmt.Println("error creating the code file: ", err)
		return err
	}
	defer f.Close()
	_, err2 := f.WriteString(code)

	if err2 != nil {
		fmt.Println(err2)
		return err2
	}

	return nil
}
func compile(code string, submissionId string, language string) error {

	cmd := getCompilingCommand(language, submissionId)
	if cmd == nil {
		fmt.Println("language is not supported")
	}
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println(cmd.Args)
		fmt.Println("Error couldn't compile file")
	}
	output := stdout.String()
	erro := stderr.String()

	if len(output) > 0 || len(erro) > 0 {
		return err
	}
	return nil
}

// returns (verdict, failed test case number, user output)
func CompileAndRun(submissionId string, problemtestcases []entities.TestCase, code string, language string) (string, int, string) {
	err := createCodeFile(code, submissionId, language)

	if err != nil {
		fmt.Println(err)
		fmt.Println("Could not parse the code into a file")
	}

	if language != "python" { // python should not compile
		CompileError := compile(code, submissionId, language)

		if CompileError != nil {
			return "Compilation Error", 0, ""
		}
	}
	fmt.Printf("%+v\n", problemtestcases)
	for i, v := range problemtestcases {

		var out bytes.Buffer
		b := []byte(v.Input)

		//Problems with path will arise.
		cmd := getExecutionCommand(language, submissionId)

		cmd.Stdin = bytes.NewBuffer(b)
		cmd.Stdout = &out

		err := cmd.Run()
		if err != nil {
			fmt.Println(cmd.Path)
			fmt.Println("Error couldn't run")
			return "Runtime Error", i, ""
		}
		output := out.String()

		/*if len(output) >= 1 {
			output = output[:len(output)-1]
		}*/
		fmt.Println(output, "this is the output")

		if output != v.Output {
			return "Wrong Answer", i, output
		}
	}
	return "Correct", 0, ""

}
