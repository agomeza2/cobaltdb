# cobaltdb
opensource serverless graphic database manager.  

cobaltdb works with our own custom SQL-like language:
Keywords
• SEARCH
• WHERE
• CATEGORY
• SHOW
• GRAPH
• FUNC
• CREATE
• DBINT
• DBFLOAT
• DBSTRING
• DBBOOL
• DBDATE
• COMMON
• WHICH
• FOR
• WHILE
• LOOP
• IF
• ELSE
• ELIF
• ALTER
• REMOVE
• DATABASE
• DATABASES
• USE
• RELATION
• NIL
• AND
• NOT
• EQ
• OR
• TRUE
• FALSE
• GROUP BY
• HOW MANY
• RETURN
• ADD

CREATE DATABASE test; //creamos la base de datos

USE test; //seleccionamos la base de datos 

//creamos directamente el nodo de una persona llamada Jo 
(people:Jo)('Jo',35,2000.4,23/04/99);

//creamos multiples nodos llamados Lili, Martyn y Constance 
(people:[Lili,Martyn,Constance])(['Lili',18,1000.3,04/05/08;
'Martyn',45,4500,05/06/85;'Constance',49,3000.5,05/05/70]);
//creamos una relacion 
CREATE RELATION Teach(String name, int classroom, DATE date); 

//asignamos una relacion entre Jo y constance, se lee como Jo ensenna a constance
(people:Jo)=>(Teach):('tutoria',101,02/02/24)=>(people:Constance);

//creamos multiple relaciones entre Martyn, Lili y Jo
(people:[Martyn,Lili,Jo]) => (TEACH):([tutoria,101,02/03/22;clase,
204,01/03/22;privada,303,01/03/22]) => (people:[Lili,Jo,Martyn]); 

//vemos los atributos de Martyn
SHOW (people:Martyn);

//Cambiamos la edad de Martyn 
ALTER (people:Martyn):(int age: 55);

//quitamos un atributo a Martyn 
REMOVE (people:Martyn):(DATE birth);

//le volvemos a colocar el atributo birth, pero aparecera como NIL
ADD (people:Martyn)(Date birth)

//vemos los nodos en comun entre Martyn y Lili
COMMON (people:Martyn) (people:Lili)

//buscamos nodos por los parametros de edad 
SEARCH INT age WHERE age = 18

//buscamos relaciones con los parametros de salon
SEARCH INT classroom WHERE classroom = 101

//buscamos nodos y los organizamos de forma ascendente 
SEARCH String name ASC;

//buscamos nodos y los organizamos de forma descendente 
SEARCH String name DESC

//buscamos nodos de la categoria people y los agrupamos por edad, default es ascendente 
SEARCH people GROUP BY age ;

//buscamos cuantos nodos de la categoria people hay 
HOW MANY people;

//buscamos cuantas relaciones de la categoria teach existen
HOW MANY Teach;
