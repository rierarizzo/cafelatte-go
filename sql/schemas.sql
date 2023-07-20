DROP DATABASE IF EXISTS CafeLatte;

CREATE DATABASE CafeLatte;

USE CafeLatte;

CREATE TABLE UserRole
(
    Code        CHAR        NOT NULL,
    Description VARCHAR(20) NOT NULL,
    PRIMARY KEY (Code)
);

INSERT INTO UserRole
VALUES ('A', 'ADMIN');
INSERT INTO UserRole
VALUES ('E', 'EMPLOYEE');
INSERT INTO UserRole
VALUES ('C', 'CLIENT');

CREATE TABLE User
(
    ID          INT AUTO_INCREMENT,
    Username    varchar(30)  NOT NULL UNIQUE,
    Name        VARCHAR(50)  NOT NULL,
    Surname     VARCHAR(50)  NOT NULL,
    PhoneNumber VARCHAR(20),
    Email       VARCHAR(50)  NOT NULL UNIQUE,
    Password    VARCHAR(100) NOT NULL,
    RoleCode    CHAR         NOT NULL,
    Status      CHAR DEFAULT 'V',
    PRIMARY KEY (ID),
    FOREIGN KEY (RoleCode) REFERENCES UserRole (Code)
);

CREATE TABLE Province
(
    ID   INT         NOT NULL,
    Name VARCHAR(50) NOT NULL,
    PRIMARY KEY (ID)
);

INSERT INTO Province
    (ID, Name)
VALUES (1, 'GUAYAS');
INSERT INTO Province
    (ID, Name)
VALUES (2, 'SANTA ELENA');

CREATE TABLE City
(
    ID         INT         NOT NULL,
    ProvinceID INT         NOT NULL,
    Name       VARCHAR(50) NOT NULL,
    PRIMARY KEY (ID, ProvinceID),
    FOREIGN KEY (ProvinceID) REFERENCES Province (ID)
);

INSERT INTO City (ID, ProvinceID, Name)
VALUES (1, 1, 'GUAYAQUIL');
INSERT INTO City (ID, ProvinceID, Name)
VALUES (2, 1, 'DURAN');
INSERT INTO City (ID, ProvinceID, Name)
VALUES (3, 1, 'YAGUACHI');
INSERT INTO City (ID, ProvinceID, Name)
VALUES (4, 1, 'MILAGRO');
INSERT INTO City (ID, ProvinceID, Name)
VALUES (1, 2, 'SALINAS');
INSERT INTO City (ID, ProvinceID, Name)
VALUES (2, 2, 'SANTA ELENA');

CREATE TABLE AddressType
(
    Code        CHAR,
    Description VARCHAR(20),
    PRIMARY KEY (Code)
);

INSERT INTO AddressType
VALUES ('D', 'DOMICILIO');
INSERT INTO AddressType
VALUES ('T', 'TRABAJO');

CREATE TABLE UserAddress
(
    ID         INT AUTO_INCREMENT,
    Type       CHAR         NOT NULL,
    UserID     INT          NOT NULL,
    ProvinceID INT          NOT NULL,
    CityID     INT          NOT NULL,
    PostalCode VARCHAR(10),
    Detail     VARCHAR(150) NOT NULL,
    Enabled    BOOL DEFAULT TRUE,
    PRIMARY KEY (ID, UserID),
    FOREIGN KEY (Type) REFERENCES AddressType (Code),
    FOREIGN KEY (UserID) REFERENCES User (ID),
    FOREIGN KEY (ProvinceID) REFERENCES Province (ID),
    FOREIGN KEY (CityID) REFERENCES City (ID)
);

CREATE TABLE CardCompany
(
    ID   INT AUTO_INCREMENT,
    Name VARCHAR(50),
    PRIMARY KEY (ID)
);

INSERT INTO CardCompany (Name)
VALUES ('AMERICAN EXPRESS');
INSERT INTO CardCompany (Name)
VALUES ('VISA');
INSERT INTO CardCompany (Name)
VALUES ('MASTERCARD');
INSERT INTO CardCompany (Name)
VALUES ('DISCOVER');

CREATE TABLE CardIssuer
(
    ID   INT AUTO_INCREMENT,
    Name VARCHAR(50),
    PRIMARY KEY (ID)
);

INSERT INTO CardIssuer (Name)
VALUES ('BANCO DE GUAYAQUIL');
INSERT INTO CardIssuer (Name)
VALUES ('BANCO DE PICHINCHA');
INSERT INTO CardIssuer (Name)
VALUES ('BANCO DEL PACIFICO');
INSERT INTO CardIssuer (Name)
VALUES ('BANCO BOLIVARIANO');

CREATE TABLE CardType
(
    Code        CHAR NOT NULL,
    Description VARCHAR(50),
    PRIMARY KEY (Code)
);

INSERT INTO CardType (Code, Description)
VALUES ('C', 'CREDITO');
INSERT INTO CardType (Code, Description)
VALUES ('D', 'DEBITO');

CREATE TABLE UserPaymentCard
(
    ID             INT AUTO_INCREMENT,
    Type           CHAR         NOT NULL,
    UserID         INT          NOT NULL,
    Company        INT          NOT NULL,
    Issuer         INT          NOT NULL,
    HolderName     VARCHAR(100) NOT NULL,
    Number         VARCHAR(100) NOT NULL,
    ExpirationDate DATE         NOT NULL,
    CVV            VARCHAR(100) NOT NULL,
    Enabled        BOOL DEFAULT TRUE,
    PRIMARY KEY (ID),
    FOREIGN KEY (Type) REFERENCES CardType (Code),
    FOREIGN KEY (UserID) REFERENCES User (ID),
    FOREIGN KEY (Company) REFERENCES CardCompany (ID),
    FOREIGN KEY (Issuer) REFERENCES CardIssuer (ID)
);
