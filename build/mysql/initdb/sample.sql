USE sample;

DROP TABLE IF EXISTS task;

CREATE TABLE task
(
    id          INT          NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name        VARCHAR(50),
    description VARCHAR(1000),
    created     DATETiME,
    updated     DATETiME,
    deadline    DATETiME,
    priority    VARCHAR(5),
    category    VARCHAR(50),
    status      VARCHAR(10),
    assignee    VARCHAR(20)
);

INSERT INTO task (name, description, created, updated, deadline, priority, category, status, assignee)
VALUES ("買い物", "カレーの材料を買う", "2023-11-30 10:10:10", "2023-11-30 10:10:10", "2023-12-10 00:00:00", "高", "家事", "未完了", "母");
VALUES ("散歩", "犬の散歩", "2023-11-30 10:10:10", "2023-11-30 10:10:10", "2023-11-30 12:00:00", "中", "その他", "未完了", "兄");
