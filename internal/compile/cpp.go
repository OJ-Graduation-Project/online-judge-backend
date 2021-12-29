package compile

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	"github.com/OJ-Graduation-Project/online-judge-backend/pkg/responses"
)

func compile(code string, submissionId string) error {
	wd, _ := os.Getwd()
	f, err := os.Create(wd + "/" + submissionId + ".cpp")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err2 := f.WriteString(code)

	if err2 != nil {
		fmt.Println(err2)
		return err2
	}

	cmd := exec.Command("g++", "-o", submissionId+".out", submissionId+".cpp")
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()
	if err != nil {
		fmt.Println("Error couldn't combile file")
	}
	output := stdout.String()
	erro := stderr.String()

	if len(output) > 0 || len(erro) > 0 {
		return err
	}
	return nil
}

func CompileAndRun(submissionId string, problemtestcases []responses.ProblemTestCases, code string, langugage string) responses.SubmissionResponse {

	CompileError := compile(code, submissionId)
	var response responses.SubmissionResponse

	if CompileError != nil {
		response.Verdict = "Compilation Error"
		return response
	}

	failed := false
	for i, v := range problemtestcases {
		if failed == true {
			break
		}
		//Problems with path will arise.
		cmd := exec.Command("./" + submissionId + ".out")

		var out bytes.Buffer
		b := []byte(v.Input)

		cmd.Stdin = bytes.NewBuffer(b)
		cmd.Stdout = &out

		err := cmd.Run()
		if err != nil {
			fmt.Println("Error couldn't run")
		}
		output := out.String()
		output = output[:len(output)-1]

		if output != v.ExpectedOutput {

			response.Verdict = "Wrong"
			response.WrongTestCase = i
			failed = true
		}

	}
	if failed == false {
		response.Verdict = "Correct"

	}
	return response

}
