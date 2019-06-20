package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

//Функция логирования Otus
//Задача: написать функцию логирования LogOtusEvent, на вход которой приходят события типа HwAccepted
// (домашняя работа принята) и HwSubmitted (студент отправил дз) для этого - спроектировать и реализовать интерфейс OtusEvent.
// Для события HwAccepted мы хотим логирровать дату, айди и грейд, для HwSubmitter - дату, id и комментарий, например:

//2019-01-01 submitted 3456 "please take a look at my homework"
//2019-01-01 accepted 3456 4


type HwAccepted struct {
	Id int
	Grade int
}

type HwSubmitted struct {
	Id int
	Code string
	Comment string
}

type OtusEvent interface {
	log() string
}

func (accept HwAccepted) log() string{
	return fmt.Sprintf("[DATE:%v] [ACTION_TYPE:%v] [ID:%v] Grade:%v\n", time.Now().Format("2006-01-02 15:04:05"), "Accepted", accept.Id, accept.Grade)
}

func (submit HwSubmitted) log() string{
	return fmt.Sprintf("[DATE:%v] [ACTION_TYPE:%v] [ID:%v] Comment:%v\n", time.Now().Format("2006-01-02 15:04:05"), "Submitted", submit.Id, submit.Comment)
}

func LogOtusEvent(e OtusEvent, w io.Writer) {
	_, err := w.Write([]byte(e.log()))
	if err != nil{
		fmt.Println(err)
	}
}
type my_writer struct {
	slice_of_bytes [][]byte
}

func (w *my_writer) Write(b []byte) (n int, err error)  {
	if _, err := os.Stat("logs.log"); os.IsNotExist(err) {
		file, err := os.Create("logs.log")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		n, err = file.Write(b)
		if err != nil{
			fmt.Println(err)
		}
		defer file.Close()

	}else {
		file, err := os.OpenFile("logs.log", os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		n, err = file.Write(b)
		if err != nil{
			fmt.Println(err)
		}
		defer file.Close()
	}
	n, err = os.Stdout.Write(b)
	if err != nil{
		fmt.Println(err)
	}

	w.slice_of_bytes = append(w.slice_of_bytes, b)

	return n, err
}

func main() {
	//writer реализующий интерфейс io.Writer умеющий печатать в файл, в stdOut и в Slice
	writer := &my_writer{slice_of_bytes:make([][]byte,0)}

	LogOtusEvent(HwAccepted{Id:1,Grade:0}, writer)
	LogOtusEvent(HwAccepted{Id:2,Grade:1}, writer)
	LogOtusEvent(HwAccepted{Id:3,Grade:2}, writer)
	LogOtusEvent(HwAccepted{Id:4,Grade:3}, writer)
	LogOtusEvent(HwSubmitted{Id:1,Code:"Code 1",Comment:"LOL"}, writer)
	LogOtusEvent(HwSubmitted{Id:2,Code:"Code 2",Comment:"is"}, writer)
	LogOtusEvent(HwSubmitted{Id:3,Code:"Code 3",Comment:"the"}, writer)
	LogOtusEvent(HwSubmitted{Id:4,Code:"Code 4",Comment:"best"}, writer)
	LogOtusEvent(HwSubmitted{Id:5,Code:"Code 5",Comment:"game!"}, writer)

	for i,val := range writer.slice_of_bytes{
		fmt.Printf("%v : %v",i, string(val))
	}


}
