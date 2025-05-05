create table if not exist tasks (
    id integer primary key autoincrement,
    user_id integer not null,
    title text not null,
    description text not null,
    date datetime not null,
    foreign key (user_id) references users (id) on delete cascade
);