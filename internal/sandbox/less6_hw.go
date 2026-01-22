package sandbox

import (
	. "fmt"
	. "github.com/Sazikoff/hexlet-struct/internal/structure"
)

func Run6() {
	draft := Draft{Title: "Winter combo"}
	draft.AddItem("Tea")
	draft.AddItem("")
	draft.AddItem("Mug")
	draft.Publish()

	Println(draft.Items) // []string{"Tea", "Mug"}
	Println(draft.IsPublished())  // true

	copyDraft := draft
	Println(copyDraft.IsPublished()) // true
}
