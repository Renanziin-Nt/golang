package game

import (
    "my-golang-cli/internal/board"
    "my-golang-cli/internal/pieces"
)

type Game struct {
    Board       *board.Board
    CurrentTurn pieces.Color
}


func NewGame(width, height int) *Game {
    return &Game{
        Board:       board.NewBoard(width, height),
        CurrentTurn: pieces.White,
    }
}

func (g *Game) SwitchTurn() {
    if g.CurrentTurn == pieces.White {
        g.CurrentTurn = pieces.Black
    } else {
        g.CurrentTurn = pieces.White
    }
}