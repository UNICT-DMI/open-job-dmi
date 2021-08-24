# open-job-dmi

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