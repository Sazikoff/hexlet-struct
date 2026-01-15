package structure

import "fmt"

type Movie struct {
	Name    string
	Release int
	Genre   []string
	Rating  float32
}


func LessonBook() {
	FF := Movie{
		Name:    "Война и мир",
		Release: 1234,
		Genre:   []string{"Ужасы", "Фантастика"},
		Rating:  4.9,
	}
	
	fmt.Printf("%v (%v) — жанры: %v — рейтинг: %.2f\n", FF.Name, FF.Release, FF.Genre, FF.Rating)
}

type Account struct {
	Id    int
	User string
	Balance  float64
}

var acc = Account{
	Id:   123,
	User: "Petrov",
	Balance:  0,
}

func Deposit(acc *Account, amount float64) {
	acc.Balance += amount
}

func LessonDeposit() {
	fmt.Println(acc.Balance)
	Deposit(&acc, 100)
	fmt.Println(acc.Balance)
}

type Parcel struct {
	ID          string
	Destination string
	Weight      float64
	Insured     bool
}

func Describe(p Parcel) string {
	// BEGIN (write your solution here)
	insured := "no"
	if p.Insured {
		insured = "yes"
	}
	return fmt.Sprintf("#%v -> %v (%.1f kg, insured: %s)", p.ID, p.Destination, p.Weight, insured)
	// END
}

// var parcel1 = Parcel{ID: "BR-7", Destination: "Berlin", Weight: 12, Insured: true}
// fmt.Println(Describe(parcel1)) // "#BR-7 -> Berlin (12.0 kg, insured: yes)"