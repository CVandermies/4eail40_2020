//Package piece handles game piece
package piece

import (
	"fmt"

	"C:\Users\User\Documents\ECAM\4e_Architecture\4eail40_2020\exercises\chess\model\coord"
	"C:\Users\User\Documents\ECAM\4e_Architecture\4eail40_2020\exercises\chess\model\player"
)

//Piece represents a game piece
type Piece interface {
	fmt.Stringer
	//Color returns the appartenance
	Color() player.Color
	//Moves returns a set of valid moves
	Moves() map[coord.ChessCoordinates]bool
}

func test() {
	//var p Piece
	//Check if move is contained in the map
	//if p.Moves[my_coord]{
		//Check if it is contained, if yes break out of the loop
}