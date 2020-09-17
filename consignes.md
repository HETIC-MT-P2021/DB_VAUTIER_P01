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
- pour chaque employé savoir avec quel magasin

## Vue magasins

- la liste des employés associés

```sql
SELECT e.employeeNumber
FROM employees e
INNER JOIN offices o ON e.officeCode = o.officeCode
WHERE o.officeCode = 1
```

