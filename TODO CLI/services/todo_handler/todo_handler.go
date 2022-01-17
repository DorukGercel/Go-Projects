package todo_handler

import (
	"TODO_CLI/definitions"
	"TODO_CLI/services/file_handler"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"
)

type Service int

type StrServiceArg interface{}

type record struct {
	id       string
	todo     string
	dateDiff string
	complete bool
}

const todoVersion = "1.0.0"

// Desired operation
const (
	Version = iota
	ListAll
	ListCompleted
	AddNew
	MarkComplete
	Delete
	None = -1
)

// List all or list completed
type listOpt int

// Handle performs the required todo action
func Handle(service Service, arg StrServiceArg) {
	switch service {
	case Version:
		printVersion()
	case ListAll:
		metas, records, err := file_handler.ReadRecord()
		if err != nil {
			printError(err)
		}
		printTable(metas, records, ListAll)
	case ListCompleted:
		metas, records, err := file_handler.ReadRecord()
		if err != nil {
			printError(err)
		}
		printTable(metas, records, ListCompleted)
	case AddNew:
		strArg, ok := arg.(string)
		if !ok {
			printError(errors.New(definitions.ErrNarg))
		}
		err := file_handler.WriteRecord(strArg)
		if err != nil {
			printError(err)
			return
		}
		printSuccess(definitions.SucAdd, strArg)
	case MarkComplete:
		strArg, ok := arg.(string)
		if !ok {
			printError(errors.New(definitions.ErrNarg))
		}
		err := file_handler.MarkRecordComplete(strArg)
		if err != nil {
			printError(err)
			return
		}
		printSuccess(definitions.SucMark, strArg)
	case Delete:
		strArg, ok := arg.(string)
		if !ok {
			printError(errors.New(definitions.ErrNarg))
		}
		err := file_handler.DeleteRecord(strArg)
		if err != nil {
			printError(err)
			return
		}
		printSuccess(definitions.SucDelete, strArg)
	}
}

// Version op
func printVersion() {
	fmt.Println("Version " + todoVersion)
}

// Print operation success
func printSuccess(msg string, arg string) {
	msgS := strings.Split(msg, " ")
	fmt.Println(strings.Title(msgS[0] + " " + arg + " " + msgS[1]))
}

// Print operation error
func printError(err error) {
	fmt.Println(strings.Title(err.Error()))
}

// Print table for records
func printTable(metas []string, records []string, opt listOpt) {
	w := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)
	fmt.Fprintln(w, "ID\tItem\tDate")
	for _, val := range records {
		if !isDeletedRecord(val) {
			rec := separateFields(val)
			if opt == ListAll {
				fmt.Fprintln(w, rec.id+"\t"+rec.todo+"\t"+rec.dateDiff)
			} else if opt == ListCompleted && rec.complete {
				fmt.Fprintln(w, rec.id+"\t"+rec.todo+"\t"+rec.dateDiff)
			}
		}
	}
	w.Flush()
	if opt == ListAll {
		n := metas[file_handler.NoRecordsIndex]
		verbStr := "are"
		itemStr := "items"
		if n == "1" {
			verbStr = "is"
			itemStr = "item"
		}
		fmt.Println("\nThere " + verbStr + " " + n + " " + itemStr + " in your todo list.")
	} else if opt == ListCompleted {
		n := metas[file_handler.NoCompRecordsIndex]
		verbStr := "are"
		itemStr := "items"
		if n == "1" {
			verbStr = "is"
			itemStr = "item"
		}
		fmt.Println("\nThere " + verbStr + " " + n + " completed " + itemStr + " in your todo list.")
	}
}

// Separate string record to its fields
func separateFields(recStr string) record {
	rec := record{}
	rs := strings.Split(recStr, file_handler.FieldSpr)
	rec.id = rs[file_handler.IdIndex]
	rec.dateDiff = getDateDiff(rs[file_handler.DateIndex])
	if rec.complete = false; rs[file_handler.CompleteIndex] == "1" {
		rec.complete = true
	}
	l, _ := strconv.Atoi(rs[file_handler.StrLenIndex])
	lenRs := len(rs)
	todo := ""
	for i := file_handler.StrIndex; i < lenRs; i++ {
		todo = todo + rs[i]
		if len(todo) == l {
			rec.todo = todo
			break
		}
	}
	return rec
}

// Calculate diff between two dates
func getDateDiff(date string) string {
	oldDate, _ := time.Parse(file_handler.DateLayout, date)
	oldY, oldM, oldD := oldDate.Date()
	newY, newM, newD := time.Now().Date()
	if yDif := newY - oldY; yDif > 0 {
		if yDif == 1 {
			return "Year ago"
		} else {
			return strconv.Itoa(yDif) + " years ago"
		}
	}
	if mDif := newM - oldM; mDif > 0 {
		if mDif == 1 {
			return "Month ago"
		}
	}
	if dDif := newD - oldD; dDif >= 7 {
		if dDif >= 7 && dDif < 14 {
			return "Week ago"
		} else {
			return strconv.Itoa(dDif/7) + " weeks ago"
		}
	}
	if dDif := newD - oldD; dDif > 0 {
		if dDif == 1 {
			return "Yesterday"
		} else {
			return strconv.Itoa(dDif) + " days ago"
		}
	}
	return "Today"
}

// Check if record is deleted
func isDeletedRecord(rec string) bool {
	return file_handler.DeletedRecMark == string(rec[0])
}
