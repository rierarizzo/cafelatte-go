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
values ('A', 'Admin');
insert into userrole
values ('E', 'Employee');
insert into userrole
values ('C', 'Client');

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
values (1, 'Guayas');
insert into province
    (ID, Name)
values (2, 'Santa Elena');

create table city
(
    ID         int         not null,
    ProvinceID int         not null,
    Name       varchar(50) not null,
    primary key (ID, ProvinceID),
    foreign key (ProvinceID) references province (ID)
);

insert into city (ID, ProvinceID, Name)
values (1, 1, 'Guayaquil');
insert into city (ID, ProvinceID, Name)
values (2, 1, 'Durán');
insert into city (ID, ProvinceID, Name)
values (3, 1, 'Yaguachi');
insert into city (ID, ProvinceID, Name)
values (4, 1, 'Milagro');
insert into city (ID, ProvinceID, Name)
values (1, 2, 'Salinas');
insert into city (ID, ProvinceID, Name)
values (2, 2, 'Santa Elena');

create table addresstype
(
    Code        char,
    Description varchar(20),
    primary key (Code)
);

insert into addresstype
values ('H', 'Home');
insert into addresstype
values ('W', 'Work');

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
values ('American Express');
insert into cardcompany (Name)
values ('Visa');
insert into cardcompany (Name)
values ('Mastercard');
insert into cardcompany (Name)
values ('Discover');

create table cardtype
(
    Code        char not null,
    Description varchar(50),
    primary key (Code)
);

insert into cardtype (Code, Description)
values ('C', 'Credit');
insert into cardtype (Code, Description)
values ('D', 'Debit');

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

create table productcategory
(
    Code        varchar(6),
    Description varchar(50) not null,
    primary key (Code)
);

insert into productcategory (Code, Description)
values ('HOTBEV', 'Hot beverage');
insert into productcategory (Code, Description)
values ('COLBEV', 'Cold beverage');
insert into productcategory (Code, Description)
values ('DESSER', 'Dessert');

create table product
(
    ID           int auto_increment,
    Name         varchar(50)   not null,
    Description  varchar(150)  not null,
    ImageURL     varchar(400),
    Price        decimal(6, 2) not null,
    CategoryCode varchar(6)    not null,
    Stock        int           not null,
    Status       bool     default true,
    CreatedAt    datetime default current_timestamp,
    UpdatedAt    datetime default current_timestamp on update current_timestamp,
    primary key (ID),
    foreign key (CategoryCode) references productcategory (Code)
);
