const API = "http://localhost:8080";

function authHeader() {
    const token = localStorage.getItem("token");
    return token ? { Authorization: `Bearer ${token}` } : {};
}

let modal = document.getElementById("userModal");

document.getElementById("addUserBtn").onclick = () => {
    openModal();
};

document.getElementById("closeModal").onclick = () => {
    modal.classList.add("hidden");
};

function openModal(user = null) {

    modal.classList.remove("hidden");

    const passwordField = document.getElementById("password");

    if (user) {
        document.getElementById("modalTitle").innerText = "Редактировать";

        document.getElementById("userId").value = user.id;
        document.getElementById("name").value = user.name;
        document.getElementById("surname").value = user.surname;
        document.getElementById("login").value = user.login;
        document.getElementById("email").value = user.email;
        document.getElementById("phone").value = user.phone_number;
        document.getElementById("role").value = user.user_role;

        passwordField.placeholder = "Новый пароль (оставьте пустым чтобы не менять)";
        passwordField.value = "";
    } else {
        document.getElementById("modalTitle").innerText = "Добавить";

        document.getElementById("userId").value = "";
        document.getElementById("name").value = "";
        document.getElementById("surname").value = "";
        document.getElementById("login").value = "";
        document.getElementById("email").value = "";
        document.getElementById("phone").value = "";
        document.getElementById("role").value = "student";

        passwordField.placeholder = "Пароль";
        passwordField.value = "";
    }
}

document.getElementById("saveUser").onclick = async () => {

    const id = document.getElementById("userId").value;
    const password = document.getElementById("password").value;

    const payload = {
        name: document.getElementById("name").value,
        surname: document.getElementById("surname").value,
        login: document.getElementById("login").value,
        email: document.getElementById("email").value,
        phone_number: document.getElementById("phone").value,
        user_role: document.getElementById("role").value
    };

    if (password) {
        payload.password = password;
    }

    try {
        if (id) {
            await axios.put(`${API}/update_user/${id}`, payload, { headers: authHeader() });
        } else {
            await axios.post(`${API}/new_user`, payload, { headers: authHeader() });
        }

        location.reload();

    } catch (e) {
        alert("Ошибка: " + e.response.data);
    }
};

document.querySelectorAll(".edit-btn").forEach(btn => {

    btn.onclick = () => {

        const row = btn.closest("tr");

        openModal({
            id: row.dataset.id,
            name: row.querySelector(".u-name").innerText,
            surname: row.querySelector(".u-surname").innerText,
            login: row.querySelector(".u-login").innerText,
            email: row.querySelector(".u-email").innerText,
            phone_number: row.querySelector(".u-phone").innerText,
            user_role: row.querySelector(".u-role").innerText
        });
    };
});

document.querySelectorAll(".delete-btn").forEach(btn => {

    btn.onclick = async () => {

        const row = btn.closest("tr");
        const id = row.dataset.id;

        if (!confirm("Удалить пользователя?")) return;

        await axios.delete(`${API}/delete_user/${id}`, { headers: authHeader() });

        row.remove();
    };
});