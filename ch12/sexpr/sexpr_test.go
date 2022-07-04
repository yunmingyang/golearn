package sexpr

import (
	"fmt"
	"log"
	"testing"
)



type Movie struct {
	Title, Subtitle string
	Year int
	Actor map[string]string
	Oscars []string
}

func TestSExpr(t *testing.T) {
	strangeLove := Movie {
		Title: "Dr.Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year: 1964,
		Actor: map[string]string{
			"Dr. Strangelove": "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley": "Peter Sellers",
			"Gen. Buck Turgidson": "George C. Scott",
			"Brig. Gen. Jack D. Ripper": "Sterling Hayden",
			`Maj. T.J. "King" Kong`: "Slim Pickens",
		},
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}

	b, err := Marshal(strangeLove)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", b)

	fmt.Println("----------------------------")

	var out Movie
	err = Unmarshal(b, &out)
	if err != nil {
		t.Fatalf("unmarshal error: %v\n", err)
	}
	fmt.Printf("%v\n", out.Actor)
	fmt.Printf("%v\n", out.Oscars)
	fmt.Printf("%v\n", out.Year)
	fmt.Printf("%v\n", out.Title)
	fmt.Printf("%v\n", out.Subtitle)
}