# Vazifa Boshqaruv Tizimi

## Loyiha Tavsifi

Ushbu loyiha Go, PostgreSQL, MongoDB, Redis, Docker, CI/CD va boshqa texnologiyalar yordamida ishlab chiqilgan vazifa boshqaruv tizimini taqdim etadi. Loyihaning maqsadi foydalanuvchilarga vazifalarni boshqarish, ro'yxatdan o'tish va autentifikatsiya qilish imkoniyatlarini taqdim etishdir.

### Funksional

- **Foydalanuvchi Ro'yxatdan O'tishi:** Foydalanuvchi ro'yxatdan o'tgandan so'ng, uning elektron pochtasiga tasdiqlash kodi yuboriladi. Tasdiqlash `POST /user/verify` endpointi orqali amalga oshiriladi va muvaffaqiyatli tasdiqlash JWT token qaytaradi.
- **Autentifikatsiya:** Foydalanuvchi tizimga kirishi uchun login va parolini taqdim etishi kerak. Muvaffaqiyatli kirgandan so'ng JWT token oladi, bu token tizimda autentifikatsiya uchun ishlatiladi.
- **Vazifalar Boshqaruvi:**
    - **Vazifa Yaratish:** Foydalanuvchi yangi vazifa yaratadi. Vazifa ma'lumotlari PostgreSQL bazasiga saqlanadi, vazifa holatlari esa MongoDB bazasida saqlanadi.
    - **Vazifani Yangilash:** Vazifa ma'lumotlarini yangilash uchun `PATCH` so'rovi ishlatiladi. Bu so'rov vazifaning ma'lumotlarini o'zgartiradi, boshqa ma'lumotlarga ta'sir qilmaydi.
    - **Vazifalarni Ko'rish:** Vazifalar dinamik offset va limit parametrlariga ega bo'lib, shuningdek, field va value bo'yicha filtrlash imkonini beradi. Agar filtrlash parametrlarisiz so'rov yuborilsa, barcha vazifalar foydalanuvchiga ko'rsatiladi.

### Texnologiyalar

- **Golang:** Backend dasturlash tili.
- **PostgreSQL:** Vazifalar ma'lumotlarini saqlash uchun.
- **MongoDB:** Vazifa holatlarini saqlash uchun.
- **Redis:** Foydalanuvchi vaqtinchalik ma'lumotlarini saqlash uchun (tasdiqlash jarayoni davomida).
- **Docker:** Loyihani konteynerlarda ishlatish.
- **Swagger:** API hujjatlarni avtomatik ravishda yaratish uchun.
- **Rate Limiting:** So'rovlarni cheklash orqali tizimni himoya qilish.
- **Casbin:** Ruxsatnoma boshqaruvi.
- **Gin:** HTTP server yaratish uchun.

### Loyihani Ishga Tushirish

1. **Loyihani Klonlash:**
    ```bash
    git clone https://github.com/D1Y0RBEKORIFJONOV/todolist
    cd todolist
    ```

2. **Docker Konteynerlarini Ishga Tushurish:**
    ```bash
    docker-compose up --build
    ```

3. **Swagger API Hujjatlarini Ko'rish:**
   API hujjatlarni ko'rish uchun Swagger UI orqali kirish mumkin. Swagger UI manzili: [Swagger UI](http://52.59.220.158:9000/swagger/index.html#/auth/post_user_register).

4. **Test:**
   API endpointlarni test qilish uchun Postman yoki boshqa HTTP mijozini ishlatishingiz mumkin.

### API Endpoints

- **Foydalanuvchi Endpoints:**
    - Ro'yxatdan o'tish: `POST /user/register`
    - Tasdiqlash: `POST /user/verify`
    - Kirish: `POST /user/login`

- **Vazifa Endpoints:**
    - Vazifa yaratish: `POST /task/create`
    - Vazifani yangilash: `PATCH /task/update/{task_id}`
    - Vazifani olish: `GET /task/{field}/{value}`
    - Vazifalar ro'yxatini olish: `GET /task/tasks`
    - Vazifani o'chirish: `DELETE /task/delete/{task_id}`

### Clean Architecture

Loyiha **Clean Architecture** prinsiplariga asoslangan bo'lib, quyidagi qatlamlardan iborat:

- **Domain Layer:** Asosiy biznes qoidalari va modellar bu qatlamda joylashgan.
- **Application Layer:** Biznes jarayonlarini va xususiyatlarni boshqarish uchun xizmatlar va interfeyslar mavjud.
- **Infrastructure Layer:** Tashqi tizimlar bilan integratsiya (masalan, ma'lumotlar bazalari, tashqi API) shu qatlamda amalga oshiriladi.
- **Presentation Layer:** Foydalanuvchi interfeysi va API endpoints ushbu qatlamda joylashgan.

Har qanday savollar yoki muammolar yuzasidan, iltimos, loyiha muallifiga murojaat qiling.

---

Bu README fayli loyihangizning asosiy qismlarini va qanday ishlashini tushuntiradi, shuningdek, qanday qilib ishga tushirish va test qilish haqida ma'lumot beradi.
