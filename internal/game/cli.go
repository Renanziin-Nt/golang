package game

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"

    "my-golang-cli/internal/board"
    "my-golang-cli/internal/pieces"
)

func askBoardDimensions(reader *bufio.Reader) (int, int) {
    fmt.Println()
    fmt.Println("+-------------------------+")
    fmt.Println("| Welcome to Unvoid Chess |")
    fmt.Println("+-------------------------+")
    fmt.Println()
    fmt.Println("Select board size to start.")
    fmt.Println("Values must be from 6 to 12 on each dimension.")
    var width, height int
    for {
        fmt.Print("Enter board width (X): ")
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)
        w, err := strconv.Atoi(input)
        if err != nil || w < 6 || w > 12 {
            fmt.Println("Invalid input. Please enter a number between 6 and 12.")
            continue
        }
        width = w
        break
    }
    for {
        fmt.Print("Enter board height (Y): ")
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)
        h, err := strconv.Atoi(input)
        if err != nil || h < 6 || h > 12 {
            fmt.Println("Invalid input. Please enter a number between 6 and 12.")
            continue
        }
        height = h
        break
    }
    fmt.Printf("\nStarting match on an %dx%d board...\n\n", width, height)
    return width, height
}

func Run() {
    reader := bufio.NewReader(os.Stdin)
    width, height := askBoardDimensions(reader)
    b := board.NewBoard(width, height)
    turn := pieces.White

    for {
        b.Display()
        fmt.Printf("Turn: %s\n", strings.Title(string(turn)))
        fmt.Print("Type a command (move <from> <to>, help, restart, exit): ")
        line, _ := reader.ReadString('\n')
        line = strings.TrimSpace(line)

        switch {
        case line == "exit":
            fmt.Println("Goodbye!")
            return
        case line == "help":
            fmt.Println("Available commands:")
            fmt.Println("  move <from> <to>   Move a piece (e.g., move A1 B3)")
            fmt.Println("  restart            Restart the match")
            fmt.Println("  help               Show this list")
            fmt.Println("  exit               Exit the game")
            continue
        case line == "restart":
            width, height = askBoardDimensions(reader)
            b = board.NewBoard(width, height)
            turn = pieces.White
            continue
        }

        parts := strings.Fields(line)
        if len(parts) == 3 && parts[0] == "move" {
            fromX, fromY, ok1 := board.ParseCoord(parts[1], b.Width, b.Height)
            toX, toY, ok2 := board.ParseCoord(parts[2], b.Width, b.Height)
            if !ok1 || !ok2 {
                fmt.Println("Invalid coordinates. Use format like A1 B3.")
                continue
            }
            if fromX == toX && fromY == toY {
                fmt.Println("Destination must be different from origin.")
                continue
            }
            p := b.Squares[fromX][fromY]
            if p == nil {
                fmt.Println("No piece at origin.")
                continue
            }
            if p.Color() != turn {
                fmt.Println("You can't move your opponent's piece.")
                continue
            }
            if !b.IsValidMove(fromX, fromY, toX, toY) {
                fmt.Println("Invalid move for this piece.")
                continue
            }
            captured, _ := b.MovePiece(fromX, fromY, toX, toY)
            if captured != nil && captured.Type() == pieces.ProductOwnerType {
                b.Display()
                fmt.Printf("%s wins! 🎉\n", strings.Title(string(turn)))
                fmt.Println("Type 'restart' to play again or 'exit' to leave.")
                for {
                    fmt.Print("> ")
                    cmd, _ := reader.ReadString('\n')
                    cmd = strings.TrimSpace(cmd)
                    if cmd == "restart" {
                        width, height = askBoardDimensions(reader)
                        b = board.NewBoard(width, height)
                        turn = pieces.White
                        break
                    } else if cmd == "exit" {
                        fmt.Println("Goodbye!")
                        return
                    } else {
                        fmt.Println("Type 'restart' to play again or 'exit' to leave.")
                    }
                }
                continue
            }

            if turn == pieces.White {
                turn = pieces.Black
            } else {
                turn = pieces.White
            }
        } else {
            fmt.Println("Unknown command. Type 'help' for options.")
        }
    }
}