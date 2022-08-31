package comdb

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Структура комментария
type Comment struct {
	ID       int    // номер комментария
	Content  string // содержание комментария
	NewsID   int    // id новости, к которой относится комментарий
	ParentID int    // id родительского комментария, если комментарий является ответом на другой
}

// Interface задаёт контракт на работу с БД
type Interface interface {
	Comments(n int) ([]Comment, error) // получение комментариев по id публикации
	AddComm(Comment) error             // добавление нового комментария в базу

}

// хранилище данных
type Store struct {
	db *pgxpool.Pool
}

// конструктор объекта хранилища
func New(constr string) (*Store, error) {
	db, err := pgxpool.Connect(context.Background(), constr)
	if err != nil {
		return nil, err
	}
	s := Store{
		db: db,
	}
	return &s, nil
}

// Comments выводит все существующие комментарии к публикации по ее id
func (s *Store) Comments(n int) ([]Comment, error) {

	rows, err := s.db.Query(context.Background(), `
	SELECT id, content, newsID, parentID FROM comments
	WHERE newsID = $1
	ORDER BY id DESC
	`,
		n,
	)
	if err != nil {
		return nil, err
	}

	var comments []Comment
	// итерирование по результату выполнения запроса
	// и сканирование каждой строки в переменную
	for rows.Next() {
		var c Comment
		err = rows.Scan(
			&c.ID,
			&c.Content,
			&c.NewsID,
			&c.ParentID,
		)
		if err != nil {
			return nil, err
		}
		// добавление переменной в массив результатов
		comments = append(comments, c)

	}
	// ВАЖНО не забыть проверить rows.Err()
	return comments, rows.Err()
}

// AddComm создает новый комментарий
func (s *Store) AddComm(c Comment) error {
	err := s.db.QueryRow(context.Background(), `
		INSERT INTO comments (content, newsID, parentID)
		VALUES ($1, $2, $3);
		`,
		c.Content,
		c.NewsID,
		c.ParentID,
	).Scan()
	return err
}
