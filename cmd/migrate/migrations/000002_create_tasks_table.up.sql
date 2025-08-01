CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    type_id INTEGER NOT NULL REFERENCES task_types(id),
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    date TIMESTAMP NOT NULL,
    total_units INTEGER, -- e.g 310 pages, 120 minutes, 10 modules, 3 sub-tasks
    progress_units INTEGER DEFAULT 0
);