document.getElementById('loginForm').addEventListener('submit', async function(event) {
    event.preventDefault();

    // Собираем данные из формы
    const username = document.getElementById('username').value;
    const password = document.getElementById('password').value;

    console.log(username)
    console.log(password)
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

            var expiresDate = new Date(Date.now() + expiresIn * 2000).toUTCString(); // преобразуем в строку

            // Записываем токен в куки
            document.cookie = `token=${accessToken}; expires=${expiresDate}; path=/; secure; samesite=strict`;
            document.cookie = `token_expires=${expiresDate}; expires=${expiresDate}; path=/; secure; samesite=strict`;
            document.cookie = `username=${username}; expires=${expiresDate}; path=/; secure; samesite-strict`;

            // Перенаправляем на домашнюю страницу
            window.location.href = '/vacancies';
        } else {
            // Если произошла ошибка, показываем сообщение
            document.getElementById('error-message').style.display = 'block';
        }
    } catch (error) {
        console.error('Ошибка отправки данных:', error);
        document.getElementById('error-message').style.display = 'block';
    }
});
