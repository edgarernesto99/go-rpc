package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
)

var subjects map[string]map[string]float64 = make(map[string]map[string]float64)
var students map[string]map[string]float64 = make(map[string]map[string]float64)

type Server struct{}

type Args struct {
	Subject string
	Student string
	Grade   float64
}

func (this *Server) SetGrade(args Args, reply *bool) error {
	_, subjectExists := subjects[args.Subject]
	_, isStudent := subjects[args.Subject][args.Student]
	_, studentExists := students[args.Student]
	if subjectExists && isStudent {
		return errors.New("El estudiante ya tiene calificación en " + args.Subject)
	}
	if !subjectExists {
		newStudent := make(map[string]float64)
		newStudent[args.Student] = args.Grade
		subjects[args.Subject] = newStudent
	} else {
		subjects[args.Subject][args.Student] = args.Grade
	}
	if !studentExists {
		newSubject := make(map[string]float64)
		newSubject[args.Subject] = args.Grade
		students[args.Student] = newSubject
	} else {
		students[args.Student][args.Subject] = args.Grade
	}
	return nil
}

func (this *Server) GetStudentAverage(student string, reply *float64) error {
	_, studentExists := students[student]
	if !studentExists {
		return errors.New("El alumno no existe")
	}
	var average float64 = 0
	var total float64 = 0
	for _, grade := range students[student] {
		average = average + grade
		total = total + 1
	}
	average = average / total
	*reply = average
	return nil
}

func (this *Server) GetGeneralAverage(args Args, reply *float64) error {
	var generalAverage float64 = 0
	var generalTotal float64 = 0
	for student := range students {
		var average float64 = 0
		total := 0
		generalTotal = generalTotal + 1
		for _, grade := range students[student] {
			average = average + grade
			total = total + 1
		}
		average = average / float64(total)
		generalAverage = generalAverage + average
	}
	if generalTotal == 0 {
		return errors.New("No hay estudiantes")
	}
	generalAverage = generalAverage / generalTotal
	*reply = generalAverage
	return nil
}

func (this *Server) GetSubjectAverage(subject string, reply *float64) error {
	_, subjectExists := subjects[subject]
	if !subjectExists {
		return errors.New("La materia no existe")
	}
	var average float64 = 0
	var total float64 = 0
	for _, grade := range subjects[subject] {
		average = average + grade
		total = total + 1
	}
	average = average / total
	*reply = average
	return nil
}

func server() {
	rpc.Register(new(Server))
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go rpc.ServeConn(c)
	}
}

func main() {
	go server()

	var input string
	fmt.Scanln(&input)
}
