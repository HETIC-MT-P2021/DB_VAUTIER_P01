# SGBD

DB_NOM_P01
Faire une api json

## Vue client

- Toutes les données
```sql
SELECT * FROM customers c WHERE c.customerNumber = 103;
```

- liste des commandes du client
- pouvoir sortir le nombre d'items commandés avec le prix total pour un client

```sql
SELECT o.orderNumber, SUM(od.quantityOrdered) AS "Items number", SUM(od.priceEach * od.quantityOrdered) AS "Total price"
FROM orders o 
INNER JOIN orderdetails od ON od.orderNumber = o.orderNumber
WHERE o.customerNumber = 103
GROUP BY o.orderNumber;

```

## Vue commmande

commande avec tous les items associés
détail de chaque item

Infos de la commande par n°
```sql
SELECT o.orderNumber, o.orderDate , c.contactFirstName, c.contactLastName, SUM(od.quantityOrdered) AS "Items number", SUM(od.priceEach * od.quantityOrdered) AS "Total price"
FROM orders o
INNER JOIN customers c ON c.customerNumber = o.customerNumber
INNER JOIN orderdetails od ON od.orderNumber = o.orderNumber
WHERE o.orderNumber = 10123
GROUP BY o.orderNumber
```

Tous les items avec le détail par n° commande
```sql
SELECT od.productCode,p.productName, p.productDescription, od.quantityOrdered, od.priceEach
FROM orders o 
INNER JOIN orderdetails od ON od.orderNumber = o.orderNumber
INNER JOIN products p ON od.productCode = p.productCode
WHERE o.orderNumber = 10124
ORDER BY od.orderLineNumber
```

## Vue employés

- liste des employés avec leurs détails
```sql
SELECT e.employeeNumber, e.firstName, e.lastName, e.jobTitle, e.email, e.officeCode, o.city, o.addressLine1 FROM employees e 
INNER JOIN offices o ON o.officeCode = e.officeCode
ORDER BY e.officeCode, e.lastName
```
- pour chaque employé savoir avec quel magasin

## Vue magasins

```sql
SELECT o.officeCode, o.addressLine1, o.addressLine2, o.city, o.postalCode, o.`state`, o.phone, o.territory
FROM offices o
WHERE o.officeCode = 1 
```
- la liste des employés associés

```sql
SELECT e.employeeNumber, e.firstName, e.lastName
FROM employees e
INNER JOIN offices o ON e.officeCode = o.officeCode
WHERE o.officeCode = 1
```

