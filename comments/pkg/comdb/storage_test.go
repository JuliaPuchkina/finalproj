package comdb

import (
	"testing"
)

func TestStore_AddComm(t *testing.T) {
	dataBase, err := New("postgres://postgres:postgres@127.0.0.1/newscomm")
	comment := Comment{
		Content:  "Task Content",
		NewsID:   0,
		ParentID: 0,
	}
	id := dataBase.AddComm(comment)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Создан комментарий с id:", id)
}

// этот тест долго думала, как писать, и в итоге, мне кажется, получилось что-то не то
func TestStore_Comments(t *testing.T) {
	dataBase, _ := New("postgres://postgres:postgres@127.0.0.1/newscomm")
	news, err := dataBase.Comments(1)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", news)

}
