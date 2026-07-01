const express = require("express");
const hbs = require("hbs");
const path = require("path");
const cookieParser = require("cookie-parser");

const api = require("./api");

const app = express();

app.set("view engine", "hbs");
app.set("views", path.join(__dirname, "views"));

hbs.registerHelper("eq", (a, b) => a === b);
hbs.registerPartials(path.join(__dirname, "views/partials"));

app.use(express.json());
app.use(express.static(path.join(__dirname, "public")));
app.use(cookieParser()); 

// ===== AUTH MIDDLEWARE =====
const jwt = require("jsonwebtoken");

async function requireAuth(req, res, next) {

    const token = req.cookies?.token;

    if (!token) {
        return res.redirect("/login");
    }

    try {

        const decoded = jwt.decode(token);

        const userId = decoded.user_id;

        // 2. получаем пользователя через API
        const user = await api.getUserById(userId, token);

        req.user = {
            id: decoded.user_id,
            user_role:  user.user_role
        };

        res.locals.user_role =  user.user_role;

        next();

    } catch (err) {

        return res.redirect("/login");
    }
}

hbs.registerHelper("formatDate", (date) => {
    if (!date) return "";

    const d = new Date(date);

    return d.toLocaleString("ru-RU", {
        year: "numeric",
        month: "2-digit",
        day: "2-digit",
        hour: "2-digit",
        minute: "2-digit"
    });
});

// ===== HOME =====
app.get("/", (req, res) => {
    res.redirect("/login");
});


// ===== EVENTS =====
app.get("/events", requireAuth, async (req, res) => {
    const token = req.cookies.token;


    const events = (await api.getAllEvents(token)) || []; 

    const calendarEvents = events.map(e => ({
        id: e.id,
        title: e.name,
        start: e.date_time,
        status: e.status,
        category: e.category_events,
        color:
        e.status === "ПРОВЕДЕНО"
            ? "#3bf64b" // зелёный
            : "#d19003" // синий
    }));

    const categories = [...new Set(
        events.map(e => e.category_events)
    )];

    res.render("events", {
        title: "Все мероприятия",
        events,
        eventsJson: JSON.stringify(calendarEvents),
        categories 
    });
});
// ===== USERS =====
app.get("/users", requireAuth, async (req, res) => {
    try {
        const token = req.cookies.token;
        const users = await api.getAllUsers(token);
        res.render("users", { 
            users, 
            title: "Пользователи" 
        });
    } catch (error) {
        console.error(error);
        res.status(500).send("Ошибка загрузки пользователей");
    }
});

// ===== CREATE EVENT PAGE =====
app.get("/create-event", requireAuth, async (req, res) => {
    const token = req.cookies.token;
    const users = await api.getAllUsers(token);

    res.render("create-event", {
        users, 
            title: "Новое мероприятие"
    });
});

app.get("/events-admin", requireAuth, async (req, res) => {
    try {
        const token = req.cookies.token;
        const events = await api.getAllEvents(token);
 
        res.render("events-admin", {
            title: "Управление мероприятиями",
            events
        });
    } catch (e) {
        console.error(e);
        res.status(500).send("Ошибка загрузки мероприятий");
    }
});


// ===== ADD PARTICIPANTS PAGE =====
app.get("/add-participants", requireAuth,async (req, res) => {
    const token = req.cookies.token;
    const students = await api.getUsersByRole("student",token);
    const organizers = await api.getUsersByRole("organizer",token);
    const responsible = await api.getUsersByRole("responsible",token);
    const groups = await api.getAllGroups(token);

    res.render("add-participants", {
        students,
        organizers,
        responsible,
        groups
    });
});

// ===== API ROUTE (group add) =====
app.post("/api/add-group", async (req, res) => {
    try {
        const token = req.cookies.token;
        const result = await api.addParticipantsByGroup(
            req.body.eventId,
            req.body.groupId,
            token
        );
        res.json(result);
    } catch (e) {
        res.status(500).json({ error: e.message });
    }
});



app.get("/event-participants-page/:id", requireAuth, async (req, res) => {
    const eventId = Number(req.params.id);
    const token = req.cookies.token;
    const event = await api.getEventById(eventId,token);
    const participants = await api.getParticipantsByEvent(eventId,token);

    res.render("event-participants", {
        event,
        participants
    });
});

const multer = require("multer");

// ===== MULTER CONFIG =====
const storage = multer.diskStorage({
    destination: (req, file, cb) => {
        cb(null, path.join(__dirname, "public", "uploads"));
    },
    filename: (req, file, cb) => {
        const unique = Date.now() + "_" + file.originalname;
        cb(null, unique);
    }
});

