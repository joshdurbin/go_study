//go:build ignore

package main

import (
	"errors"
	"fmt"
	"strings"
)

// REPOSITORY PATTERN
// ==================
// Intent: decouple business logic from data access mechanics.
// The business layer works against an interface; the storage implementation
// can be swapped (Postgres → SQLite → in-memory for tests) without touching logic.
//
// This is table-stakes for any infra tool with non-trivial data access.
// Your redislab tool, ProxySQL config agent, GCP reservation tooling —
// all of these would benefit from this separation.
//
// Go-specific advantage: interface satisfaction is implicit, so you can define
// the repository interface in the *domain* package and have the storage package
// implement it. The domain package never imports the storage package (inversion).
//
// Pattern for interviews: always define the interface first, then the implementation.
// Shows you're thinking about boundaries before code.

// ── Domain types (no storage dependencies) ──────────────────────────────────

type Cluster struct {
	ID       string
	Name     string
	Host     string
	Port     int
	Replicas int
	Tags     map[string]string
}

var ErrNotFound = errors.New("not found")
var ErrAlreadyExists = errors.New("already exists")

// ── Repository interface — defined in the domain, implemented by storage ─────

type ClusterRepository interface {
	Create(c Cluster) error
	Get(id string) (Cluster, error)
	Update(c Cluster) error
	Delete(id string) error
	List() ([]Cluster, error)
	FindByTag(key, value string) ([]Cluster, error)
}

// ── In-memory implementation (used for tests and dev) ────────────────────────

type MemoryClusterRepository struct {
	clusters map[string]Cluster
}

func NewMemoryClusterRepository() *MemoryClusterRepository {
	return &MemoryClusterRepository{clusters: make(map[string]Cluster)}
}

func (r *MemoryClusterRepository) Create(c Cluster) error {
	if _, exists := r.clusters[c.ID]; exists {
		return fmt.Errorf("cluster %s: %w", c.ID, ErrAlreadyExists)
	}
	r.clusters[c.ID] = c
	return nil
}

func (r *MemoryClusterRepository) Get(id string) (Cluster, error) {
	c, ok := r.clusters[id]
	if !ok {
		return Cluster{}, fmt.Errorf("cluster %s: %w", id, ErrNotFound)
	}
	return c, nil
}

func (r *MemoryClusterRepository) Update(c Cluster) error {
	if _, ok := r.clusters[c.ID]; !ok {
		return fmt.Errorf("cluster %s: %w", c.ID, ErrNotFound)
	}
	r.clusters[c.ID] = c
	return nil
}

func (r *MemoryClusterRepository) Delete(id string) error {
	if _, ok := r.clusters[id]; !ok {
		return fmt.Errorf("cluster %s: %w", id, ErrNotFound)
	}
	delete(r.clusters, id)
	return nil
}

func (r *MemoryClusterRepository) List() ([]Cluster, error) {
	result := make([]Cluster, 0, len(r.clusters))
	for _, c := range r.clusters {
		result = append(result, c)
	}
	return result, nil
}

func (r *MemoryClusterRepository) FindByTag(key, value string) ([]Cluster, error) {
	var result []Cluster
	for _, c := range r.clusters {
		if v, ok := c.Tags[key]; ok && v == value {
			result = append(result, c)
		}
	}
	return result, nil
}

// ── Business logic layer — only knows about the interface, not the storage impl ──

type ClusterService struct {
	repo ClusterRepository // interface, not concrete type
}

func NewClusterService(repo ClusterRepository) *ClusterService {
	return &ClusterService{repo: repo}
}

func (s *ClusterService) Provision(id, name, host string, port, replicas int) error {
	if strings.TrimSpace(name) == "" {
		return fmt.Errorf("cluster name cannot be empty")
	}
	if replicas < 1 {
		return fmt.Errorf("replicas must be >= 1")
	}
	return s.repo.Create(Cluster{
		ID:       id,
		Name:     name,
		Host:     host,
		Port:     port,
		Replicas: replicas,
		Tags:     make(map[string]string),
	})
}

func (s *ClusterService) ScaleReplicas(id string, newCount int) error {
	c, err := s.repo.Get(id)
	if err != nil {
		return fmt.Errorf("ScaleReplicas: %w", err)
	}
	if newCount < 1 {
		return fmt.Errorf("replica count must be >= 1")
	}
	c.Replicas = newCount
	return s.repo.Update(c)
}

func (s *ClusterService) TagCluster(id, key, value string) error {
	c, err := s.repo.Get(id)
	if err != nil {
		return fmt.Errorf("TagCluster: %w", err)
	}
	c.Tags[key] = value
	return s.repo.Update(c)
}

func (s *ClusterService) GetProductionClusters() ([]Cluster, error) {
	return s.repo.FindByTag("env", "production")
}

func main() {
	// Wire up with in-memory repo (swap for SQLiteClusterRepository in prod)
	repo := NewMemoryClusterRepository()
	svc := NewClusterService(repo)

	// Business operations — zero storage code visible here
	svc.Provision("cluster-1", "primary", "db1.internal", 3306, 3)
	svc.Provision("cluster-2", "analytics", "db2.internal", 3306, 2)
	svc.Provision("cluster-3", "staging", "db3.internal", 3307, 1)

	svc.TagCluster("cluster-1", "env", "production")
	svc.TagCluster("cluster-2", "env", "production")
	svc.TagCluster("cluster-3", "env", "staging")

	svc.ScaleReplicas("cluster-1", 5)

	prod, _ := svc.GetProductionClusters()
	fmt.Println("production clusters:")
	for _, c := range prod {
		fmt.Printf("  %s (%s) replicas=%d\n", c.Name, c.ID, c.Replicas)
	}

	// Error handling flows naturally through the interface
	err := svc.Provision("cluster-1", "duplicate", "x", 3306, 1)
	fmt.Println("duplicate:", errors.Is(err, ErrAlreadyExists)) // true

	_, err = svc.GetProductionClusters()
	fmt.Println("err:", err) // nil

	// Key insight: to switch to a real SQLite or Postgres backend,
	// you implement ClusterRepository on a new struct.
	// ClusterService and all business logic stay completely unchanged.
}
