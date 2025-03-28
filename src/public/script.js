document.getElementById("addPersonForm").addEventListener("submit", async function (e) {
    e.preventDefault();
    
    const nombre = document.getElementById("nombre").value;
    const edad = document.getElementById("edad").value;
    const sexo = document.getElementById("sexo").value === "true"; // Convertir a booleano

    const response = await fetch("http://localhost:8080/addPerson", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ nombre, edad, sexo })
    });

    if (response.ok) {
        alert("Persona agregada exitosamente");
    }
});

setInterval(async () => {
    const response = await fetch("http://localhost:8080/newPersonIsAdded");
    const data = await response.json();
    if (data.newData) {
        alert("¡Se ha agregado una nueva persona!");
        updateGenderCount(); 
    }
}, 5000);


async function getGenderCount() {
    const response = await fetch("http://localhost:8080/countGender");
    const data = await response.json();
    document.getElementById("hombres").innerText = `Hombres: ${data.hombres}`;
    document.getElementById("mujeres").innerText = `Mujeres: ${data.mujeres}`;
    
    // Mantener la conexión abierta
    setTimeout(getGenderCount, 5000); 
}

getGenderCount();
