document.querySelectorAll("tr.clickable").forEach(row => {
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

            // Точечно обновляем строку без перезагрузки
            const statusCell = row.querySelector(".status");
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