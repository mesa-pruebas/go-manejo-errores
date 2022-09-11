package main

import "fmt"
import "os"

func RecordTaskDone(taskName string){
	f, err := os.OpenFile("task.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
		panic(err)
	}
	_, err = f.WriteString(taskName)
    if err != nil {
	    panic(err)
    }
    defer f.Close()
}

func WritingTest(path string) error {
	fmt.Println("writing:", path)
	data := []byte("This is the data for "+path)
	err := os.WriteFile(path, data, 0644)
	if(err == nil){
		RecordTaskDone("wrote:" + path)
	}
	return err;
}

func main() {
	fmt.Println("Starting tests")
	os.WriteFile("task.log", []byte("\n==== report ====\n"), 0644)
	var err error
	if err == nil { err = WritingTest("data.txt") }
	if err == nil { err = WritingTest(fmt.Sprintf("nonexistent-dir%cdata.txt", os.PathSeparator)) }
	if err == nil { err = WritingTest(fmt.Sprintf("other-nonexistent-dir%cdata.txt", os.PathSeparator)) }
	report, _ := os.ReadFile("task.log")
	fmt.Println(string(report))
}

