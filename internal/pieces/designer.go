package pieces

type Designer struct {
    color Color
}

func NewDesigner(color Color) *Designer {
    return &Designer{color: color}
}

func (d *Designer) Type() PieceType { return DesignerType }
func (d *Designer) Color() Color    { return d.color }
func (d *Designer) Symbol() string {
    if d.color == White {
        return "♘" // Cavalo branco
    }
    return "♞" // Cavalo preto
}

func (d *Designer) ValidMoves(board [][]Piece, x, y int) [][2]int {
    moves := [][2]int{}
    offsets := [][2]int{
        {2, 1}, {1, 2}, {-1, 2}, {-2, 1},
        {-2, -1}, {-1, -2}, {1, -2}, {2, -1},
    }
    height := len(board)
    width := len(board[0])
    for _, off := range offsets {
        nx, ny := x+off[0], y+off[1]
        if nx >= 0 && nx < height && ny >= 0 && ny < width {
            if board[nx][ny] == nil || board[nx][ny].Color() != d.color {
                moves = append(moves, [2]int{nx, ny})
            }
        }
    }
    return moves
}

func (d *Designer) CanCapture(board [][]Piece, fromX, fromY, toX, toY int) bool {
    for _, move := range d.ValidMoves(board, fromX, fromY) {
        if move[0] == toX && move[1] == toY {
            return true
        }
    }
    return false
}