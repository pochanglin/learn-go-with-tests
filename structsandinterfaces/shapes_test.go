package structsandinterfaces

import (
	"testing"
)

func TestPerimeter(t *testing.T) {

	t.Run("in rectangle", func(t *testing.T) {
		r := Rectangle{
			width:  10.0,
			height: 10.0,
		}
		got := r.Perimeter()
		want := 40.0
		AssertMsg(t, got, want)
	})

	t.Run("in circle", func(t *testing.T) {
		c := Circle{
			radius: 10.0,
		}
		got := c.Perimeter()
		want := 40.0
		AssertMsg(t, got, want)
	})
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{
			"Rectangle",
			Rectangle{
				width:  10.0,
				height: 10.0,
			},
			10.0,
		},
		{
			"Circle",
			Circle{
				radius: 10.0,
			},
			314.1592653589793,
		},
	}

	for _, tt := range areaTests {
		// using tt.name from the case to use it as the `t.Run` test name
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})
	}
}

func AssertMsg(t testing.TB, got float64, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
		// The f is for our float64 and the .2 means print 2 decimal places
	}
}
