CREATE TABLE users(
uid int NOT NULL AUTO_INCREMENT,
user_name VARCHAR(32),
pass_word VARCHAR(32),
PRIMARY KEY (uid)
)
ALTER TABLE users AUTO_INCREMENT=100000;

CREATE TABLE users_token (
    token VARCHAR(64),
    uid int,
	CONSTRAINT FK_users_token FOREIGN KEY (uid)
    REFERENCES users(uid)
);

CREATE TABLE messages(
msg_id int NOT NULL AUTO_INCREMENT,
sender_id int,
receiver_id int,
conversatio_type TINYINT,
message_body_type TINYINT,
content VARCHAR(1024),
PRIMARY KEY (msg_id)
);
ALTER TABLE messages AUTO_INCREMENT=10000000;
