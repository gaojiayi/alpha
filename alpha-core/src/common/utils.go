package common

import (
	"alpha-core/src/config"
	"bytes"
	"crypto/rand"
	"fmt"
	"log"
	"os/exec"
)

func ExecPythonCommand(args []string) string {
	var outInfo bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(config.Python_home, args...)
	cmd.Stdout = &outInfo
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf(err.Error(), stderr.String())
	}
	log.Println("command:", "config.Python_home", "args:", args, "output:", outInfo.String())
	return outInfo.String()
}

func GetUUID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}
