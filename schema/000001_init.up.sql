CREATE TABLE users 
(
    id              serial          not null unique,
    username        varchar(255)    not null,
    password_hash   varchar(255)    not null,
    name            varchar(255)    not null,
    img_url         varchar(255)
);

CREATE TABLE chats 
(
    id              serial                                         not null unique,
    user1_id        int references users (id) on delete cascade    not null,
    user2_id        int references users (id) on delete cascade    not null
);

CREATE TABLE messages 
(
    id              serial                                         not null unique,
    chat_id         int references chats (id) on delete cascade    not null,
    sender_id       int references users (id) on delete cascade    not null,
    content         varchar(255)                                   not null,
    created_at      timestamp                                      not null,
    state           int                                            not null,
    type            int                                            not null
);

CREATE TABLE profiles 
(
    id              serial                                         not null unique,
    user_id         int references users (id) on delete cascade    not null,
    descript        varchar(255),
    age             int,
    country         varchar(255),
    city            varchar(255)
);