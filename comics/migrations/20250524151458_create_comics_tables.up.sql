CREATE TABLE genres (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE comics (
    id SERIAL PRIMARY KEY,
    author_id INTEGER NOT NULL,
    translator_id INTEGER NOT NULL,
    artist_id INTEGER NOT NULL,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    cover_image VARCHAR(255),
    status VARCHAR(100),
    comic_release_date TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    views INTEGER DEFAULT 0,
    rating DOUBLE PRECISION DEFAULT 0
);

CREATE TABLE alt_titles (
    id SERIAL PRIMARY KEY,
    comic_id INTEGER REFERENCES comics(id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL
);

CREATE TABLE comic_genres (
    comic_id INTEGER REFERENCES comics(id) ON DELETE CASCADE,
    genre_id INTEGER REFERENCES genres(id) ON DELETE CASCADE,
    PRIMARY KEY (comic_id, genre_id)
);

CREATE TABLE chapters (
    id SERIAL PRIMARY KEY,
    comic_id INTEGER REFERENCES comics(id) ON DELETE CASCADE,
    title VARCHAR(255),
    number DOUBLE PRECISION,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    likes INTEGER DEFAULT 0,
    dislikes INTEGER DEFAULT 0
);

CREATE TABLE pages (
    id SERIAL PRIMARY KEY,
    chapter_id INTEGER REFERENCES chapters(id) ON DELETE CASCADE,
    image_url VARCHAR(255),
    page_num INTEGER
);

CREATE TABLE comments (
    id SERIAL PRIMARY KEY,
    page_id INTEGER REFERENCES pages(id) ON DELETE CASCADE,
    user_id INTEGER NOT NULL,
    user_name VARCHAR(255),
    comment TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    likes INTEGER DEFAULT 0,
    dislikes INTEGER DEFAULT 0
);

CREATE TABLE replies (
    id SERIAL PRIMARY KEY,
    comment_id INTEGER REFERENCES comments(id) ON DELETE CASCADE,
    user_id INTEGER NOT NULL,
    user_name VARCHAR(255),
    reply TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    likes INTEGER DEFAULT 0,
    dislikes INTEGER DEFAULT 0
);
