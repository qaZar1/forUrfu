document.getElementById('loginForm').addEventListener('submit', async function(event) {
    event.preventDefault();

    // Собираем данные из формы
    const username = document.getElementById('username').value;
    const password = document.getElementById('password').value;

    // Создаем объект FormData для отправки данных на сервер
    const formData = new FormData();
    formData.append('username', username);
    formData.append('password', password);

    

    try {
        // Отправляем данные на сервер через Fetch API
        const response = await fetch('/submitToLogin', {
            method: 'POST',
            body: formData
        });

        if (response.ok) {
            var data = await response.json();
            var accessToken = data.access_token;
            var expiresIn = data.expires_in;
            var tokenType = data.token_type;

            var expiresDate = new Date(Date.now() + expiresIn * 2000).toUTCString(); // преобразуем в строку

            // Записываем токен в куки
            document.cookie = `token_employers=${accessToken}; expires=${expiresDate}; path=/; secure; samesite=strict`;
            document.cookie = `token_employers_expires=${expiresDate}; expires=${expiresDate}; path=/; secure; samesite=strict`;
            document.cookie = `username_employers=${username}; expires=${expiresDate}; path=/; secure; samesite-strict`;

            // Перенаправляем на домашнюю страницу
            window.location.href = '/employers/vacancies';
        } else {
            // Если произошла ошибка, показываем сообщение
            document.getElementById('error-message').style.display = 'block';
        }
    } catch (error) {
        console.error('Ошибка отправки данных:', error);
        document.getElementById('error-message').style.display = 'block';
    }
});