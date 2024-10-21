// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"net/http"
// 	"os"
// 	"time"

// 	"github.com/fatih/color"
// )

// var (
// 	red    = color.New(color.FgHiRed).SprintFunc()
// 	green  = color.New(color.FgHiGreen).SprintFunc()
// 	purple = []func(...interface{}) string{
// 		color.New(color.FgHiMagenta).SprintFunc(),
// 		color.New(color.FgMagenta).SprintFunc(),
// 		color.New(color.FgHiCyan).SprintFunc(),
// 		color.New(color.FgCyan).SprintFunc(),
// 	}
// )

// func printTitle(text string) {
// 	for i, char := range text {
// 		colorIndex := i % len(purple)
// 		fmt.Print(purple[colorIndex](string(char)))
// 	}
// 	fmt.Println()
// }

// func logFoundURL(url string) {
// 	fmt.Println(green("Found URL:") + " " + url)
// }

// func scanner(url string, w *os.File) {
// 	resp, err := http.Get(url + "/.git/HEAD")
// 	if err != nil {
// 		fmt.Println("Error retrieving url: " + url)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != 404 {
// 		logFoundURL(url)
// 		_, err := w.WriteString(url + "\n")
// 		if err != nil {
// 			fmt.Println("Error inserting url: " + url + " into valid_git_urls.txt")
// 		}
// 	}
// }

// func main() {
// 	if len(os.Args) != 2 {
// 		fmt.Println(red("Please provide a url list!"))
// 		os.Exit(0)
// 	}

// 	helpText := "git scanner made by erg0sum"

// 	file, err := os.Open(os.Args[1])
// 	if err != nil {
// 		fmt.Println(red("Failed to load file: " + os.Args[1] + ". Is it in the same directory?"))
// 		os.Exit(1)
// 	}
// 	defer file.Close()

// 	outputFile, err := os.OpenFile("valid_git_urls.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
// 	if err != nil {
// 		fmt.Println("Error opening file valid_git_urls.txt")
// 		os.Exit(1)
// 	}
// 	defer outputFile.Close()

// 	printTitle(helpText)

// 	fmt.Println(green("Loaded file: " + file.Name()))
// 	time.Sleep(time.Second)
// 	fmt.Println(green("Starting the scanner..."))
// 	time.Sleep(time.Second)
// 	s := bufio.NewScanner(file)
// 	for s.Scan() {
// 		scanner(s.Text(), outputFile)
// 	}
// }
