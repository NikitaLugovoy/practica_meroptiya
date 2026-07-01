document.getElementById("submitBtn").addEventListener("click", async () => {

    const eventId = document.getElementById("eventId").value;
    const resultText = document.getElementById("resultText").value;

    if (!eventId || !resultText.trim()) {
        alert("Заполните все поля");
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

        document.getElementById("message").textContent =
            "Результат сохранён, статус мероприятия изменён.";

        document.getElementById("resultText").value = "";

    } catch (err) {
        console.error(err);
        alert("Ошибка: " + err.message);
    }
});