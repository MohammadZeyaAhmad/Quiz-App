package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fileName := flag.String("csv", "problems.csv", "csv problem set")

	_ = fileName
	file, err := os.Open(*fileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s",*fileName))
	}
	reader:=csv.NewReader(file)
	lines,err:=reader.ReadAll()
	if(err!=nil){
		exit("failed to parse the CSV file")
	}
	
	problems:=parseLines(lines);
	correct :=0
	timeLimit:=flag.Int("limit",30,"This is the time for the quiz in seconds")
    flag.Parse()
	timer:=time.NewTimer(time.Duration(*timeLimit)*time.Second)
	
  
	for i,p:=range problems{
		fmt.Printf("Problem #%d: %s=\n",i+1,p.question)
		answerCh:=make(chan string)
		go func(){
           var answer string
		 fmt.Scanf("%s\n",&answer)
		 answerCh<-answer
		}()
		
		select
		{
		case <-timer.C:
			fmt.Printf("You scored %d out of %d points",correct,len(problems))
			return;
	    case answer:=<-answerCh:
			if(answer==p.answer){
				fmt.Println("Correct answer")
				correct++
			}else{
				fmt.Println("Oops your answer was incorrect")
			}
          
         
		}
		
	}
	fmt.Printf("\nYou scored %d out of %d points",correct,len(problems))
}

type problem struct {
	question string
	answer string
}

func parseLines(lines [][]string)[]problem {
	problems:=make([]problem, len(lines))
	for i,line:= range lines {
		problems[i]=problem{
          question:line[0],
		  answer:strings.TrimSpace(line[1]),
		}
	}
	return problems
}
func exit(msg string){
  fmt.Println(msg)
  os.Exit(1)
}
