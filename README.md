CobaltDB is a Graphical Database

cobaltdb works with our own custom gnu/linux-like language:

Keywords
• ls
• cd 
• rm
• common
• cat
• touch
• grep
• modify
• import
• export
• db 
• node 
• relation
• user
• admin 
•nodes 
•relations

touch user Jhon admin //crear usuario 
touch user Oswald  
--despues de esto el gestor te pedira logeo --- 

ls db; //vemos las bases de datos 
touch db test; //creamos la base de datos
cd test 
touch node {category:people,name:jo,properties:{salary:3000,age:20}} //creamos directamente el nodo de una persona llamada Jo 

//creamos multiples nodos llamados Lili, Martyn y Constance 
touch nodes [{category:people,name:Lili,properties:{salary:5000,age:40}},{category:people,name:Martyn,properties:{salary:7000,age:90}},{category:people,name:Constance,properties:{salary:9000,age:50}}]
//creamos una relacion 
touch relation {category:teach,name:tutoria,properties:{classroom:202,time:"9:00"},source:0,target:89}// con los ID de los nodos de origen y destino 

//creamos multiple relaciones entre Martyn, Lili y Jo 
touch relations [{category:teach,name:tutoria2,properties:{classroom:302,time:"8:00"},source:1,target:99},{category:teach,name:tutoria3,properties:{classroom:202,time:"10:00"},source:5,target:89},{category:teach,name:tutoria4,properties:{classroom:502,time:"7:00"},source:0,target:99}]
//vemos los atributos de Martyn
cat node ID:10 

cat node name:Martyn //muestra no solo un nodo sino todos los llamados Martyn, asi con mas atributos

cat relation tutoria //muestra las relaciones con nombre tutoria 
//Cambiamos la edad de Martyn 
modify node ID:10 name:Oswald

modify node name:Martyn age:26 

//quitamos un atributo a Martyn 
rm node name:Martyn age 
rm relation classroom:205 ID:89 
//le volvemos a colocar el atributo age, pero aparecera como NIL
add node name:Martyn age:23 
add relation ID:89 classroom:35 

//vemos los nodos en comun entre Martyn y Lili
common ID:10 ID:20 
rm db test //Eliminamos la base de datos 
grep age:26 //busca entre todos los nodos y las relaciones, los nodos o relaciones con la propiedad 
ls nodes //lista de todos los nodos 
ls relations //lista de todas las relaciones 
LIST NODES 
IMPORT ../salary_country.xlsx 
EXPORT DATABASE test
