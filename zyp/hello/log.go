package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	//file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer file.Close()
	//log.SetOutput(file)
	//log.Println("hey , asd")
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Print("Hey! I'm a log message!")
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	log.Debug().
		Int("EmployeeID", 1001).
		Msg("Getting employee information")

	log.Debug().
		Str("Name", "John").
		Send()
}
