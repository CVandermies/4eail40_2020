package coord

import "testing"

func TestNewCartesian(t *testing.T) {
	c := NewCartesian(1, 2)
	if (c.x != 1) || (c.y != 2) {
		//test success
	} else {
		t.Errorf("Expected 1 and 2 as coordinate")
	}

}

func TestCartesian_Coord(t *testing.T) {
	c := NewCartesian(1, 2)

	tests := map[int]int{
		0: 1,
		1:2,
	}
	
	for n, want := range tests{
		t.Run(string(rune(n)), func(t *testing.T){
		got, err := c.Coord(n)
		if err != nil {
			t.Error(err)
		}
		if got != want {
			t.Errorf("expected %d, but got %d", want, got)
		}

		})
}
	
t.Run("err")