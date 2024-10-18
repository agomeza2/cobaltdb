import * as d3 from "https://cdn.jsdelivr.net/npm/d3@7/+esm";

async function obtenerDatos() {
  try {
      const response = await fetch('/api/datos');
      const datos = await response.json();
      return datos;
  } catch (error) {
      console.error('Error al obtener los datos:', error);
  }
}

obtenerDatos().then(graph => {
  const width = 928;
  const height = 600;

  const categories = Array.from(new Set(graph.links.map(d => d.category))); // Asignar categorías desde los links
  const color = d3.scaleOrdinal(categories, d3.schemeCategory10);

  const links = graph.links.map(d => ({...d}));
  console.log(links[1])
  const nodes = graph.nodes.map(d => ({...d}));

  // Simulación con fuerzas
  const simulation = d3.forceSimulation(nodes)
    .force("link", d3.forceLink(links).id(d => d.ID).distance(100)) // Usamos "name" como ID
    .force("charge", d3.forceManyBody().strength(-350))  // Fuerza negativa para dispersar nodos
    .force("x", d3.forceX())  // Aplicamos fuerza en el eje X
    .force("y", d3.forceY())  // Aplicamos fuerza en el eje Y
    .on("tick", ticked);

  const svg = d3.create("svg")
    .attr("viewBox", [-width / 2, -height / 2, width, height])
    .attr("width", width)
    .attr("height", height)
    .attr("style", "max-width: 100%; height: auto; font: 12px sans-serif;");

  // Definir flechas en los enlaces (arcos)
  svg.append("defs").selectAll("marker")
    .data(categories)
    .join("marker")
      .attr("id", d => `arrow-${d}`)
      .attr("viewBox", "0 -5 10 10")
      .attr("refX", 15)
      .attr("refY", -0.5)
      .attr("markerWidth", 6)
      .attr("markerHeight", 6)
      .attr("orient", "auto")
    .append("path")
      .attr("fill", color)
      .attr("d", "M0,-5L10,0L0,5");

  // Creación de los enlaces (con categorías y flechas)
  const link = svg.append("g")
      .attr("fill", "none")
      .attr("stroke-width", 1.5)
    .selectAll("path")
    .data(links)
    .join("path")
      .attr("stroke", d => color(d.category)) // Usamos la categoría para el color
      .attr("marker-end", d => `url(#arrow-${d.category})`); // Flechas por categoría

  // Creación de los nodos
  const node = svg.append("g")
      .attr("fill", "currentColor")
      .attr("stroke-linecap", "round")
      .attr("stroke-linejoin", "round")
    .selectAll("g")
    .data(nodes)
    .join("g")
      .on("click", (event, d) => tarjetaNodo(event, d)) // Pasar evento correctamente
      .call(d3.drag()
        .on("start", dragstarted)
        .on("drag", dragged)
        .on("end", dragended));

  // Agregar el círculo en cada nodo
  node.append("circle")
      .attr("stroke", "white")
      .attr("stroke-width", 1.5)
      .attr("r", 4); // Tamaño de los nodos

  // Agregar el texto en cada nodo
  node.append("text")
      .attr("x", 8)  // Posición del texto respecto al nodo
      .attr("y", "0.31em")
      .text(d => d.name) // Mostrar el nombre de cada nodo
    .clone(true).lower()
      .attr("fill", "none")
      .attr("stroke", "white")
      .attr("stroke-width", 3); // Contorno blanco para el texto

  // Función para dibujar enlaces curvados
  function linkArc(d) {
    const dx = d.target.x- d.source.x,
          dy = d.target.y - d.source.y,
          dr = Math.sqrt(dx * dx + dy * dy); // Radio para la curvatura
    return `M${d.source.x},${d.source.y}A${dr},${dr} 0 0,1 ${d.target.x},${d.target.y}`;
  }

  // Función para limitar los nodos dentro de los márgenes del SVG
  function limitPosition(node) {
    node.x = Math.max(-width / 2, Math.min(width / 2, node.x));  // Limita en el eje X
    node.y = Math.max(-height / 2, Math.min(height / 2, node.y));  // Limita en el eje Y
  }

  // Actualización de la posición de nodos y enlaces en cada tick de la simulación
  function ticked() {
    link.attr("d", linkArc); // Enlaces curvados
    
    node.each(limitPosition)  // Limitar la posición de los nodos
        .attr("transform", d => `translate(${d.x},${d.y})`); // Posición de los nodos
  }

  function dragstarted(event) {
    if (!event.active) simulation.alphaTarget(0.3).restart();
    event.subject.fx = event.subject.x;
    event.subject.fy = event.subject.y;
  }

  function dragged(event) {
    event.subject.fx = event.x;
    event.subject.fy = event.y;
  }

  function dragended(event) {
    if (!event.active) simulation.alphaTarget(0);
    event.subject.fx = null;
    event.subject.fy = null;
  }

  // Función para mostrar la tarjeta de información del nodo
  const card = d3.select("body").append("div")
    .attr("class", "card")
    .style("position", "absolute")
    .style("background", "#fff")
    .style("border", "1px solid #ccc")
    .style("padding", "10px")
    .style("display", "none");

  function tarjetaNodo(event, d) {
    card.style("left", `${event.pageX + 10}px`) // Usar event.pageX y event.pageY
        .style("top", `${event.pageY + 10}px`)
        .style("display", "block")
        .html(`<strong>Nombre:</strong> ${d.name}<br><strong>ID:</strong> ${d.ID}<br><strong>Propiedades:</strong> ${JSON.stringify(d.properties)}`);
    
    setTimeout(() => card.style("display", "none"), 3000); // Ocultar tarjeta después de 3 segundos
  }

  // Añadir el SVG al contenedor
  container.append(svg.node());

  // Crear la leyenda para los colores de los enlaces
  const legend = d3.select("body").append("div")
    .attr("class", "legend")
    .style("position", "absolute")
    .style("top", "20px")
    .style("right", "20px")
    .style("background", "#fff")
    .style("border", "1px solid #ccc")
    .style("padding", "10px");

  legend.append("h3").text("Leyenda de colores");
  categories.forEach(category => {
    legend.append("div")
      .style("display", "flex")
      .style("align-items", "center")
      .html(`<div style="width: 20px; height: 20px; background:${color(category)}; margin-right: 5px;"></div> ${category}`);
  });

}).catch(err => {
  console.error("Error cargando el archivo JSON:", err);
});
