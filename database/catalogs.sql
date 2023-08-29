use cafelatte;

-- User roles
insert into UserRole
    (Code, Description)
values ('A', 'Admin'),
       ('E', 'Employee'),
       ('C', 'Client');

-- Provinces
insert into Province
    (Id, Name)
values (1, 'Guayas'),
       (2, 'Santa Elena');

-- Cities
insert into City
    (Id, ProvinceId, Name)
values (1, 1, 'Guayaquil'),
       (2, 1, 'Durán'),
       (3, 1, 'Yaguachi'),
       (4, 1, 'Milagro'),
       (1, 2, 'Salinas'),
       (2, 2, 'Santa Elena');

-- Address types
insert into AddressType
    (Code, Description)
values ('H', 'Home'),
       ('W', 'Work');

-- Card companies
insert into CardCompany
    (Id, Name)
values (1, 'American Express'),
       (2, 'Visa'),
       (3, 'Mastercard'),
       (4, 'Discover');

-- Card types
insert into CardType
    (Code, Description)
values ('C', 'Credit'),
       ('D', 'Debit');

-- Product categories
insert into ProductCategory
    (Code, Description)
values ('HOTBEV', 'Hot beverage'),
       ('COLBEV', 'Cold beverage'),
       ('DESSER', 'Dessert');

--
insert into PurchaseOrderStatus
    (Code, Description)
values ('PE', 'Pending'),
       ('PR', 'Processed'),
       ('SU', 'Submitted'),
       ('CO', 'Completed'),
       ('CA', 'Canceled');
