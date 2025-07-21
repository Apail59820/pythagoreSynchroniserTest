# Pythagore Synchroniser

Cet outil Go automatise l'envoi de vos factures vers la plateforme **FNE** (Facture Normalisée Électronique) de la DGI. Depuis 2025, toutes les factures doivent être digitalisées et certifiées via la FNE (articles 384-385 du CGI, LPF).

## Fonctionnement

1. Récupération des nouvelles factures dans la base PostgreSQL `invoice_test`.
2. Conversion des factures au format requis par la FNE.
3. Envoi sécurisé de la facture à l'API FNE pour obtenir la référence et le token de vérification.
4. Sauvegarde de ces métadonnées dans `invoice_metadata.json` pour audit ou génération de QR code.
5. Stockage de l'identifiant de la dernière facture synchronisée dans `sync_state.json` afin de reprendre automatiquement la synchronisation.

La tâche s'exécute périodiquement (10 s par défaut) pour ne traiter que les nouvelles factures.

## Configuration

Définissez les variables d'environnement suivantes (un fichier `.env` peut être utilisé durant le développement) :

- `DB_USER` – utilisateur PostgreSQL
- `DB_PASSWORD` – mot de passe
- `DB_HOST` – hôte de la base
- `DB_PORT` – port d'écoute
- `DB_NAME` – nom de la base (par défaut `invoice_test`)
- `SYNC_INTERVAL` – intervalle entre deux synchronisations en secondes (défaut : `10`)
- `FNE_API_URL` – URL de l'API FNE
- `FNE_API_TOKEN` – jeton d'authentification pour l'API

## Lancement

```bash
go run ./...
```

Pour exécuter les tests unitaires :

```bash
go test ./...
```

Ce projet simplifie la mise en conformité avec la facturation électronique obligatoire en Côte d'Ivoire.
