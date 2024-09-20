import * as d3 from "https://cdn.jsdelivr.net/npm/d3@7/+esm";
/* global d3 */

// se capturan los datos del servidor 
async function obtenerDatos() {
    try {
        const response = await fetch('/api/datos');
        const datos = await response.json();
        return datos
    } catch (error) {
        console.error('Error al obtener los datos:', error);
    }
}

let canvas = d3.select("#network"),
    width = canvas.attr("width"),
    height = canvas.attr("height"),
    ctx = canvas.node().getContext("2d"),
    radio = 20,
    simulation = d3.forceSimulation()
        .force("x", d3.forceX(width / 2))
        .force("y", d3.forceY(height / 2))
        .force("collide", d3.forceCollide(radio + 1))
        .force("charge", d3.forceManyBody().strength(-400))
        .force("link", d3.forceLink().id(d => d.name));

let datosGrafo;

obtenerDatos().then((datos) => {
    console.log(datos)
    datosGrafo = datos
});


obtenerDatos().then(graph => {
    simulation.nodes(graph.nodes);
    simulation.force("link").links(graph.links);
    simulation.on("tick", update);

    canvas.call(d3.drag()
        .container(canvas.node())
        .subject(dragsubject)
        .on("start", dragstarted)
        .on("drag", dragged)
        .on("end", dragended));

    function update() {
        ctx.clearRect(0, 0, width, height);

        ctx.beginPath();
        graph.links.forEach(printLink);
        ctx.stroke();

        ctx.beginPath();
        graph.nodes.forEach(printNode);
        ctx.fill();
    }

    function dragsubject(event) {
        return simulation.find(event.x, event.y);
    }

    function printNode(persona) {
        ctx.moveTo(persona.x + radio, persona.y);
        ctx.arc(persona.x, persona.y, radio, 0, Math.PI * 2);
    }

    function printLink(l) {
        ctx.moveTo(l.source.x, l.source.y);
        ctx.lineTo(l.target.x, l.target.y);
    }

}).catch(err => {
    console.error("Error cargando el archivo JSON:", err);
});

function dragstarted(event) {
    if (!event.active) simulation.alphaTarget(0.3).restart();
    event.subject.fx = event.subject.x;
    event.subject.fy = event.subject.y;
    console.log(event.subject);
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
