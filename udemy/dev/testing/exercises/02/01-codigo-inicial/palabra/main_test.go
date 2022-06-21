package palabra

import (
	"01-codigo-inicial/cita"
	"fmt"
	"reflect"
	"testing"
)

func TestConteoUso(t *testing.T) {
	expRes := map[string]int{
		"Ahora,":         1,
		"Amor":           1,
		"Autenticidad":   1,
		"Autoestima":     1,
		"Cuando":         9,
		"De":             1,
		"Hoy":            9,
		"Humildad":       1,
		"Madurez":        1,
		"No":             1,
		"Pero":           1,
		"Plenitud":       1,
		"Propio":         1,
		"Respeto":        1,
		"Saber":          1,
		"Sencillez":      1,
		"Simplicidad":    1,
		"Todo":           1,
		"Vivir":          1,
		"Y":              1,
		"a":              7,
		"abajo.":         1,
		"abandoné":       1,
		"aceptar":        1,
		"acontece,":      1,
		"acontece.":      1,
		"actitud":        1,
		"al":             2,
		"alguna":         1,
		"aliado.":        1,
		"amé":            9,
		"angustia,":      1,
		"aquello":        1,
		"así":            1,
		"atormentarme":   1,
		"aun":            1,
		"caos":           1,
		"chocan,":        1,
		"circunstancia,": 1,
		"coloco":         1,
		"comencé":        3,
		"comprendí":      1,
		"contra":         1,
		"contribuye":     1,
		"corazón,":       1,
		"correcta,":      1,
		"correcto,":      2,
		"cosa":           1,
		"crecimiento.":   1,
		"cualquier":      2,
		"cuando":         2,
		"cuestionarnos,": 1,
		"de":             23,
		"debemos":        1,
		"decepcionarme.": 1,
		"dejé":           2,
		"del":            1,
		"descubrí":       1,
		"desear":         1,
		"deseo,":         1,
		"desistí":        3,
		"diferente,":     1,
		"donde":          1,
		"día":            1,
		"egoísmo.":       1,
		"el":             7,
		"ella":           1,
		"emocional,":     1,
		"empujara":       1,
		"en":             5,
		"encuentro":      1,
		"entonces,":      1,
		"erré":           1,
		"es":             4,
		"esa":            1,
		"eso":            8,
		"estaba":         1,
		"estrellas.":     1,
		"está":           1,
		"es…":            5,
		"exacto,":        1,
		"forzar":         1,
		"fuera":          1,
		"fuese":          1,
		"futuro.":        2,
		"gran":           1,
		"grandes":        1,
		"gusta,":         1,
		"hacer":          1,
		"hacia":          1,
		"hago":           1,
		"hasta":          1,
		"hecho":          1,
		"hora":           1,
		"inclusive":      1,
		"inicio":         1,
		"la":             7,
		"las":            1,
		"librarme":       1,
		"libre":          1,
		"llama…":         3,
		"llamó":          1,
		"lo":             4,
		"los":            2,
		"lugar":          1,
		"mantengo":       1,
		"mayoría":        1,
		"me":             12,
		"mega-proyectos": 1,
		"menos":          1,
		"mente":          1,
		"mi":             8,
		"miedo":          1,
		"mis":            1,
		"mismo.":         1,
		"momento":        1,
		"momento,":       1,
		"nacer":          1,
		"no":             4,
		"nombre":         1,
		"nombre…":        1,
		"o":              2,
		"ofensivo":       1,
		"para":           1,
		"pasado,":        1,
		"percibir":       2,
		"percibí":        1,
		"persona":        1,
		"persona,":       1,
		"personas,":      1,
		"planes,":        1,
		"planetas":       1,
		"por":            1,
		"preocupándome":  1,
		"preparada,":     1,
		"presente,":      1,
		"propias":        1,
		"propio":         1,
		"pude":           2,
		"puede":          1,
		"que":            20,
		"quedarme":       1,
		"querer":         1,
		"quiero,":        1,
		"razón":          1,
		"razón,":         1,
		"realizar":       1,
		"relajarme.":     1,
		"reviviendo":     1,
		"ritmo.":         1,
		"sabiendo":       1,
		"saludable:":     1,
		"se":             3,
		"servicio":       1,
		"señal":          1,
		"siempre":        1,
		"sino":           1,
		"situaciones":    1,
		"situación,":     1,
		"suelen":         1,
		"sufrimiento":    1,
		"sé":             4,
		"sólo":           1,
		"temer":          1,
		"tener":          2,
		"tiempo":         1,
		"tiene":          2,
		"todo":           2,
		"tratar":         1,
		"un":             2,
		"una":            2,
		"valioso":        1,
		"veces.":         1,
		"verdad":         1,
		"verdad,":        8,
		"verdades.":      1,
		"vez.":           1,
		"vida":           2,
		"vivo":           1,
		"voy":            1,
		"y":              14,
		"yo":             2,
	}

	r := ConteoUso(cita.Cuando)

	if !reflect.DeepEqual(expRes, r) {
		t.Error("Expected", expRes, "GOT", r)
	}
}

func ExampleConteoUso() {
	fmt.Println(ConteoUso("welcome to jungle"))
	// Output:
	// map[jungle:1 to:1 welcome:1]
}

func BenchmarkConteoUso(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConteoUso(cita.Cuando)
	}
}

func TestConteo(t *testing.T) {
	expRes := 3
	r := Conteo("welcome to jungle")

	if r != expRes {
		t.Error("Expected", expRes, "GOT", r)
	}
}

func ExampleConteo() {
	fmt.Println(Conteo("welcome to jungle"))
	// Output
	// 3
}

func BenchmarkConteo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Conteo("welcome to jungle")
	}
}
