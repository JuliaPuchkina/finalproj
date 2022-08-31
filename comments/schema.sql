DROP TABLE IF EXISTS comments;

-- комментарии
CREATE TABLE comments (
    id SERIAL PRIMARY KEY,
	content TEXT NOT NULL UNIQUE, -- текст комментария
    newsID BIGINT, -- id новости
	parentID BIGINT -- id родительского комментария 
);