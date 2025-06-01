package pieces

type Color string

const (
    White Color = "white"
    Black Color = "black"
)

type PieceType string

const (
    ProductOwnerType PieceType = "ProductOwner"
    DeveloperType    PieceType = "Developer"
    DesignerType     PieceType = "Designer"
)

type Piece interface {
    Type() PieceType
    Color() Color
    Symbol() string
    ValidMoves(board [][]Piece, x, y int) [][2]int
    CanCapture(board [][]Piece, fromX, fromY, toX, toY int) bool
}

func PieceSymbol(p Piece) string {
    switch p.Type() {
    case ProductOwnerType:
        if p.Color() == White {
            return "♙"
        }
        return "♟"
    case DeveloperType:
        if p.Color() == White {
            return "♖"
        }
        return "♜"
    case DesignerType:
        if p.Color() == White {
            return "♗"
        }
        return "♝"
    }
    return "?"
}