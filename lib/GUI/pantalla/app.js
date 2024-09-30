import * as d3 from "https://cdn.jsdelivr.net/npm/d3@7/+esm";


async function obtenerDatos() {
  try {
      const response = await fetch('/api/datos');
      const datos = await response.json();
      return datos
  } catch (error) {
      console.error('Error al obtener los datos:', error);
  }
}

 
obtenerDatos().then(graph => {
    // Specify the dimensions of the chart.
const width = 1000;
const height = 1000;

// Specify the color scale.
const color = d3.scaleOrdinal(d3.schemeCategory10);

// The force simulation mutates links and nodes, so create a copy
// so that re-evaluating this cell produces the same result.
const links = graph.links.map(d => ({...d}));
const nodes = graph.nodes.map(d => ({...d}));

// Create a simulation with several forces.
const simulation = d3.forceSimulation(nodes)
    .force("link", d3.forceLink(links).id(d => d.name).distance(100))
    .force("charge", d3.forceManyBody())
    .force("center", d3.forceCenter(width / 2, height / 2))
    .on("tick", ticked);

// Create the SVG container.
const svg = d3.create("svg")
    .attr("width", width)
    .attr("height", height)
    .attr("viewBox", [0, 0, width, height])
    .attr("style", "max-width: 100%; max-height: 100%;");

// Add a line for each link, and a circle for each node.
const link = svg.append("g")
    .attr("stroke", "#999")
    .attr("stroke-opacity", 0.6)
  .selectAll()
  .data(links)
  .join("line")
    .attr("stroke-width", 1.5);

const node = svg.append("g")
  .selectAll()
  .data(nodes)
  .join("g")  // Agrupa los nodos y los textos
        .call(d3.drag()
            .on("start", dragstarted)
            .on("drag", dragged)
            .on("end", dragended));

  node.append("circle")
    .attr("r", 30)
    .attr("fill", d => color(d.category));

node.append("text")
     .attr("x", 0)
    .attr("y", 4)  // Ajusta para centrar verticalmente el texto
    .attr("text-anchor", "middle")
    .attr("font-size", "15px")
    .attr("font-weight", "200") 
    .attr("fill", "#fff")
    .text(d => d.name);

node.append("title")
   .attr("x", 0)
   .attr("y", 4)  // Ajusta para centrar verticalmente el texto
   .attr("text-anchor", "middle")
   .attr("font-size", "15px")
   .attr("font-weight", "200") 
   .attr("fill", "#fff")
   .text(d =>  d.category );



// Set the position attributes of links and nodes each time the simulation ticks.
function ticked() {
  link
      .attr("x1", d => d.source.x  = Math.max(10, Math.min(width - 10, d.source.x)))
      .attr("y1", d => d.source.y =  Math.max(10, Math.min(height - 10, d.source.y)))
      .attr("x2", d => d.target.x  = Math.max(10, Math.min(width - 10, d.target.x)))
      .attr("y2", d => d.target.y = Math.max(10, Math.min(height - 10, d.target.y)));
  

      node.attr("transform", d => `translate(${d.x}, ${d.y})`);
}

// Reheat the simulation when drag starts, and fix the subject position.
function dragstarted(event) {
  if (!event.active) simulation.alphaTarget(0.3).restart();
  event.subject.fx = event.subject.x;
  event.subject.fy = event.subject.y;
  tarjetaNodo(event.subject)
}

// Update the subject (dragged node) position during drag.
function dragged(event) {
  event.subject.fx = event.x;
  event.subject.fy = event.y;
}

// Restore the target alpha so the simulation cools after dragging ends.
// Unfix the subject position now that it’s no longer being dragged.
function dragended(event) {
  if (!event.active) simulation.alphaTarget(0);
  event.subject.fx = null;
  event.subject.fy = null;
}

// When this cell is re-run, stop the previous simulation. (This doesn’t
// really matter since the target alpha is zero and the simulation will
// stop naturally, but it’s a good practice.)

function tarjetaNodo(objNodo){
  const x = objNodo.x;
  const y = objNodo.y;
  // Coloca la tarjeta en las coordenadas donde se hizo clic
  card.style.left = `${x}px`;
  card.style.top = `${y}px`;

  card.textContent = "nombre:"+ objNodo.name +"Id:"+ objNodo.ID +"propiedades"+ objNodo.properties

  // Muestra la tarjeta
  card.classList.remove('hidden');

  // Opcional: oculta la tarjeta después de 2 segundos
  setTimeout(() => {
      card.classList.add('hidden');
  }, 2000);
} 


container.append(svg.node());
}).catch(err => {
    console.error("Error cargando el archivo JSON:", err);
});

