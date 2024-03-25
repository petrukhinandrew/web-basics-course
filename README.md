# HW 1 + 3

Кажется, получилось сделать задание полностью. Запускался через `npm start`

# Пояснение про сертификаты

Так как это отдельная возня для Ubuntu, распишу: сначала надо запустить `root-cert.sh` - получаем два файла, `rootCA.crt` надо добавить в trust и (в частности для firefox) импортировать [как тут](https://www.ibm.com/docs/en/devops-test-hub/10.5.3?topic=lists-importing-certificate-authority-into-mozilla-firefox-browser), потом `server-cert.sh` (там стоит поменять под себя конфиг) и можно запускаться