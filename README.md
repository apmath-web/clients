# Client-service
Micro service for handling user information

## Interfaces
- ViewModels
  - в данные модели преобразуется json из запроса.
  - в них происходит валидация данных
- Validation
  - харнение валидационных сообщений
- Converter 
  - преобразование между моделями
- Repository 
  - хранение моделей
- ApplicationModels
  - внутренняя модель приложения
  - данные полчаемые из базы данных преобразуются в эту модель