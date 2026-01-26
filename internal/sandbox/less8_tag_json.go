package sandbox

import (
	"encoding/json"
	"fmt"
)

type InvoiceDTO struct {
ID int  `json:"id"`
Customer string  `json:"customer"`
Amount float64  `json:"amount"`
Comment string  `json:"comment,omitempty"`
Secret string  `json:"-"`
}

type User struct {
	Name string `json:"username"`
	Age  int    `json:"years"`
}

func EncodeInvoice(inv InvoiceDTO) ([]byte, error) {
	return json.Marshal(inv)
}

func DecodeInvoice(payload []byte) (InvoiceDTO, error) {
	var i InvoiceDTO
	err := json.Unmarshal(payload, &i)
	return i, err
}

func Run8() {
	jsonStr := `{"username":"Мария","years":25}`
	fmt.Printf("%T", jsonStr, )
	var u User
	json.Unmarshal([]byte(jsonStr), &u)

	fmt.Println(u)

}
