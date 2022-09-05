package Test

import "testing"

/*func TestSuma(t *testing.T) {
	prueba := Suma(3, 6)
	if prueba != 9 {
		t.Errorf("suma incorrecta,tiene %d se esperaba %d", prueba, 10)
	}

}*/

func TestSuma(t *testing.T) {

	tabla := []struct {
		a int
		b int
		c int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{25, 25, 50},
	}

	for _, item := range tabla {
		total := Suma(item.a, item.b)
		if total != item.c {
			t.Errorf("suma incorrecta,tiene %d se esperaba %d", total, item.c)
		}

	}

}
