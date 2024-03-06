-- +goose Up
-- +goose StatementBegin
CREATE TABLE `user` (
    id INT NOT NULL AUTO_INCREMENT, 
    createdAt datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    updatedAt datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
    firstName varchar(50) NOT NULL,
    lastName varchar(50) NOT NULL,
    username varchar(30) NOT NULL,
    email varchar(100) NOT NULL,
    password varchar(255) NOT NULL,
    birthDate datetime NOT NULL,
    isActive tinyint NOT NULL DEFAULT 0,
    activationCode varchar(255) NOT NULL,
    UNIQUE INDEX IDX_U_username (username),
    UNIQUE INDEX IDX_U_email (email),
    PRIMARY KEY (id)
);

-- Set auto increment start number. 286 was choosen randomly
ALTER TABLE user AUTO_INCREMENT=286;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IDX_U_username ON  user;
DROP INDEX IDX_U_email ON user;
DROP TABLE user;
-- +goose StatementEnd
