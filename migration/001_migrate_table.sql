CREATE TABLE users(
    id serial primary key not null,
    username varchar(255) not null unique, 
    email varchar(255) not null unique,  
    password varchar(255) not null,
    Age int not null,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now(),

);

CREATE TABLE photos(
    id serial primary key not null,
    title varchar(255) not null, 
    url text not null,  
    caption text,
    photo_url text not null,
    user_id  int not null, 
    created_at timestamp not null default now(),
    updated_at timestamp not null default now(),
 
    constraint fk_photo_user_id 
        foreign key (user_id) 
        references users(id)
);

CREATE TABLE comments(
    id serial primary key not null,
    message text not null, 
    user_id  int not null, 
    photo_id  int not null, 
    message text not null,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now(),
   
    constraint fk_comments_photo_id 
        foreign key (photo_id) 
        references photos(id),
    constraint fk_comments_user_id 
        foreign key (user_id) 
        references users(id)
);

CREATE TABLE social_medias(
    id serial primary key not null,
    name varchar(255) not null, 
    social_media_url text not null, 
    user_id  int not null, 
    created_at timestamp not null default now(),
    updated_at timestamp not null default now(),
    
    constraint fk_social_medias_user_id 
        foreign key (user_id) 
        references users(id)
);