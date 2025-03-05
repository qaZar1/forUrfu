document.getElementById("registerForm").addEventListener("submit", async function(event) {
    event.preventDefault(); // Останавливаем стандартное поведение формы

    // Собираем данные из полей формы
    const name = document.getElementById("name").value;
    const username = document.getElementById("username").value;
    const password = document.getElementById("password").value;
    const resume = document.getElementById("resume").value;

    if (name == "" || username == "" || password == "" || resume == "" ){
        showMessage("Не все поля заполнены", false)
        return
    }

    if (!validateURL("resume")) {
        showMessage("Неправильно указана ссылка на резюме", false)
        return
    }

    // Создаем объект FormData
    const formData = new FormData();
    formData.append("name", name);
    formData.append("username", username);
    formData.append("password", password);
    formData.append("resume", resume);

    try {
        // Отправляем данные на сервер
        const response = await fetch("/submit", {
            method: "POST",
            body: formData
        });

        if (response.status === 204) {
            redirectWithCountdown("/vacancies", 5);
        } else if (response.status === 400) {
            showMessage("Ошибка: пользователь с таким username уже существует. Попробуйте другой!", false);
        } else if (response.status === 401) {
            showMessage("Ошибка: у вас нет прав для выполнения этого запроса.", false);
        } else if (response.status === 409) {
            showMessage("Ошибка: пользователь с таким username уже существует.", false);
        } else if (response.status === 500) {
            showMessage("Ошибка: произошла внутренняя ошибка сервера. Попробуйте позже.", false);
        } else {
            showMessage("Ошибка: неизвестная ошибка. Попробуйте позднее.", false);
        }
    } catch (error) {
        console.error("Ошибка при отправке данных:", error);
        showMessage("Ошибка сети: не удалось связаться с сервером.", false);
    }
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

function validateURL(inputId) {
    const inputField = document.getElementById(inputId);
    const urlPattern = /^(https?:\/\/)?([\w-]+\.)+[\w-]{2,}(:\d+)?(\/[^\s]*)?$/i;

    if (!urlPattern.test(inputField.value.trim())) {
        showMessage("Введите корректную ссылку!", false);
        inputField.value = ""; // Очищает поле, если введено что-то некорректное
        return false;
    }
    return true;
}

function redirectWithCountdown(url, seconds = 5) {
    let timeLeft = Math.min(seconds, 5);

    function updateMessage() {
        showMessage(`Вы будете автоматически перенаправлены через ${timeLeft} секунд...`, true);
        if (timeLeft <= 0) {
            window.location.href = url;
        } else {
            timeLeft--;
            setTimeout(updateMessage, 1000);
        }
    }

    updateMessage();
}