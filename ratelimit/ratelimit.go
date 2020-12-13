package ratelimit

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

var GlobalRateLimit string

type User struct {
	RemAddress string
	ReqRemain  string
}

func (u *User) MakeNewLog() error {

	err := os.Chdir("data")
	fileName := u.RemAddress + "_log.tmp"

	emptyFile, err := os.Create(fileName)
	if err != nil {
		// fmt.Printf("MakeNewLog: Error creating log for user %s\n%v\n", u.RemAddress, err)
		// fmt.Printf("Warning: Rate limiting may be diabled.\n")
	}

	t := time.Now()
	tStr := t.String()

	w := bufio.NewWriter(emptyFile)
	w.WriteString(tStr + " " + u.RemAddress + " -1" + "\n")

	fmt.Printf("Created new log for user %s.\n", u.RemAddress)

	emptyFile.Close()
	
	fileErr := os.Chdir("..")
	if fileErr != nil {
		fmt.Printf("ReadFile: Error switching to root dir %v\n", fileErr)
	}

	return err
}

func (u *User) CheckLog() (int, error) {
	//open the file
	readFile(u.RemAddress)
	//check last log time and reqRemain
	//if time > 1 min ago; clear all log
	//if time < min ago:
	//	if rate < limit:
	//		write new log and close file
	//	if rate = limit:
	// display wait message

	fileName := u.RemAddress + "_log.tmp"
	_, err := os.Open(fileName)
	// lastEntry := 0
	reqRemaining := 0

	return reqRemaining, err
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

func readFile(fname string) {
	fileErr := os.Chdir("data")
	if fileErr != nil {
		fmt.Printf("ReadFile: Error switching to data dir %v\n", fileErr)
	}

	file, err := os.Open(fname)
	if err != nil {
		fmt.Printf("Cant find file '%v'.\n%v\n", fname, err)
	}
	defer file.Close()

	buf := make([]byte, 62)
	stat, err := os.Stat(fname)
	start := stat.Size() - 62
	_, err = file.ReadAt(buf, start)
	if err == nil {
		fmt.Printf("Printing file.......\n%s\n", buf)
	}

	fileErr = os.Chdir("..")
	if fileErr != nil {
		fmt.Printf("ReadFile: Error switching to root dir %v\n", fileErr)
	}
}
