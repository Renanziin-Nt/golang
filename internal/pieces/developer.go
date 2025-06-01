package pieces

type Developer struct {
    color Color
}

func NewDeveloper(color Color) *Developer {
    return &Developer{color: color}
}

func (d *Developer) Type() PieceType { return DeveloperType }
func (d *Developer) Color() Color    { return d.color }
func (d *Developer) Symbol() string {
    if d.color == White {
        return "♖"
    }
    return "♜"
}


func (d *Developer) ValidMoves(board [][]Piece, x, y int) [][2]int {
    moves := [][2]int{}
    dirs := [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
    height := len(board)
    width := len(board[0])
    
    for _, dir := range dirs {
        for dist := 1; dist <= 3; dist++ {
            nx, ny := x+dir[0]*dist, y+dir[1]*dist
            if nx < 0 || nx >= height || ny < 0 || ny >= width {
                break
            }
            if board[nx][ny] == nil {
                moves = append(moves, [2]int{nx, ny})
            } else {
                if board[nx][ny].Color() != d.color {
                    moves = append(moves, [2]int{nx, ny})
                }
                break 
            }
        }
    }
    return moves
}

func (d *Developer) CanCapture(board [][]Piece, fromX, fromY, toX, toY int) bool {

    dx, dy := toX-fromX, toY-fromY
    if dx == 0 && dy == 0 {
        return false
    }
    steps := max(abs(dx), abs(dy))
    if steps < 2 || steps > 3 {
        return false
    }
    stepX, stepY := dx/steps, dy/steps
    for i := 1; i < steps; i++ {
        ix, iy := fromX+stepX*i, fromY+stepY*i
        if board[ix][iy] != nil {
            return false 
        }
    }
    return true
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}