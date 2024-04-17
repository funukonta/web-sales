CREATE TABLE IF NOT EXISTS Sales (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) REFERENCES Users(email),
    transactiondate DATE,
    sales_type VARCHAR(50),
    nominal NUMERIC(15, 2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


insert into Sales (email,transactiondate,sales_type,nominal) values
('evanroy36@gmail.com','2024-04-16','JASA',10000),
('evanroy36@gmail.com','2024-04-15','Barang',20000);