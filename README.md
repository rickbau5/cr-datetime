# cr-datetime
Simple Golang service that returns the current UTC time in ISO 8601 format.

## Deploy to Cloud Run (from Source)
```bash
$ gcloud run deploy cr-datetime \
    --project <project-id> \
    --source . \
    --region us-central1 \
    --allow-unauthenticated
```