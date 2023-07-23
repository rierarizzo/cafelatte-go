drop database if exists cafelatte;

create database cafelatte;

use cafelatte;

create table userrole
(
    Code        char        not null,
    Description varchar(20) not null,
    primary key (Code)
);

insert into userrole
values ('A', 'ADMIN');
insert into userrole
values ('E', 'EMPLOYEE');
insert into userrole
values ('C', 'CLIENT');

create table user
(
    ID          int auto_increment,
    Username    varchar(30)  not null unique,
    Name        varchar(50)  not null,
    Surname     varchar(50)  not null,
    PhoneNumber varchar(20),
    Email       varchar(50)  not null unique,
    Password    varchar(100) not null,
    RoleCode    char         not null,
    Status      bool     default true,
    CreatedAt   datetime default current_timestamp,
    UpdatedAt   datetime default current_timestamp on update current_timestamp,
    primary key (ID),
    foreign key (RoleCode) references userrole (Code)
);

create table province
(
    ID   int         not null,
    Name varchar(50) not null,
    primary key (ID)
);

insert into province
    (ID, Name)
values (1, 'GUAYAS');
insert into province
    (ID, Name)
values (2, 'SANTA ELENA');

create table city
(
    ID         int         not null,
    ProvinceID int         not null,
    Name       varchar(50) not null,
    primary key (ID, ProvinceID),
    foreign key (ProvinceID) references province (ID)
);

insert into city (ID, ProvinceID, Name)
values (1, 1, 'GUAYAQUIL');
insert into city (ID, ProvinceID, Name)
values (2, 1, 'DURAN');
insert into city (ID, ProvinceID, Name)
values (3, 1, 'YAGUACHI');
insert into city (ID, ProvinceID, Name)
values (4, 1, 'MILAGRO');
insert into city (ID, ProvinceID, Name)
values (1, 2, 'SALINAS');
insert into city (ID, ProvinceID, Name)
values (2, 2, 'SANTA ELENA');

create table addresstype
(
    Code        char,
    Description varchar(20),
    primary key (Code)
);

insert into addresstype
values ('D', 'DOMICILIO');
insert into addresstype
values ('T', 'TRABAJO');

create table useraddress
(
    ID         int auto_increment,
    Type       char         not null,
    UserID     int          not null,
    ProvinceID int          not null,
    CityID     int          not null,
    PostalCode varchar(10),
    Detail     varchar(150) not null,
    Status     bool     default true,
    CreatedAt  datetime default current_timestamp,
    UpdatedAt  datetime default current_timestamp on update current_timestamp,
    primary key (ID, UserID),
    foreign key (Type) references addresstype (Code),
    foreign key (UserID) references user (ID),
    foreign key (ProvinceID) references province (ID),
    foreign key (CityID) references city (ID)
);

create table cardcompany
(
    ID   int auto_increment,
    Name varchar(50),
    primary key (ID)
);

insert into cardcompany (Name)
values ('AMERICAN EXPRESS');
insert into cardcompany (Name)
values ('VISA');
insert into cardcompany (Name)
values ('MASTERCARD');
insert into cardcompany (Name)
values ('DISCOVER');

create table cardtype
(
    Code        char not null,
    Description varchar(50),
    primary key (Code)
);

insert into cardtype (Code, Description)
values ('C', 'CREDITO');
insert into cardtype (Code, Description)
values ('D', 'DEBITO');

create table userpaymentcard
(
    ID              int auto_increment,
    Type            char         not null,
    UserID          int          not null,
    Company         int          not null,
    HolderName      varchar(100) not null,
    Number          varchar(100) not null,
    ExpirationYear  int          not null,
    ExpirationMonth int          not null,
    CVV             varchar(100) not null,
    Status          bool     default true,
    CreatedAt       datetime default current_timestamp,
    UpdatedAt       datetime default current_timestamp on update current_timestamp,
    primary key (ID),
    foreign key (Type) references cardtype (Code),
    foreign key (UserID) references user (ID),
    foreign key (Company) references cardcompany (ID)
);
