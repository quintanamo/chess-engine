package models

import (
	"testing"
)

func TestInitialize(t *testing.T) {
	b := Board{}
	b.Initialize()
	if b.whitePawnsPos != 0b0000000011111111000000000000000000000000000000000000000000000000 {
		t.Errorf("White Pawns Incorrectly Positioned")
	}
	if b.whiteKnightsPos != 0b0100001000000000000000000000000000000000000000000000000000000000 {
		t.Errorf("White Knights Incorrectly Positioned")
	}
	if b.whiteBishopsPos != 0b0010010000000000000000000000000000000000000000000000000000000000 {
		t.Errorf("White Bishops Incorrectly Positioned")
	}
	if b.whiteRooksPos != 0b1000000100000000000000000000000000000000000000000000000000000000 {
		t.Errorf("White Rooks Incorrectly Positioned")
	}
	if b.whiteQueensPos != 0b0001000000000000000000000000000000000000000000000000000000000000 {
		t.Errorf("White Queen Incorrectly Positioned")
	}
	if b.whiteKingPos != 0b0000100000000000000000000000000000000000000000000000000000000000 {
		t.Errorf("White King Incorrectly Positioned")
	}
	if b.blackPawnsPos != 0b0000000000000000000000000000000000000000000000001111111100000000 {
		t.Errorf("Black Pawns Incorrectly Positioned")
	}
	if b.blackKnightsPos != 0b0000000000000000000000000000000000000000000000000000000001000010 {
		t.Errorf("Black Knights Incorrectly Positioned")
	}
	if b.blackBishopsPos != 0b0000000000000000000000000000000000000000000000000000000000100100 {
		t.Errorf("Black Bishops Incorrectly Positioned")
	}
	if b.blackRooksPos != 0b0000000000000000000000000000000000000000000000000000000010000001 {
		t.Errorf("Black Rooks Incorrectly Positioned")
	}
	if b.blackQueensPos != 0b0000000000000000000000000000000000000000000000000000000000010000 {
		t.Errorf("Black Queen Incorrectly Positioned")
	}
	if b.blackKingPos != 0b0000000000000000000000000000000000000000000000000000000000001000 {
		t.Errorf("Black King Incorrectly Positioned")
	}
}

func TestIsValidMove(t *testing.T) {
	b := Board{}
	b.Initialize()
	b.whitePawnsPos = 0b0000000011111110100000000000000000000000000000000000000000000000

	// test pushing pawn to square occupied by pawn
	m := Move{
		PAWN,
		0b0000000010000000000000000000000000000000000000000000000000000000,
		0b0000000000000000100000000000000000000000000000000000000000000000,
		true}
	if b.IsValidMove(m) {
		t.Errorf("TestIsValidMove failed for case where pawn moves to another pawn's position")
	}

	// test pushing pawn to open square
	m = Move{
		PAWN,
		0b0000000001000000000000000000000000000000000000000000000000000000,
		0b0000000000000000010000000000000000000000000000000000000000000000,
		true}
	if !b.IsValidMove(m) {
		t.Errorf("TestIsValidMove failed for case where pawn moves forward once to an open square")
	}

	// test pushing pawn sidways (illegal)
	m = Move{
		PAWN,
		0b0000000000000000100000000000000000000000000000000000000000000000,
		0b0000000000000000010000000000000000000000000000000000000000000000,
		true}
	if b.IsValidMove(m) {
		t.Errorf("TestIsValidMove failed for case where pawn moves sideways")
	}
}
