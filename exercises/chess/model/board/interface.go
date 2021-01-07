package board

import (
	"fmt"
	"C:\Users\User\Documents\ECAM\4e_Architecture\4eail40_2020\exercises\chess\model\piece"
	"C:\Users\User\Documents\ECAM\4e_Architecture\4eail40_2020\exercises\chess\model\coord"
)

//Board is an interface to a chess board.
//It defines methods for handling pieces on it.
type Board interface {
	fmt.Stringer
	//PieceAt retrieves piece at give coordinates.
	//returns nil if no piece was found.
	PieceAt(at coord.ChessCoordinates) piece.Piece
	//MovePiece moves a piece from given coordinates to given coordinates
	//Returns an error if destination was occupied.
	MovePiece(from, to coord.ChessCoordinates) error
}

type Builder interface {
	AddPiece(p piece.Piece, at coord.ChessCoordinates) Builder
	Build() Board
}