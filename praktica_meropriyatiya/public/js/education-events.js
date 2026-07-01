document.addEventListener("DOMContentLoaded", () => {

    const participants = JSON.parse(
        document.getElementById("participantsData").textContent || "[]"
    );

    const userId = Number(window.USER_ID);

    function isJoined(eventId) {
        return participants.some(p =>
            Number(p.event_id) === eventId &&
            Number(p.user_id) === userId
        );
    }

    document.querySelectorAll(".join-btn").forEach(btn => {

        const eventId = Number(btn.dataset.eventId);
        const card = btn.closest(".education-card");

        // уже участвует при загрузке
        if (isJoined(eventId)) {
            btn.disabled = true;
            btn.innerText = "✓ Уже участвуете";
            card.classList.add("is-joined");
        }

        btn.addEventListener("click", async () => {

            if (btn.disabled) return;

            const confirmJoin = confirm("Вы точно хотите участвовать в мероприятии?");
            if (!confirmJoin) return;

            const message = btn.parentElement.querySelector(".join-message");

            try {

                const res = await fetch(`/join-event/${eventId}`, {
                    method: "POST"
                });

                const data = await res.json();

                if (data.success) {

                    btn.disabled = true;
                    btn.innerText = "✓ Уже участвуете";

                    card.classList.add("is-joined");

                    message.innerText = "Вы записаны на мероприятие";

                } else if (data.reason === "ALREADY_JOINED") {

                    btn.disabled = true;
                    btn.innerText = "✓ Уже участвуете";

                    card.classList.add("is-joined");

                    message.innerText = "Вы уже участвуете";
                }

            } catch (e) {
                message.innerText = "Ошибка записи";
            }
        });
    });

});

