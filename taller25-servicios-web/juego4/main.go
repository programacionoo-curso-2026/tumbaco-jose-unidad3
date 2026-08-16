package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sort"
	"sync"
)

// ---------- Preguntas ----------

type Pregunta struct {
	Pregunta string `json:"pregunta"`
	OpcionA  string `json:"opcion_a"`
	OpcionB  string `json:"opcion_b"`
	OpcionC  string `json:"opcion_c"`
	Solucion string `json:"solucion"`
}

// Banco de preguntas
var preguntas = []Pregunta{
	{
		Pregunta: "for i &blank 0; i &lt; 10; i++",
		OpcionA:  "=",
		OpcionB:  "<",
		OpcionC:  ":=",
		Solucion: "C",
	},
	{
		Pregunta: `var nombre &blank "Sofia"`,
		OpcionA:  "=",
		OpcionB:  "<",
		OpcionC:  ":=",
		Solucion: "A",
	},
	{
		Pregunta: "fmt.&blank(\"Hola\")",
		OpcionA:  "Println",
		OpcionB:  "Scanln",
		OpcionC:  "Printf",
		Solucion: "A",
	},
	{
		Pregunta: "if edad &blank 18 {",
		OpcionA:  "=",
		OpcionB:  "==",
		OpcionC:  ":=",
		Solucion: "B",
	},

	// ---------- Preguntas agregadas por el estudiante ----------

	{
		Pregunta: "fmt.&blank(\"Hola\")",
		OpcionA:  "Println",
		OpcionB:  "Scanln",
		OpcionC:  "Printf",
		Solucion: "A",
	},
	{
		Pregunta: "var edad &blank 20",
		OpcionA:  ":=",
		OpcionB:  "=",
		OpcionC:  "==",
		Solucion: "B",
	},
}

// ---------- Resultados (ranking en memoria) ----------

type Resultado struct {
	Nombre      string `json:"nombre"`
	TiempoMs    int64  `json:"tiempo_ms"`
	TiempoTexto string `json:"tiempo_texto"`
}

var (
	mu         sync.Mutex
	resultados []Resultado
)

func formatearTiempo(ms int64) string {
	totalSeg := ms / 1000
	min := totalSeg / 60
	seg := totalSeg % 60

	return fmt.Sprintf("%02d:%02d", min, seg)
}

func guardarResultadoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var entrada struct {
		Nombre   string `json:"nombre"`
		TiempoMs int64  `json:"tiempo_ms"`
	}

	if err := json.NewDecoder(r.Body).Decode(&entrada); err != nil {
		http.Error(w, "json inválido", http.StatusBadRequest)
		return
	}

	nuevo := Resultado{
		Nombre:      entrada.Nombre,
		TiempoMs:    entrada.TiempoMs,
		TiempoTexto: formatearTiempo(entrada.TiempoMs),
	}

	mu.Lock()
	resultados = append(resultados, nuevo)

	sort.Slice(resultados, func(i, j int) bool {
		return resultados[i].TiempoMs < resultados[j].TiempoMs
	})

	mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(nuevo)
}

func rankingHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resultados)
}

