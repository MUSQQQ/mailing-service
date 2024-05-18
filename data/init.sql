CREATE TABLE IF NOT EXISTS mailing_details (
    id SERIAL PRIMARY KEY,
    mailing_id INT NOT NULL,
    email VARCHAR (255) NOT NULL,
    title VARCHAR(255),
    content VARCHAR(255),
    insert_time DATE NOT NULL default CURRENT_DATE
)
