CobaltDB is a Graphical Database
cobaltdb works with our own custom SQL-like language:
Keywords
• SHOW
• COMMON
• ALTER
• REMOVE
• DATABASE
• DATABASES
• RELATION
• IN
• HOW MANY
•NODE
•SELECT
•CREATE
•NODES 
•RELATIONS
•IMPORT 
•EXPORT
•LIST
•ADD 

LIST DATABASES; //vemos las bases de datos 
CREATE DATABASE test; //creamos la base de datos

SELECT DATABASE test; //seleccionamos la base de datos 

//creamos directamente el nodo de una persona llamada Jo 
CREATE NODE {category:people,name:jo,properties:{salary:3000,age:20}}

//creamos multiples nodos llamados Lili, Martyn y Constance 
CREATE NODES [{category:people,name:Lili,properties:{salary:5000,age:40}},{category:people,name:Martyn,properties:{salary:7000,age:90}},{category:people,name:Constance,properties:{salary:9000,age:50}}]
//creamos una relacion 
CREATE RELATION {category:teach,name:tutoria,properties:{classroom:202,time:"9:00"},source:0,target:89}// con los ID de los nodos de origen y destino 

//creamos multiple relaciones entre Martyn, Lili y Jo 
CREATE RELATIONS [{category:teach,name:tutoria2,properties:{classroom:302,time:"8:00"},source:1,target:99},{category:teach,name:tutoria3,properties:{classroom:202,time:"10:00"},source:5,target:89},{category:teach,name:tutoria4,properties:{classroom:502,time:"7:00"},source:0,target:99}]
//vemos los atributos de Martyn
SHOW NODE ID:10 

SHOW NODE name:Martyn //muestra no solo un nodo sino todos los llamados Martyn, asi con mas atributos

SHOW RELATION tutoria //muestra la relacion tutoria 
//Cambiamos la edad de Martyn 
ALTER NODE ID:10 

ALTER NODE name:Martyn

ALTER RELATION ID:89
//quitamos un atributo a Martyn 
REMOVE age IN NODE name:Martyn
REMOVE classroom IN Relation ID:89 
//le volvemos a colocar el atributo age, pero aparecera como NIL
ADD age IN NODE name:Martyn
ADD classroom IN RELATION ID:89 

//vemos los nodos en comun entre Martyn y Lili
COMMON name:Martyn name:Lili

//buscamos cuantos nodos de la categoria people hay 
HOW MANY people

//buscamos cuantas relaciones de la categoria teach existen
HOW MANY Teach

REMOVE DATABASE test

LIST RELATIONS 
LIST NODES 
IMPORT ../salary_country.xlsx 
EXPORT DATABASE test