package pieces

type ProductOwner struct {
    color Color
}

func NewProductOwner(color Color) *ProductOwner {
    return &ProductOwner{color: color}
}

func (p *ProductOwner) Type() PieceType { return ProductOwnerType }
func (p *ProductOwner) Color() Color    { return p.color }
func (p *ProductOwner) Symbol() string {
    if p.color == White {
        return "♔" 
    }
    return "♚" 
}


func (p *ProductOwner) ValidMoves(board [][]Piece, x, y int) [][2]int {
    moves := [][2]int{}
    dirs := [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
    height := len(board)
    width := len(board[0])
    
    for _, d := range dirs {
        nx, ny := x+d[0], y+d[1]
        if nx >= 0 && nx < height && ny >= 0 && ny < width {
            if board[nx][ny] == nil || board[nx][ny].Color() != p.color {
                moves = append(moves, [2]int{nx, ny})
            }
        }
    }
    return moves
}

func (p *ProductOwner) CanCapture(board [][]Piece, fromX, fromY, toX, toY int) bool {
    return true 
}