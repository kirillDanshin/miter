package main

import (
	"fmt"
	"os"
	"strconv"

	cli "gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "miter"
	app.Usage = "Calculates the ratio of given proportions."
	app.Version = "0.0.1"

	app.Action = func(c *cli.Context) error {
		if len(c.Args()) <= 1 || len(c.Args()) > 2 {
			fmt.Println("Invalid number of parameters.")
		} else {
			width, errw := strconv.Atoi(c.Args()[0])
			if errw != nil {
				fmt.Println("The width is invalid.")
			} else {
				height, errh := strconv.Atoi(c.Args()[1])
				if errh != nil {
					fmt.Println("The height is invalid.")
				} else {
					gcd := gcd(width, height)
					fmt.Println("The ratio is: " + strconv.Itoa(width/gcd) + ":" + strconv.Itoa(height/gcd))
					fmt.Println("The percentage is: " + strconv.FormatFloat(float64(height)/float64(width)*float64(100), 'f', 6, 64) + "%")
				}
			}
		}

		return nil
	}

	app.Run(os.Args)
}

func gcd(w, h int) int {
	for h != 0 {
		w, h = h, w%h
	}
	return w
}
