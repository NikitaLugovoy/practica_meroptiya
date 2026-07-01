document.addEventListener("DOMContentLoaded", () => {

    let currentStatus = "all";
    let currentCategory = "all";

    let allEvents = [];
    let calendar;

    const statusButtons = document.querySelectorAll(".status-btn");
    const categoryButtons = document.querySelectorAll(".category-btn");

    // =====================
    // STATUS FILTER
    // =====================
    statusButtons.forEach(btn => {

        btn.addEventListener("click", () => {

            currentStatus = btn.dataset.status || "all";

            statusButtons.forEach(b => b.classList.remove("active"));
            btn.classList.add("active");

            applyFilters();
        });
    });

    // =====================
    // CATEGORY FILTER
    // =====================
    categoryButtons.forEach(btn => {

    btn.addEventListener("click", () => {

        currentCategory = btn.dataset.category || "all";

        categoryButtons.forEach(b => b.classList.remove("active"));
        btn.classList.add("active");

        applyFilters();
    });
});

    function applyFilters() {

        let filtered = allEvents;

        if (currentStatus !== "all") {
            filtered = filtered.filter(e => e.status === currentStatus);
        }

        if (currentCategory !== "all") {
            filtered = filtered.filter(e => e.category === currentCategory);
        }

        calendar.removeAllEvents();
        calendar.addEventSource(filtered);
    }

    // =====================
    // INIT CALENDAR
    // =====================
    allEvents = JSON.parse(
        document.getElementById("eventsData").textContent
    );

    calendar = new FullCalendar.Calendar(
        document.getElementById("calendar"),
        {
            initialView: "dayGridMonth",
            locale: "ru",
            timeZone: "local",
            events: allEvents,

            eventDidMount(info) {

                const category = info.event.extendedProps.category;

                const colors = {
                    "КОНФЕРЕНЦИЯ": "#0004ff",
                    "СУББОТНИК": "#00fff2",
                    "ОЛИМПИАДА": "#e100ff",
                    "ПРОФЕССИОНАЛЬНАЯ": "#e9ff85",
                    "ТРЕННИНГ": "#ec4899",
                    "СПОРТИВНАЯ": "#852d2d"
                };

                const color = colors[category] || "#3b82f6";

                info.el.style.backgroundColor = color;
                info.el.style.borderColor = color;
            },

            dateClick(info) {

                const d = info.date;

                let dayEvents = allEvents.filter(event => {

                    const eDate = new Date(event.start);

                    return (
                        eDate.getFullYear() === d.getFullYear() &&
                        eDate.getMonth() === d.getMonth() &&
                        eDate.getDate() === d.getDate()
                    );
                });

                if (currentStatus !== "all") {
                    dayEvents = dayEvents.filter(e => e.status === currentStatus);
                }

                if (currentCategory !== "all") {
                    dayEvents = dayEvents.filter(e => e.category === currentCategory);
                }

                renderDayEvents(
                    dayEvents,
                    d.toLocaleDateString("ru-RU")
                );
            }
        }
    );

    calendar.render();
});

function renderDayEvents(dayEvents, date) {

    document.getElementById("selectedDateTitle")
        .innerText = `Мероприятия на ${date}`;

    const container = document.getElementById("dayEvents");

    if (!dayEvents.length) {
        container.innerHTML = "<p>Нет мероприятий</p>";
        return;
    }

    container.innerHTML = dayEvents.map(event => {

        const isDone = event.status === "ПРОВЕДЕНО";

        return `
            <a href="/event-page/${event.id}" class="event-card-sidebar ${isDone ? "done" : "planned"}">

                <div class="event-card-status">
                    <span class="badge ${isDone ? "done" : "planned"}">
                        ${event.status}
                    </span>
                </div>

                <h3 class="event-card-title">${event.title}</h3>

                <div class="event-card-date">
                    📅 ${new Date(event.start).toLocaleString("ru-RU")}
                </div>

            </a>
        `;
    }).join("");
}

