document.getElementById("addPersonForm").addEventListener("submit", async function (e) {
    e.preventDefault();
    
    const nombre = document.getElementById("nombre").value.trim();
    const edad = parseInt(document.getElementById("edad").value, 10); // Convertir a número
    const sexo = document.getElementById("sexo").value === "true"; // Convertir a booleano

    if (!nombre || isNaN(edad)) {
        alert("Por favor, ingresa datos válidos.");
        return;
    }

    const response = await fetch("http://localhost:8085/addPerson", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ nombre, edad, sexo })
    });

    if (response.ok) {
        alert("Persona agregada exitosamente");

        const latestIDResponse = await fetch("http://localhost:8085/newPersonIsAdded?lastID=0");
        const latestIDData = await latestIDResponse.json();
        
        if (latestIDData.newData) {
            localStorage.setItem("lastID", latestIDData.latestID); // Guardar nuevo lastID
        }
    } else {
        const errorText = await response.text();
        console.error("Error en POST /addPerson:", errorText);
    }
});


setInterval(async () => {
    let lastID = localStorage.getItem("lastID");
    lastID = lastID ? parseInt(lastID, 10) : 0; // Si no existe, usa 0

    const response = await fetch(`http://localhost:8085/newPersonIsAdded?lastID=${lastID}`);
    
    if (response.ok) {
        const data = await response.json();
        if (data.newData) {
            alert("¡Se ha agregado una nueva persona!");
            localStorage.setItem("lastID", data.latestID); // Actualizar lastID
            updateGenderCount();
        }
    } else {
        console.error("Error en GET /newPersonIsAdded:", await response.text());
    }
}, 5000);


async function getGenderCount() {
    try {
        const response = await fetch("http://localhost:8085/countGender");
        if (!response.ok) throw new Error("Respuesta no válida del servidor");

        const data = await response.json();
        document.getElementById("hombres").innerText = `Hombres: ${data.hombres || 0}`;
        document.getElementById("mujeres").innerText = `Mujeres: ${data.mujeres || 0}`;
    } catch (error) {
        console.error("Error en getGenderCount:", error);
    }

    setTimeout(getGenderCount, 5000);
}

getGenderCount();

