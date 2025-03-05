function saveResume() {
    // Получаем данные из поля резюме
    const resume = document.getElementById('resume').value;

    if (!validateURL("resume")) {
        showMessage("Неправильно указана ссылка на резюме", false)
        return
    }

    const formData = new FormData();
    formData.append('resume', resume);

    // Отправляем данные на бэкенд с использованием fetch
    fetch('/saveResume', {
        method: 'POST',
        body: formData
    })
    .then(response => {
        if (response.ok) {
            showMessage("Резюме успешно сохранено!", true);
        } else {
            showMessage("Ошибка при сохранении резюме", false);
        }
    })
    .catch(error => {
        console.error('Ошибка:', error);
        showErrorMessage('Ошибка при соединении с сервером');
    });
}

function logout() {
    const cookies = document.cookie.split("; ");
    // Устанавливаем истекший срок для каждого cookie
    cookies.forEach(cookie => {
        const [name] = cookie.split("=");
        document.cookie = `${name}=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;`;
    });

    // Перенаправление на серверный маршрут выхода
    window.location.href = '/login';
}

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

document.addEventListener("DOMContentLoaded", function () {
    document.querySelectorAll(".rating").forEach(ratingEl => {
        let rating = parseFloat(ratingEl.dataset.rating);
        let stars = ratingEl.querySelectorAll(".star");

        stars.forEach(star => {
            let value = parseInt(star.dataset.value); // Целочисленное значение звезды
            star.classList.remove("filled", "half");

            if (value <= Math.floor(rating)) {
                star.classList.add("filled"); // Полная звезда
            } else if (value === Math.ceil(rating) && rating % 1 >= 0.5) {
                star.classList.add("half"); // Половинка звезды (например, 3.5, 4.5)
            }
        });
    });
});