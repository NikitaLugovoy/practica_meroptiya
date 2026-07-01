document.getElementById("registerBtn").onclick = async () => {
        const errorMsg = document.getElementById("errorMsg");
        const successMsg = document.getElementById("successMsg");

        errorMsg.classList.add("hidden");
        successMsg.classList.add("hidden");

        const payload = {
            name:         document.getElementById("name").value.trim(),
            surname:      document.getElementById("surname").value.trim(),
            login:        document.getElementById("login").value.trim(),
            password:     document.getElementById("password").value.trim(),
            email:        document.getElementById("email").value.trim(),
            phone_number: document.getElementById("phone").value.trim(),
            url_avatar:   document.getElementById("url_avatar").value.trim(),
            user_role:    "student"
        };

        if (!payload.name || !payload.surname || !payload.login || !payload.password || !payload.email) {
            errorMsg.innerText = "Заполните все обязательные поля";
            errorMsg.classList.remove("hidden");
            return;
        }

        try {
            const res = await fetch("http://localhost:8080/new_user", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify(payload)
            });

            if (!res.ok) {
                const text = await res.text();
                errorMsg.innerText = text;
                errorMsg.classList.remove("hidden");
                return;
            }

            successMsg.innerText = "Аккаунт создан! Перенаправляем...";
            successMsg.classList.remove("hidden");

            setTimeout(() => {
                window.location.href = "/login";
            }, 1500);

        } catch (e) {
            errorMsg.innerText = "Ошибка соединения с сервером";
            errorMsg.classList.remove("hidden");
        }
    };