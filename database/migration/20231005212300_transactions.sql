-- +goose Up
-- +goose StatementBegin
CREATE TABLE transactions(
    id_transaction varchar(36) PRIMARY KEY,
    id_consumer varchar(36) not null ,
    contract_number varchar(40) not null,
    otr decimal(10,2) not null,
    fee_admin decimal(10, 2) not null ,
    installment_amount int not null,
    total_interest decimal(10, 2) not null,
    asset_name varchar(100) not null,
    transaction_date date not null,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    KEY contract_number(contract_number),
    KEY otr(otr),
    KEY fee_admin(fee_admin),
    KEY installment_amount(installment_amount),
    KEY total_interest(total_interest),
    KEY asset_name(asset_name),
    KEY transaction_date(transaction_date),
    KEY created_at(created_at),
    KEY updated_at(updated_at)
)ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `transactions`;
-- +goose StatementEnd
