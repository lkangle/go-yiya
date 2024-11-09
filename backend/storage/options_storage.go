package storage

import (
	"log"
	"sync"
	"yiya-v2/backend/consts"
	"yiya-v2/backend/types"

	"gopkg.in/yaml.v3"
)

type OptionStorage struct {
	storage *localStorage
	mutex   sync.Mutex
	option  *types.SystemOptions
}

func NewOptionStorage() *OptionStorage {
	storage := NewLocalStore("options.yaml")
	log.Printf("[options path]: %s\n", storage.ConfPath)
	return &OptionStorage{
		storage: storage,
		option:  nil,
	}
}

func (p *OptionStorage) get() (ret types.SystemOptions) {
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

func (p *OptionStorage) save(pf *types.SystemOptions) error {
	b, err := yaml.Marshal(pf)
	if err != nil {
		return err
	}

	if err = p.storage.Store(b); err != nil {
		return err
	}
	return nil
}

func (p *OptionStorage) Get() (ret types.SystemOptions) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	ret = p.get()
	ret.WindowWidth = max(ret.WindowWidth, consts.MIN_WINDOW_WIDTH)
	ret.WindowHeight = max(ret.WindowHeight, consts.MIN_WINDOW_HEIGHT)
	return
}

func (p *OptionStorage) Save(pf *types.SystemOptions) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	return p.save(pf)
}

func (p *OptionStorage) Default() types.SystemOptions {
	return types.SystemOptions{
		WindowWidth:   consts.DEFAULT_WINDOW_WIDTH,
		WindowHeight:  consts.DEFAULT_WINDOW_HEIGHT,
		IsAllowsOnTop: false,
	}
}