const upload = multer({ storage });

// ===== UPLOAD IMAGE PAGE =====
app.get("/upload-image",requireAuth, async (req, res) => {
    res.render("upload-image");
});

// ===== UPLOAD IMAGE ACTION =====
app.post("/api/upload-image", upload.single("image"), async (req, res) => {
    try {
        const eventId = Number(req.body.eventId);

        if (!req.file) {
            return res.status(400).json({ error: "Файл не передан" });
        }

        const urlImage = `/uploads/${req.file.filename}`;
        const token = req.cookies.token;
        const result = await api.addEventImage(eventId, urlImage,token);

        res.json({ ...result, url_image: urlImage });

    } catch (e) {
        console.error(e);
        res.status(500).json({ error: e.message });
    }
});

app.get("/event-images-page/:id",requireAuth, async (req, res) => {
    const eventId = Number(req.params.id);
    const token = req.cookies.token;
    const event = await api.getEventById(eventId,token);
    const images = await api.getEventImagesByEvent(eventId,token);

    res.render("event-images", {
        event,
        images
    });
});

// ===== ADD EVENT RESULT PAGE =====
app.get("/add-result", requireAuth,async (req, res) => {
    res.render("add-result");
});

// ===== ADD EVENT RESULT ACTION =====
app.post("/api/add-result", async (req, res) => {
    try {
        const eventId = Number(req.body.eventId);
        const resultText = req.body.result;
        const token = req.cookies.token;
        const saved = await api.addEventResult(eventId, resultText,token);
        const statusToggled = await api.toggleEventStatus(eventId,token);

        res.json({ saved, statusToggled });

    } catch (e) {
        console.error(e);
        res.status(500).json({ error: e.message });
    }
});


// ===== TOGGLE PARTICIPANT STATUS (AJAX) =====
app.post("/api/toggle-status", async (req, res) => {
    try {
        const participantId = Number(req.body.participantId);
        const token = req.cookies.token;
        const result = await api.toggleParticipantStatus(participantId,token);
        res.json(result);
    } catch (e) {
        console.error(e);
        res.status(500).json({ error: e.message });
    }
});

app.get("/event-page/:id",requireAuth, async (req, res) => {
    const eventId = Number(req.params.id);
    const token = req.cookies.token;

    const event = await api.getEventById(eventId,token);
    const participants = await api.getParticipantsByEvent(eventId,token) || [];;
    const images = await api.getEventImagesByEvent(eventId,token);
    const results = await api.getEventResultsByEvent(eventId,token);

    const allStudents = await api.getUsersByRole("student", token);

    const participantUserIds = new Set(
    participants.map(p => p.user_id)
);

const students = allStudents.filter(
    student => !participantUserIds.has(student.id)
);

    const groups = await api.getAllGroups(token);

    const isResponsible =
    req.user.user_role === "responsible" &&
    event.responsible_id === req.user.id;


    res.render("event-page", {
            title: event.name,
            event,
            participants,
            images,
            results,
            students,
            groups,
            isResponsible
        });
});

app.get("/groups",requireAuth, async (req, res) => {
    const token = req.cookies.token;
    const groups = await api.getAllGroups(token);

    res.render("groups", {
        title: "Группы",
        groups
    });
});

app.get("/user-groups",requireAuth, async (req, res) => {
    try {
        const token = req.cookies.token;
        const userGroups = await api.getAllUserGroups(token);
        const users = await api.getAllUsers(token);
        const groups = await api.getAllGroups(token);

        res.render("user-groups", {
            title: "Пользователи и группы",
            userGroups,
            users,
            groups
        });

    } catch (e) {
        console.error(e);
        res.status(500).send("Ошибка загрузки user-groups");
    }
});


app.get("/event-stats/:id",requireAuth, async (req, res) => {

    const eventId = Number(req.params.id);
    const token = req.cookies.token;
    const event = await api.getEventById(eventId,token);
    const participants = await api.getParticipantsByEvent(eventId,token);
    const groups = await api.getAllGroups(token);
    const userGroups = await api.getAllUserGroups(token);

    const total = participants.length;
    const came = participants.filter(p => p.participants_status === "ПРИШЁЛ").length;
    const absent = total - came;

    // user -> group map
    const userToGroup = {};
    userGroups.forEach(ug => {
        userToGroup[ug.user_id] = ug.group_id;
    });

    const groupStats = groups.map(g => {

        const groupParticipants = participants.filter(p =>
            userToGroup[p.user_id] === g.id
        );

        const groupTotal = groupParticipants.length;
        const groupCame = groupParticipants.filter(p =>
            p.participants_status === "ПРИШЁЛ"
        ).length;

        return {
            group_name: g.name,
            total: groupTotal,
            came: groupCame,
            absent: groupTotal - groupCame
        };
    }).filter(g => g.total > 0);

    res.render("event-stats", {
        title: "Статистика",
        event,
        total,
        came,
        absent,
        groupStats
    });
});


