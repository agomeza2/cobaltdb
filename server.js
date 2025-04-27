// server.js
const express = require('express');
const fs = require('fs');
const path = require('path');

const app = express();
const PORT = 5500;

// Middleware para permitir el acceso desde el frontend
app.use(express.json());
app.use(express.static('GUI/pantalla')); // Servir archivos estáticos como HTML, CSS, etc.

// Endpoint para enviar los datos combinados al frontend
app.get('/api/datos', async (req, res) => {
  try {
    const cNodos = 'db/Alice/test/Nodes';
    const cLinks = 'db/Alice/test/Relations';
    
    // Leer y combinar los archivos (basado en la lógica previa)
    const leerArchivosDesdeCarpeta = async (carpeta) => {
      const archivos = fs.readdirSync(carpeta);
      let resultado = [];
      for (const archivo of archivos) {
        if (path.extname(archivo) === '.json') {
          const datos = JSON.parse(fs.readFileSync(path.join(carpeta, archivo), 'utf8'));
          resultado = resultado.concat(datos);
        }
      }
      return resultado;
    };

    const nodes = await leerArchivosDesdeCarpeta(cNodos);
    const links = await leerArchivosDesdeCarpeta(cLinks);
    
    const datos = {
      nodes,
      links
    };
    
    res.json(datos); // Enviar los datos como JSON al frontend
  } catch (err) {
    res.status(500).json({ error: 'Error al leer los archivos' });
  }
});

// Iniciar el servidor
app.listen(PORT, () => {
  console.log(`Servidor corriendo en http://localhost:${PORT}`);
});
