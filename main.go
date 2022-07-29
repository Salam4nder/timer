package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	minCmd := flag.NewFlagSet("min", flag.ExitOnError)
	secCmd := flag.NewFlagSet("sec", flag.ExitOnError)
	inputMinutes := minCmd.Int("v", 0, "Set a timer for x minutes")
	inputSeconds := secCmd.Int("v", 0, "Set a timer for x seconds")

	switch os.Args[1] {
	case "min":
		err := minCmd.Parse(os.Args[2:])
		check(err)
		if *inputMinutes == 0 {
			log.Fatal("Incorrect input. Expecting: min -v <number>")
		}
		timer := time.NewTimer(time.Minute * time.Duration(*inputMinutes))
		fmt.Printf(
			"Timer set for %d minute(s) at %s\n", *inputMinutes,
			time.Now().Format("2006-01-02 15:04:05"))

		countdown(*inputMinutes, "min")
		for range timer.C {
			fmt.Printf(
				"Timer has finished at %s\n",
				time.Now().Format("2006-01-02 15:04:05"))
			os.Exit(1)
		}

	case "sec":
		err := secCmd.Parse(os.Args[2:])
		check(err)
		if *inputSeconds == 0 {
			log.Fatal("Incorrect input. Expecting: sec -v <seconds>")
		}
		timer := time.NewTimer(time.Second * time.Duration(*inputSeconds))
		fmt.Printf(
			"Timer set for %d second(s) at %s\n", *inputSeconds,
			time.Now().Format("2006-01-02 15:04:05"))

		countdown(*inputSeconds, "sec")
		for range timer.C {
			fmt.Printf("Timer has finished at %s\n",
				time.Now().Format("2006-01-02 15:04:05"))
			os.Exit(1)
		}
	default:
		log.Fatal(
			"Please specify a timer type. Expected : min or sec -a <number>")
	}
}

func countdown(value int, operation string) {
	switch operation {
	case "min":
		for i := value; i > 0; i-- {
			fmt.Printf("%d minute(s) remaining\n", i)
			time.Sleep(time.Minute)
		}
	case "sec":
		for i := value; i > 0; i-- {
			fmt.Printf("%d second(s) remaining\n", i)
			time.Sleep(time.Second)
		}
	default:
		log.Fatal("Incorrect type argument. Expected: min or sec.")
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
