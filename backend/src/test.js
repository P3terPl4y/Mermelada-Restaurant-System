// test.js
const url = 'http://localhost:2345/api/user/reservar';

// Caso 1: reservación válida
const reservaValida = {
    nombre: "Lucía Fernández",
    cantidad_personas: 3,
    personalizacion: "cumpleaños",
    telefono: "5551122334",
    fecha: "2026-04-28"
};

// Caso 2: personalización inválida (debería fallar)
const reservaInvalida = {
    nombre: "Pedro",
    cantidad_personas: 2,
    personalizacion: "boda",   // no está en el CHECK
    telefono: "5550001111",
    fecha: "2026-05-01"
};

const probar = async (data, descripcion) => {
    try {
        const res = await fetch(url, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(data)
        });
        const json = await res.json().catch(() => res.text());
        console.log(`${descripcion}: ${res.status}`, json);
    } catch (err) {
        console.error(`${descripcion}: error de red`, err.message);
    }
};

(async () => {
    await probar(reservaValida, 'Reserva válida');
    await probar(reservaInvalida, 'Reserva inválida (CHECK)');

    // Consulta por GET
    try {
        const res = await fetch('http://localhost:2345/api/admin/info');
        const data = await res.json();
        console.log('Listado de reservaciones:', data);
    } catch (err) {
        console.error('Error en GET:', err.message);
    }
})();