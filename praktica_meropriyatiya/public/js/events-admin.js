document.addEventListener("DOMContentLoaded", () => {

    const API = "http://localhost:8080";

    function authHeader() {
        const token = localStorage.getItem("token");
        return token ? { Authorization: `Bearer ${token}` } : {};
    }

    // ===== ПРОСМОТР =====
    const viewModal = document.getElementById("viewEventModal");
    const closeViewBtn = document.getElementById("closeViewEvent");

    function formatDateRu(isoString) {
        if (!isoString) return "—";
        const d = new Date(isoString);
        if (isNaN(d.getTime())) return isoString;
        return d.toLocaleString("ru-RU", {
            year: "numeric",
            month: "2-digit",
            day: "2-digit",
            hour: "2-digit",
            minute: "2-digit"
        });
    }


    closeViewBtn.addEventListener("click", () => {
        viewModal.classList.add("hidden");
    });

    // ===== РЕДАКТИРОВАНИЕ =====
    const editModal = document.getElementById("editEventModal");
    const closeEditBtn = document.getElementById("closeEditEvent");
    const editForm = document.getElementById("editEventForm");

    const editIdInput = document.getElementById("editEventId");
    const organizerSelect = document.getElementById("editOrganizerSelect");
    const responsibleSelect = document.getElementById("editResponsibleSelect");

    // приводим ISO-дату к формату для <input type="datetime-local">
    function toDatetimeLocalValue(isoString) {
        if (!isoString) return "";
        const d = new Date(isoString);
        if (isNaN(d.getTime())) return "";

        const pad = n => String(n).padStart(2, "0");
        const yyyy = d.getFullYear();
        const mm = pad(d.getMonth() + 1);
        const dd = pad(d.getDate());
        const hh = pad(d.getHours());
        const mi = pad(d.getMinutes());

        return `${yyyy}-${mm}-${dd}T${hh}:${mi}`;
    }

    async function loadUsersByRole(role, selectEl, selectedId) {
        const res = await axios.get(`${API}/users_by_role/${role}`, { headers: authHeader() });
        const users = res.data;

        const placeholder = selectEl.querySelector('option[value=""]');
        selectEl.innerHTML = "";
        if (placeholder) selectEl.appendChild(placeholder);

        users.forEach(u => {
            const option = document.createElement("option");
            option.value = u.id;
            option.textContent = `${u.name} ${u.surname} (id:${u.id})`;
            selectEl.appendChild(option);
        });

        if (selectedId) {
            selectEl.value = selectedId;
        }
    }

    
    closeEditBtn.addEventListener("click", () => {
        editModal.classList.add("hidden");
    });

    editForm.addEventListener("submit", async (e) => {
        e.preventDefault();

        const id = editIdInput.value;

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
            await axios.put(`${API}/update_event/${id}`, data, { headers: authHeader() });
            location.reload();
        } catch (err) {
            console.error(err);
            const serverMsg = err.response?.data || err.message;
            alert(`Ошибка сохранения: ${serverMsg}`);
        }
    });

    // ===== УДАЛЕНИЕ =====
    document.querySelectorAll(".delete-event-btn").forEach(btn => {
        btn.addEventListener("click", async () => {

            const row = btn.closest("tr");
            const id = row.dataset.id;
            const name = row.dataset.name;

            if (!confirm(`Удалить мероприятие «${name}»?`)) return;

            try {
                await axios.delete(`${API}/delete_event/${id}`, { headers: authHeader() });
                row.remove();
            } catch (err) {
                console.error(err);
                const serverMsg = err.response?.data || err.message;
                alert(`Ошибка удаления: ${serverMsg}`);
            }
        });
    });

});