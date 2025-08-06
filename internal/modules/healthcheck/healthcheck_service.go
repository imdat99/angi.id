package healthcheck

import (
	"context"
	"errors"
	"runtime"

	"angi.id/internal/container"
	"angi.id/internal/shared"
	"github.com/sirupsen/logrus"
)

type HealthCheckService interface {
	DBcheck() error
	MemoryHeapCheck() error
	RedisCheck() error
}
type healthCheckService struct {
	Log       *logrus.Logger
	container *container.Container
}

func HewHealthCheckService(
	ctn *container.Container) HealthCheckService {
	return &healthCheckService{
		Log:       shared.Log,
		container: ctn,
	}
}
func (m *healthCheckService) DBcheck() error {
	if err := m.container.DbPool.Ping(); err != nil {
		m.Log.Errorf("failed to ping the database: %v", err)
		return err
	}

	return nil
}

func (m *healthCheckService) MemoryHeapCheck() error {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats) // Collect memory statistics

	heapAlloc := memStats.HeapAlloc            // Heap memory currently allocated
	heapThreshold := uint64(300 * 1024 * 1024) // Example threshold: 300 MB

	m.Log.Infof("Heap Memory Allocation: %v bytes", heapAlloc)

	// If the heap allocation exceeds the threshold, return an error
	if heapAlloc > heapThreshold {
		m.Log.Errorf("Heap memory usage exceeds threshold: %v bytes", heapAlloc)
		return errors.New("heap memory usage too high")
	}

	return nil
}

func (m *healthCheckService) RedisCheck() error {
	ctx := context.Background()
	status := m.container.RedisClient.Ping(ctx)
	if err := status.Err(); err != nil {
		m.Log.Errorf("failed to ping Redis: %v", err)
		return err
	}

	return nil
}
