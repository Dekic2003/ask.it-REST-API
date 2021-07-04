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
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp,

    primary key (id),
    foreign key (question_id) references questions(id),
    foreign key (author_id) references users(id)
);
CREATE TABLE question_reactions(
                       question_id int unsigned not null,
                       author_id int unsigned not null,

                       rating bool,

                       created_at timestamp default current_timestamp,
                       updated_at timestamp default current_timestamp on update current_timestamp,

                       primary key (question_id,author_id),
                       foreign key (question_id) references questions(id),
                       foreign key (author_id) references users(id)
);
CREATE TABLE answer_reactions(
                     answer_id int unsigned not null,
                     author_id int unsigned not null,

                     rating bool,

                     created_at timestamp default current_timestamp,
                     updated_at timestamp default current_timestamp on update current_timestamp,

                     primary key (answer_id,author_id),
                     foreign key (answer_id) references answers(id),
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
    foreign key (question_id) references questions(id),
    foreign key (question_author_id) references questions(author_id),
    foreign key (answer_author_id) references answers(author_id)
);
ALTER TABLE answer_reactions DROP CONSTRAINT answer_reactions.answer_reactions_ibfk_1;
ALTER TABLE notifications DROP CONSTRAINT notifications.notifications_ibfk_3;
ALTER TABLE question_reactions DROP CONSTRAINT question_reactions.question_reactions_ibfk_1;
ALTER TABLE notifications DROP CONSTRAINT notifications.notifications_ibfk_1;
ALTER TABLE notifications DROP CONSTRAINT notifications.notifications_ibfk_2;