const pageTemplate = `<!DOCTYPE html>
<html lang="es">
<head>
<meta charset="UTF-8">
<title>Juego de Go</title>

<style>
body {
	font-family: -apple-system, Arial, sans-serif;
	background: #f2f2f2;
	display: flex;
	justify-content: center;
	align-items: center;
	min-height: 100vh;
	margin: 0;
}

.card {
	background: #fff;
	padding: 32px 40px;
	border-radius: 12px;
	box-shadow: 0 4px 12px rgba(0,0,0,0.15);
	width: 520px;
	box-sizing: border-box;
}

#pantalla-inicio {
	display: flex;
	flex-direction: column;
	align-items: center;
	gap: 18px;
}

#pantalla-inicio h2 {
	margin: 0;
	color: #333;
}

#inputNombre {
	width: 100%;
	box-sizing: border-box;
	font-size: 20px;
	padding: 12px;
	border-radius: 8px;
	border: 1px solid #ccc;
	text-align: center;
}

#btnIniciar {
	width: 100%;
	font-size: 22px;
	padding: 14px;
	border: none;
	border-radius: 8px;
	background: #007acc;
	color: #fff;
	cursor: pointer;
}

#btnIniciar:hover {
	background: #005f99;
}

#pantalla-pregunta {
	display: none;
	flex-direction: column;
	align-items: center;
}

#progreso {
	font-size: 18px;
	color: #555;
	margin-bottom: 12px;
}

#temporizador {
	font-size: 22px;
	font-weight: bold;
	margin-bottom: 15px;
	color: #333;
}

.code {
	width: 100%;
	box-sizing: border-box;
	min-height: 160px;
	padding: 20px;
	font-family: 'Courier New', monospace;
	font-size: 24px;
	background: #1e1e1e;
	color: #d4d4d4;
	border-radius: 8px;
	margin-bottom: 20px;
	display: flex;
	justify-content: center;
	align-items: center;
	text-align: center;
}

.blank {
	font-weight: bold;
	color: #ffcc00;
}

.blank.correcto {
	color: #4caf50;
}

.blank.incorrecto {
	color: #e74c3c;
}

.variable {
	width: 100%;
	font-size: 22px;
	color: #555;
	margin-bottom: 20px;
	min-height: 40px;
	display: flex;
	justify-content: center;
	align-items: center;
	text-align: center;
}

.buttons {
	display: flex;
	justify-content: center;
	gap: 14px;
	flex-wrap: wrap;
}

.buttons button {
	width: 140px;
	height: 100px;
	display: flex;
	justify-content: center;
	align-items: center;
	font-family: 'Courier New', monospace;
	font-size: 30px;
	padding: 0;
	border: none;
	border-radius: 8px;
	background: #007acc;
	color: #fff;
	cursor: pointer;
}

.buttons button:hover {
	background: #005f99;
}

#pantalla-final {
	display: none;
	flex-direction: column;
	align-items: center;
	text-align: center;
	gap: 14px;
}

#pantalla-final h2 {
	margin: 0;
	color: #333;
}

#resultadoTiempo {
	font-size: 22px;
	font-weight: bold;
}

#ranking {
	width: 100%;
	text-align: left;
	background: #f7f7f7;
	border-radius: 8px;
	padding: 16px 20px;
	box-sizing: border-box;
}

#ranking h3 {
	margin-top: 0;
}

#listaRanking {
	padding-left: 25px;
}

#listaRanking li {
	margin-bottom: 8px;
}
</style>
</head>

<body>

<div class="card">

	<!-- Pantalla inicial -->
	<div id="pantalla-inicio">
		<h2>¿Cuál es tu nombre?</h2>

		<input
			id="inputNombre"
			type="text"
			placeholder="Escribe tu nombre"
		>

		<button id="btnIniciar" onclick="iniciar()">
			Iniciar
		</button>
	</div>

	<!-- Pantalla de preguntas -->
	<div id="pantalla-pregunta">

		<div id="progreso"></div>

		<div id="temporizador">
			Tiempo: 00:00
		</div>

		<div class="code" id="code"></div>

		<div class="variable" id="variable">
			Elige una opción
		</div>

		<div class="buttons">
			<button id="btnA"></button>
			<button id="btnB"></button>
			<button id="btnC"></button>
		</div>

	</div>

	<!-- Pantalla final -->
	<div id="pantalla-final">

		<h2 id="tituloFinal"></h2>

		<div id="resultadoTiempo"></div>

		<div id="ranking">
			<h3>Ranking</h3>
			<ol id="listaRanking"></ol>
		</div>

	</div>

</div>

<script>
	const preguntas = {{.Preguntas}};

	let nombre = '';
	let indice = 0;
	let tiempoInicio = 0;
	let intervalo = null;

	function iniciar() {

		const valor =
			document.getElementById('inputNombre').value.trim();

		if (!valor) {
			document.getElementById('inputNombre').focus();
			return;
		}

		nombre = valor;

		tiempoInicio = Date.now();

		intervalo = setInterval(actualizarTiempo, 1000);

		document.getElementById('pantalla-inicio').style.display = 'none';
		document.getElementById('pantalla-pregunta').style.display = 'flex';

		cargarPregunta(indice);
	}

	function actualizarTiempo() {

		const tiempoMs = Date.now() - tiempoInicio;

		const totalSeg = Math.floor(tiempoMs / 1000);

		const min = Math.floor(totalSeg / 60);

		const seg = totalSeg % 60;

		document.getElementById('temporizador').textContent =
			'Tiempo: ' +
			String(min).padStart(2, '0') +
			':' +
			String(seg).padStart(2, '0');
	}

	function cargarPregunta(i) {

		const p = preguntas[i];

		document.getElementById('progreso').textContent =
			'Pregunta ' + (i + 1) +
			' de ' + preguntas.length;

		const html = p.pregunta.replace(
			'&blank',
			'<span id="blank" class="blank">___</span>'
		);

		document.getElementById('code').innerHTML = html;

		document.getElementById('variable').textContent =
			'Elige una opción';

		const btnA = document.getElementById('btnA');
		const btnB = document.getElementById('btnB');
		const btnC = document.getElementById('btnC');

		btnA.textContent = p.opcion_a;
		btnB.textContent = p.opcion_b;
		btnC.textContent = p.opcion_c;

		btnA.onclick = () => elegir(p.opcion_a, 'A');
		btnB.onclick = () => elegir(p.opcion_b, 'B');
		btnC.onclick = () => elegir(p.opcion_c, 'C');
	}

	function elegir(valor, letra) {

		const p = preguntas[indice];

		const blank = document.getElementById('blank');

		blank.textContent = valor;

		if (letra === p.solucion) {

			blank.classList.remove('incorrecto');

			blank.classList.add('correcto');

			document.getElementById('variable').textContent =
				'¡Correcto!';

			setTimeout(() => {

				indice++;

				if (indice < preguntas.length) {

					cargarPregunta(indice);

				} else {

					finalizarJuego();

				}

			}, 700);

		} else {

			blank.classList.remove('correcto');

			blank.classList.add('incorrecto');

			document.getElementById('variable').textContent =
				'Incorrecto, intenta de nuevo';
		}
	}

	function finalizarJuego() {

		clearInterval(intervalo);

		const tiempoMs = Date.now() - tiempoInicio;

		document.getElementById('pantalla-pregunta').style.display = 'none';

		document.getElementById('pantalla-final').style.display = 'flex';

		document.getElementById('tituloFinal').textContent =
			'¡Felicidades, ' + nombre + '!';

		const totalSeg = Math.floor(tiempoMs / 1000);

		const min = Math.floor(totalSeg / 60);

		const seg = totalSeg % 60;

		const tiempoTexto =
			String(min).padStart(2, '0') +
			':' +
			String(seg).padStart(2, '0');

		document.getElementById('resultadoTiempo').textContent =
			'Tiempo: ' + tiempoTexto;

		guardarResultado(nombre, tiempoMs);
	}

	async function guardarResultado(nombre, tiempoMs) {

		try {

			await fetch('/resultado', {

				method: 'POST',

				headers: {
					'Content-Type': 'application/json'
				},

				body: JSON.stringify({
					nombre: nombre,
					tiempo_ms: tiempoMs
				})
			});

			cargarRanking();

		} catch (error) {

			console.error('Error guardando resultado:', error);
		}
	}

	async function cargarRanking() {

		try {

			const respuesta = await fetch('/ranking');

			const ranking = await respuesta.json();

			const lista =
				document.getElementById('listaRanking');

			lista.innerHTML = '';

			ranking.forEach((resultado, index) => {

				const li = document.createElement('li');

				li.textContent =
					(index + 1) +
					'. ' +
					resultado.nombre +
					' - ' +
					resultado.tiempo_texto;

				lista.appendChild(li);
			});

		} catch (error) {

			console.error('Error cargando ranking:', error);
		}
	}
</script>

</body>
</html>`

var tmpl = template.Must(
	template.New("page").Parse(pageTemplate),
)

func handler(w http.ResponseWriter, r *http.Request) {

	data, err := json.Marshal(preguntas)

	if err != nil {
		http.Error(
			w,
			"error generando preguntas",
			http.StatusInternalServerError,
		)

		return
	}

	err = tmpl.Execute(
		w,
		map[string]template.JS{
			"Preguntas": template.JS(data),
		},
	)

	if err != nil {
		log.Println("Error ejecutando plantilla:", err)
	}
}

func main() {

	http.HandleFunc("/", handler)

	http.HandleFunc(
		"/resultado",
		guardarResultadoHandler,
	)

	http.HandleFunc(
		"/ranking",
		rankingHandler,
	)

	log.Println(
		"Servidor corriendo en http://localhost:8080",
	)

	log.Fatal(
		http.ListenAndServe(":8080", nil),
	)
}
