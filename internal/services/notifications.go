package services

import (
	"bytes"
	"context"
	"ecostars-fake-api/internal/domain" // Asumo que este es el import correcto para tu repo
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time" // <-- FIX: Importar 'time'
)

// Interfaz para tu repositorio de suscripciones
// (Esto es una buena práctica en lugar de usar la implementación concreta)
type SubscriptionRepository interface {
	GetAllSubscriptions(ctx context.Context) ([]domain.Subscription, error)
	// (domain.Subscription debe tener un campo 'URL' string)
}

type NotificationService struct {
	subscriptionRepo SubscriptionRepository // Usar la interfaz
	httpClient       *http.Client
	jobChannel       chan domain.MetricItem
	httpJobChannel   chan httpJob
	workerPoolSize   int
}

type httpJob struct {
	Url     string
	Payload []byte
}

func NewNotificationService(subscriptionRepo SubscriptionRepository, workerPoolSize int) *NotificationService {
	return &NotificationService{
		subscriptionRepo: subscriptionRepo,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		jobChannel:     make(chan domain.MetricItem, 10),
		httpJobChannel: make(chan httpJob, 2000),
		workerPoolSize: workerPoolSize,
	}
}

func (n *NotificationService) Start(ctx context.Context, wg *sync.WaitGroup) {
	log.Println("Notification Service starting...")
	wg.Add(1 + n.workerPoolSize)
	go n.getSubscribersWorker(ctx, wg)
	for i := 0; i < n.workerPoolSize; i++ {
		go n.httpSenderWorker(ctx, wg, i+1)
	}
}

func (n *NotificationService) NotifyMetricUpdate(metricItem domain.MetricItem) {
	select {
	case n.jobChannel <- metricItem:
		log.Printf("Enqueued notification job for metric item ID %d", metricItem.ID)
	default:
		log.Printf("Notification job channel full, dropping notification for metric item ID %d", metricItem.ID)
	}
}

func (n *NotificationService) getSubscribersWorker(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	log.Println("getSubscribersWorker started")

	for {
		select {
		case <-ctx.Done(): // Primero comprobar si nos han cancelado
			log.Println("getSubscribersWorker shutting down")
			return

		case metricItem := <-n.jobChannel: // Segundo, procesar trabajo
			// (Ya no se necesita el 'select' interno)
			payload, err := json.Marshal(metricItem)
			if err != nil {
				log.Printf("failed to marshal metric item: %v", err)
				continue
			}

			// (Asumo que tu domain.Subscription tiene un campo 'URL')
			subscribers, err := n.subscriptionRepo.GetAllSubscriptions(ctx)
			if err != nil {
				log.Printf("failed to get subscribers: %v", err)
				continue
			}
			if len(subscribers) == 0 {
				continue // No subscribers to notify
			}

			log.Printf("Fanning-out notification for metric %s to %d subscribers", metricItem.ItemType, len(subscribers))
			for _, sub := range subscribers {
				job := httpJob{
					Url:     sub.URL, // Asegúrate que 'sub' tenga el campo 'URL'
					Payload: payload,
				}
				// Enviar al siguiente canal, pero respetando el apagado
				select {
				case n.httpJobChannel <- job:
					// trabajo encolado para los httpSenders
				case <-ctx.Done():
					log.Println("getSubscribersWorker shutting down, job not sent.")
					return
				}
			}
		}
	}
	// --------------------------------------------------------
}

func (n *NotificationService) httpSenderWorker(ctx context.Context, wg *sync.WaitGroup, workerID int) {
	defer wg.Done()
	log.Printf("httpSenderWorker %d started", workerID)

	for {
		select {
		case job := <-n.httpJobChannel:
			log.Printf("[Worker %d] Sending notification to %s", workerID, job.Url)

			req, err := http.NewRequestWithContext(ctx, "POST", job.Url, bytes.NewBuffer(job.Payload))
			if err != nil {
				log.Printf("[Worker %d] ERROR creating request for %s: %v", workerID, job.Url, err)
				continue
			}
			req.Header.Set("Content-Type", "application/json")

			resp, err := n.httpClient.Do(req)
			if err != nil {
				log.Printf("[Worker %d] ERROR sending POST to %s: %v", workerID, job.Url, err)
				continue
			}

			if resp.StatusCode < 200 || resp.StatusCode >= 300 {
				log.Printf("[Worker %d] FAILED notification to %s, status: %s", workerID, job.Url, resp.Status)
			}
			// ¡Importante! Cerrar el body para reutilizar la conexión
			resp.Body.Close()

		case <-ctx.Done():
			log.Printf("httpSenderWorker %d shutting down.", workerID)
			return
		}
	}
}
