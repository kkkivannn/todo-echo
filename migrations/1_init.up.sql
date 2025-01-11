-- Создаем таблицу задач
CREATE TABLE IF NOT EXISTS Tasks
(
    id             INTEGER PRIMARY KEY AUTOINCREMENT,                        -- Автоинкрементируемый первичный ключ
    title          TEXT NOT NULL,                                            -- Заголовок задачи
    body           TEXT NOT NULL DEFAULT '',                                 -- Описание задачи
    created_at     TIMESTAMP     DEFAULT CURRENT_TIMESTAMP,                  -- Дата создания задачи
    task_status_id INTEGER,                                                  -- Ссылка на статус задачи
    FOREIGN KEY (task_status_id) REFERENCES Statuses (id) ON DELETE SET NULL -- Внешний ключ на таблицу Statuses
);

-- Создаем таблицу статусов
CREATE TABLE IF NOT EXISTS Statuses
(
    id     INTEGER PRIMARY KEY AUTOINCREMENT, -- Автоинкрементируемый первичный ключ
    status TEXT NOT NULL                      -- Название статуса задачи
);
