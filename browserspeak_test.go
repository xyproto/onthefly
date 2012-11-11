package browserspeak

import "testing"

func TestSVG(t *testing.T) {
	svg := NewExampleSVG()
	svg.GetHTML(false)
	//s := svg.GetHTML(false)
	t.Log("hi")
	//t.Errorf("%s\n", s)
	//const in, out = 4, 2
	//if x := Sqrt(in); x != out {
	//	t.Errorf("Sqrt(%v) = %v, want %v", in, x, out)
	//}
}

func TestSVG2(t *testing.T) {
	svg := NewExampleSVG2()
	s := svg.GetHTML(false)
	t.Errorf("%s\n", s)
}
