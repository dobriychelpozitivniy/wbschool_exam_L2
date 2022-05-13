package pattern

import "fmt"

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

func main() {
	rps := NewRPS()

	rps.FirstPlayerMove()
}

type State interface {
	StartNewGame() error
	FirstPlayerMove() error
	SecondPlayerMove() error
	GetWinner() error
}

type RockPaperScissors struct {
	newGame          State
	firstPlayerMove  State
	secondPlayerMove State
	twoPlayersMoved  State
	currentState     State
	firstPlayer      string
	secondPlayer     string
	winner           string
}

func NewRPS() *RockPaperScissors {
	rps := RockPaperScissors{
		firstPlayer:  "",
		secondPlayer: "",
		winner:       "",
	}

	ng := newGame{
		rps: &rps,
	}

	fpm := firstPlayerMove{
		rps: &rps,
	}

	spm := secondPlayerMove{
		rps: &rps,
	}

	tpm := twoPlayersMoved{
		rps: &rps,
	}

	rps.setState(&ng)

	rps.firstPlayerMove = &fpm
	rps.secondPlayerMove = &spm
	rps.twoPlayersMoved = &tpm
	rps.newGame = &ng

	return &rps
}

func (rps *RockPaperScissors) setState(s State) {
	rps.currentState = s
}

func (rps *RockPaperScissors) StartNewGame() error {
	return rps.currentState.StartNewGame()
}

func (rps *RockPaperScissors) FirstPlayerMove() error {
	return rps.currentState.FirstPlayerMove()
}

func (rps *RockPaperScissors) SecondPlayerMove() error {
	return rps.currentState.SecondPlayerMove()
}

func (rps *RockPaperScissors) GetWinner() error {
	return rps.currentState.GetWinner()
}

type newGame struct {
	rps *RockPaperScissors
}

func (s *newGame) StartNewGame() error {
	return fmt.Errorf("Игра уже началась")
}

func (s *newGame) FirstPlayerMove() error {
	s.rps.firstPlayer = "Rock" // Тут должен быть генератор рандомных значений
	s.rps.setState(s.rps.firstPlayerMove)

	return nil
}

func (s *newGame) SecondPlayerMove() error {
	s.rps.secondPlayer = "Paper" // Тут должен быть генератор рандомных значений
	s.rps.setState(s.rps.secondPlayerMove)

	return nil
}

func (s *newGame) GetWinner() error {
	return fmt.Errorf("Еще нет победителя")
}

type firstPlayerMove struct {
	rps *RockPaperScissors
}

func (s *firstPlayerMove) StartNewGame() error {
	return fmt.Errorf("Игра уже началась")
}

func (s *firstPlayerMove) FirstPlayerMove() error {
	return fmt.Errorf("Первый игрок уже ходил")
}

func (s *firstPlayerMove) SecondPlayerMove() error {
	s.rps.secondPlayer = "Paper"                 // Тут должен быть генератор рандомных значений
	if s.rps.secondPlayer == s.rps.firstPlayer { // Тут должен быть определятор победителя
		s.rps.winner = "Second player"
	}

	s.rps.setState(s.rps.twoPlayersMoved)

	return nil
}

func (s *firstPlayerMove) GetWinner() error {
	return fmt.Errorf("Еще нет победителя")
}

type secondPlayerMove struct {
	rps *RockPaperScissors
}

func (s *secondPlayerMove) StartNewGame() error {
	return fmt.Errorf("Игра уже началась")
}

func (s *secondPlayerMove) FirstPlayerMove() error {
	s.rps.firstPlayer = "Paper"                  // Тут должен быть генератор рандомных значений
	if s.rps.secondPlayer == s.rps.firstPlayer { // Тут должен быть определятор победителя
		s.rps.winner = "First player"
	}

	s.rps.setState(s.rps.twoPlayersMoved)

	return nil
}

func (s *secondPlayerMove) SecondPlayerMove() error {
	return fmt.Errorf("Второй игрок уже ходил")
}

func (s *secondPlayerMove) GetWinner() error {
	return fmt.Errorf("Еще нет победителя")
}

type twoPlayersMoved struct {
	rps *RockPaperScissors
}

func (s *twoPlayersMoved) StartNewGame() error {
	s.rps.setState(s.rps.newGame)

	return nil
}

func (s *twoPlayersMoved) FirstPlayerMove() error {
	return fmt.Errorf("Первый игрок уже ходил")
}

func (s *twoPlayersMoved) SecondPlayerMove() error {
	return fmt.Errorf("Второй игрок уже ходил")
}

func (s *twoPlayersMoved) GetWinner() error {
	fmt.Println(s.rps.winner)

	return nil
}
