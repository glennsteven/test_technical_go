-- +goose Up
-- +goose StatementBegin
CREATE TABLE consumer_limits (
    id_limit varchar(36) PRIMARY KEY,
    id_consumer varchar(36) not null,
    tenor int not null ,
    limit_amount decimal(10,2) not null default 0,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    KEY id_consumer(id_consumer),
    KEY tenor(tenor),
    KEY limit_amount(limit_amount),
    KEY created_at(created_at),
    KEY updated_at(updated_at)
)ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `consumer_limits`;
-- +goose StatementEnd
