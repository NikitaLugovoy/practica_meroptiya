const API = "http://localhost:8080";

function authHeader() {
    const token = localStorage.getItem("token");
    return token ? { Authorization: `Bearer ${token}` } : {};
}

document.querySelectorAll(".delete-participant-btn").forEach(btn => {
    btn.addEventListener("click", async (e) => {

        e.stopPropagation(); // важно, чтобы не срабатывал toggle статуса

        const row = btn.closest(".floating-row");
        const participantId = row.dataset.id;

        if (!confirm("Удалить участника с мероприятия?")) return;

        try {
            const res = await fetch(`/api/delete-participant`, {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({ participantId })
            });

            if (!res.ok) throw new Error("Ошибка удаления");

            row.remove(); // сразу убираем из UI

        } catch (err) {
            console.error(err);
            alert("Не удалось удалить участника");
        }
    });
});

// ---------------- FLOATING PARTICIPANTS PANEL ----------------
const floatingHeader = document.getElementById("floatingHeader");
const floatingPanel = document.getElementById("floatingParticipants");

if (floatingHeader) {
    floatingHeader.addEventListener("click", () => {
        floatingPanel.classList.toggle("collapsed");
    });
}

document.querySelectorAll(".floating-row.clickable").forEach(row => {
    row.addEventListener("click", async () => {

        const participantId = row.dataset.id;

        try {
            const res = await fetch("/api/toggle-status", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({ participantId })
            });

            if (!res.ok) {
                throw new Error("Ошибка переключения статуса");
            }

            const statusCell = row.querySelector(".floating-status");
            const isCame = row.classList.contains("came");

            if (isCame) {
                row.classList.remove("came");
                row.classList.add("absent");
                statusCell.textContent = "ОТСУТСТВОВАЛ";
            } else {
                row.classList.remove("absent");
                row.classList.add("came");
                statusCell.textContent = "ПРИШЁЛ";
            }

        } catch (err) {
            console.error(err);
            alert("Не удалось изменить статус");
        }
    });
});

const toggleBtn = document.getElementById("toggleResultFormBtn");
const formBlock = document.getElementById("resultFormBlock");

if (toggleBtn) {
    toggleBtn.addEventListener("click", () => {
        const isHidden = formBlock.style.display === "none";
        formBlock.style.display = isHidden ? "block" : "none";
    });
}

const submitResultBtn = document.getElementById("submitResultBtn");

if (submitResultBtn) {
    submitResultBtn.addEventListener("click", async () => {

        const eventId = document.getElementById("resultSection").dataset.eventId;
        const resultText = document.getElementById("resultText").value;

        if (!resultText.trim()) {
            alert("Введите текст результата");
            return;
        }

        try {
            const res = await fetch("/api/add-result", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({ eventId, result: resultText })
            });

            const data = await res.json();

            if (!res.ok) {
                throw new Error(data.error || "Ошибка сохранения");
            }

            location.reload();

        } catch (err) {
            console.error(err);
            alert("Ошибка: " + err.message);
        }
    });
}

const toggleImageBtn = document.getElementById("toggleImageFormBtn");
const imageFormBlock = document.getElementById("imageFormBlock");

if (toggleImageBtn) {
    toggleImageBtn.addEventListener("click", () => {
        const isHidden = imageFormBlock.style.display === "none";
        imageFormBlock.style.display = isHidden ? "block" : "none";
    });
}

const submitImageBtn = document.getElementById("submitImageBtn");

if (submitImageBtn) {
    submitImageBtn.addEventListener("click", async () => {

        const eventId = document.getElementById("imageSection").dataset.eventId;
        const fileInput = document.getElementById("imageInput");

        if (!fileInput.files.length) {
            alert("Выберите файл");
            return;
        }

        const formData = new FormData();
        formData.append("eventId", eventId);
        formData.append("image", fileInput.files[0]);

        try {
            const res = await fetch("/api/upload-image", {
                method: "POST",
                body: formData
            });

            const data = await res.json();

            if (!res.ok) {
                throw new Error(data.error || "Ошибка загрузки");
            }

            location.reload();

        } catch (err) {
            console.error(err);
            alert("Ошибка: " + err.message);
        }
    });
}

// ---------------- ADD PARTICIPANTS BLOCK ----------------
const toggleAddParticipantBtn = document.getElementById("toggleAddParticipantBtn");
const addParticipantFormBlock = document.getElementById("addParticipantFormBlock");

if (toggleAddParticipantBtn) {
    toggleAddParticipantBtn.addEventListener("click", () => {
        const isHidden = addParticipantFormBlock.style.display === "none";
        addParticipantFormBlock.style.display = isHidden ? "block" : "none";
    });
}

const participantMode = document.getElementById("participantMode");

if (participantMode) {
    participantMode.addEventListener("change", () => {

        document.getElementById("singleBlock").style.display = "none";
        document.getElementById("multiBlock").style.display = "none";
        document.getElementById("groupBlock").style.display = "none";

        if (participantMode.value === "single") {
            document.getElementById("singleBlock").style.display = "block";
        }
        if (participantMode.value === "multiple") {
            document.getElementById("multiBlock").style.display = "block";
        }
        if (participantMode.value === "group") {
            document.getElementById("groupBlock").style.display = "block";
        }
    });
}

const submitParticipantBtn = document.getElementById("submitParticipantBtn");

if (submitParticipantBtn) {
    submitParticipantBtn.addEventListener("click", async () => {

        const eventId = Number(document.getElementById("addParticipantSection").dataset.eventId);
        const modeValue = participantMode.value;

        try {

            if (modeValue === "single") {
                await axios.post(`${API}/new_event_participant`, {
                    event_id: eventId,
                    user_id: Number(document.getElementById("singleUser").value)
                }, { headers: authHeader() });
            }

            if (modeValue === "multiple") {
                const ids = Array.from(
                    document.getElementById("multiUsers").selectedOptions
                ).map(o => Number(o.value));

                for (const id of ids) {
                    await axios.post(`${API}/new_event_participant`, {
                        event_id: eventId,
                        user_id: id
                    }, { headers: authHeader() });
                }
            }

            if (modeValue === "group") {
                await axios.post(`${API}/add_event_participants_by_group`, {
                    event_id: eventId,
                    group_id: Number(document.getElementById("groupSelect").value)
                }, { headers: authHeader() });
            }

            location.reload();

        } catch (err) {
            console.error(err);
            alert("Ошибка добавления участников");
        }
    });
}

// ---------------- PHOTO MODAL ----------------
const photoModal = document.getElementById("photoModal");
const photoModalImg = document.getElementById("photoModalImg");
const photoModalClose = document.getElementById("photoModalClose");

document.querySelectorAll(".event-photo-clickable").forEach(img => {
    img.addEventListener("click", () => {
        photoModalImg.src = img.dataset.full || img.src;
        photoModal.classList.add("open");
    });
});

function closePhotoModal() {
    photoModal.classList.remove("open");
    photoModalImg.src = "";
}

if (photoModalClose) {
    photoModalClose.addEventListener("click", closePhotoModal);
}

if (photoModal) {
    photoModal.addEventListener("click", (e) => {
        if (e.target === photoModal) closePhotoModal();
    });
}

document.addEventListener("keydown", (e) => {
    if (e.key === "Escape") closePhotoModal();
});