echo ".exit" | sqlite3 database.db 


sqlite-utils query database.db "CREATE TABLE IF NOT EXISTS nwita (
    id INTEGER PRIMARY KEY,
    content TEXT NOT NULL
    )"