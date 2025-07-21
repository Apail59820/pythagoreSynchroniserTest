package metrics

import "time"

var (
	// StartTime enregistre l'heure de démarrage de l'application.
	StartTime = time.Now()
	// LastSync est l'heure de la dernière synchronisation réussie.
	LastSync time.Time
	// LastSyncDuration est la durée de la dernière synchronisation.
	LastSyncDuration time.Duration
)

// RecordSync met à jour les informations sur la dernière synchronisation.
func RecordSync(d time.Duration) {
	LastSync = time.Now()
	LastSyncDuration = d
}
