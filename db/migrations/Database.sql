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
CREATE TABLE questions(
    id int unsigned not null auto_increment,
    author_id int unsigned not null,
    question text,
    likes int unsigned not null DEFAULT 0,
    dislikes int unsigned not null DEFAULT 0,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp,

    primary key (id),
    foreign key (author_id) references users(id)
);
CREATE TABLE answers(
    id int unsigned not null auto_increment,
    question_id int unsigned not null,
    author_id int unsigned not null,
    answer text,
    likes int unsigned not null DEFAULT 0,
    dislikes int unsigned not null DEFAULT 0,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp,

    primary key (id),
    foreign key (question_id) references question(id),
    foreign key (author_id) references users(id)
);
CREATE TABLE question_reactions(
                       question_id int unsigned not null,
                       author_id int unsigned not null,

                       rating bool,

                       created_at timestamp default current_timestamp,
                       updated_at timestamp default current_timestamp on update current_timestamp,

                       primary key (question_id,author_id),
                       foreign key (question_id) references question(id),
                       foreign key (author_id) references users(id)
);
CREATE TABLE answer_reactions(
                     answer_id int unsigned not null,
                     author_id int unsigned not null,

                     rating bool,

                     created_at timestamp default current_timestamp,
                     updated_at timestamp default current_timestamp on update current_timestamp,

                     primary key (answer_id,author_id),
                     foreign key (answer_id) references answer(id),
                     foreign key (author_id) references users(id)
);
CREATE TABLE notifications(
    id int unsigned not null auto_increment,
    question_id int unsigned,
    question_author_id int unsigned,
    answer_author_id int unsigned,
    isRead bool DEFAULT false,

    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp,

    primary key (id),
    foreign key (question_id) references question(id),
    foreign key (question_author_id) references question(author_id),
    foreign key (answer_author_id) references answer(author_id)
)