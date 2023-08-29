drop database if exists cafelatte;

create database cafelatte;

use cafelatte;

create table UserRole
(
    Code        char        not null,
    Description varchar(20) not null,
    primary key (Code)
);

create table User
(
    Id          int auto_increment,
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
    primary key (Id),
    foreign key (RoleCode) references UserRole (Code),
    index IdxUsername (Username), -- Índice en Username para búsquedas por nombre de usuario
    index IdxEmail (Email),       -- Índice en Email para búsquedas por correo electrónico
    index IdxStatus (Status)      -- Índice en Status para búsquedas por estado
);

create table Province
(
    Id   int         not null,
    Name varchar(50) not null,
    primary key (Id)
);

create table City
(
    Id         int         not null,
    ProvinceId int         not null,
    Name       varchar(50) not null,
    primary key (Id, ProvinceId),
    foreign key (ProvinceId) references Province (Id)
);

create table AddressType
(
    Code        char,
    Description varchar(20),
    primary key (Code)
);

create table UserAddress
(
    Id         int auto_increment,
    Type       char         not null,
    UserId     int          not null,
    ProvinceId int          not null,
    CityId     int          not null,
    PostalCode varchar(10),
    Detail     varchar(150) not null,
    Status     bool     default true,
    CreatedAt  datetime default current_timestamp,
    UpdatedAt  datetime default current_timestamp on update current_timestamp,
    primary key (Id, UserId),
    foreign key (Type) references AddressType (Code),
    foreign key (UserId) references User (Id),
    foreign key (ProvinceId) references Province (Id),
    foreign key (CityId) references City (Id),
    index IdxStatus (Status) -- Índice en Status para búsquedas por estado
);

create table CardCompany
(
    Id   int auto_increment,
    Name varchar(50),
    primary key (Id)
);

create table CardType
(
    Code        char not null,
    Description varchar(50),
    primary key (Code)
);

create table UserPaymentCard
(
    Id              int auto_increment,
    Type            char         not null,
    UserId          int          not null,
    Company         int          not null,
    HolderName      varchar(100) not null,
    Number          varchar(100) not null,
    ExpirationYear  int          not null,
    ExpirationMonth int          not null,
    CVV             varchar(100) not null,
    Status          bool     default true,
    CreatedAt       datetime default current_timestamp,
    UpdatedAt       datetime default current_timestamp on update current_timestamp,
    primary key (Id),
    foreign key (Type) references CardType (Code),
    foreign key (UserId) references User (Id),
    foreign key (Company) references CardCompany (Id),
    index IdxStatus (Status) -- Índice en Status para búsquedas por estado
);

create table ProductCategory
(
    Code        varchar(6),
    Description varchar(50) not null,
    primary key (Code)
);

create table Product
(
    Id           int auto_increment,
    Name         varchar(50)    not null,
    Description  varchar(150)   not null,
    ImageUrl     varchar(400),
    Price        decimal(10, 2) not null,
    CategoryCode varchar(6)     not null,
    Stock        int            not null,
    Status       bool     default true,
    CreatedAt    datetime default current_timestamp,
    UpdatedAt    datetime default current_timestamp on update current_timestamp,
    primary key (Id),
    foreign key (CategoryCode) references ProductCategory (Code)
);

create table PurchaseOrderStatus
(
    Code        char(2),
    Description varchar(50) not null,
    primary key (Code)
);

create table PurchaseOrder
(
    Id                int auto_increment,
    UserId            int not null,
    ShippingAddressId int not null,
    PaymentMethodId   int not null,
    Notes             varchar(200),
    TotalAmount       decimal(10, 2) default 0.00,
    OrderedAt         datetime       default current_timestamp,
    OrderStatus       char(2)        default 'PE',
    CreatedAt         datetime       default current_timestamp,
    UpdatedAt         datetime       default current_timestamp on update current_timestamp,
    primary key (Id),
    foreign key (userId) references User (Id),
    foreign key (ShippingAddressId) references UserAddress (Id),
    foreign key (PaymentMethodId) references UserPaymentCard (Id),
    foreign key (OrderStatus) references PurchaseOrderStatus (Code),
    index IdxUserId (UserId),          -- Índice en UserId para búsquedas por usuario
    index IdxOrderedAt (OrderedAt),    -- Índice en OrderedAt para búsquedas por fecha de orden
    index IdxOrderStatus (OrderStatus) -- Índice en OrderStatus para búsquedas por estado
);

create table ProductInOrder
(
    Id        int auto_increment,
    OrderId   int not null,
    ProductId int not null,
    Quantity  int not null,
    CreatedAt datetime default current_timestamp,
    UpdatedAt datetime default current_timestamp on update current_timestamp,
    primary key (Id),
    foreign key (OrderId) references PurchaseOrder (Id),
    foreign key (ProductId) references Product (Id),
    index IdxOrderId (OrderId) -- Índice en OrderId para búsquedas por orden
);
