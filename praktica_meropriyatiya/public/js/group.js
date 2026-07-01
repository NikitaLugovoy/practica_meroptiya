document.addEventListener("DOMContentLoaded", () => {

    const modal = document.getElementById("groupModal");

    const addBtn = document.getElementById("addGroupBtn");
    const closeBtn = document.getElementById("closeGroupModal");
    const saveBtn = document.getElementById("saveGroup");

    const API = "http://localhost:8080";

    function authHeader() {
        const token = localStorage.getItem("token");
        return token ? { Authorization: `Bearer ${token}` } : {};
    }

    function openModal(group = null) {

        modal.classList.remove("hidden");

        if (group) {
            document.getElementById("groupModalTitle").innerText = "Редактировать группу";

            document.getElementById("groupId").value = group.id;
            document.getElementById("groupName").value = group.name;
        } else {
            document.getElementById("groupModalTitle").innerText = "Добавить группу";

            document.getElementById("groupId").value = "";
            document.getElementById("groupName").value = "";
        }
    }

    addBtn.addEventListener("click", () => openModal());

    closeBtn.addEventListener("click", () => {
        modal.classList.add("hidden");
    });

    saveBtn.addEventListener("click", async () => {

        const id = document.getElementById("groupId").value;
        const name = document.getElementById("groupName").value;

        if (!name) return alert("Введите название");

        if (id) {
            await axios.put(`${API}/update_group/${id}`, { name }, { headers: authHeader() });
        } else {
            await axios.post(`${API}/new_group`, { name }, { headers: authHeader() });
        }

        location.reload();
    });

    document.querySelectorAll(".edit-group-btn").forEach(btn => {
        btn.addEventListener("click", () => {

            const row = btn.closest("tr");

            openModal({
                id: row.dataset.id,
                name: row.querySelector(".g-name").innerText
            });
        });
    });

    document.querySelectorAll(".delete-group-btn").forEach(btn => {
        btn.addEventListener("click", async () => {

            const row = btn.closest("tr");
            const id = row.dataset.id;

            if (!confirm("Удалить группу?")) return;

            await axios.delete(`${API}/delete_group/${id}`, { headers: authHeader() });

            row.remove();
        });
    });

});