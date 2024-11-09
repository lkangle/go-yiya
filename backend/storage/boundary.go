package storage

import (
	"log"
	"sync"
	"yiya-v2/backend/consts"
	"yiya-v2/backend/types"

	"gopkg.in/yaml.v3"
)

type BoundaryStorage struct {
	storage *localStorage
	mutex   sync.Mutex
}

func NewBoundaryStorage() *BoundaryStorage {
	storage := NewLocalStore("boundary.yaml")
	log.Printf("boundary path: %s\n", storage.ConfPath)
	return &BoundaryStorage{
		storage: storage,
	}
}

func (p *BoundaryStorage) Default() types.WindowBoundary {
	return types.WindowBoundary{
		WindowWidth:  consts.DEFAULT_WINDOW_WIDTH,
		WindowHeight: consts.DEFAULT_WINDOW_HEIGHT,
	}
}

func (p *BoundaryStorage) get() (ret types.WindowBoundary) {
	ret = p.Default()
	b, err := p.storage.Load()
	if err != nil {
		return
	}

	if err = yaml.Unmarshal(b, &ret); err != nil {
		ret = p.Default()
		return
	}
	return
}

func (p *BoundaryStorage) save(pf *types.WindowBoundary) error {
	b, err := yaml.Marshal(pf)
	if err != nil {
		return err
	}

	if err = p.storage.Store(b); err != nil {
		return err
	}
	return nil
}

func (p *BoundaryStorage) Get() (ret types.WindowBoundary) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	ret = p.get()
	ret.WindowWidth = max(ret.WindowWidth, consts.MIN_WINDOW_WIDTH)
	ret.WindowHeight = max(ret.WindowHeight, consts.MIN_WINDOW_HEIGHT)
	return
}

func (p *BoundaryStorage) Save(pf *types.WindowBoundary) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	return p.save(pf)
}

func (p *BoundaryStorage) Restore() types.WindowBoundary {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	pf := p.Default()
	p.save(&pf)
	return pf
}
