use test;

CREATE TABLE users (
  id int(10) unsigned not null primary key auto_increment,
  name varchar(255) not null default '',
  display_name varchar(255) not null default '',
  email varchar(255) not null default ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE projects (
  id int(10) unsigned not null auto_increment primary key,
  name varchar(255) not null default '',
  description varchar(255) not null default ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE issues (
  id int(10) unsigned not null auto_increment primary key,
  project_id int(10) unsigned not null,
  name varchar(255) not null default '',
  status boolean not null default false,
  foreign key (project_id) references projects (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE comments (
  id int(10) unsigned not null auto_increment primary key,
  issue_id int(10) unsigned not null,
  user_id int(10) unsigned not null,
  comment text not null,
  foreign key (issue_id) references issues (id),
  foreign key (user_id) references users (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;