# Pythagore Synchroniser

Petit projet Go pour se connecter à PostgreSQL et exécuter périodiquement une action de synchronisation.

## Variables d'environnement

- `DB_USER`
- `DB_PASSWORD`
- `DB_HOST`
- `DB_PORT`
- `DB_NAME`
- `SYNC_INTERVAL` (optionnel, en secondes, valeur par défaut : `10`)

Le programme mémorise l'identifiant de la dernière facture synchronisée dans
le fichier `sync_state.json` afin de n'importer que les nouvelles factures.

Créez éventuellement un fichier `.env` pour définir ces variables lors du développement.

## Lancement

```bash
go run ./...
```
