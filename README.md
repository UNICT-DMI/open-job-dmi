# open-job-dmi
[![CodeFactor](https://www.codefactor.io/repository/github/unict-dmi/open-job-dmi/badge)](https://www.codefactor.io/repository/github/unict-dmi/open-job-dmi)

## Live Demo
- [Telegram channel OpenJobDMI](https://t.me/OpenJobDMI) & [open-job-dmi website](https://open-job-dmi.unictdevs.com)


## Run with Docker Compose
```yaml
version: '2'
services: 
  open-job-dmi:
    build: .
    environment:
      - TELEGRAM_TOKEN=
      - CHANNEL_ID=
      - ADMIN_GROUP_ID=
      - RECAPTCHA_SITE_KEY=
      - RECAPTCHA_SECRET=
    ports:
      - 8080:8080
```
