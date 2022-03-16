package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func writeLineToFile(file os.File, s string) {
	_, err := file.WriteString(s + "\n")

	if err != nil {
		fmt.Println(err)
		return
	}
}

func writeProgramStart(file os.File) {
	writeLineToFile(file, "#include <stdio.h>\n")
	writeLineToFile(file, "int main()")
	writeLineToFile(file, "{")
	writeLineToFile(file, "char array[30000] = {0};")
	writeLineToFile(file, "char *ptr = array;")
}

func main() {
	code_string := "++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++."
	code_slice := strings.Split(code_string, "")

	file, err := os.Create("transpiled.c")

	if err != nil {
		fmt.Println(err)
		return
	}

	writeProgramStart(*file)

	for _, token := range code_slice {
		if strings.Compare(token, ">") == 0 {
			writeLineToFile(*file, "++ptr;")
		} else if strings.Compare(token, "<") == 0 {
			writeLineToFile(*file, "--ptr;")
		} else if strings.Compare(token, "+") == 0 {
			writeLineToFile(*file, "++*ptr;")
		} else if strings.Compare(token, "-") == 0 {
			writeLineToFile(*file, "--*ptr;")
		} else if strings.Compare(token, ".") == 0 {
			writeLineToFile(*file, "putchar(*ptr);")
		} else if strings.Compare(token, ",") == 0 {
			writeLineToFile(*file, "*ptr = getchar();")
		} else if strings.Compare(token, "[") == 0 {
			writeLineToFile(*file, "while (*ptr) {")
		} else if strings.Compare(token, "]") == 0 {
			writeLineToFile(*file, "}")
		} else {
			fmt.Println("Invalid token", token, "detected, aborting ...")
			return
		}
	}

	writeLineToFile(*file, "}")

	cmd := exec.Command("gcc", "transpiled.c")
	err = cmd.Run()

	if err != nil {
		log.Fatal(err)
	}

	return
}
