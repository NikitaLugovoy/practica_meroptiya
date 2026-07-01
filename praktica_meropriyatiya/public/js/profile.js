const API = "http://localhost:8080";

function authHeader() {
    const token = localStorage.getItem("token");
    return token ? { Authorization: `Bearer ${token}` } : {};
}

document.getElementById("saveProfileBtn").onclick = async () => {

    const id = window.USER_ID;

    const password = document.getElementById("password").value;

    const payload = {
        name: document.getElementById("name").value,
        surname: document.getElementById("surname").value,
        login: document.getElementById("login").value,
        email: document.getElementById("email").value,
        phone_number: document.getElementById("phone_number").value,
        url_avatar: document.getElementById("url_avatar").value,
        user_role: document.getElementById("role").value // важно: теперь фиксируем поле
    };

    if (password && password.trim() !== "") {
        payload.password = password;
    }

    try {
        const res = await axios.put(
            `${API}/update_user/${id}`,
            payload,
            { headers: authHeader() }
        );

        console.log("UPDATED:", res.data);

        location.reload();

    } catch (e) {
        console.log("ERROR:", e.response?.data || e.message);
        alert("Ошибка сохранения");
    }
};