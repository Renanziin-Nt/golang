# Unvoid Chess

Unvoid Chess is a custom chess-like game implemented in Go. It features unique pieces and movement rules designed for a technical test. The game satisfies requirements for correctness, code readability, maintainability, extensibility, and efficiency.

## Table of Contents

- [Overview](#overview)
- [Project Structure](#project-structure)
- [Game Rules and Piece Movements](#game-rules-and-piece-movements)
- [Initial Piece Positions](#initial-piece-positions)
- [How to Run the Application](#how-to-run-the-application)
- [Usage Examples](#usage-examples)

## Overview

- **ProductOwner (PO):**  
  Acts like a king (moves one square in any direction).  
  Symbols: ♔ (white) and ♚ (black).

- **Developer (Tower):**  
  Acts as a rook, moving horizontally or vertically up to 3 squares.  
  Symbols: ♖ (white) and ♜ (black).

- **Designer (Knight):**  
  Moves in an 'L' shape (similar to the chess knight).  
  Symbols: ♘ (white) and ♞ (black).

## Project Structure

```
Unvoid-Chess/
│
├── cmd/
│   └── main.go            // Application entry point: simply calls game.Run()
│
├── internal/
│   ├── board/
│   │   └── board.go       // Board representation, initialization, display, and coordinate parsing
│   │
│   ├── game/
│   │   ├── cli.go         // Main game loop that handles commands (move, help, restart, exit)
│   │   └── game.go        // Game structure holding the board and current turn information
│   │
│   └── pieces/
│       ├── piece.go       // Piece interface and type/constant definitions (colors, piece types)
│       ├── product_owner.go  // ProductOwner piece: moves one square in any direction
│       ├── developer.go      // Developer piece: acts as a Tower (rook) moving horizontally/vertically up to 3 squares
│       └── designer.go       // Designer piece: acts as a Knight, moving in "L"-shape
│
├── go.mod                 // Go module file
└── README.md              // Project documentation (this file)
```

## Game Rules and Piece Movements

### ProductOwner (PO)
- **Movement:** Can move one square in any direction (horizontal, vertical, or diagonal).
- **Capture:** Can capture any opponent's piece on the destination square.
  
### Developer (Tower)
- **Role:** Represents the Developer piece, which now functions as a Tower.
- **Movement:** Moves horizontally or vertically up to 3 squares. It stops if its path is blocked; it may capture an enemy piece on its way if it is the first encountered piece.
  
### Designer (Knight)
- **Role:** Represents the Designer piece, which now functions as a Knight.
- **Movement:** Moves in an 'L' shape – two squares in one direction and then one square perpendicularly. It can jump over pieces, but the destination must be vacant or contain an enemy piece.

## Initial Piece Positions

- **White Pieces (Bottom Row):**
  - **ProductOwner (PO):** Starts at **A1**
  - **Developer (Tower):** Starts at **B1**
  - **Designer (Knight):** Starts at **C1**

- **Black Pieces (Top Row):**  
  (Configured symmetrically on the opposite side)
  - **ProductOwner (PO):** Starts at **H6**
  - **Developer (Tower):** Starts at **G6**
  - **Designer (Knight):** Starts at **F6**

> **Note:** The board dimensions (width and height) are selected by the player at the start of the game. In these examples, the board is set to 8×6.

## How to Run the Application

1. **Prerequisites:**  
   Make sure you have Go installed. See [Go installation instructions](https://golang.org/doc/install).

2. **Clone the Repository:**  
   Clone or download this codebase to your local machine.

3. **Compile and Run:**  
   Open a terminal (or PowerShell) and navigate to the project root, then run:
   ```
   go run cmd/main.go
   ```
   
4. **Board Setup:**  
   On start, you will be prompted to enter the board dimensions. Enter numbers between 6 and 12. For example:
   ```
   Enter board width (X): 8
   Enter board height (Y): 6
   Starting match on an 8x6 board...
   ```

## Usage Examples

### Moving a Piece
Use the following command format:
```
move <origin> <destination>
```
For example, to move the white ProductOwner from A1 to B2:
```
move A1 B2
```

### Other Commands
- **help:**  
  Displays available commands.
  
- **restart:**  
  Restarts the game and asks for board dimensions again.
  
- **exit:**  
  Exits the application.

### Example Sequence for a White Victory
Given that white pieces are set at A1 (PO), B1 (Tower/Developer), C1 (Knight/Designer) and black pieces are at H6, G6, F6 respectively:
```
move A1 B2    // White PO moves diagonally
move H6 G5    // Black PO moves diagonally
move B2 C3    // White PO moves forward
move G5 F4    // Black PO continues movement
move C3 D4    // White PO advances
move F4 E4    // Black PO moves sideways
move D4 E5    // White PO positions
move E4 F4    // Black PO shifts
move E5 F6    // White PO captures Black PO and wins!
```

### Testing Developer and Designer Movements
- **Developer (Tower - Developer):**  
  From any cell (e.g., D3), test vertical/horizontal moves:
  ```
  move D3 D6   // Move vertically up to 3 squares (if path is clear)
  move D3 A3   // Move horizontally
  ```
  
- **Designer (Knight - Designer):**  
  If the Designer is in B2, test L-shaped movement:
  ```
  move B2 C4   // Moves like a knight; ensure destination is valid
  ```

## Final Notes

- **Validation:**  
  If you attempt an invalid move (e.g., more than one square for the ProductOwner or moving beyond allowed limits for the Tower/Knight), the game will output "Invalid move for this piece."

- **Extensibility:**  
  The code is modular. New pieces or movement rules can be added by implementing the `Piece` interface in the `internal/pieces` package.

- **Feedback:**  
  The game will show the board after each move and provide prompt messages to keep you informed.

Enjoy testing Unvoid Chess!
