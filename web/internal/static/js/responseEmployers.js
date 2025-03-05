document.addEventListener("DOMContentLoaded", function () {
    const ratingButton = document.getElementById("rating");
    const ratingModalEl = document.getElementById("ratingModal");
    const submitRating = document.getElementById("submitRating");
    const ratingInput = document.getElementById("ratingInput");

    if (!ratingButton || !ratingModalEl || !submitRating || !ratingInput) {
        console.error("Ошибка: Один из элементов (кнопка, модальное окно, поле ввода) не найден!");
        return;
    }

    const ratingModal = new bootstrap.Modal(ratingModalEl);

    // Открытие модального окна
    ratingButton.addEventListener("click", function () {
        ratingInput.value = ""; // Очистить поле перед показом
        ratingModal.show();
    });


    // Отправка рейтинга на сервер
    submitRating.addEventListener("click", async function () {
        const ratingValue = parseInt(ratingInput.value, 10);
        const seekerUsername = ratingButton.getAttribute("data-seeker-username");

        if (isNaN(ratingValue) || ratingValue < 1 || ratingValue > 5) {
            ("Введите число от 1 до 5!");
            return;
        }

        const formData = new FormData();
        formData.append("rating", ratingValue);
        formData.append("seeker-username", seekerUsername);

        try{
            const response = await fetch("/rate-seeker", {
                method: "POST",
                body: formData,
            });

            if (response.status === 200) {
                showMessage("Спасибо за вашу оценку!", true);
            } else {
                showMessage("Ошибка при отправке оценки. Попробуйте еще раз.", false);
            }
        }
        catch(error){
            console.error("Ошибка при отправке запроса:", error);
            showMessage("Ошибка сети. Попробуйте позже.", false);
        }
    });
});

// Функция для отображения сообщений (ошибок или успеха)
function showMessage(message, isSuccess = false) {
    const messageBox = document.getElementById("message-box");
    messageBox.textContent = message;
    
    // Убираем все возможные классы
    messageBox.classList.remove("d-none", "alert-danger", "alert-success");
    
    // Добавляем нужный класс
    messageBox.classList.add(isSuccess ? "alert-success" : "alert-danger");

    // Показываем сообщение
    messageBox.classList.remove("d-none");
}