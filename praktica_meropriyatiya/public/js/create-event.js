const API = "http://localhost:8080";

function authHeader() {
    const token = localStorage.getItem("token");
    return token ? { Authorization: `Bearer ${token}` } : {};
}

async function loadUsersByRole(role, selectId) {
    const res = await axios.get(`${API}/users_by_role/${role}`, { headers: authHeader() });
    const users = res.data;

    const select = document.getElementById(selectId);
    select.innerHTML = "";

    users.forEach(u => {
        const option = document.createElement("option");
        option.value = u.id;
        option.textContent = `${u.name} ${u.surname} (id:${u.id})`;
        select.appendChild(option);
    });
}

document.addEventListener("DOMContentLoaded", async () => {
    await loadUsersByRole("organizer", "organizerSelect");
    await loadUsersByRole("responsible", "responsibleSelect");
});

document.getElementById("eventForm").addEventListener("submit", async (e) => {
    e.preventDefault();

    const data = {
        name: e.target.name.value,
        description: e.target.description.value,
        date_time: new Date(e.target.date_time.value).toISOString(),
        location: e.target.location.value,
        category_events: e.target.category_events.value,
        status: e.target.status.value,
        organizer_id: Number(e.target.organizer_id.value),
        responsible_id: Number(e.target.responsible_id.value),
    };

    try {
        await axios.post(`${API}/new_event`, data, { headers: authHeader() });
        alert("Event создан");
    } catch (err) {
        console.error(err);
        alert("Ошибка создания события");
    }
});