create table user_account(id int not null auto_increment primary key, domain varchar(20) not null, password varchar(20) not null);


create table IM_USER(ID int not null auto_increment primary key, 
    DOMAIN varchar(100) not null,
    PASSWORD varchar(50) not null, 
    FULL_NAME varchar(255) default ' ',
    EMAIL varchar(255) default ' ', 
    FIRST_NAME varchar(255) default ' ', 
    LAST_NAME varchar(255) default ' ', 
    PHONE_NUMBER varchar(255) default ' ', 
    ADDRESS varchar(255) default ' ', 
    IS_ACTIVE varchar(1) default '1');

create table IM_JWT_MANAGEMENT (ID int not null auto_increment primary key,
  TOKEN varchar(1000) not null,
  USER_ID int not null,
  CREATED_TIME TIMESTAMP
);

create table IM_ROLE(ID int not null auto_increment primary key, 
    ROLE_NAME varchar(100) not null,
    ROLE_DESCRIPTION varchar(50) not null, 
    ROLE_MANAGEMENT_DOMAIN varchar(255) default ' ',
    ROLE_ALIAS varchar(255) default ' ', 
    IS_ACTIVE varchar(1) default '1');
