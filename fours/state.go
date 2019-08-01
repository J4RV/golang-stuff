package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"
)

var (
	ErrColIndexOutOfRange = errors.New("column index out of range")
	ErrColumnFull         = errors.New("that column is full")
)

type state struct {
	board           [][]cell
	currPlayerColor cell
}

func New() *state {
	board := make([][]cell, BoardHorizontalSize)
	for i := 0; i < BoardHorizontalSize; i++ {
		board[i] = make([]cell, BoardVerticalSize)
	}
	firstPlayer := cell(rand.Intn(2) + 1)
	return &state{board, firstPlayer}
}

func (s *state) row(row int) []cell {
	res := make([]cell, BoardHorizontalSize)
	for colIndex := range s.board {
		res[colIndex] = s.board[colIndex][row]
	}
	return res
}

func (s *state) col(col int) []cell {
	return s.board[col]
}

func (s *state) dropCoin(col int, color cell) error {
	lowestEmptyRowIndex := -1
	column := s.col(col)

	for row := range column {
		if column[row] == EMPTY {
			lowestEmptyRowIndex = row
			break
		}
	}

	if lowestEmptyRowIndex == -1 {
		return ErrColumnFull
	}

	s.board[col][lowestEmptyRowIndex] = color
	return nil
}

func (s *state) nextPlayer() {
	switch s.currPlayerColor {
	case WHITE:
		s.currPlayerColor = BLACK
	case BLACK:
		s.currPlayerColor = WHITE
	default:
		log.Panicf("invalid currPlayerColor %d", s.currPlayerColor)
	}
}

func (s *state) String() string {
	boardString := ""
	for rowIndex := BoardVerticalSize - 1; rowIndex >= 0; rowIndex-- {
		boardString += fmt.Sprintf("%v\n", s.row(rowIndex))
	}

	boardString += " "
	for i := 0; i < BoardHorizontalSize; i++ {
		boardString += strconv.Itoa(i) + " "
	}
	boardString += "\n"

	return fmt.Sprintf("%vCurrent player: %s", boardString, s.currPlayerColor.EnumString())
}

func (s *state) Play(col int) error {
	if col < 0 || col >= BoardHorizontalSize {
		return ErrColIndexOutOfRange
	}

	err := s.dropCoin(col, s.currPlayerColor)
	if err != nil {
		return err
	}
	s.nextPlayer()

	return nil
}

func (s *state) IsFinished() bool {
	return false // TODO
}