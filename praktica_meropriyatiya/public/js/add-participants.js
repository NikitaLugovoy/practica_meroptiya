const API = "http://localhost:8080";

function authHeader() {
    const token = localStorage.getItem("token");
    return token ? { Authorization: `Bearer ${token}` } : {};
}

// ---------------- SWITCH UI ----------------
const mode = document.getElementById("mode");

mode.addEventListener("change", () => {

    document.getElementById("singleBlock").style.display = "none";
    document.getElementById("multiBlock").style.display = "none";
    document.getElementById("groupBlock").style.display = "none";

    if (mode.value === "single") {
        document.getElementById("singleBlock").style.display = "block";
    }

    if (mode.value === "multiple") {
        document.getElementById("multiBlock").style.display = "block";
    }

    if (mode.value === "group") {
        document.getElementById("groupBlock").style.display = "block";
    }
});

// ---------------- SUBMIT ----------------
document.getElementById("submitBtn").addEventListener("click", async () => {

    const eventId = Number(document.getElementById("eventId").value);
    const modeValue = mode.value;

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

        alert("Готово");

    } catch (err) {
        console.error(err);
        alert("Ошибка");
    }
});

async function addGroupToEvent(eventId, groupId) {
    const res = await fetch("/api/add-group", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            ...authHeader()
        },
        body: JSON.stringify({ eventId, groupId })
    });

    const data = await res.json();
    console.log(data);
}