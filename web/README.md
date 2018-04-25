# BibTex Web

Веб-приложение, написанное на TypeScript + Bootstrap.

Для сборки используется browserify и npm-модули.

Модули описаны в package.json(по идее этого достаточно, чтобы скачать и заимпортить все нужные библиотеки).

Команда сборки(в fish):  
`tsc main.ts; and browserify main.js -o bundle.js`

В директории лежит файл package.json. Предполагается, что все библиотеки будут находиться в папке src/dist.

Следует скопировать package.json в dist, и вызвать
`npm install --prefix dist/`, находять в `src`.

main.js лежит в исходной директории ради консистентности в путях к библиотекам с main.ts.