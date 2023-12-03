USE sample;

DROP TABLE IF EXISTS todo;

CREATE TABLE todo
(
    id          INT           PRIMARY KEY AUTO_INCREMENT,
    title       VARCHAR(50)   NOT NULL,
    description VARCHAR(1000),
    created     DATETIME,
    updated     DATETIME,
    deadline    DATETIME,
    priority    VARCHAR(5),
    category    VARCHAR(50),
    status      VARCHAR(10),
    assignee    VARCHAR(20)
);

INSERT INTO todo (title, description, created, updated, deadline, priority, category, status, assignee)
VALUES ("買い物", "カレーの材料を買う", "2023-11-30 10:10:10", "2023-11-30 10:10:10", "2023-12-10 00:00:00", "高", "家事", "未完了", "母");
INSERT INTO todo (title, description, created, updated, deadline, priority, category, status, assignee)
VALUES ("散歩", "犬の散歩", "2023-11-30 10:10:10", "2023-11-30 10:10:10", "2023-11-30 12:00:00", "中", "その他", "未完了", "兄");
