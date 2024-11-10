package storage

import (
	"log"
	"sync"
	"yiya-v2/backend/consts"
	"yiya-v2/backend/types"

	"gopkg.in/yaml.v3"
)

type OptionsStorage struct {
	storage *localStorage
	mutex   sync.Mutex
	option  *types.AppOptions
}

func NewOptionStorage() *OptionsStorage {
	storage := NewLocalStore("options.yaml")
	log.Printf("[options path]: %s\n", storage.ConfPath)
	return &OptionsStorage{
		storage: storage,
		option:  nil,
	}
}

func (p *OptionsStorage) get() (ret types.AppOptions) {
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

func (p *OptionsStorage) save(pf *types.AppOptions) error {
	b, err := yaml.Marshal(pf)
	if err != nil {
		return err
	}

	if err = p.storage.Store(b); err != nil {
		return err
	}
	return nil
}

func (p *OptionsStorage) Get() (ret types.AppOptions) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	ret = p.get()
	ret.System.WindowWidth = max(ret.System.WindowWidth, consts.MIN_WINDOW_WIDTH)
	ret.System.WindowHeight = max(ret.System.WindowHeight, consts.MIN_WINDOW_HEIGHT)
	return
}

func (p *OptionsStorage) Save(pf *types.AppOptions) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	return p.save(pf)
}

func (p *OptionsStorage) Default() types.AppOptions {
	return types.AppOptions{
		System: types.SystemOptions{
			WindowWidth:   consts.DEFAULT_WINDOW_WIDTH,
			WindowHeight:  consts.DEFAULT_WINDOW_HEIGHT,
			IsAlwaysOnTop: false,
		},
		Compress: types.CompressOptions{
			Quality:   consts.BEST_QUALITY,
			Override:  true,
			NewSuffix: "yiya",
		},
	}
}
