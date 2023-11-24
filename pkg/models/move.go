package models

type Move struct {
	piece          Piece
	startingPos    uint64
	destinationPos uint64
	isWhiteMoving  bool
}
