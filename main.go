package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type question struct {
	question string
	answer   int
}

type score struct {
	correct        int
	incorrect      int
	totalQuestions int
}

func (s *score) incrementCorrect() {
	s.correct++
}

func (s *score) incrementIncorrect() {
	s.incorrect++
}

func (s *score) printScore() {
	fmt.Printf("You got %v right and %v wrong out of a possible %v \n",
		s.correct,
		s.incorrect,
		s.totalQuestions)
}

func main() {
	questions, err := getQuestionsFromCSV("./problems.csv")
	if err != nil {
		panic(err)
	}

	playerScore := score{0, 0, len(*questions)}
	askQuestions(questions, &playerScore)

	playerScore.printScore()
}

func getQuestionsFromCSV(path string) (*[]question, error) {
	file, err := readFile(path)
	if err != nil {
		return nil, err
	}

	questions, err := parseCSV(file)
	if err != nil {
		return nil, err
	}

	return questions, nil
}

func readFile(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func parseCSV(input string) (*[]question, error) {
	reader := csv.NewReader(strings.NewReader(input))
	questions := []question{}

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		ans, err := strconv.Atoi(line[1])
		if err != nil {
			return nil, err
		}

		questions = append(questions, question{question: line[0], answer: ans})
	}

	return &questions, nil
}

func askQuestions(questions *[]question, pScore *score) {
	for i, q := range *questions {
		fmt.Printf("%v: %s\n", i+1, q.question)
		answer := getUsersAnswer()
		if answer != q.answer {
			pScore.incrementIncorrect()
			continue
		}

		pScore.incrementCorrect()
	}
	return
}

func getUsersAnswer() int {
	reader := bufio.NewReader(os.Stdin)

	response, _ := reader.ReadString('\n')

	answer, err := formatResponse(response)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Please enter a valid answer!")
		return getUsersAnswer()
	}

	return answer

}

func formatResponse(input string) (int, error) {
	// convert CRLF to LF
	text := strings.Replace(input, "\n", "", -1)
	return strconv.Atoi(text)
}
