# Pythagore Synchroniser

Petit projet Go pour se connecter à PostgreSQL et exécuter périodiquement une action de synchronisation.

## Variables d'environnement

- `DB_USER`
- `DB_PASSWORD`
- `DB_HOST`
- `DB_PORT`
- `DB_NAME`
- `SYNC_INTERVAL` (optionnel, en secondes, valeur par défaut : `10`)

Créez éventuellement un fichier `.env` pour définir ces variables lors du développement.

## Lancement

```bash
go run ./...
```
