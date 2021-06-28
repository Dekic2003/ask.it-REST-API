CREATE TABLE users (
    id int unsigned not null auto_increment,
    name varchar(255),
    surname varchar(255),
    email varchar (255),
    password varchar(255),
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp,

    primary key (id)
);
CREATE TABLE question(
    id int unsigned not null auto_increment,
    author_id int unsigned not null,
    question text,
    likes int unsigned,
    dislikes int unsigned,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp,

    primary key (id),
    foreign key (author_id) references users(id)
);
CREATE TABLE answer(
    id int unsigned not null auto_increment,
    question_id int unsigned not null,
    author_id int unsigned not null,
    likes int unsigned,
    dislikes int unsigned,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp,

    primary key (id),
    foreign key (question_id) references question(id),
    foreign key (author_id) references users(id)
);
