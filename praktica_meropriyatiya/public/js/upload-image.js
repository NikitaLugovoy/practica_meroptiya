document.getElementById("uploadBtn").addEventListener("click", async () => {

    const eventId = document.getElementById("eventId").value;
    const fileInput = document.getElementById("imageInput");

    if (!fileInput.files.length) {
        alert("Выберите файл");
        return;
    }

    const formData = new FormData();
    formData.append("eventId", eventId);
    formData.append("image", fileInput.files[0]);

    try {
        const res = await fetch("/api/upload-image", {
            method: "POST",
            body: formData
        });

        const data = await res.json();

        if (!res.ok) {
            throw new Error(data.error || "Ошибка загрузки");
        }

        document.getElementById("result").innerHTML =
            `Загружено: <img src="${data.url_image}" width="150">`;

    } catch (err) {
        console.error(err);
        alert("Ошибка: " + err.message);
    }
});