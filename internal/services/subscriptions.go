package services

import (
	"context"
	"ecostars-fake-api/internal/domain"
	"ecostars-fake-api/internal/repo"
)

type SubscriptionService struct {
	SubscriptionRepository *repo.SubscriptionRepository
}

func NewSubscriptionService(subscriptionRepo *repo.SubscriptionRepository) *SubscriptionService {
	return &SubscriptionService{
		SubscriptionRepository: subscriptionRepo,
	}
}

func (s *SubscriptionService) GetAllSubscriptions(ctx context.Context) ([]domain.Subscription, error) {
	return s.SubscriptionRepository.GetAllSubscriptions(ctx)
}

func (s *SubscriptionService) GetSubscriptionByID(ctx context.Context, id uint) (domain.Subscription, error) {
	return s.SubscriptionRepository.GetSubscriptionByID(ctx, id)
}

func (s *SubscriptionService) CreateSubscription(ctx context.Context, newSubscription *domain.SubscriptionCreateRequest) (domain.Subscription, error) {
	return s.SubscriptionRepository.CreateSubscription(ctx, newSubscription)
}

func (s *SubscriptionService) DeleteSubscription(ctx context.Context, id uint) error {
	return s.SubscriptionRepository.DeleteSubscription(ctx, id)
}
