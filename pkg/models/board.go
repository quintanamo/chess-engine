package models

type Board struct {
	whitePawnsPos   uint64
	whiteKnightsPos uint64
	whiteBishopsPos uint64
	whiteRooksPos   uint64
	whiteQueensPos  uint64
	whiteKingPos    uint64
	blackPawnsPos   uint64
	blackKnightsPos uint64
	blackBishopsPos uint64
	blackRooksPos   uint64
	blackQueensPos  uint64
	blackKingPos    uint64
}

type Piece int

const (
	PAWN   Piece = 0
	KNIGHT Piece = 1
	BISHOP Piece = 2
	ROOK   Piece = 3
	QUEEN  Piece = 4
	KING   Piece = 5
)

/*
Initialize the board with each piece's initial positions.
Each of the 64 bits represents a square on the board.
*/
func (b *Board) Initialize() {
	b.whitePawnsPos = 0b0000000011111111000000000000000000000000000000000000000000000000
	b.whiteKnightsPos = 0b0100001000000000000000000000000000000000000000000000000000000000
	b.whiteBishopsPos = 0b0010010000000000000000000000000000000000000000000000000000000000
	b.whiteRooksPos = 0b1000000100000000000000000000000000000000000000000000000000000000
	b.whiteQueensPos = 0b0001000000000000000000000000000000000000000000000000000000000000
	b.whiteKingPos = 0b0000100000000000000000000000000000000000000000000000000000000000
	b.blackPawnsPos = 0b0000000000000000000000000000000000000000000000001111111100000000
	b.blackKnightsPos = 0b0000000000000000000000000000000000000000000000000000000001000010
	b.blackBishopsPos = 0b0000000000000000000000000000000000000000000000000000000000100100
	b.blackRooksPos = 0b0000000000000000000000000000000000000000000000000000000010000001
	b.blackQueensPos = 0b0000000000000000000000000000000000000000000000000000000000010000
	b.blackKingPos = 0b0000000000000000000000000000000000000000000000000000000000001000
}

func (b Board) IsValidMove(m Move) bool {
	if m.isWhiteMoving {
		if ((m.destinationPos & b.whitePawnsPos) > 0) ||
			((m.destinationPos & b.whiteKnightsPos) > 0) ||
			((m.destinationPos & b.whiteBishopsPos) > 0) ||
			((m.destinationPos & b.whiteRooksPos) > 0) ||
			((m.destinationPos & b.whiteQueensPos) > 0) ||
			((m.destinationPos & b.whiteKingPos) > 0) {
			return false
		}
	} else {
		if ((m.destinationPos & b.blackPawnsPos) > 0) ||
			((m.destinationPos & b.blackKnightsPos) > 0) ||
			((m.destinationPos & b.blackBishopsPos) > 0) ||
			((m.destinationPos & b.blackRooksPos) > 0) ||
			((m.destinationPos & b.blackQueensPos) > 0) ||
			((m.destinationPos & b.blackKingPos) > 0) {
			return false
		}
	}
	return true
}

func (b *Board) MakeMove(m Move) {
	if m.isWhiteMoving {
		switch m.piece {
		case PAWN:
			b.whitePawnsPos = b.whitePawnsPos &^ m.startingPos
			b.whitePawnsPos = b.whitePawnsPos | m.destinationPos
		case KNIGHT:
			b.whiteKnightsPos = b.whiteKnightsPos &^ m.startingPos
			b.whiteKnightsPos = b.whiteKnightsPos | m.destinationPos
		case BISHOP:
			b.whiteBishopsPos = b.whiteBishopsPos &^ m.startingPos
			b.whiteBishopsPos = b.whiteBishopsPos | m.destinationPos
		case ROOK:
			b.whiteRooksPos = b.whiteRooksPos &^ m.startingPos
			b.whiteRooksPos = b.whiteRooksPos | m.destinationPos
		case QUEEN:
			b.whiteQueensPos = b.whiteQueensPos &^ m.startingPos
			b.whiteQueensPos = b.whiteQueensPos | m.destinationPos
		case KING:
			b.whiteKingPos = b.whiteKingPos &^ m.startingPos
			b.whiteKingPos = b.whiteKingPos | m.destinationPos
		}
	} else {
		switch m.piece {
		case PAWN:
			b.blackPawnsPos = b.blackPawnsPos &^ m.startingPos
			b.blackPawnsPos = b.blackPawnsPos | m.destinationPos
		case KNIGHT:
			b.blackKnightsPos = b.blackKnightsPos &^ m.startingPos
			b.blackKnightsPos = b.blackKnightsPos | m.destinationPos
		case BISHOP:
			b.blackBishopsPos = b.blackBishopsPos &^ m.startingPos
			b.blackBishopsPos = b.blackBishopsPos | m.destinationPos
		case ROOK:
			b.blackRooksPos = b.blackRooksPos &^ m.startingPos
			b.blackRooksPos = b.blackRooksPos | m.destinationPos
		case QUEEN:
			b.blackQueensPos = b.blackQueensPos &^ m.startingPos
			b.blackQueensPos = b.blackQueensPos | m.destinationPos
		case KING:
			b.blackKingPos = b.blackKingPos &^ m.startingPos
			b.blackKingPos = b.blackKingPos | m.destinationPos
		}
	}
}
