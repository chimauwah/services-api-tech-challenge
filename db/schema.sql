-- name: create-employee-table
CREATE TABLE Employee (
  id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
  enter_ts datetime NOT NULL,
  last_change_ts datetime DEFAULT NULL,
  active INTEGER NOT NULL,
  first_name varchar(255) DEFAULT NULL,
  last_name varchar(255) DEFAULT NULL,
  address1 varchar(255) DEFAULT NULL,
  address2 varchar(255) DEFAULT NULL,
  city varchar(255) DEFAULT NULL,
  state varchar(2) DEFAULT NULL,
  zip varchar(20) DEFAULT NULL,
  cell_phone varchar(25) DEFAULT NULL,
  home_phone varchar(25) DEFAULT NULL,
  picture blob,
  title varchar(255) DEFAULT NULL,
  role INTEGER DEFAULT NULL,
  ip_phone varchar(25) DEFAULT NULL,
  samaccountname varchar(255) DEFAULT NULL,
  mail varchar(255) DEFAULT NULL,
  primary_pa varchar(255) DEFAULT NULL,
  secondary_pa varchar(255) DEFAULT NULL,
  office varchar(255) DEFAULT NULL,
  manager_dn varchar(255) DEFAULT NULL,
  travel_pref varchar(255) DEFAULT NULL,
  manager_samaccountname varchar(255) DEFAULT NULL,
  last_hash integer DEFAULT NULL,
  image_hash varchar(100) DEFAULT NULL,
  nick_name varchar(255) DEFAULT NULL,
  client_loc varchar(255) DEFAULT NULL
);

-- name: create-coreskill-table
CREATE TABLE CoreSkill (
  id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
  enter_ts datetime NOT NULL,
  last_change_ts datetime DEFAULT NULL,
  active INTEGER NOT NULL,
  skill varchar(255) NOT NULL,
  sequence INTEGER NOT NULL,
  employee_id INTEGER DEFAULT NULL,
  proficiency varchar(255) DEFAULT NULL,
  -- UNIQUE KEY id2 (id),
  CONSTRAINT FK182F35D2476C24D4 FOREIGN KEY (employee_id) REFERENCES Employee (id)
);