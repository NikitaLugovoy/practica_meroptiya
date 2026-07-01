document.addEventListener("DOMContentLoaded", () => {

    const modal = document.getElementById("ugModal");

    const addBtn = document.getElementById("addUGBtn");
    const closeBtn = document.getElementById("closeUG");
    const saveBtn = document.getElementById("saveUG");

    const titleEl = document.getElementById("ugTitle");
    const idInput = document.getElementById("ugId");

    const userSelect = document.getElementById("ugUser");
    const userGroupWrap = document.getElementById("ugUserGroup");
    const groupSelect = document.getElementById("ugGroup");

    const studentsBox = document.getElementById("ugStudentsBox");
    const studentsSelect = document.getElementById("ugStudentsSelect");
    const noStudentsMsg = document.getElementById("ugNoStudents");

    const API = "http://localhost:8080";

    let mode = "create"; // "create" | "edit"

    function authHeader() {
        const token = localStorage.getItem("token");
        return token ? { Authorization: `Bearer ${token}` } : {};
    }

    // id студентов, у которых уже есть связь с группой (из таблицы userGroups, отрендеренной сервером)
    function getStudentsWithGroup() {
        const ids = new Set();
        document.querySelectorAll("#userGroupsTable tr[data-user-id]").forEach(row => {
            ids.add(row.dataset.userId);
        });
        return ids;
    }

    // показать option доступных студентов (роль student + нет группы), скрыть остальные
    function renderAvailableStudents() {
        const busyIds = getStudentsWithGroup();
        let visibleCount = 0;

        Array.from(studentsSelect.options).forEach(opt => {
            const isStudent = opt.dataset.role === "student";
            const isFree = !busyIds.has(opt.value);

            const show = isStudent && isFree;
            opt.hidden = !show;
            opt.disabled = !show;
            opt.selected = false;

            if (show) visibleCount++;
        });

        noStudentsMsg.classList.toggle("hidden", visibleCount > 0);
    }

    function openModal(ug = null) {

        modal.classList.remove("hidden");

        if (ug) {
            mode = "edit";

            titleEl.innerText = "Редактировать связь";

            idInput.value = ug.id;

            userGroupWrap.classList.remove("hidden");
            userSelect.value = ug.userId;

            groupSelect.value = ug.groupId;

            studentsBox.classList.add("hidden");

        } else {
            mode = "create";

            titleEl.innerText = "Добавить связь";

            idInput.value = "";
            userSelect.value = "";
            userGroupWrap.classList.add("hidden");

            groupSelect.value = "";

            studentsBox.classList.add("hidden");
        }
    }

    addBtn.addEventListener("click", () => openModal());

    closeBtn.addEventListener("click", () => {
        modal.classList.add("hidden");
    });

    // при выборе группы (только в режиме создания) показываем студентов
    groupSelect.addEventListener("change", () => {
        if (mode !== "create") return;

        if (groupSelect.value) {
            renderAvailableStudents();
            studentsBox.classList.remove("hidden");
        } else {
            studentsBox.classList.add("hidden");
        }
    });

    saveBtn.addEventListener("click", async () => {

        const groupId = groupSelect.value;

        if (!groupId) return alert("Выберите группу");

        if (mode === "edit") {

            const id = idInput.value;
            const userId = userSelect.value;

            if (!userId) return alert("Выберите пользователя");

            try {
                await axios.put(`${API}/update_user_group/${id}`, {
                    user_id: Number(userId),
                    group_id: Number(groupId)
                }, { headers: authHeader() });

                location.reload();
            } catch (e) {
                const serverMsg = e.response?.data || e.message;
                alert(`Ошибка при сохранении: ${serverMsg}`);
            }

            return;
        }

        // mode === "create" — массовое добавление выбранных студентов
        const selectedOptions = Array.from(studentsSelect.selectedOptions);

        if (selectedOptions.length === 0) return alert("Выберите хотя бы одного студента");

        for (const opt of selectedOptions) {
            try {
                await axios.post(`${API}/new_user_group`, {
                    user_id: Number(opt.value),
                    group_id: Number(groupId)
                }, { headers: authHeader() });
            } catch (e) {
                const serverMsg = e.response?.data || e.message;
                alert(`Ошибка при добавлении студента (id=${opt.value}): ${serverMsg}`);
                return;
            }
        }

        location.reload();
    });

    document.querySelectorAll(".edit-ug-btn").forEach(btn => {
        btn.addEventListener("click", () => {

            const row = btn.closest("tr");

            openModal({
                id: row.dataset.id,
                userId: row.dataset.userId,
                groupId: row.dataset.groupId
            });
        });
    });

    document.querySelectorAll(".delete-ug-btn").forEach(btn => {
        btn.addEventListener("click", async () => {

            const row = btn.closest("tr");
            const id = row.dataset.id;

            if (!confirm("Удалить связь?")) return;

            await axios.delete(`${API}/delete_user_group/${id}`, { headers: authHeader() });

            row.remove();
        });
    });

});
