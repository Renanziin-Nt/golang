package board

import (
    "fmt"
    "strings"
    "strconv" 
    "my-golang-cli/internal/pieces"
)

type Board struct {
    Squares [][]pieces.Piece
    Width   int
    Height  int
}

func NewBoard(width, height int) *Board {
    b := &Board{
        Squares: make([][]pieces.Piece, height),
        Width:   width,
        Height:  height,
    }
    for i := range b.Squares {
        b.Squares[i] = make([]pieces.Piece, width)
    }
    b.Initialize()
    return b
}

func (b *Board) Initialize() {

    for i := range b.Squares {
        for j := range b.Squares[i] {
            b.Squares[i][j] = nil
        }
    }
    

    b.Squares[b.Height-1][0] = pieces.NewProductOwner(pieces.White)
    b.Squares[b.Height-1][1] = pieces.NewDeveloper(pieces.White)
    b.Squares[b.Height-1][2] = pieces.NewDesigner(pieces.White)
    

    b.Squares[0][b.Width-1] = pieces.NewProductOwner(pieces.Black)
    b.Squares[0][b.Width-2] = pieces.NewDeveloper(pieces.Black)
    b.Squares[0][b.Width-3] = pieces.NewDesigner(pieces.Black)
}

func (b *Board) Display() {

    fmt.Print("   ")
    for c := 0; c < b.Width; c++ {
        fmt.Printf("%c ", 'A'+c)
    }
    fmt.Println()
    for i := 0; i < b.Height; i++ {
        fmt.Printf("%2d ", b.Height-i)
        for j := 0; j < b.Width; j++ {
            if b.Squares[i][j] == nil {
                fmt.Print(". ")
            } else {
                fmt.Print(b.Squares[i][j].Symbol() + " ")
            }
        }
        fmt.Println()
    }
}

func (b *Board) FindPiece(pt pieces.PieceType, color pieces.Color) (int, int, bool) {
    for i := 0; i < b.Height; i++ {
        for j := 0; j < b.Width; j++ {
            p := b.Squares[i][j]
            if p != nil && p.Type() == pt && p.Color() == color {
                return i, j, true
            }
        }
    }
    return -1, -1, false
}


func (b *Board) MovePiece(fromX, fromY, toX, toY int) (captured pieces.Piece, ok bool) {
    p := b.Squares[fromX][fromY]
    if p == nil {
        return nil, false
    }
    captured = b.Squares[toX][toY]
    b.Squares[toX][toY] = p
    b.Squares[fromX][fromY] = nil
    return captured, true
}


func (b *Board) IsValidMove(fromX, fromY, toX, toY int) bool {
    p := b.Squares[fromX][fromY]
    if p == nil {
        return false
    }
    valid := false
    for _, mv := range p.ValidMoves(b.Squares, fromX, fromY) {
        if mv[0] == toX && mv[1] == toY {
            valid = true
            break
        }
    }
    return valid
}


func ParseCoord(coord string, width, height int) (int, int, bool) {
    if len(coord) < 2 {
        return 0, 0, false
    }
    col := int(strings.ToUpper(coord[:1])[0] - 'A')
    row, err := strconv.Atoi(coord[1:])
    if err != nil {
        return 0, 0, false
    }
    rowIdx := height - row
    if col < 0 || col >= width || rowIdx < 0 || rowIdx >= height {
        return 0, 0, false
    }
    return rowIdx, col, true
}