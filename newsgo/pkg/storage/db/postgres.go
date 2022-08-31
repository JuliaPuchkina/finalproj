package db

import (
	"context"
	storage "finalproj/newsgo/pkg/storage"
	"fmt"
	"math"

	"github.com/jackc/pgx/v4/pgxpool"
)

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

// количество новостей на странице по-дурацки пока будет здесь
var num float64 = 15

var page struct {
	pages    int // количество страниц
	curPage  int // номер текущей страницы
	newsPage int // количество элементов на страницу
}

// Posts выводит все существующие публикации
func (s *Store) Posts(n int) ([]storage.Post, error) {
	f := float64(n)
	rows, err := s.db.Query(context.Background(), `
	SELECT id, title, content, published, link FROM news
	ORDER BY published DESC
	LIMIT $1
	OFFSET $2;
	`,
		num,
		num*(f-1),
	)
	if err != nil {
		return nil, err
	}

	var posts []storage.Post
	// итерирование по результату выполнения запроса
	// и сканирование каждой строки в переменную
	for rows.Next() {
		var p storage.Post
		err = rows.Scan(
			&p.ID,
			&p.Title,
			&p.Content,
			&p.PubTime,
			&p.Link,
		)
		if err != nil {
			return nil, err
		}
		// добавление переменной в массив результатов
		posts = append(posts, p)

	}

	// считаем кол-во публикаций, чтобы потом посчитать кол-во страниц
	numRows, err := s.db.Query(context.Background(), `
	SELECT count(*) FROM news;
	`,
		num,
		num*(f-1),
	)
	if err != nil {
		return nil, err
	}

	// считаем кол-во страниц и добавляем данные в структуру объекта пагинации
	var numNews float64
	err = numRows.Scan(&numNews)
	page.pages = int(math.Ceil(numNews / num))
	page.curPage = n
	page.newsPage = int(num)
	if err != nil {
		return nil, err
	}

	fmt.Println(page)

	// ВАЖНО не забыть проверить rows.Err()
	return posts, rows.Err()
}

// AddPost создает новую публикацию
func (s *Store) AddPost(p storage.Post) error {
	err := s.db.QueryRow(context.Background(), `
		INSERT INTO news (title, content, published, link)
		VALUES ($1, $2, $3, $4);
		`,
		p.Title,
		p.Content,
		p.PubTime,
		p.Link,
	).Scan()
	return err
}

// NewsTitle осуществляет поиск по названию новости
func (s *Store) NewsTitle(line string) ([]storage.Post, error) {
	rows, err := s.db.Query(context.Background(), `
	SELECT id, title, content, published, link FROM news
	WHERE title ILIKE '%$1%';
	`,
		line,
	)
	if err != nil {
		return nil, err
	}

	var posts []storage.Post
	// итерирование по результату выполнения запроса
	// и сканирование каждой строки в переменную
	for rows.Next() {
		var p storage.Post
		err = rows.Scan(
			&p.ID,
			&p.Title,
			&p.Content,
			&p.PubTime,
			&p.Link,
		)
		if err != nil {
			return nil, err
		}
		// добавление переменной в массив результатов
		posts = append(posts, p)

	}
	// ВАЖНО не забыть проверить rows.Err()
	return posts, rows.Err()
}
