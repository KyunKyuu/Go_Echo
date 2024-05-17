create table users(
    user_id serial primary key,
    name varchar(255) not null,
    email varchar(255) not null,
    password varchar(255) not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    deleted_at timestamp default current_timestamp
  
);

create table address(
    address_id serial primary key,
    city varchar(255) not null,
    province varchar(255) not null,
    postal_code varchar(255) not null,
    user_id_fk integer not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    CONSTRAINT fk_user FOREIGN KEY (user_id_fk) REFERENCES users(user_id)
    
);

