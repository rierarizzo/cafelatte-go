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
    foreign key (RoleCode) references UserRole (Code),
    index IdxUsername (Username), -- Índice en Username para búsquedas por nombre de usuario
    index IdxEmail (Email),       -- Índice en Email para búsquedas por correo electrónico
    index IdxStatus (Status)      -- Índice en Status para búsquedas por estado
);

create table Province
(
    ID   int         not null,
    Name varchar(50) not null,
    primary key (ID)
);

create table City
(
    ID         int         not null,
    ProvinceID int         not null,
    Name       varchar(50) not null,
    primary key (ID, ProvinceID),
    foreign key (ProvinceID) references Province (ID)
);

create table AddressType
(
    Code        char,
    Description varchar(20),
    primary key (Code)
);

create table UserAddress
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
    foreign key (Type) references AddressType (Code),
    foreign key (UserID) references User (ID),
    foreign key (ProvinceID) references Province (ID),
    foreign key (CityID) references City (ID),
    index IdxStatus (Status) -- Índice en Status para búsquedas por estado
);

create table CardCompany
(
    ID   int auto_increment,
    Name varchar(50),
    primary key (ID)
);

create table CardType
(
    Code        char not null,
    Description varchar(50),
    primary key (Code)
);

create table UserPaymentCard
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
    foreign key (Type) references CardType (Code),
    foreign key (UserID) references User (ID),
    foreign key (Company) references CardCompany (ID),
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
    ID           int auto_increment,
    Name         varchar(50)    not null,
    Description  varchar(150)   not null,
    ImageURL     varchar(400),
    Price        decimal(10, 2) not null,
    CategoryCode varchar(6)     not null,
    Stock        int            not null,
    Status       bool     default true,
    CreatedAt    datetime default current_timestamp,
    UpdatedAt    datetime default current_timestamp on update current_timestamp,
    primary key (ID),
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
    ID                int auto_increment,
    UserID            int not null,
    ShippingAddressID int not null,
    PaymentMethodID   int not null,
    Notes             varchar(200),
    TotalAmount       decimal(10, 2) default 0.00,
    OrderedAt         datetime       default current_timestamp,
    OrderStatus       char(2)        default 'PE',
    CreatedAt         datetime       default current_timestamp,
    UpdatedAt         datetime       default current_timestamp on update current_timestamp,
    primary key (ID),
    foreign key (userID) references User (ID),
    foreign key (ShippingAddressID) references UserAddress (ID),
    foreign key (PaymentMethodID) references UserPaymentCard (ID),
    foreign key (OrderStatus) references PurchaseOrderStatus (Code),
    index IdxUserID (UserID),          -- Índice en UserID para búsquedas por usuario
    index IdxOrderedAt (OrderedAt),    -- Índice en OrderedAt para búsquedas por fecha de orden
    index IdxOrderStatus (OrderStatus) -- Índice en OrderStatus para búsquedas por estado
);

create table PurchasedProduct
(
    ID        int auto_increment,
    OrderID   int not null,
    ProductID int not null,
    Quantity  int not null,
    CreatedAt datetime default current_timestamp,
    UpdatedAt datetime default current_timestamp on update current_timestamp,
    primary key (ID),
    foreign key (OrderID) references PurchaseOrder (ID),
    foreign key (ProductID) references Product (ID),
    index IdxOrderID (OrderID) -- Índice en OrderID para búsquedas por orden
);
