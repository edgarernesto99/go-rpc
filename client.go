package main

import (
	"fmt"
	"net/rpc"
	"bufio"
	"os"
)

type Args struct {
	Subject string
	Student string
	Grade float64
}

func client() {
	c, err := rpc.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	var op int64
	scaner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println(" ========================================= ")
		fmt.Println("|                    MENU                 |")
		fmt.Println("|=========================================|")
		fmt.Println("| 1. Agregar calificación de una materia  |")
		fmt.Println("| 2. Mostrar promedio de un alumno        |")
		fmt.Println("| 3. Mostrar promedio general             |")
		fmt.Println("| 4. Mostrar promedio de una materia      |")
		fmt.Println("| 0. Salir                                |")
		fmt.Println(" ========================================= ")
		fmt.Print("Opción: ")
		fmt.Scanln(&op)

		switch op {
		case 1:
			var subject, student string
			var grade float64
			fmt.Print("Materia: ")
			scaner.Scan()
			subject = scaner.Text()
			fmt.Print("Estudiante: ")
			scaner.Scan()
			student = scaner.Text()
			fmt.Print("Calificación: ")
			fmt.Scanln(&grade)

			var result bool
			err = c.Call("Server.SetGrade", Args{subject, student, grade}, &result)
			if err != nil {
				fmt.Println(err)
			}
		case 2:
			var student string
			fmt.Print("Estudiante: ")
			scaner.Scan()
			student = scaner.Text()

			var result float64
			err = c.Call("Server.GetStudentAverage", student, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Promedio de", student, ":", result)
			}
		case 3:
			var result float64
			err = c.Call("Server.GetGeneralAverage", Args{}, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Promedio General: ", result)
			}
		case 4:
			var subject string
			fmt.Print("Materia: ")
			scaner.Scan()
			subject = scaner.Text()

			var result float64
			err = c.Call("Server.GetSubjectAverage", subject, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Promedio de", subject, ":", result)
			}
		case 0:
			return
		}
	}
}

func main() {
	client()
}