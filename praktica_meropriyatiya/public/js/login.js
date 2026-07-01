document.getElementById("loginBtn").onclick = async () => {
        const login = document.getElementById("login").value.trim();
        const password = document.getElementById("password").value.trim();
        const errorMsg = document.getElementById("errorMsg");

        errorMsg.classList.add("hidden");

        if (!login || !password) {
            errorMsg.innerText = "Заполните все поля";
            errorMsg.classList.remove("hidden");
            return;
        }

        try {
            const res = await fetch("http://localhost:8080/login", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({ login, password })
            });

            if (!res.ok) {
                const text = await res.text();
                errorMsg.innerText = text;
                errorMsg.classList.remove("hidden");
                return;
            }

            
            const data = await res.json();
            localStorage.setItem("token", data.token);
            document.cookie = `token=${data.token}; path=/`;
            window.location.href = "/events";

        } catch (e) {
            errorMsg.innerText = "Ошибка соединения с сервером";
            errorMsg.classList.remove("hidden");
        }
    };