// ===== AUTH =====
app.get("/login", (req, res) => {
    res.render("login", { title: "Вход" });
});

app.get("/register", (req, res) => {
    res.render("register", { title: "Регистрация" });
});

// ===== EDUCATIONAL EVENTS =====
app.get("/education-events", requireAuth, async (req, res) => {

    const token = req.cookies.token;

    const events = await api.getAllEvents(token) || [];

    const filteredEvents = events.filter(event =>
        event.category_events === "УЧЕБНЫЕ" ||
        event.category_events === "ОЛИМПИАДА"
    );

    const participants = await api.getAllParticipants(token);
    res.render("education-events", {
        title: "Учебные мероприятия",
        participants: JSON.stringify(participants),
        events: filteredEvents,
        userId: req.user.id
    });
});

app.post("/join-event/:id", requireAuth, async (req, res) => {

    try {

        const token = req.cookies.token;

        const eventId = Number(req.params.id);
        const userId = Number(req.user.id);

        const event = await api.getEventById(eventId, token);

        if (!event) {
            return res.status(404).json({
                success: false,
                reason: "EVENT_NOT_FOUND"
            });
        }

        // 1. блок если мероприятие завершено
        if (event.status === "ПРОВЕДЕНО") {
            return res.json({
                success: false,
                reason: "EVENT_DONE"
            });
        }

        let participants = await api.getParticipantsByEvent(eventId, token);

        // защита от null/undefined/не массива
        if (!Array.isArray(participants)) {
            participants = [];
        }

        // 2. проверка: уже участвует ли пользователь
        const alreadyJoined = participants.some(p =>
            Number(p.user_id) === userId
        );

        if (alreadyJoined) {
            return res.json({
                success: false,
                reason: "ALREADY_JOINED"
            });
        }

        // 3. защита от гонки (два клика / два запроса)
        try {
            await api.addParticipant(eventId, userId, token);
        } catch (e) {

            // если backend уже не дал создать (unique constraint)
            return res.json({
                success: false,
                reason: "ALREADY_JOINED"
            });
        }

        return res.json({
            success: true
        });

    } catch (err) {
        console.log(err);

        return res.status(500).json({
            success: false,
            reason: "SERVER_ERROR"
        });
    }
});


app.post("/api/delete-participant", async (req, res) => {
    try {
        const token = req.cookies.token;
        const participantId = Number(req.body.participantId);

        const result = await api.deleteParticipant(participantId, token);

        res.json(result);
    } catch (e) {
        console.error(e);
        res.status(500).json({ error: e.message });
    }
});

app.get("/profile", requireAuth, async (req, res) => {
    const token = req.cookies.token;

    const user = await api.getUserById(req.user.id, token);

    res.render("profile", {
        title: "Личный кабинет",
        user
    });
});

app.get("/my-events", requireAuth, async (req, res) => {

    const token = req.cookies.token;
    const userId = req.user.id;

    const user = await api.getUserById(userId, token);
    const events = await api.getAllEvents(token) || [];
    const participants = await api.getAllParticipants(token);

    let myEvents = [];

    // =====================
    // STUDENT: только ПРИШЁЛ
    // =====================
    if (user.user_role === "student") {

        const eventIds = participants
            .filter(p =>
                Number(p.user_id) === userId &&
                p.participants_status === "ПРИШЁЛ"
            )
            .map(p => p.event_id);

        myEvents = events.filter(e => eventIds.includes(e.id));
    }

    // =====================
    // ORGANIZER
    // =====================
    if (user.user_role === "organizer") {
        myEvents = events.filter(e => e.organizer_id === userId);
    }

    // =====================
    // RESPONSIBLE
    // =====================
    if (user.user_role === "responsible") {
        myEvents = events.filter(e => e.responsible_id === userId);
    }

    // =====================
    // ADMIN
    // =====================
    if (user.user_role === "admin") {
        myEvents = events;
    }

    const totalEvents = myEvents.length;

    const completedEvents = myEvents.filter(
        e => e.status === "ПРОВЕДЕНО"
    ).length;

    res.render("my-events", {
        title: "Мои мероприятия",
        events: myEvents,
        role: user.user_role,
        totalEvents,
        completedEvents
    });
});



app.listen(3000, () => {
    console.log("Node frontend running on http://localhost:3000");
});