const fs = require('fs');
const path = require('path');

// Función para leer archivos JSON desde una carpeta
async function leerArchivosDesdeCarpeta(carpeta) {
  const archivos = fs.readdirSync(carpeta); // Lee el contenido de la carpeta
  let resultado = [];

  for (const archivo of archivos) {
    if (path.extname(archivo) === '.json') { // Filtra solo los archivos .json
      const datos = JSON.parse(fs.readFileSync(path.join(carpeta, archivo), 'utf8'));
      resultado = resultado.concat(datos); // O combinar según tu necesidad
    }
  }

  return resultado;
}

async function getDatos(){
    const cNodos = './db/Alex/test/Nodes'; 
    const cLinks = './db/Alex/test/Relations'; 
    
    try {
        datos = {
            nodes:[],
            links:[]
    
        }
        // Esperar a que se lean los archivos
        datos.nodes = await leerArchivosDesdeCarpeta(cNodos);
        datos.links = await leerArchivosDesdeCarpeta(cLinks);
    
        // Imprimir los resultados una vez completada la lectura
        //console.log('Datos de nodes:', datos.nodes);
        //console.log('Datos de links:', datos.links);
        
        return datos; 
    
      } catch (err) {
        console.error('Error leyendo las carpetas:', err);
      }
}

export {getDatos};


